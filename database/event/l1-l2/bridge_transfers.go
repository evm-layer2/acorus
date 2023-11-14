package l1_l2

import (
	"errors"
	"fmt"
	"github.com/cornerstone-labs/acorus/event/processors/op-stack/mantle/op-bindings/predeploys"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cornerstone-labs/acorus/database/utils"
)

var (
	ZERO_ADDRESS = common.HexToAddress("0x0000000000000000000000000000000000000000")

	ETHDepositTokenPair    = TokenPair{LocalTokenAddress: ZERO_ADDRESS, RemoteTokenAddress: predeploys.BVM_ETHAddr}
	ETHWithdrawalTokenPair = TokenPair{LocalTokenAddress: predeploys.BVM_ETHAddr, RemoteTokenAddress: ZERO_ADDRESS}
	//TODO update the MNT ADDRESS
	MNTDepositTokenPair    = TokenPair{LocalTokenAddress: ZERO_ADDRESS, RemoteTokenAddress: ZERO_ADDRESS}
	MNTWithdrawalTokenPair = TokenPair{LocalTokenAddress: ZERO_ADDRESS, RemoteTokenAddress: ZERO_ADDRESS}
)

/**
 * Types
 */

type TokenPair struct {
	LocalTokenAddress  common.Address `gorm:"serializer:bytes"`
	RemoteTokenAddress common.Address `gorm:"serializer:bytes"`
}

type BridgeTransfer struct {
	CrossDomainMessageHash *common.Hash `gorm:"serializer:bytes"`

	Tx        Transaction `gorm:"embedded"`
	TokenPair TokenPair   `gorm:"embedded"`
}

type L1BridgeDeposit struct {
	BridgeTransfer        `gorm:"embedded"`
	TransactionSourceHash common.Hash `gorm:"primaryKey;serializer:bytes"`
}

type L1BridgeDepositWithTransactionHashes struct {
	L1BridgeDeposit L1BridgeDeposit `gorm:"embedded"`

	L1BlockHash       common.Hash `gorm:"serializer:bytes"`
	L1TransactionHash common.Hash `gorm:"serializer:bytes"`
	L2TransactionHash common.Hash `gorm:"serializer:bytes"`
}

type L2BridgeWithdrawal struct {
	BridgeTransfer            `gorm:"embedded"`
	TransactionWithdrawalHash common.Hash `gorm:"primaryKey;serializer:bytes"`
}

type L2BridgeWithdrawalWithTransactionHashes struct {
	L2BridgeWithdrawal L2BridgeWithdrawal `gorm:"embedded"`
	L2TransactionHash  common.Hash        `gorm:"serializer:bytes"`
	L2BlockHash        common.Hash        `gorm:"serializer:bytes"`

	ProvenL1TransactionHash    common.Hash `gorm:"serializer:bytes"`
	FinalizedL1TransactionHash common.Hash `gorm:"serializer:bytes"`
}

type BridgeTransfersView interface {
	L1BridgeDeposit(common.Hash) (*L1BridgeDeposit, error)
	L1BridgeDepositWithFilter(BridgeTransfer) (*L1BridgeDeposit, error)
	L1BridgeDepositsByAddress(common.Address, string, int) (*L1BridgeDepositsResponse, error)

	GetL1FirstDepositTimestamp() uint64
	GetL1EndDepositTimestamp(uint64, int) uint64
	GetL2FirstWithdrawTimestamp() uint64
	GetL2EndWithdrawTimestamp(uint64, int) uint64
	StoreL1BridgeDepositsByTimestamp(uint64, uint64) ([]L1BridgeDepositWithTransactionHashes, error)
	StoreL2BridgeWithdrawsByTimestamp(uint64, uint64) ([]L2BridgeWithdrawalWithTransactionHashes, error)

	L2BridgeWithdrawal(common.Hash) (*L2BridgeWithdrawal, error)
	L2BridgeWithdrawalWithFilter(BridgeTransfer) (*L2BridgeWithdrawal, error)
	L2BridgeWithdrawalsByAddress(common.Address, string, int) (*L2BridgeWithdrawalsResponse, error)
}

type BridgeTransfersDB interface {
	BridgeTransfersView

	StoreL1BridgeDeposits([]L1BridgeDeposit) error
	StoreL2BridgeWithdrawals([]L2BridgeWithdrawal) error
}

/**
 * Implementation
 */

type bridgeTransfersDB struct {
	gorm *gorm.DB
}

func NewBridgeTransfersDB(db *gorm.DB) BridgeTransfersDB {
	return &bridgeTransfersDB{gorm: db}
}

/**
 * Tokens Bridged (Deposited) from L1
 */

func (db *bridgeTransfersDB) StoreL1BridgeDeposits(deposits []L1BridgeDeposit) error {
	result := db.gorm.CreateInBatches(&deposits, utils.BatchInsertSize)
	return result.Error
}

func (db *bridgeTransfersDB) L1BridgeDeposit(txSourceHash common.Hash) (*L1BridgeDeposit, error) {
	var deposit L1BridgeDeposit
	result := db.gorm.Where(&L1BridgeDeposit{TransactionSourceHash: txSourceHash}).Take(&deposit)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &deposit, nil
}

// L1BridgeDepositWithFilter queries for a bridge deposit with set fields in the `BridgeTransfer` filter
func (db *bridgeTransfersDB) L1BridgeDepositWithFilter(filter BridgeTransfer) (*L1BridgeDeposit, error) {
	var deposit L1BridgeDeposit
	result := db.gorm.Where(&filter).Take(&deposit)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &deposit, nil
}

type L1BridgeDepositsResponse struct {
	Deposits    []L1BridgeDepositWithTransactionHashes
	Cursor      string
	HasNextPage bool
}

// L1BridgeDepositsByAddress retrieves a list of deposits initiated by the specified address,
// coupled with the L1/L2 transaction hashes that complete the bridge transaction.
func (db *bridgeTransfersDB) L1BridgeDepositsByAddress(address common.Address, cursor string, limit int) (*L1BridgeDepositsResponse, error) {
	if limit <= 0 {
		return nil, fmt.Errorf("limit must be greater than 0")
	}

	cursorClause := ""
	if cursor != "" {
		sourceHash := common.HexToHash(cursor)
		txDeposit := new(L1TransactionDeposit)
		result := db.gorm.Model(&L1TransactionDeposit{}).Where(&L1TransactionDeposit{SourceHash: sourceHash}).Take(txDeposit)
		if result.Error != nil || errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find transaction with supplied cursor source hash %s: %w", sourceHash, result.Error)
		}
		cursorClause = fmt.Sprintf("l1_transaction_deposits.timestamp <= %d", txDeposit.Tx.Timestamp)
	}

	// TODO join with l1_bridged_tokens and l2_bridged_tokens
	ethAddressString := predeploys.BVM_ETHAddr.String()

	// Coalesce l1 transaction deposits that are simply ETH sends
	ethTransactionDeposits := db.gorm.Model(&L1TransactionDeposit{})
	ethTransactionDeposits = ethTransactionDeposits.Where(&Transaction{FromAddress: address}).Where("data = '0x' AND eth_amount > 0")
	ethTransactionDeposits = ethTransactionDeposits.Joins("INNER JOIN l1_contract_events ON l1_contract_events.guid = initiated_l1_event_guid")
	ethTransactionDeposits = ethTransactionDeposits.Select(`
from_address, to_address, eth_amount, data, source_hash AS transaction_source_hash,
l2_transaction_hash, l1_contract_events.transaction_hash AS l1_transaction_hash, l1_contract_events.block_hash as l1_block_hash,
l1_transaction_deposits.timestamp, NULL AS cross_domain_message_hash, ? AS local_token_address, ? AS remote_token_address`, ethAddressString, ethAddressString)
	ethTransactionDeposits = ethTransactionDeposits.Order("timestamp DESC").Limit(limit + 1)
	if cursorClause != "" {
		ethTransactionDeposits = ethTransactionDeposits.Where(cursorClause)
	}

	depositsQuery := db.gorm.Model(&L1BridgeDeposit{})
	depositsQuery = depositsQuery.Where(&Transaction{FromAddress: address})
	depositsQuery = depositsQuery.Joins("INNER JOIN l1_transaction_deposits ON l1_transaction_deposits.source_hash = transaction_source_hash")
	depositsQuery = depositsQuery.Joins("INNER JOIN l1_contract_events ON l1_contract_events.guid = l1_transaction_deposits.initiated_l1_event_guid")
	depositsQuery = depositsQuery.Select(`
l1_bridge_deposits.from_address, l1_bridge_deposits.to_address, l1_bridge_deposits.eth_amount, l1_bridge_deposits.data, transaction_source_hash,
l2_transaction_hash, l1_contract_events.transaction_hash AS l1_transaction_hash, l1_contract_events.block_hash as l1_block_hash,
l1_bridge_deposits.timestamp, cross_domain_message_hash, local_token_address, remote_token_address`)
	depositsQuery = depositsQuery.Order("timestamp DESC").Limit(limit + 1)
	if cursorClause != "" {
		depositsQuery = depositsQuery.Where(cursorClause)
	}

	query := db.gorm.Table("(?) AS deposits", depositsQuery)
	query = query.Joins("UNION (?)", ethTransactionDeposits)
	query = query.Select("*").Order("timestamp DESC").Limit(limit + 1)
	deposits := []L1BridgeDepositWithTransactionHashes{}
	result := query.Find(&deposits)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	nextCursor := ""
	hasNextPage := false
	if len(deposits) > limit {
		hasNextPage = true
		nextCursor = deposits[limit].L1BridgeDeposit.TransactionSourceHash.String()
		deposits = deposits[:limit]
	}

	response := &L1BridgeDepositsResponse{Deposits: deposits, Cursor: nextCursor, HasNextPage: hasNextPage}
	return response, nil
}

func (db *bridgeTransfersDB) GetL1FirstDepositTimestamp() uint64 {

	var deposit L1BridgeDeposit
	db.gorm.Table("l1_bridge_deposits").Order("timestamp ASC").First(&deposit)
	if deposit.Tx.Timestamp == 0 {
		return 1
	}

	return deposit.Tx.Timestamp
}

func (db *bridgeTransfersDB) GetL1EndDepositTimestamp(startTimestamp uint64, insertLoop int) uint64 {

	var endTimestamp uint64
	db.gorm.Table("l1_bridge_deposits").Select("timestamp").Where("timestamp > ?", startTimestamp).Offset(insertLoop - 1).Limit(1).Order("timestamp ASC").Find(&endTimestamp)

	if endTimestamp == 0 {
		db.gorm.Table("l1_bridge_deposits").Select("timestamp").Order("timestamp DESC").Limit(1).Find(&endTimestamp)
	}

	return endTimestamp
}

func (db *bridgeTransfersDB) GetL2FirstWithdrawTimestamp() uint64 {

	var withdraw L2BridgeWithdrawal
	db.gorm.Table("l2_bridge_withdrawals").Order("timestamp ASC").First(&withdraw)
	if withdraw.Tx.Timestamp == 0 {
		return 1
	}

	return withdraw.Tx.Timestamp
}

func (db *bridgeTransfersDB) GetL2EndWithdrawTimestamp(startTimestamp uint64, insertLoop int) uint64 {

	var endTimestamp uint64
	db.gorm.Table("l2_bridge_withdrawals").Select("timestamp").Where("timestamp > ?", startTimestamp).Offset(insertLoop - 1).Limit(1).Order("timestamp ASC").Find(&endTimestamp)

	if endTimestamp == 0 {
		db.gorm.Table("l2_bridge_withdrawals").Select("timestamp").Order("timestamp DESC").Limit(1).Find(&endTimestamp)
	}

	return endTimestamp
}

func (db *bridgeTransfersDB) StoreL1BridgeDepositsByTimestamp(start uint64, end uint64) ([]L1BridgeDepositWithTransactionHashes, error) {

	ethAddressString := predeploys.BVM_ETHAddr.String()

	// Coalesce l1 transaction deposits that are simply ETH sends
	ethTransactionDeposits := db.gorm.Model(&L1TransactionDeposit{})
	ethTransactionDeposits = ethTransactionDeposits.Where("data = '0x' AND eth_amount > 0").Where("l1_transaction_deposits.timestamp >= ?", start).Where("l1_transaction_deposits.timestamp <= ?", end)
	ethTransactionDeposits = ethTransactionDeposits.Joins("INNER JOIN l1_contract_events ON l1_contract_events.guid = initiated_l1_event_guid")
	ethTransactionDeposits = ethTransactionDeposits.Select(`
from_address, to_address, eth_amount, data, source_hash AS transaction_source_hash,
l2_transaction_hash, l1_contract_events.transaction_hash AS l1_transaction_hash, l1_contract_events.block_hash as l1_block_hash,
l1_transaction_deposits.timestamp, NULL AS cross_domain_message_hash, ? AS local_token_address, ? AS remote_token_address`, ethAddressString, ethAddressString)
	ethTransactionDeposits = ethTransactionDeposits.Order("timestamp DESC")

	depositsQuery := db.gorm.Model(&L1BridgeDeposit{})
	depositsQuery = depositsQuery.Where("l1_bridge_deposits.timestamp >= ?", start).Where("l1_bridge_deposits.timestamp <= ?", end)
	depositsQuery = depositsQuery.Joins("INNER JOIN l1_transaction_deposits ON l1_transaction_deposits.source_hash = transaction_source_hash")
	depositsQuery = depositsQuery.Joins("INNER JOIN l1_contract_events ON l1_contract_events.guid = l1_transaction_deposits.initiated_l1_event_guid")
	depositsQuery = depositsQuery.Select(`
l1_bridge_deposits.from_address, l1_bridge_deposits.to_address, l1_bridge_deposits.eth_amount, l1_bridge_deposits.data, transaction_source_hash,
l2_transaction_hash, l1_contract_events.transaction_hash AS l1_transaction_hash, l1_contract_events.block_hash as l1_block_hash,
l1_bridge_deposits.timestamp, cross_domain_message_hash, local_token_address, remote_token_address`)
	depositsQuery = depositsQuery.Order("timestamp DESC")

	query := db.gorm.Table("(?) AS deposits", depositsQuery)
	query = query.Joins("UNION (?)", ethTransactionDeposits)
	query = query.Select("*").Order("timestamp DESC")
	deposits := []L1BridgeDepositWithTransactionHashes{}
	result := query.Find(&deposits)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return deposits, nil
}

func (db *bridgeTransfersDB) StoreL2BridgeWithdrawsByTimestamp(start uint64, end uint64) ([]L2BridgeWithdrawalWithTransactionHashes, error) {

	ethAddressString := predeploys.BVM_ETHAddr.String()

	// Coalesce l2 transaction withdrawals that are simply ETH sends
	ethTransactionWithdrawals := db.gorm.Model(&L2TransactionWithdrawal{})
	ethTransactionWithdrawals = ethTransactionWithdrawals.Where(`data = '0x' AND eth_amount > 0`).Where("l2_transaction_withdrawals.timestamp >= ?", start).Where("l2_transaction_withdrawals.timestamp <= ?", end)
	ethTransactionWithdrawals = ethTransactionWithdrawals.Joins("INNER JOIN l2_contract_events ON l2_contract_events.guid = l2_transaction_withdrawals.initiated_l2_event_guid")
	ethTransactionWithdrawals = ethTransactionWithdrawals.Joins("LEFT JOIN l1_contract_events AS proven_l1_events ON proven_l1_events.guid = l2_transaction_withdrawals.proven_l1_event_guid")
	ethTransactionWithdrawals = ethTransactionWithdrawals.Joins("LEFT JOIN l1_contract_events AS finalized_l1_events ON finalized_l1_events.guid = l2_transaction_withdrawals.finalized_l1_event_guid")
	ethTransactionWithdrawals = ethTransactionWithdrawals.Select(`
from_address, to_address, eth_amount, data, withdrawal_hash AS transaction_withdrawal_hash,
l2_contract_events.transaction_hash AS l2_transaction_hash, l2_contract_events.block_hash as l2_block_hash, proven_l1_events.transaction_hash AS proven_l1_transaction_hash, finalized_l1_events.transaction_hash AS finalized_l1_transaction_hash,
l2_transaction_withdrawals.timestamp, NULL AS cross_domain_message_hash, ? AS local_token_address, ? AS remote_token_address`, ethAddressString, ethAddressString)
	ethTransactionWithdrawals = ethTransactionWithdrawals.Order("timestamp DESC")

	withdrawalsQuery := db.gorm.Model(&L2BridgeWithdrawal{})
	withdrawalsQuery = withdrawalsQuery.Where("l2_bridge_withdrawals.timestamp >= ?", start).Where("l2_bridge_withdrawals.timestamp <= ?", end)
	withdrawalsQuery = withdrawalsQuery.Joins("INNER JOIN l2_transaction_withdrawals ON withdrawal_hash = l2_bridge_withdrawals.transaction_withdrawal_hash")
	withdrawalsQuery = withdrawalsQuery.Joins("INNER JOIN l2_contract_events ON l2_contract_events.guid = l2_transaction_withdrawals.initiated_l2_event_guid")
	withdrawalsQuery = withdrawalsQuery.Joins("LEFT JOIN l1_contract_events AS proven_l1_events ON proven_l1_events.guid = l2_transaction_withdrawals.proven_l1_event_guid")
	withdrawalsQuery = withdrawalsQuery.Joins("LEFT JOIN l1_contract_events AS finalized_l1_events ON finalized_l1_events.guid = l2_transaction_withdrawals.finalized_l1_event_guid")
	withdrawalsQuery = withdrawalsQuery.Select(`
l2_bridge_withdrawals.from_address, l2_bridge_withdrawals.to_address, l2_bridge_withdrawals.eth_amount, l2_bridge_withdrawals.data, transaction_withdrawal_hash,
l2_contract_events.transaction_hash AS l2_transaction_hash, l2_contract_events.block_hash as l2_block_hash, proven_l1_events.transaction_hash AS proven_l1_transaction_hash, finalized_l1_events.transaction_hash AS finalized_l1_transaction_hash,
l2_bridge_withdrawals.timestamp, cross_domain_message_hash, local_token_address, remote_token_address`)
	withdrawalsQuery = withdrawalsQuery.Order("timestamp DESC")

	query := db.gorm.Table("(?) AS withdrawals", withdrawalsQuery)
	query = query.Joins("UNION (?)", ethTransactionWithdrawals)
	query = query.Select("*").Order("timestamp DESC")
	withdrawals := []L2BridgeWithdrawalWithTransactionHashes{}

	// (3) Execute query and process results
	result := query.Find(&withdrawals)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return withdrawals, nil
}

/**
 * Tokens Bridged (Withdrawn) from L2
 */

func (db *bridgeTransfersDB) StoreL2BridgeWithdrawals(withdrawals []L2BridgeWithdrawal) error {
	result := db.gorm.CreateInBatches(&withdrawals, utils.BatchInsertSize)
	return result.Error
}

func (db *bridgeTransfersDB) L2BridgeWithdrawal(txWithdrawalHash common.Hash) (*L2BridgeWithdrawal, error) {
	var withdrawal L2BridgeWithdrawal
	result := db.gorm.Where(&L2BridgeWithdrawal{TransactionWithdrawalHash: txWithdrawalHash}).Take(&withdrawal)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &withdrawal, nil
}

// L2BridgeWithdrawalWithFilter queries for a bridge withdrawal with set fields in the `BridgeTransfer` filter
func (db *bridgeTransfersDB) L2BridgeWithdrawalWithFilter(filter BridgeTransfer) (*L2BridgeWithdrawal, error) {
	var withdrawal L2BridgeWithdrawal
	result := db.gorm.Where(filter).Take(&withdrawal)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &withdrawal, nil
}

type L2BridgeWithdrawalsResponse struct {
	Withdrawals []L2BridgeWithdrawalWithTransactionHashes
	Cursor      string
	HasNextPage bool
}

// L2BridgeDepositsByAddress retrieves a list of deposits initiated by the specified address, coupled with the L1/L2 transaction hashes
// that complete the bridge transaction. The hashes that correspond with the Bedrock multi-step withdrawal process are also surfaced
func (db *bridgeTransfersDB) L2BridgeWithdrawalsByAddress(address common.Address, cursor string, limit int) (*L2BridgeWithdrawalsResponse, error) {
	if limit <= 0 {
		return nil, fmt.Errorf("limit must be greater than 0")
	}
	fmt.Println("cursorcursorcursorcursor======", cursor)
	// (1) Generate cursor clause provided a cursor tx hash
	cursorClause := ""
	if cursor != "" {
		withdrawalHash := common.HexToHash(cursor)
		var txWithdrawal L2TransactionWithdrawal
		result := db.gorm.Model(&L2TransactionWithdrawal{}).Where(&L2TransactionWithdrawal{WithdrawalHash: withdrawalHash}).Take(&txWithdrawal)
		if result.Error != nil || errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("unable to find transaction with supplied cursor withdrawal hash %s: %w", withdrawalHash, result.Error)
		}
		cursorClause = fmt.Sprintf("l2_transaction_withdrawals.timestamp <= %d", txWithdrawal.Tx.Timestamp)
	}
	fmt.Println("cursorClause====", cursorClause)

	// (2) Generate query for fetching ETH withdrawal data
	// This query is a UNION (A | B) of two sub-queries:
	//   - (A) ETH sends from L2 to L1
	//   - (B) Bridge withdrawals from L2 to L1

	// TODO join with l1_bridged_tokens and l2_bridged_tokens
	ethAddressString := predeploys.BVM_ETHAddr.String()

	// Coalesce l2 transaction withdrawals that are simply ETH sends
	ethTransactionWithdrawals := db.gorm.Model(&L2TransactionWithdrawal{})
	ethTransactionWithdrawals = ethTransactionWithdrawals.Where(&Transaction{FromAddress: address}).Where(`data = '0x' AND eth_amount > 0`)
	ethTransactionWithdrawals = ethTransactionWithdrawals.Joins("INNER JOIN l2_contract_events ON l2_contract_events.guid = l2_transaction_withdrawals.initiated_l2_event_guid")
	ethTransactionWithdrawals = ethTransactionWithdrawals.Joins("LEFT JOIN l1_contract_events AS proven_l1_events ON proven_l1_events.guid = l2_transaction_withdrawals.proven_l1_event_guid")
	ethTransactionWithdrawals = ethTransactionWithdrawals.Joins("LEFT JOIN l1_contract_events AS finalized_l1_events ON finalized_l1_events.guid = l2_transaction_withdrawals.finalized_l1_event_guid")
	ethTransactionWithdrawals = ethTransactionWithdrawals.Select(`
from_address, to_address, eth_amount, data, withdrawal_hash AS transaction_withdrawal_hash,
l2_contract_events.transaction_hash AS l2_transaction_hash, l2_contract_events.block_hash as l2_block_hash, proven_l1_events.transaction_hash AS proven_l1_transaction_hash, finalized_l1_events.transaction_hash AS finalized_l1_transaction_hash,
l2_transaction_withdrawals.timestamp, NULL AS cross_domain_message_hash, ? AS local_token_address, ? AS remote_token_address`, ethAddressString, ethAddressString)
	ethTransactionWithdrawals = ethTransactionWithdrawals.Order("timestamp DESC").Limit(limit + 1)
	if cursorClause != "" {
		ethTransactionWithdrawals = ethTransactionWithdrawals.Where(cursorClause)
	}

	withdrawalsQuery := db.gorm.Model(&L2BridgeWithdrawal{})
	withdrawalsQuery = withdrawalsQuery.Where(&Transaction{FromAddress: address})
	withdrawalsQuery = withdrawalsQuery.Joins("INNER JOIN l2_transaction_withdrawals ON withdrawal_hash = l2_bridge_withdrawals.transaction_withdrawal_hash")
	withdrawalsQuery = withdrawalsQuery.Joins("INNER JOIN l2_contract_events ON l2_contract_events.guid = l2_transaction_withdrawals.initiated_l2_event_guid")
	withdrawalsQuery = withdrawalsQuery.Joins("LEFT JOIN l1_contract_events AS proven_l1_events ON proven_l1_events.guid = l2_transaction_withdrawals.proven_l1_event_guid")
	withdrawalsQuery = withdrawalsQuery.Joins("LEFT JOIN l1_contract_events AS finalized_l1_events ON finalized_l1_events.guid = l2_transaction_withdrawals.finalized_l1_event_guid")
	withdrawalsQuery = withdrawalsQuery.Select(`
l2_bridge_withdrawals.from_address, l2_bridge_withdrawals.to_address, l2_bridge_withdrawals.eth_amount, l2_bridge_withdrawals.data, transaction_withdrawal_hash,
l2_contract_events.transaction_hash AS l2_transaction_hash, l2_contract_events.block_hash as l2_block_hash, proven_l1_events.transaction_hash AS proven_l1_transaction_hash, finalized_l1_events.transaction_hash AS finalized_l1_transaction_hash,
l2_bridge_withdrawals.timestamp, cross_domain_message_hash, local_token_address, remote_token_address`)
	withdrawalsQuery = withdrawalsQuery.Order("timestamp DESC").Limit(limit + 1)
	if cursorClause != "" {
		withdrawalsQuery = withdrawalsQuery.Where(cursorClause)
	}

	query := db.gorm.Table("(?) AS withdrawals", withdrawalsQuery)
	query = query.Joins("UNION (?)", ethTransactionWithdrawals)
	query = query.Select("*").Order("timestamp DESC").Limit(limit + 1)
	withdrawals := []L2BridgeWithdrawalWithTransactionHashes{}

	// (3) Execute query and process results
	result := query.Find(&withdrawals)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	nextCursor := ""
	hasNextPage := false
	if len(withdrawals) > limit {
		hasNextPage = true
		nextCursor = withdrawals[limit].L2BridgeWithdrawal.TransactionWithdrawalHash.String()
		withdrawals = withdrawals[:limit]
	}

	response := &L2BridgeWithdrawalsResponse{Withdrawals: withdrawals, Cursor: nextCursor, HasNextPage: hasNextPage}
	return response, nil
}