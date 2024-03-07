// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"
	"github.com/cornerstone-labs/acorus/event/mantle/op-bindings/solc"
)

const L1MantleTokenStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/local/TestMantleToken.sol:L1MantleToken\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/local/TestMantleToken.sol:L1MantleToken\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/local/TestMantleToken.sol:L1MantleToken\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1027_storage\"},{\"astId\":1003,\"contract\":\"contracts/local/TestMantleToken.sol:L1MantleToken\",\"label\":\"_balances\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1004,\"contract\":\"contracts/local/TestMantleToken.sol:L1MantleToken\",\"label\":\"_allowances\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1005,\"contract\":\"contracts/local/TestMantleToken.sol:L1MantleToken\",\"label\":\"_totalSupply\",\"offset\":0,\"slot\":\"53\",\"type\":\"t_uint256\"},{\"astId\":1006,\"contract\":\"contracts/local/TestMantleToken.sol:L1MantleToken\",\"label\":\"_name\",\"offset\":0,\"slot\":\"54\",\"type\":\"t_string_storage\"},{\"astId\":1007,\"contract\":\"contracts/local/TestMantleToken.sol:L1MantleToken\",\"label\":\"_symbol\",\"offset\":0,\"slot\":\"55\",\"type\":\"t_string_storage\"},{\"astId\":1008,\"contract\":\"contracts/local/TestMantleToken.sol:L1MantleToken\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"56\",\"type\":\"t_array(t_uint256)1024_storage\"},{\"astId\":1009,\"contract\":\"contracts/local/TestMantleToken.sol:L1MantleToken\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_array(t_uint256)1027_storage\"},{\"astId\":1010,\"contract\":\"contracts/local/TestMantleToken.sol:L1MantleToken\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_address\"},{\"astId\":1011,\"contract\":\"contracts/local/TestMantleToken.sol:L1MantleToken\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)1026_storage\"},{\"astId\":1012,\"contract\":\"contracts/local/TestMantleToken.sol:L1MantleToken\",\"label\":\"_HASHED_NAME\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_bytes32\"},{\"astId\":1013,\"contract\":\"contracts/local/TestMantleToken.sol:L1MantleToken\",\"label\":\"_HASHED_VERSION\",\"offset\":0,\"slot\":\"202\",\"type\":\"t_bytes32\"},{\"astId\":1014,\"contract\":\"contracts/local/TestMantleToken.sol:L1MantleToken\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"203\",\"type\":\"t_array(t_uint256)1027_storage\"},{\"astId\":1015,\"contract\":\"contracts/local/TestMantleToken.sol:L1MantleToken\",\"label\":\"_nonces\",\"offset\":0,\"slot\":\"253\",\"type\":\"t_mapping(t_address,t_struct(Counter)1029_storage)\"},{\"astId\":1016,\"contract\":\"contracts/local/TestMantleToken.sol:L1MantleToken\",\"label\":\"_PERMIT_TYPEHASH_DEPRECATED_SLOT\",\"offset\":0,\"slot\":\"254\",\"type\":\"t_bytes32\"},{\"astId\":1017,\"contract\":\"contracts/local/TestMantleToken.sol:L1MantleToken\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"255\",\"type\":\"t_array(t_uint256)1026_storage\"},{\"astId\":1018,\"contract\":\"contracts/local/TestMantleToken.sol:L1MantleToken\",\"label\":\"_delegates\",\"offset\":0,\"slot\":\"304\",\"type\":\"t_mapping(t_address,t_address)\"},{\"astId\":1019,\"contract\":\"contracts/local/TestMantleToken.sol:L1MantleToken\",\"label\":\"_checkpoints\",\"offset\":0,\"slot\":\"305\",\"type\":\"t_mapping(t_address,t_array(t_struct(Checkpoint)1028_storage)dyn_storage)\"},{\"astId\":1020,\"contract\":\"contracts/local/TestMantleToken.sol:L1MantleToken\",\"label\":\"_totalSupplyCheckpoints\",\"offset\":0,\"slot\":\"306\",\"type\":\"t_array(t_struct(Checkpoint)1028_storage)dyn_storage\"},{\"astId\":1021,\"contract\":\"contracts/local/TestMantleToken.sol:L1MantleToken\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"307\",\"type\":\"t_array(t_uint256)1025_storage\"},{\"astId\":1022,\"contract\":\"contracts/local/TestMantleToken.sol:L1MantleToken\",\"label\":\"mintCapNumerator\",\"offset\":0,\"slot\":\"354\",\"type\":\"t_uint256\"},{\"astId\":1023,\"contract\":\"contracts/local/TestMantleToken.sol:L1MantleToken\",\"label\":\"nextMint\",\"offset\":0,\"slot\":\"355\",\"type\":\"t_uint256\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_struct(Checkpoint)1028_storage)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"struct ERC20VotesUpgradeable.Checkpoint[]\",\"numberOfBytes\":\"32\"},\"t_array(t_uint256)1024_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[45]\",\"numberOfBytes\":\"1440\"},\"t_array(t_uint256)1025_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[47]\",\"numberOfBytes\":\"1504\"},\"t_array(t_uint256)1026_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1027_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_address,t_address)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e address)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_address\"},\"t_mapping(t_address,t_array(t_struct(Checkpoint)1028_storage)dyn_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e struct ERC20VotesUpgradeable.Checkpoint[])\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_array(t_struct(Checkpoint)1028_storage)dyn_storage\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_struct(Counter)1029_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e struct CountersUpgradeable.Counter)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_struct(Counter)1029_storage\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_string_storage\":{\"encoding\":\"bytes\",\"label\":\"string\",\"numberOfBytes\":\"32\"},\"t_struct(Checkpoint)1028_storage\":{\"encoding\":\"inplace\",\"label\":\"struct ERC20VotesUpgradeable.Checkpoint\",\"numberOfBytes\":\"32\"},\"t_struct(Counter)1029_storage\":{\"encoding\":\"inplace\",\"label\":\"struct CountersUpgradeable.Counter\",\"numberOfBytes\":\"32\"},\"t_uint224\":{\"encoding\":\"inplace\",\"label\":\"uint224\",\"numberOfBytes\":\"28\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1MantleTokenStorageLayout = new(solc.StorageLayout)

var L1MantleTokenDeployedBin = "0x608060405234801561001057600080fd5b50600436106102415760003560e01c806379cc679011610145578063a9f8ad04116100bd578063da35a26f1161008c578063e40172b311610071578063e40172b31461053b578063f1127ed814610562578063f2fde38b146105b457600080fd5b8063da35a26f146104e2578063dd62ed3e146104f557600080fd5b8063a9f8ad04146104a7578063c3cda520146104b2578063cf665443146104c5578063d505accf146104cf57600080fd5b80638e539e8c116101145780639ab24eb0116100f95780639ab24eb01461046e578063a457c2d714610481578063a9059cbb1461049457600080fd5b80638e539e8c1461045357806395d89b411461046657600080fd5b806379cc6790146104065780637ecebe001461041957806389110e5d1461042c5780638da5cb5b1461043557600080fd5b80633a46b1a8116101d85780635c19a95c116101a75780636fcfff451161018c5780636fcfff45146103a057806370a08231146103c8578063715018a6146103fe57600080fd5b80635c19a95c146103835780636561e2111461039657600080fd5b80633a46b1a8146102eb57806340c10f19146102fe57806342966c6814610311578063587cde1e1461032457600080fd5b806323b872dd1161021457806323b872dd146102ae578063313ce567146102c15780633644e515146102d057806339509351146102d857600080fd5b806306fdde0314610246578063095ea7b31461026457806318160ddd146102875780631ae7f5f314610299575b600080fd5b61024e6105c7565b60405161025b9190612e00565b60405180910390f35b610277610272366004612e9c565b610659565b604051901515815260200161025b565b6035545b60405190815260200161025b565b6102ac6102a7366004612ec6565b610671565b005b6102776102bc366004612edf565b610749565b6040516012815260200161025b565b61028b61076d565b6102776102e6366004612e9c565b61077c565b61028b6102f9366004612e9c565b6107c8565b6102ac61030c366004612e9c565b61086a565b6102ac61031f366004612ec6565b61094a565b61035e610332366004612f1b565b73ffffffffffffffffffffffffffffffffffffffff908116600090815261013060205260409020541690565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161025b565b6102ac610391366004612f1b565b610957565b61028b6101625481565b6103b36103ae366004612f1b565b610961565b60405163ffffffff909116815260200161025b565b61028b6103d6366004612f1b565b73ffffffffffffffffffffffffffffffffffffffff1660009081526033602052604090205490565b6102ac610997565b6102ac610414366004612e9c565b6109ab565b61028b610427366004612f1b565b6109c4565b61028b61271081565b60975473ffffffffffffffffffffffffffffffffffffffff1661035e565b61028b610461366004612ec6565b6109ef565b61024e610a66565b61028b61047c366004612f1b565b610a75565b61027761048f366004612e9c565b610b42565b6102776104a2366004612e9c565b610c13565b61028b6301e1338081565b6102ac6104c0366004612f47565b610c21565b61028b6101635481565b6102ac6104dd366004612f9f565b610d98565b6102ac6104f0366004613009565b610f57565b61028b610503366004613035565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260346020908152604080832093909416825291909152205490565b61028b7f000000000000000000000000000000000000000000000000000000000000000081565b61057561057036600461305f565b611229565b60408051825163ffffffff1681526020928301517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16928101929092520161025b565b6102ac6105c2366004612f1b565b6112d0565b6060603680546105d69061309f565b80601f01602080910402602001604051908101604052809291908181526020018280546106029061309f565b801561064f5780601f106106245761010080835404028352916020019161064f565b820191906000526020600020905b81548152906001019060200180831161063257829003601f168201915b5050505050905090565b600033610667818585611719565b5060019392505050565b6106796118cc565b7f0000000000000000000000000000000000000000000000000000000000000000811115610701576040517fba96c68c000000000000000000000000000000000000000000000000000000008152600481018290527f000000000000000000000000000000000000000000000000000000000000000060248201526044015b60405180910390fd5b610162805490829055604080518281526020810184905233917fe2ee754bdb1a4ec4a5ecd3f810e4e7ca817cbbc379c89ff4e7a8b4dc6841a766910160405180910390a25050565b60003361075785828561194d565b610762858585611a1e565b506001949350505050565b6000610777611cd7565b905090565b33600081815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919061066790829086906107c390879061311b565b611719565b6000438210610833576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f4552433230566f7465733a20626c6f636b206e6f7420796574206d696e65640060448201526064016106f8565b73ffffffffffffffffffffffffffffffffffffffff83166000908152610131602052604090206108639083611d52565b9392505050565b6108726118cc565b60006127106101625461088460355490565b61088e9190613133565b6108989190613170565b9050808211156108de576040517f5d84733f00000000000000000000000000000000000000000000000000000000815260048101839052602481018290526044016106f8565b6101635442101561092957610163546040517fea14abd200000000000000000000000000000000000000000000000000000000815242600482015260248101919091526044016106f8565b6109376301e133804261311b565b610163556109458383611e39565b505050565b6109543382611e43565b50565b6109543382611e4d565b73ffffffffffffffffffffffffffffffffffffffff81166000908152610131602052604081205461099190611637565b92915050565b61099f6118cc565b6109a96000611eed565b565b6109b682338361194d565b6109c08282611e43565b5050565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260fd6020526040812054610991565b6000438210610a5a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f4552433230566f7465733a20626c6f636b206e6f7420796574206d696e65640060448201526064016106f8565b61099161013283611d52565b6060603780546105d69061309f565b73ffffffffffffffffffffffffffffffffffffffff8116600090815261013160205260408120548015610b1a5773ffffffffffffffffffffffffffffffffffffffff8316600090815261013160205260409020610ad36001836131ab565b81548110610ae357610ae36131c2565b60009182526020909120015464010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16610b1d565b60005b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff169392505050565b33600081815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919083811015610c06576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f00000000000000000000000000000000000000000000000000000060648201526084016106f8565b6107628286868403611719565b600033610667818585611a1e565b83421115610c8b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f4552433230566f7465733a207369676e6174757265206578706972656400000060448201526064016106f8565b604080517fe48329057bfd03d55e49b547132e39cffd9c1820ad7b9d4c5307691425d15adf602082015273ffffffffffffffffffffffffffffffffffffffff8816918101919091526060810186905260808101859052600090610d1290610d0a9060a00160405160208183030381529060405280519060200120611f64565b858585611fcd565b9050610d1d81611ff5565b8614610d85576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4552433230566f7465733a20696e76616c6964206e6f6e63650000000000000060448201526064016106f8565b610d8f8188611e4d565b50505050505050565b83421115610e02576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332305065726d69743a206578706972656420646561646c696e6500000060448201526064016106f8565b60007f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9888888610e318c611ff5565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090506000610e9982611f64565b90506000610ea982878787611fcd565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610f40576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f45524332305065726d69743a20696e76616c6964207369676e6174757265000060448201526064016106f8565b610f4b8a8a8a611719565b50505050505050505050565b600054610100900460ff1615808015610f775750600054600160ff909116105b80610f915750303b158015610f91575060005460ff166001145b61101d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016106f8565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561107b57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b82158061109c575073ffffffffffffffffffffffffffffffffffffffff8216155b156110d3576040517ff57dc27600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111476040518060400160405280600681526020017f4d616e746c6500000000000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f4d4e54000000000000000000000000000000000000000000000000000000000081525061202a565b61114f6120cb565b611157612162565b6111956040518060400160405280600681526020017f4d616e746c650000000000000000000000000000000000000000000000000000815250612201565b61119d6120cb565b6111a78284611e39565b6111b56301e133804261311b565b610163556111c282611eed565b801561094557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b604080518082019091526000808252602082015273ffffffffffffffffffffffffffffffffffffffff8316600090815261013160205260409020805463ffffffff841690811061127b5761127b6131c2565b60009182526020918290206040805180820190915291015463ffffffff8116825264010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16918101919091529392505050565b6112d86118cc565b73ffffffffffffffffffffffffffffffffffffffff811661137b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016106f8565b61095481611eed565b61138e8282611451565b6035547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff101561143c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f4552433230566f7465733a20746f74616c20737570706c79207269736b73206f60448201527f766572666c6f77696e6720766f7465730000000000000000000000000000000060648201526084016106f8565b61144b610132611579836122d7565b50505050565b73ffffffffffffffffffffffffffffffffffffffff82166114ce576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016106f8565b80603560008282546114e0919061311b565b909155505073ffffffffffffffffffffffffffffffffffffffff82166000908152603360205260408120805483929061151a90849061311b565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a36109c0600083836124b9565b6000610863828461311b565b60007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff821115611633576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203260448201527f323420626974730000000000000000000000000000000000000000000000000060648201526084016106f8565b5090565b600063ffffffff821115611633576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203360448201527f322062697473000000000000000000000000000000000000000000000000000060648201526084016106f8565b73ffffffffffffffffffffffffffffffffffffffff83811660009081526101306020526040808220548584168352912054610945929182169116836124c4565b600061086382846131ab565b73ffffffffffffffffffffffffffffffffffffffff83166117bb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f726573730000000000000000000000000000000000000000000000000000000060648201526084016106f8565b73ffffffffffffffffffffffffffffffffffffffff821661185e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f737300000000000000000000000000000000000000000000000000000000000060648201526084016106f8565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526034602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b60975473ffffffffffffffffffffffffffffffffffffffff1633146109a9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016106f8565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152603460209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461144b5781811015611a11576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016106f8565b61144b8484848403611719565b73ffffffffffffffffffffffffffffffffffffffff8316611ac1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f647265737300000000000000000000000000000000000000000000000000000060648201526084016106f8565b73ffffffffffffffffffffffffffffffffffffffff8216611b64576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f657373000000000000000000000000000000000000000000000000000000000060648201526084016106f8565b73ffffffffffffffffffffffffffffffffffffffff831660009081526033602052604090205481811015611c1a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e6365000000000000000000000000000000000000000000000000000060648201526084016106f8565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260336020526040808220858503905591851681529081208054849290611c5e90849061311b565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051611cc491815260200190565b60405180910390a361144b8484846124b9565b60006107777f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f611d0660c95490565b60ca546040805160208101859052908101839052606081018290524660808201523060a082015260009060c0016040516020818303038152906040528051906020012090509392505050565b8154600090815b81811015611db6576000611d6d828461266b565b905084868281548110611d8257611d826131c2565b60009182526020909120015463ffffffff161115611da257809250611db0565b611dad81600161311b565b91505b50611d59565b8115611e0f5784611dc86001846131ab565b81548110611dd857611dd86131c2565b60009182526020909120015464010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16611e12565b60005b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1695945050505050565b6109c08282611384565b6109c08282612686565b73ffffffffffffffffffffffffffffffffffffffff82811660008181526101306020818152604080842080546033845282862054949093528787167fffffffffffffffffffffffff00000000000000000000000000000000000000008416811790915590519190951694919391928592917f3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f9190a461144b8284836124c4565b6097805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000610991611f71611cd7565b836040517f19010000000000000000000000000000000000000000000000000000000000006020820152602281018390526042810182905260009060620160405160208183030381529060405280519060200120905092915050565b6000806000611fde8787878761269f565b91509150611feb816127b7565b5095945050505050565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260fd602052604090208054600181018255905b50919050565b600054610100900460ff166120c1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106f8565b6109c08282612a0b565b600054610100900460ff166109a9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106f8565b600054610100900460ff166121f9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106f8565b6109a9612abb565b600054610100900460ff16612298576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106f8565b610954816040518060400160405280600181526020017f3100000000000000000000000000000000000000000000000000000000000000815250612b5b565b82546000908190801561233757856122f06001836131ab565b81548110612300576123006131c2565b60009182526020909120015464010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1661233a565b60005b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16925061236883858763ffffffff16565b91506000811180156123a6575043866123826001846131ab565b81548110612392576123926131c2565b60009182526020909120015463ffffffff16145b15612430576123b482611585565b866123c06001846131ab565b815481106123d0576123d06131c2565b9060005260206000200160000160046101000a8154817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff02191690837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1602179055506124b0565b85604051806040016040528061244543611637565b63ffffffff16815260200161245985611585565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90811690915282546001810184556000938452602093849020835194909301519091166401000000000263ffffffff909316929092179101555b50935093915050565b6109458383836116cd565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141580156125005750600081115b156109455773ffffffffffffffffffffffffffffffffffffffff8316156125b65773ffffffffffffffffffffffffffffffffffffffff831660009081526101316020526040812081906125569061170d856122d7565b915091508473ffffffffffffffffffffffffffffffffffffffff167fdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a72483836040516125ab929190918252602082015260400190565b60405180910390a250505b73ffffffffffffffffffffffffffffffffffffffff8216156109455773ffffffffffffffffffffffffffffffffffffffff8216600090815261013160205260408120819061260790611579856122d7565b915091508373ffffffffffffffffffffffffffffffffffffffff167fdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724838360405161265c929190918252602082015260400190565b60405180910390a25050505050565b600061267a6002848418613170565b6108639084841661311b565b6126908282612c0c565b61144b61013261170d836122d7565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156126d657506000905060036127ae565b8460ff16601b141580156126ee57508460ff16601c14155b156126ff57506000905060046127ae565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015612753573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166127a7576000600192509250506127ae565b9150600090505b94509492505050565b60008160048111156127cb576127cb6131f1565b036127d35750565b60018160048111156127e7576127e76131f1565b0361284e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016106f8565b6002816004811115612862576128626131f1565b036128c9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016106f8565b60038160048111156128dd576128dd6131f1565b0361296a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016106f8565b600481600481111561297e5761297e6131f1565b03610954576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016106f8565b600054610100900460ff16612aa2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106f8565b6036612aae838261329d565b506037610945828261329d565b600054610100900460ff16612b52576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106f8565b6109a933611eed565b600054610100900460ff16612bf2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106f8565b81516020928301208151919092012060c99190915560ca55565b73ffffffffffffffffffffffffffffffffffffffff8216612caf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f730000000000000000000000000000000000000000000000000000000000000060648201526084016106f8565b73ffffffffffffffffffffffffffffffffffffffff821660009081526033602052604090205481811015612d65576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f636500000000000000000000000000000000000000000000000000000000000060648201526084016106f8565b73ffffffffffffffffffffffffffffffffffffffff83166000908152603360205260408120838303905560358054849290612da19084906131ab565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a3610945836000846124b9565b600060208083528351808285015260005b81811015612e2d57858101830151858201604001528201612e11565b81811115612e3f576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114612e9757600080fd5b919050565b60008060408385031215612eaf57600080fd5b612eb883612e73565b946020939093013593505050565b600060208284031215612ed857600080fd5b5035919050565b600080600060608486031215612ef457600080fd5b612efd84612e73565b9250612f0b60208501612e73565b9150604084013590509250925092565b600060208284031215612f2d57600080fd5b61086382612e73565b803560ff81168114612e9757600080fd5b60008060008060008060c08789031215612f6057600080fd5b612f6987612e73565b95506020870135945060408701359350612f8560608801612f36565b92506080870135915060a087013590509295509295509295565b600080600080600080600060e0888a031215612fba57600080fd5b612fc388612e73565b9650612fd160208901612e73565b95506040880135945060608801359350612fed60808901612f36565b925060a0880135915060c0880135905092959891949750929550565b6000806040838503121561301c57600080fd5b8235915061302c60208401612e73565b90509250929050565b6000806040838503121561304857600080fd5b61305183612e73565b915061302c60208401612e73565b6000806040838503121561307257600080fd5b61307b83612e73565b9150602083013563ffffffff8116811461309457600080fd5b809150509250929050565b600181811c908216806130b357607f821691505b602082108103612024577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561312e5761312e6130ec565b500190565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561316b5761316b6130ec565b500290565b6000826131a6577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b6000828210156131bd576131bd6130ec565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b601f82111561094557600081815260208120601f850160051c810160208610156132765750805b601f850160051c820191505b8181101561329557828155600101613282565b505050505050565b815167ffffffffffffffff8111156132b7576132b7613220565b6132cb816132c5845461309f565b8461324f565b602080601f83116001811461331e57600084156132e85750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555613295565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b8281101561336b5788860151825594840194600190910190840161334c565b50858210156133a757878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b0190555056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L1MantleTokenStorageLayoutJSON), L1MantleTokenStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1MantleToken"] = L1MantleTokenStorageLayout
	deployedBytecodes["L1MantleToken"] = L1MantleTokenDeployedBin
}
