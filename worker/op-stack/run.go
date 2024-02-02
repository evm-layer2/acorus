package op_stack

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/cornerstone-labs/acorus/common/global_const"
	"github.com/cornerstone-labs/acorus/common/tasks"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/event"
	"github.com/cornerstone-labs/acorus/database/worker"
)

type WorkerProcessor struct {
	db      *database.DB
	chainId string
	tasks   tasks.Group
}

func NewWorkerProcessor(db *database.DB, chainId string, shutdown context.CancelCauseFunc) (*WorkerProcessor, error) {
	workerProcessor := WorkerProcessor{
		db:      db,
		chainId: chainId,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in worker processor: %w", err))
		}},
	}
	return &workerProcessor, nil
}

func (b *WorkerProcessor) Start(ctx context.Context) error {
	ticker := time.NewTicker(time.Second)
	b.tasks.Go(func() error {
		for range ticker.C {
			if err := b.db.L2ToL1.UpdateTimeLeft(b.chainId); err != nil {
				log.Println(err.Error())
			}
		}
		return nil
	})
	tickerRun := time.NewTicker(time.Second * 5)
	b.tasks.Go(func() error {
		for range tickerRun.C {
			err := b.syncL2ToL1StateRoot()
			if err != nil {
				return err
			}
		}
		return nil
	})
	return nil
}

func (b *WorkerProcessor) syncL2ToL1StateRoot() error {
	blockNumber, err := b.db.StateRoots.GetLatestStateRootL2BlockNumber(b.chainId)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	if blockNumber == 0 {
		return nil
	}

	blockTimestamp, err := b.db.Blocks.BlockTimeStampByNum(b.chainId, blockNumber)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = b.db.L2ToL1.UpdateReadyForProvedStatus(b.chainId, blockTimestamp, 1)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	log.Println("update ready for proven status success")
	return nil
}

func (b *WorkerProcessor) markedL1ToL2Finalized() error {
	log.Println("start marked l1 to l2 finalized")
	finalizedList, err := b.db.RelayMessage.RelayMessageUnRelatedList(strconv.FormatUint(global_const.OpChinId, 10))
	if err != nil {
		return err
	}
	var depositL2ToL1List []worker.L1ToL2
	var needMarkDepositList []event.RelayMessage
	for i := range finalizedList {
		finalized := finalizedList[i]
		l1l2Tx := worker.L1ToL2{
			L2TransactionHash: finalized.RelayTransactionHash,
			L1BlockNumber:     finalized.BlockNumber,
		}
		withdrawTx, _ := b.db.L1ToL2.L1ToL2TransactionDeposit(strconv.FormatUint(global_const.OpChinId, 10), finalized.MessageHash)
		if withdrawTx != nil {
			depositL2ToL1List = append(depositL2ToL1List, l1l2Tx)
			needMarkDepositList = append(needMarkDepositList, finalized)
		}
	}
	if err := b.db.Transaction(func(tx *database.DB) error {
		if len(depositL2ToL1List) > 0 {
			if err := b.db.L1ToL2.MarkL1ToL2TransactionDepositFinalized(global_const.ChainId, depositL2ToL1List); err != nil {
				log.Println("Marked l2 to l1 transaction withdraw proven fail", "err", err)
				return err
			}
			if err := b.db.RelayMessage.MarkedRelayMessageRelated(strconv.FormatUint(global_const.OpChinId, 10), needMarkDepositList); err != nil {
				log.Println("Marked withdraw proven related fail", "err", err)
				return err
			}
			log.Println("marked deposit transaction success", "deposit size", len(depositL2ToL1List), "marked size", len(needMarkDepositList))
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (b *WorkerProcessor) markedL2ToL1Proven() error {
	log.Println("start marked l2 to l1 proven")
	provenList, err := b.db.WithdrawProven.WithdrawProvenUnRelatedList(strconv.FormatUint(global_const.OpChinId, 10))
	if err != nil {
		return err
	}
	var withdrawL2ToL1List []worker.L2ToL1
	var withdrawL2ToL1ListV0 []worker.L2ToL1
	var needMarkWithdrawList []event.WithdrawProven
	var needMarkWithdrawListV0 []event.WithdrawProven
	for i := range provenList {
		provenTxn := provenList[i]
		l2l1Tx := worker.L2ToL1{
			WithdrawTransactionHash: provenTxn.WithdrawHash,
			L1ProveTxHash:           provenTxn.ProvenTransactionHash,
			L1BlockNumber:           provenTxn.BlockNumber,
		}
		withdrawTx, _ := b.db.L2ToL1.L2ToL1TransactionWithdrawal(strconv.FormatUint(global_const.OpChinId, 10), provenTxn.WithdrawHash)
		if withdrawTx != nil {
			if withdrawTx.Version != 0 {
				withdrawL2ToL1List = append(withdrawL2ToL1List, l2l1Tx)
				needMarkWithdrawList = append(needMarkWithdrawList, provenTxn)
			} else {
				withdrawL2ToL1ListV0 = append(withdrawL2ToL1ListV0, l2l1Tx)
				needMarkWithdrawListV0 = append(needMarkWithdrawListV0, provenTxn)
			}
		}
	}
	if err := b.db.Transaction(func(tx *database.DB) error {
		if len(withdrawL2ToL1List) > 0 {
			if err := b.db.L2ToL1.MarkL2ToL1TransactionWithdrawalProven(strconv.FormatUint(global_const.OpChinId, 10), withdrawL2ToL1List); err != nil {
				log.Println("Marked l2 to l1 transaction withdraw proven fail", "err", err)
				return err
			}
			if err := b.db.WithdrawProven.MarkedWithdrawProvenRelated(strconv.FormatUint(global_const.OpChinId, 10), needMarkWithdrawList); err != nil {
				log.Println("Marked withdraw proven related fail", "err", err)
				return err
			}
			log.Println("marked proven transaction success", "withdraw size", len(provenList), "marked size", len(needMarkWithdrawList))
		}
		if len(withdrawL2ToL1ListV0) > 0 {
			if err := b.db.L2ToL1.MarkL2ToL1TransactionWithdrawalProven(strconv.FormatUint(global_const.OpChinId, 10), withdrawL2ToL1ListV0); err != nil {
				log.Println("Marked l2 to l1 transaction withdraw proven fail", "err", err)
				return err
			}
			if err := b.db.WithdrawProven.MarkedWithdrawProvenRelated(strconv.FormatUint(global_const.OpChinId, 10), needMarkWithdrawListV0); err != nil {
				log.Println("Marked withdraw proven related fail", "err", err)
				return err
			}
			log.Println("marked proven v0 transaction success", "withdraw size", len(provenList), "marked size", len(needMarkWithdrawList))
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (b *WorkerProcessor) markedL2ToL1Finalized() error {
	log.Println("start marked l2 to l1 finalized")
	withdrawList, err := b.db.WithdrawFinalized.WithdrawFinalizedUnRelatedList(strconv.FormatUint(global_const.OpChinId, 10))
	if err != nil {
		log.Println("fetch withdraw finalized un-related list fail", "err", err)
		return err
	}
	var withdrawL2ToL1List []worker.L2ToL1
	var withdrawL2ToL1ListV0 []worker.L2ToL1
	var needMarkWithdrawList []event.WithdrawFinalized
	var needMarkWithdrawListV0 []event.WithdrawFinalized
	for i := range withdrawList {
		finalizedTxn := withdrawList[i]
		l2l1Tx := worker.L2ToL1{
			WithdrawTransactionHash: finalizedTxn.WithdrawHash,
			L1FinalizeTxHash:        finalizedTxn.FinalizedTransactionHash,
			L1BlockNumber:           finalizedTxn.BlockNumber,
		}
		withdrawTx, _ := b.db.L2ToL1.L2ToL1TransactionWithdrawal(strconv.FormatUint(global_const.OpChinId, 10), finalizedTxn.WithdrawHash)
		if withdrawTx != nil {
			if withdrawTx != nil {
				if withdrawTx.Version != 0 {
					withdrawL2ToL1List = append(withdrawL2ToL1List, l2l1Tx)
					needMarkWithdrawList = append(needMarkWithdrawList, finalizedTxn)
				} else {
					withdrawL2ToL1ListV0 = append(withdrawL2ToL1ListV0, l2l1Tx)
					needMarkWithdrawListV0 = append(needMarkWithdrawListV0, finalizedTxn)
				}
			}
		}
	}
	if err := b.db.Transaction(func(tx *database.DB) error {
		if len(withdrawL2ToL1List) > 0 {
			if err := b.db.L2ToL1.MarkL2ToL1TransactionWithdrawalFinalized(strconv.FormatUint(global_const.OpChinId, 10), withdrawL2ToL1List); err != nil {
				log.Println("Marked l2 to l1 transaction withdraw finalized fail", "err", err)
				return err
			}
			if err := b.db.WithdrawFinalized.MarkedWithdrawFinalizedRelated(strconv.FormatUint(global_const.OpChinId, 10), needMarkWithdrawList); err != nil {
				log.Println("Marked withdraw finalized related fail", "err", err)
				return err
			}
			log.Println("marked finalized transaction success", "withdraw size", len(withdrawList), "marked size", len(needMarkWithdrawList))
		}
		if len(withdrawL2ToL1ListV0) > 0 {
			if err := b.db.L2ToL1.MarkL2ToL1TransactionWithdrawalFinalizedV0(strconv.FormatUint(global_const.OpChinId, 10), withdrawL2ToL1ListV0); err != nil {
				log.Println("Marked l2 to l1 transaction withdraw proven fail", "err", err)
				return err
			}
			if err := b.db.WithdrawFinalized.MarkedWithdrawFinalizedRelated(strconv.FormatUint(global_const.OpChinId, 10), needMarkWithdrawListV0); err != nil {
				log.Println("Marked withdraw proven related fail", "err", err)
				return err
			}
			log.Println("marked proven v0 transaction success", "withdraw size", len(withdrawL2ToL1ListV0), "marked size", len(needMarkWithdrawList))
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
