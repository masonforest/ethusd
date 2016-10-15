// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// EthUSDABI is the input ABI used to generate the binding from.
const EthUSDABI = `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[{"name":"success","type":"bool"}],"type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"success","type":"bool"}],"type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"balances","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[{"name":"amountInCents","type":"uint256"}],"name":"withdraw","outputs":[{"name":"success","type":"bool"}],"type":"function"},{"constant":false,"inputs":[{"name":"myid","type":"bytes32"},{"name":"result","type":"string"},{"name":"proof","type":"bytes"}],"name":"__callback","outputs":[],"type":"function"},{"constant":true,"inputs":[],"name":"exchangeRate","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[],"name":"purchase","outputs":[{"name":"success","type":"bool"}],"type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[{"name":"success","type":"bool"}],"type":"function"},{"constant":false,"inputs":[{"name":"delay","type":"uint256"}],"name":"updateExchangeRate","outputs":[],"type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"remaining","type":"uint256"}],"type":"function"},{"inputs":[{"name":"_name","type":"string"},{"name":"_symbol","type":"string"}],"type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Purchase","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Withdraw","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_from","type":"address"},{"indexed":true,"name":"_to","type":"address"},{"indexed":false,"name":"_value","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_owner","type":"address"},{"indexed":true,"name":"_spender","type":"address"},{"indexed":false,"name":"_value","type":"uint256"}],"name":"Approval","type":"event"}]`

// EthUSDBin is the compiled bytecode used for deploying new contracts.
const EthUSDBin = `0x60606040526040516115dc3803806115dc833981016040528051608051908201910161007f7f110000000000000000000000000000000000000000000000000000000000000060008054600160a060020a031614156102a05761029e60005b600060006103b4731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed5b3b90565b61016860006103a781604060405190810160405280600381526020017f55524c0000000000000000000000000000000000000000000000000000000000815260200150608060405190810160405280604481526020017f6a736f6e2868747470733a2f2f706f6c6f6e6965782e636f6d2f7075626c696381526020017f3f636f6d6d616e643d72657475726e5469636b6572292e555344545f4554482e81526020017f6c61737400000000000000000000000000000000000000000000000000000000815260200150600080548190600160a060020a03168114156104e9576104e7600061005e565b8160066000509080519060200190828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106101cf57805160ff19168380011785555b506101ff9291505b8082111561025857600081556001016101bb565b828001600101855582156101b3579182015b828111156101b35782518260005055916020019190600101906101e1565b50508060076000509080519060200190828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061025c57805160ff19168380011785555b5061028c9291506101bb565b5090565b8280016001018555821561024c579182015b8281111561024c57825182600050559160200191906001019061026e565b50505050610e4b806107916000396000f35b505b600060009054906101000a9004600160a060020a0316600160a060020a03166338cc4831604051817c01000000000000000000000000000000000000000000000000000000000281526004018090506020604051808303816000876161da5a03f11561000257505060408051805160018054600160a060020a031916909117908190557f688dcfd70000000000000000000000000000000000000000000000000000000082527fff00000000000000000000000000000000000000000000000000000000000000851660048301529151600160a060020a0392909216925063688dcfd7916024808301926000929190829003018183876161da5a03f1156100025750505050565b5050565b5060005b919050565b11156103e8575060008054600160a060020a031916731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed17905560016103af565b6000610407739efbea6358bed926b293d2ce63a730d6d98d43dd61007b565b111561043e575060008054739efbea6358bed926b293d2ce63a730d6d98d43dd600160a060020a03199190911617905560016103af565b600061045d7320e12a1f859b3feae5fb2a0a32c18f5a65555bbf61007b565b11156104945750600080547320e12a1f859b3feae5fb2a0a32c18f5a65555bbf600160a060020a03199190911617905560016103af565b60006104b3739a1d6e5c6c8d081ac45c6af98b74a42442afba6061007b565b11156103ab575060008054600160a060020a031916739a1d6e5c6c8d081ac45c6af98b74a42442afba6017905560016103af565b505b600060009054906101000a9004600160a060020a0316600160a060020a03166338cc4831604051817c01000000000000000000000000000000000000000000000000000000000281526004018090506020604051808303816000876161da5a03f115610002575050604051805160018054600160a060020a031916909117908190557f524f388900000000000000000000000000000000000000000000000000000000825260206004838101828152895160248601528951600160a060020a0394909416955063524f3889948a949193849360449290920192868201929091829185918391869160009190601f850104600302600f01f150905090810190601f16801561060a5780820380516001836020036101000a031916815260200191505b50925050506020604051808303816000876161da5a03f11561000257505060405151915050670de0b6b3a764000062030d403a020181111561065357600091505b509392505050565b600160009054906101000a9004600160a060020a0316600160a060020a031663adf59f9982878787604051857c01000000000000000000000000000000000000000000000000000000000281526004018084815260200180602001806020018381038352858181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f16801561070a5780820380516001836020036101000a031916815260200191505b508381038252848181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f1680156107635780820380516001836020036101000a031916815260200191505b509550505050505060206040518083038185886185025a03f11561000257505060405151935061064b91505056606060405236156100ae5760e060020a600035046306fdde0381146100b0578063095ea7b31461010e57806318160ddd1461018357806323b872dd1461018c57806327e235e3146102795780632e1a7d4d1461029157806338bbfa50146102f65780633ba0b9a9146103bd57806364edfbf0146103c657806370a082311461049257806395d89b41146104b7578063a9059cbb14610515578063b9e205ae146105bc578063dd62ed3e146106a7575b005b6040805160068054602060026001831615610100026000190190921691909104601f81018290048202840182019094528383526106db939083018282801561079a5780601f1061076f5761010080835404028352916020019161079a565b61074960043560243533600160a060020a03908116600081815260056020908152604080832094871680845294825280832086905580518681529051929493927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060015b92915050565b61075d60035481565b610749600435602435604435600160a060020a0383166000908152600460205260408120548290108015906101df575060056020908152604080832033600160a060020a03168452909152812054829010155b80156101eb5750600082115b156107a257600160a060020a03838116600081815260046020908152604080832080548801905588851680845281842080548990039055600583528184203390961684529482529182902080548790039055815186815291519293927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35060016107a6565b61075d60043560086020526000908152604090205481565b61074960043533600160a060020a031660009081526008602052604081205481908390106107ad5760408120805484900390556107b8600254600354600091670de0b6b3a764000091020430600160a060020a031631101561092d57506104b0610932565b60408051602060248035600481810135601f81018590048502860185019096528585526100ae9581359591946044949293909201918190840183828082843750506040805160209735808a0135601f81018a90048a0283018a01909352828252969897606497919650602491909101945090925082915084018382808284375094965050505050505061082b60008054600160a060020a03168114156109375761093560005b60006000610d10731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed5b3b90565b61075d60025481565b6107496000600060003411156109185760408051348152905133600160a060020a0316917f2499a5330ab0979cc612135e7883ebc3cd5c9f7a8508f042540c34723348f632919081900360200190a2506040805160025433600160a060020a0316808552600860209081528486208054670de0b6b3a7640000349590950294909404938401905560038054840190558284529351919390927f2499a5330ab0979cc612135e7883ebc3cd5c9f7a8508f042540c34723348f632929081900390910190a26001915061091d565b61075d600435600160a060020a0381166000908152600860205260409020545b919050565b6040805160078054602060026001831615610100026000190190921691909104601f81018290048202840182019094528383526106db939083018282801561079a5780601f1061076f5761010080835404028352916020019161079a565b61074960043560243533600160a060020a03166000908152600460205260408120548290108015906105475750600082115b156109215733600160a060020a03908116600081815260046020908152604080832080548890039055938716808352918490208054870190558351868152935191937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a350600161017d565b6100ae6004355b61092981604060405190810160405280600381526020017f55524c0000000000000000000000000000000000000000000000000000000000815260200150608060405190810160405280604481526020017f6a736f6e2868747470733a2f2f706f6c6f6e6965782e636f6d2f7075626c696381526020017f3f636f6d6d616e643d72657475726e5469636b6572292e555344545f4554482e81526020017f6c61737400000000000000000000000000000000000000000000000000000000815260200150600080548190600160a060020a0316811415610a9a57610a98600061039c565b61075d600435602435600160a060020a0382811660009081526005602090815260408083209385168352929052205461017d565b60405180806020018281038252838181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f16801561073b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b604080519115158252519081900360200190f35b60408051918252519081900360200190f35b820191906000526020600020905b81548152906001019060200180831161077d57829003601f168201915b505050505081565b5060005b9392505050565b600091505b50919050565b604051908404670de0b6b3a764000002915033600160a060020a031690600090839082818181858883f1505060038054919091039055506040805183815290517f2499a5330ab0979cc612135e7883ebc3cd5c9f7a8508f042540c34723348f6329181900360200190a2600191506107b2565b600160a060020a031633600160a060020a031614151561084a57610002565b6109068260026040805160208101909152600090819052828180805b83518110156108e857603060f860020a028482815181101561000257016020015160f860020a9081900402600160f860020a031916108015906108d35750603960f860020a028482815181101561000257016020015160f860020a9081900402600160f860020a03191611155b15610a07578115610a68578560001415610a5f575b60008611156108fb57600a86900a909202915b509095945050505050565b60025561091360006105c3565b505050565b600091505b5090565b50600061017d565b5050565b506002545b90565b505b600060009054906101000a9004600160a060020a0316600160a060020a03166338cc48316040518160e060020a0281526004018090506020604051808303816000876161da5a03f11561000257505060408051805160018054600160a060020a031916909117908190557fc281d19e0000000000000000000000000000000000000000000000000000000082529151600160a060020a0392909216925063c281d19e91600482810192602092919082900301816000876161da5a03f1156100025750506040515191506109329050565b8381815181101561000257016020015160f860020a9081900402600160f860020a0319167f2e000000000000000000000000000000000000000000000000000000000000001415610a5757600191505b600101610866565b60001995909501945b600a83029250825060308482815181101561000257016020015160f860020a908190048102040390920191610a57565b505b600060009054906101000a9004600160a060020a0316600160a060020a03166338cc48316040518160e060020a0281526004018090506020604051808303816000876161da5a03f115610002575050604051805160018054600160a060020a031916909117908190557f524f388900000000000000000000000000000000000000000000000000000000825260206004838101828152895160248601528951600160a060020a0394909416955063524f3889948a949193849360449290920192868201929091829185918391869160009190601f850104600302600f01f150905090810190601f168015610ba25780820380516001836020036101000a031916815260200191505b50925050506020604051808303816000876161da5a03f11561000257505060405151915050670de0b6b3a764000062030d403a0201811115610beb57600091505b509392505050565b600160009054906101000a9004600160a060020a0316600160a060020a031663adf59f99828787876040518560e060020a0281526004018084815260200180602001806020018381038352858181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f168015610c895780820380516001836020036101000a031916815260200191505b508381038252848181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f168015610ce25780820380516001836020036101000a031916815260200191505b509550505050505060206040518083038185886185025a03f115610002575050604051519350610be3915050565b1115610d44575060008054600160a060020a031916731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed17905560016104b2565b6000610d63739efbea6358bed926b293d2ce63a730d6d98d43dd6103b9565b1115610d9a575060008054739efbea6358bed926b293d2ce63a730d6d98d43dd600160a060020a03199190911617905560016104b2565b6000610db97320e12a1f859b3feae5fb2a0a32c18f5a65555bbf6103b9565b1115610df05750600080547320e12a1f859b3feae5fb2a0a32c18f5a65555bbf600160a060020a03199190911617905560016104b2565b6000610e0f739a1d6e5c6c8d081ac45c6af98b74a42442afba606103b9565b1115610e43575060008054600160a060020a031916739a1d6e5c6c8d081ac45c6af98b74a42442afba6017905560016104b2565b5060006104b256`

// DeployEthUSD deploys a new Ethereum contract, binding an instance of EthUSD to it.
func DeployEthUSD(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string) (common.Address, *types.Transaction, *EthUSD, error) {
	parsed, err := abi.JSON(strings.NewReader(EthUSDABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthUSDBin), backend, _name, _symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthUSD{EthUSDCaller: EthUSDCaller{contract: contract}, EthUSDTransactor: EthUSDTransactor{contract: contract}}, nil
}

// EthUSD is an auto generated Go binding around an Ethereum contract.
type EthUSD struct {
	EthUSDCaller     // Read-only binding to the contract
	EthUSDTransactor // Write-only binding to the contract
}

// EthUSDCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthUSDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthUSDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthUSDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthUSDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthUSDSession struct {
	Contract     *EthUSD           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthUSDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthUSDCallerSession struct {
	Contract *EthUSDCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EthUSDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthUSDTransactorSession struct {
	Contract     *EthUSDTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthUSDRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthUSDRaw struct {
	Contract *EthUSD // Generic contract binding to access the raw methods on
}

// EthUSDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthUSDCallerRaw struct {
	Contract *EthUSDCaller // Generic read-only contract binding to access the raw methods on
}

// EthUSDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthUSDTransactorRaw struct {
	Contract *EthUSDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthUSD creates a new instance of EthUSD, bound to a specific deployed contract.
func NewEthUSD(address common.Address, backend bind.ContractBackend) (*EthUSD, error) {
	contract, err := bindEthUSD(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthUSD{EthUSDCaller: EthUSDCaller{contract: contract}, EthUSDTransactor: EthUSDTransactor{contract: contract}}, nil
}

// NewEthUSDCaller creates a new read-only instance of EthUSD, bound to a specific deployed contract.
func NewEthUSDCaller(address common.Address, caller bind.ContractCaller) (*EthUSDCaller, error) {
	contract, err := bindEthUSD(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &EthUSDCaller{contract: contract}, nil
}

// NewEthUSDTransactor creates a new write-only instance of EthUSD, bound to a specific deployed contract.
func NewEthUSDTransactor(address common.Address, transactor bind.ContractTransactor) (*EthUSDTransactor, error) {
	contract, err := bindEthUSD(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &EthUSDTransactor{contract: contract}, nil
}

// bindEthUSD binds a generic wrapper to an already deployed contract.
func bindEthUSD(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthUSDABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthUSD *EthUSDRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthUSD.Contract.EthUSDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthUSD *EthUSDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthUSD.Contract.EthUSDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthUSD *EthUSDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthUSD.Contract.EthUSDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthUSD *EthUSDCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthUSD.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthUSD *EthUSDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthUSD.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthUSD *EthUSDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthUSD.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_EthUSD *EthUSDCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EthUSD.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_EthUSD *EthUSDSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _EthUSD.Contract.Allowance(&_EthUSD.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_EthUSD *EthUSDCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _EthUSD.Contract.Allowance(&_EthUSD.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_EthUSD *EthUSDCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EthUSD.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_EthUSD *EthUSDSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _EthUSD.Contract.BalanceOf(&_EthUSD.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_EthUSD *EthUSDCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _EthUSD.Contract.BalanceOf(&_EthUSD.CallOpts, _owner)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_EthUSD *EthUSDCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EthUSD.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_EthUSD *EthUSDSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _EthUSD.Contract.Balances(&_EthUSD.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_EthUSD *EthUSDCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _EthUSD.Contract.Balances(&_EthUSD.CallOpts, arg0)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() constant returns(uint256)
func (_EthUSD *EthUSDCaller) ExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EthUSD.contract.Call(opts, out, "exchangeRate")
	return *ret0, err
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() constant returns(uint256)
func (_EthUSD *EthUSDSession) ExchangeRate() (*big.Int, error) {
	return _EthUSD.Contract.ExchangeRate(&_EthUSD.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() constant returns(uint256)
func (_EthUSD *EthUSDCallerSession) ExchangeRate() (*big.Int, error) {
	return _EthUSD.Contract.ExchangeRate(&_EthUSD.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_EthUSD *EthUSDCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _EthUSD.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_EthUSD *EthUSDSession) Name() (string, error) {
	return _EthUSD.Contract.Name(&_EthUSD.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_EthUSD *EthUSDCallerSession) Name() (string, error) {
	return _EthUSD.Contract.Name(&_EthUSD.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_EthUSD *EthUSDCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _EthUSD.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_EthUSD *EthUSDSession) Symbol() (string, error) {
	return _EthUSD.Contract.Symbol(&_EthUSD.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_EthUSD *EthUSDCallerSession) Symbol() (string, error) {
	return _EthUSD.Contract.Symbol(&_EthUSD.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_EthUSD *EthUSDCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EthUSD.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_EthUSD *EthUSDSession) TotalSupply() (*big.Int, error) {
	return _EthUSD.Contract.TotalSupply(&_EthUSD.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_EthUSD *EthUSDCallerSession) TotalSupply() (*big.Int, error) {
	return _EthUSD.Contract.TotalSupply(&_EthUSD.CallOpts)
}

// __callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(myid bytes32, result string, proof bytes) returns()
func (_EthUSD *EthUSDTransactor) __callback(opts *bind.TransactOpts, myid [32]byte, result string, proof []byte) (*types.Transaction, error) {
	return _EthUSD.contract.Transact(opts, "__callback", myid, result, proof)
}

// __callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(myid bytes32, result string, proof bytes) returns()
func (_EthUSD *EthUSDSession) __callback(myid [32]byte, result string, proof []byte) (*types.Transaction, error) {
	return _EthUSD.Contract.__callback(&_EthUSD.TransactOpts, myid, result, proof)
}

// __callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(myid bytes32, result string, proof bytes) returns()
func (_EthUSD *EthUSDTransactorSession) __callback(myid [32]byte, result string, proof []byte) (*types.Transaction, error) {
	return _EthUSD.Contract.__callback(&_EthUSD.TransactOpts, myid, result, proof)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_EthUSD *EthUSDTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EthUSD.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_EthUSD *EthUSDSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EthUSD.Contract.Approve(&_EthUSD.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_EthUSD *EthUSDTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EthUSD.Contract.Approve(&_EthUSD.TransactOpts, _spender, _value)
}

// Purchase is a paid mutator transaction binding the contract method 0x64edfbf0.
//
// Solidity: function purchase() returns(success bool)
func (_EthUSD *EthUSDTransactor) Purchase(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthUSD.contract.Transact(opts, "purchase")
}

// Purchase is a paid mutator transaction binding the contract method 0x64edfbf0.
//
// Solidity: function purchase() returns(success bool)
func (_EthUSD *EthUSDSession) Purchase() (*types.Transaction, error) {
	return _EthUSD.Contract.Purchase(&_EthUSD.TransactOpts)
}

// Purchase is a paid mutator transaction binding the contract method 0x64edfbf0.
//
// Solidity: function purchase() returns(success bool)
func (_EthUSD *EthUSDTransactorSession) Purchase() (*types.Transaction, error) {
	return _EthUSD.Contract.Purchase(&_EthUSD.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_EthUSD *EthUSDTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EthUSD.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_EthUSD *EthUSDSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EthUSD.Contract.Transfer(&_EthUSD.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_EthUSD *EthUSDTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EthUSD.Contract.Transfer(&_EthUSD.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_EthUSD *EthUSDTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EthUSD.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_EthUSD *EthUSDSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EthUSD.Contract.TransferFrom(&_EthUSD.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_EthUSD *EthUSDTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EthUSD.Contract.TransferFrom(&_EthUSD.TransactOpts, _from, _to, _value)
}

// UpdateExchangeRate is a paid mutator transaction binding the contract method 0xb9e205ae.
//
// Solidity: function updateExchangeRate(delay uint256) returns()
func (_EthUSD *EthUSDTransactor) UpdateExchangeRate(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _EthUSD.contract.Transact(opts, "updateExchangeRate", delay)
}

// UpdateExchangeRate is a paid mutator transaction binding the contract method 0xb9e205ae.
//
// Solidity: function updateExchangeRate(delay uint256) returns()
func (_EthUSD *EthUSDSession) UpdateExchangeRate(delay *big.Int) (*types.Transaction, error) {
	return _EthUSD.Contract.UpdateExchangeRate(&_EthUSD.TransactOpts, delay)
}

// UpdateExchangeRate is a paid mutator transaction binding the contract method 0xb9e205ae.
//
// Solidity: function updateExchangeRate(delay uint256) returns()
func (_EthUSD *EthUSDTransactorSession) UpdateExchangeRate(delay *big.Int) (*types.Transaction, error) {
	return _EthUSD.Contract.UpdateExchangeRate(&_EthUSD.TransactOpts, delay)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amountInCents uint256) returns(success bool)
func (_EthUSD *EthUSDTransactor) Withdraw(opts *bind.TransactOpts, amountInCents *big.Int) (*types.Transaction, error) {
	return _EthUSD.contract.Transact(opts, "withdraw", amountInCents)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amountInCents uint256) returns(success bool)
func (_EthUSD *EthUSDSession) Withdraw(amountInCents *big.Int) (*types.Transaction, error) {
	return _EthUSD.Contract.Withdraw(&_EthUSD.TransactOpts, amountInCents)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amountInCents uint256) returns(success bool)
func (_EthUSD *EthUSDTransactorSession) Withdraw(amountInCents *big.Int) (*types.Transaction, error) {
	return _EthUSD.Contract.Withdraw(&_EthUSD.TransactOpts, amountInCents)
}

// ExchangeRateUpdaterABI is the input ABI used to generate the binding from.
const ExchangeRateUpdaterABI = `[{"constant":false,"inputs":[{"name":"myid","type":"bytes32"},{"name":"result","type":"string"},{"name":"proof","type":"bytes"}],"name":"__callback","outputs":[],"type":"function"},{"constant":true,"inputs":[],"name":"exchangeRate","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[{"name":"delay","type":"uint256"}],"name":"updateExchangeRate","outputs":[],"type":"function"},{"inputs":[],"type":"constructor"}]`

// ExchangeRateUpdaterBin is the compiled bytecode used for deploying new contracts.
const ExchangeRateUpdaterBin = `0x60606040526100627f110000000000000000000000000000000000000000000000000000000000000060008054600160a060020a0316141561015b5761015960005b6000600061026f731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed5b3b90565b61014b600061026281604060405190810160405280600381526020017f55524c0000000000000000000000000000000000000000000000000000000000815260200150608060405190810160405280604481526020017f6a736f6e2868747470733a2f2f706f6c6f6e6965782e636f6d2f7075626c696381526020017f3f636f6d6d616e643d72657475726e5469636b6572292e555344545f4554482e81526020017f6c61737400000000000000000000000000000000000000000000000000000000815260200150600080548190600160a060020a03168114156103a4576103a26000610041565b6108038061064c6000396000f35b505b600060009054906101000a9004600160a060020a0316600160a060020a03166338cc4831604051817c01000000000000000000000000000000000000000000000000000000000281526004018090506020604051808303816000876161da5a03f11561000257505060408051805160018054600160a060020a031916909117908190557f688dcfd70000000000000000000000000000000000000000000000000000000082527fff00000000000000000000000000000000000000000000000000000000000000851660048301529151600160a060020a0392909216925063688dcfd7916024808301926000929190829003018183876161da5a03f1156100025750505050565b5050565b5060005b919050565b11156102a3575060008054600160a060020a031916731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed179055600161026a565b60006102c2739efbea6358bed926b293d2ce63a730d6d98d43dd61005e565b11156102f9575060008054739efbea6358bed926b293d2ce63a730d6d98d43dd600160a060020a031991909116179055600161026a565b60006103187320e12a1f859b3feae5fb2a0a32c18f5a65555bbf61005e565b111561034f5750600080547320e12a1f859b3feae5fb2a0a32c18f5a65555bbf600160a060020a031991909116179055600161026a565b600061036e739a1d6e5c6c8d081ac45c6af98b74a42442afba6061005e565b1115610266575060008054600160a060020a031916739a1d6e5c6c8d081ac45c6af98b74a42442afba60179055600161026a565b505b600060009054906101000a9004600160a060020a0316600160a060020a03166338cc4831604051817c01000000000000000000000000000000000000000000000000000000000281526004018090506020604051808303816000876161da5a03f115610002575050604051805160018054600160a060020a031916909117908190557f524f388900000000000000000000000000000000000000000000000000000000825260206004838101828152895160248601528951600160a060020a0394909416955063524f3889948a949193849360449290920192868201929091829185918391869160009190601f850104600302600f01f150905090810190601f1680156104c55780820380516001836020036101000a031916815260200191505b50925050506020604051808303816000876161da5a03f11561000257505060405151915050670de0b6b3a764000062030d403a020181111561050e57600091505b509392505050565b600160009054906101000a9004600160a060020a0316600160a060020a031663adf59f9982878787604051857c01000000000000000000000000000000000000000000000000000000000281526004018084815260200180602001806020018381038352858181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f1680156105c55780820380516001836020036101000a031916815260200191505b508381038252848181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f16801561061e5780820380516001836020036101000a031916815260200191505b509550505050505060206040518083038185886185025a03f11561000257505060405151935061050691505056606060405260e060020a600035046338bbfa5081146100315780633ba0b9a9146100f8578063b9e205ae14610101575b005b60408051602060248035600481810135601f810185900485028601850190965285855261002f9581359591946044949293909201918190840183828082843750506040805160209735808a0135601f81018a90048a0283018a0190935282825296989760649791965060249190910194509092508291508401838280828437509496505050505050506101fe60008054600160a060020a03168114156102f1576102ef60005b600060006106d0731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed5b3b90565b6101ec60025481565b61002f6004355b6102eb81604060405190810160405280600381526020017f55524c0000000000000000000000000000000000000000000000000000000000815260200150608060405190810160405280604481526020017f6a736f6e2868747470733a2f2f706f6c6f6e6965782e636f6d2f7075626c696381526020017f3f636f6d6d616e643d72657475726e5469636b6572292e555344545f4554482e81526020017f6c61737400000000000000000000000000000000000000000000000000000000815260200150600080548190600160a060020a03168114156104515761044f60006100d7565b60408051918252519081900360200190f35b600160a060020a031633600160a060020a031614151561021d57610002565b6102d98260026040805160208101909152600090819052828180805b83518110156102bb57603060f860020a028482815181101561000257016020015160f860020a9081900402600160f860020a031916108015906102a65750603960f860020a028482815181101561000257016020015160f860020a9081900402600160f860020a03191611155b156103be57811561041f578560001415610416575b60008611156102ce57600a86900a909202915b509095945050505050565b6002556102e66000610108565b505050565b5050565b505b600060009054906101000a9004600160a060020a0316600160a060020a03166338cc48316040518160e060020a0281526004018090506020604051808303816000876161da5a03f11561000257505060408051805160018054600160a060020a031916909117908190557fc281d19e0000000000000000000000000000000000000000000000000000000082529151600160a060020a0392909216925063c281d19e91600482810192602092919082900301816000876161da5a03f1156100025750506040515191505090565b8381815181101561000257016020015160f860020a9081900402600160f860020a0319167f2e00000000000000000000000000000000000000000000000000000000000000141561040e57600191505b600101610239565b60001995909501945b600a83029250825060308482815181101561000257016020015160f860020a90819004810204039092019161040e565b505b600060009054906101000a9004600160a060020a0316600160a060020a03166338cc48316040518160e060020a0281526004018090506020604051808303816000876161da5a03f115610002575050604051805160018054600160a060020a031916909117908190557f524f388900000000000000000000000000000000000000000000000000000000825260206004838101828152895160248601528951600160a060020a0394909416955063524f3889948a949193849360449290920192868201929091829185918391869160009190601f850104600302600f01f150905090810190601f1680156105595780820380516001836020036101000a031916815260200191505b50925050506020604051808303816000876161da5a03f11561000257505060405151915050670de0b6b3a764000062030d403a02018111156105a257600091505b509392505050565b600160009054906101000a9004600160a060020a0316600160a060020a031663adf59f99828787876040518560e060020a0281526004018084815260200180602001806020018381038352858181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f1680156106405780820380516001836020036101000a031916815260200191505b508381038252848181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f1680156106995780820380516001836020036101000a031916815260200191505b509550505050505060206040518083038185886185025a03f11561000257505060405151935061059a915050565b5060005b919050565b1115610704575060008054600160a060020a031916731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed17905560016106cb565b6000610723739efbea6358bed926b293d2ce63a730d6d98d43dd6100f4565b111561075a575060008054739efbea6358bed926b293d2ce63a730d6d98d43dd600160a060020a03199190911617905560016106cb565b60006107797320e12a1f859b3feae5fb2a0a32c18f5a65555bbf6100f4565b11156107b05750600080547320e12a1f859b3feae5fb2a0a32c18f5a65555bbf600160a060020a03199190911617905560016106cb565b60006107cf739a1d6e5c6c8d081ac45c6af98b74a42442afba606100f4565b11156106c7575060008054600160a060020a031916739a1d6e5c6c8d081ac45c6af98b74a42442afba6017905560016106cb56`

// DeployExchangeRateUpdater deploys a new Ethereum contract, binding an instance of ExchangeRateUpdater to it.
func DeployExchangeRateUpdater(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExchangeRateUpdater, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeRateUpdaterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExchangeRateUpdaterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExchangeRateUpdater{ExchangeRateUpdaterCaller: ExchangeRateUpdaterCaller{contract: contract}, ExchangeRateUpdaterTransactor: ExchangeRateUpdaterTransactor{contract: contract}}, nil
}

// ExchangeRateUpdater is an auto generated Go binding around an Ethereum contract.
type ExchangeRateUpdater struct {
	ExchangeRateUpdaterCaller     // Read-only binding to the contract
	ExchangeRateUpdaterTransactor // Write-only binding to the contract
}

// ExchangeRateUpdaterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeRateUpdaterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeRateUpdaterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeRateUpdaterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeRateUpdaterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeRateUpdaterSession struct {
	Contract     *ExchangeRateUpdater // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ExchangeRateUpdaterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeRateUpdaterCallerSession struct {
	Contract *ExchangeRateUpdaterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ExchangeRateUpdaterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeRateUpdaterTransactorSession struct {
	Contract     *ExchangeRateUpdaterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ExchangeRateUpdaterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeRateUpdaterRaw struct {
	Contract *ExchangeRateUpdater // Generic contract binding to access the raw methods on
}

// ExchangeRateUpdaterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeRateUpdaterCallerRaw struct {
	Contract *ExchangeRateUpdaterCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeRateUpdaterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeRateUpdaterTransactorRaw struct {
	Contract *ExchangeRateUpdaterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchangeRateUpdater creates a new instance of ExchangeRateUpdater, bound to a specific deployed contract.
func NewExchangeRateUpdater(address common.Address, backend bind.ContractBackend) (*ExchangeRateUpdater, error) {
	contract, err := bindExchangeRateUpdater(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExchangeRateUpdater{ExchangeRateUpdaterCaller: ExchangeRateUpdaterCaller{contract: contract}, ExchangeRateUpdaterTransactor: ExchangeRateUpdaterTransactor{contract: contract}}, nil
}

// NewExchangeRateUpdaterCaller creates a new read-only instance of ExchangeRateUpdater, bound to a specific deployed contract.
func NewExchangeRateUpdaterCaller(address common.Address, caller bind.ContractCaller) (*ExchangeRateUpdaterCaller, error) {
	contract, err := bindExchangeRateUpdater(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeRateUpdaterCaller{contract: contract}, nil
}

// NewExchangeRateUpdaterTransactor creates a new write-only instance of ExchangeRateUpdater, bound to a specific deployed contract.
func NewExchangeRateUpdaterTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeRateUpdaterTransactor, error) {
	contract, err := bindExchangeRateUpdater(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ExchangeRateUpdaterTransactor{contract: contract}, nil
}

// bindExchangeRateUpdater binds a generic wrapper to an already deployed contract.
func bindExchangeRateUpdater(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeRateUpdaterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeRateUpdater *ExchangeRateUpdaterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExchangeRateUpdater.Contract.ExchangeRateUpdaterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeRateUpdater *ExchangeRateUpdaterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeRateUpdater.Contract.ExchangeRateUpdaterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeRateUpdater *ExchangeRateUpdaterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeRateUpdater.Contract.ExchangeRateUpdaterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeRateUpdater *ExchangeRateUpdaterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExchangeRateUpdater.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeRateUpdater *ExchangeRateUpdaterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeRateUpdater.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeRateUpdater *ExchangeRateUpdaterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeRateUpdater.Contract.contract.Transact(opts, method, params...)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() constant returns(uint256)
func (_ExchangeRateUpdater *ExchangeRateUpdaterCaller) ExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExchangeRateUpdater.contract.Call(opts, out, "exchangeRate")
	return *ret0, err
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() constant returns(uint256)
func (_ExchangeRateUpdater *ExchangeRateUpdaterSession) ExchangeRate() (*big.Int, error) {
	return _ExchangeRateUpdater.Contract.ExchangeRate(&_ExchangeRateUpdater.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() constant returns(uint256)
func (_ExchangeRateUpdater *ExchangeRateUpdaterCallerSession) ExchangeRate() (*big.Int, error) {
	return _ExchangeRateUpdater.Contract.ExchangeRate(&_ExchangeRateUpdater.CallOpts)
}

// __callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(myid bytes32, result string, proof bytes) returns()
func (_ExchangeRateUpdater *ExchangeRateUpdaterTransactor) __callback(opts *bind.TransactOpts, myid [32]byte, result string, proof []byte) (*types.Transaction, error) {
	return _ExchangeRateUpdater.contract.Transact(opts, "__callback", myid, result, proof)
}

// __callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(myid bytes32, result string, proof bytes) returns()
func (_ExchangeRateUpdater *ExchangeRateUpdaterSession) __callback(myid [32]byte, result string, proof []byte) (*types.Transaction, error) {
	return _ExchangeRateUpdater.Contract.__callback(&_ExchangeRateUpdater.TransactOpts, myid, result, proof)
}

// __callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(myid bytes32, result string, proof bytes) returns()
func (_ExchangeRateUpdater *ExchangeRateUpdaterTransactorSession) __callback(myid [32]byte, result string, proof []byte) (*types.Transaction, error) {
	return _ExchangeRateUpdater.Contract.__callback(&_ExchangeRateUpdater.TransactOpts, myid, result, proof)
}

// UpdateExchangeRate is a paid mutator transaction binding the contract method 0xb9e205ae.
//
// Solidity: function updateExchangeRate(delay uint256) returns()
func (_ExchangeRateUpdater *ExchangeRateUpdaterTransactor) UpdateExchangeRate(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _ExchangeRateUpdater.contract.Transact(opts, "updateExchangeRate", delay)
}

// UpdateExchangeRate is a paid mutator transaction binding the contract method 0xb9e205ae.
//
// Solidity: function updateExchangeRate(delay uint256) returns()
func (_ExchangeRateUpdater *ExchangeRateUpdaterSession) UpdateExchangeRate(delay *big.Int) (*types.Transaction, error) {
	return _ExchangeRateUpdater.Contract.UpdateExchangeRate(&_ExchangeRateUpdater.TransactOpts, delay)
}

// UpdateExchangeRate is a paid mutator transaction binding the contract method 0xb9e205ae.
//
// Solidity: function updateExchangeRate(delay uint256) returns()
func (_ExchangeRateUpdater *ExchangeRateUpdaterTransactorSession) UpdateExchangeRate(delay *big.Int) (*types.Transaction, error) {
	return _ExchangeRateUpdater.Contract.UpdateExchangeRate(&_ExchangeRateUpdater.TransactOpts, delay)
}

// OraclizeAddrResolverIABI is the input ABI used to generate the binding from.
const OraclizeAddrResolverIABI = `[{"constant":false,"inputs":[],"name":"getAddress","outputs":[{"name":"_addr","type":"address"}],"type":"function"}]`

// OraclizeAddrResolverIBin is the compiled bytecode used for deploying new contracts.
const OraclizeAddrResolverIBin = `0x`

// DeployOraclizeAddrResolverI deploys a new Ethereum contract, binding an instance of OraclizeAddrResolverI to it.
func DeployOraclizeAddrResolverI(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OraclizeAddrResolverI, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclizeAddrResolverIABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OraclizeAddrResolverIBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OraclizeAddrResolverI{OraclizeAddrResolverICaller: OraclizeAddrResolverICaller{contract: contract}, OraclizeAddrResolverITransactor: OraclizeAddrResolverITransactor{contract: contract}}, nil
}

// OraclizeAddrResolverI is an auto generated Go binding around an Ethereum contract.
type OraclizeAddrResolverI struct {
	OraclizeAddrResolverICaller     // Read-only binding to the contract
	OraclizeAddrResolverITransactor // Write-only binding to the contract
}

// OraclizeAddrResolverICaller is an auto generated read-only Go binding around an Ethereum contract.
type OraclizeAddrResolverICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeAddrResolverITransactor is an auto generated write-only Go binding around an Ethereum contract.
type OraclizeAddrResolverITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeAddrResolverISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OraclizeAddrResolverISession struct {
	Contract     *OraclizeAddrResolverI // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OraclizeAddrResolverICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OraclizeAddrResolverICallerSession struct {
	Contract *OraclizeAddrResolverICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// OraclizeAddrResolverITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OraclizeAddrResolverITransactorSession struct {
	Contract     *OraclizeAddrResolverITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// OraclizeAddrResolverIRaw is an auto generated low-level Go binding around an Ethereum contract.
type OraclizeAddrResolverIRaw struct {
	Contract *OraclizeAddrResolverI // Generic contract binding to access the raw methods on
}

// OraclizeAddrResolverICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OraclizeAddrResolverICallerRaw struct {
	Contract *OraclizeAddrResolverICaller // Generic read-only contract binding to access the raw methods on
}

// OraclizeAddrResolverITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OraclizeAddrResolverITransactorRaw struct {
	Contract *OraclizeAddrResolverITransactor // Generic write-only contract binding to access the raw methods on
}

// NewOraclizeAddrResolverI creates a new instance of OraclizeAddrResolverI, bound to a specific deployed contract.
func NewOraclizeAddrResolverI(address common.Address, backend bind.ContractBackend) (*OraclizeAddrResolverI, error) {
	contract, err := bindOraclizeAddrResolverI(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OraclizeAddrResolverI{OraclizeAddrResolverICaller: OraclizeAddrResolverICaller{contract: contract}, OraclizeAddrResolverITransactor: OraclizeAddrResolverITransactor{contract: contract}}, nil
}

// NewOraclizeAddrResolverICaller creates a new read-only instance of OraclizeAddrResolverI, bound to a specific deployed contract.
func NewOraclizeAddrResolverICaller(address common.Address, caller bind.ContractCaller) (*OraclizeAddrResolverICaller, error) {
	contract, err := bindOraclizeAddrResolverI(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &OraclizeAddrResolverICaller{contract: contract}, nil
}

// NewOraclizeAddrResolverITransactor creates a new write-only instance of OraclizeAddrResolverI, bound to a specific deployed contract.
func NewOraclizeAddrResolverITransactor(address common.Address, transactor bind.ContractTransactor) (*OraclizeAddrResolverITransactor, error) {
	contract, err := bindOraclizeAddrResolverI(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &OraclizeAddrResolverITransactor{contract: contract}, nil
}

// bindOraclizeAddrResolverI binds a generic wrapper to an already deployed contract.
func bindOraclizeAddrResolverI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclizeAddrResolverIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OraclizeAddrResolverI *OraclizeAddrResolverIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OraclizeAddrResolverI.Contract.OraclizeAddrResolverICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OraclizeAddrResolverI *OraclizeAddrResolverIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.OraclizeAddrResolverITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OraclizeAddrResolverI *OraclizeAddrResolverIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.OraclizeAddrResolverITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OraclizeAddrResolverI *OraclizeAddrResolverICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OraclizeAddrResolverI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OraclizeAddrResolverI *OraclizeAddrResolverITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OraclizeAddrResolverI *OraclizeAddrResolverITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a paid mutator transaction binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() returns(_addr address)
func (_OraclizeAddrResolverI *OraclizeAddrResolverITransactor) GetAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeAddrResolverI.contract.Transact(opts, "getAddress")
}

// GetAddress is a paid mutator transaction binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() returns(_addr address)
func (_OraclizeAddrResolverI *OraclizeAddrResolverISession) GetAddress() (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.GetAddress(&_OraclizeAddrResolverI.TransactOpts)
}

// GetAddress is a paid mutator transaction binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() returns(_addr address)
func (_OraclizeAddrResolverI *OraclizeAddrResolverITransactorSession) GetAddress() (*types.Transaction, error) {
	return _OraclizeAddrResolverI.Contract.GetAddress(&_OraclizeAddrResolverI.TransactOpts)
}

// OraclizeIABI is the input ABI used to generate the binding from.
const OraclizeIABI = `[{"constant":false,"inputs":[{"name":"_datasource","type":"string"},{"name":"gaslimit","type":"uint256"}],"name":"getPrice","outputs":[{"name":"_dsprice","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[{"name":"_datasource","type":"string"}],"name":"getPrice","outputs":[{"name":"_dsprice","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[{"name":"_coupon","type":"string"}],"name":"useCoupon","outputs":[],"type":"function"},{"constant":false,"inputs":[{"name":"_proofType","type":"bytes1"}],"name":"setProofType","outputs":[],"type":"function"},{"constant":false,"inputs":[{"name":"_timestamp","type":"uint256"},{"name":"_datasource","type":"string"},{"name":"_arg1","type":"string"},{"name":"_arg2","type":"string"}],"name":"query2","outputs":[{"name":"_id","type":"bytes32"}],"type":"function"},{"constant":false,"inputs":[{"name":"_timestamp","type":"uint256"},{"name":"_datasource","type":"string"},{"name":"_arg1","type":"string"},{"name":"_arg2","type":"string"},{"name":"_gaslimit","type":"uint256"}],"name":"query2_withGasLimit","outputs":[{"name":"_id","type":"bytes32"}],"type":"function"},{"constant":false,"inputs":[{"name":"_timestamp","type":"uint256"},{"name":"_datasource","type":"string"},{"name":"_arg","type":"string"}],"name":"query","outputs":[{"name":"_id","type":"bytes32"}],"type":"function"},{"constant":true,"inputs":[],"name":"cbAddress","outputs":[{"name":"","type":"address"}],"type":"function"},{"constant":false,"inputs":[{"name":"_timestamp","type":"uint256"},{"name":"_datasource","type":"string"},{"name":"_arg","type":"string"},{"name":"_gaslimit","type":"uint256"}],"name":"query_withGasLimit","outputs":[{"name":"_id","type":"bytes32"}],"type":"function"},{"constant":false,"inputs":[{"name":"_gasPrice","type":"uint256"}],"name":"setCustomGasPrice","outputs":[],"type":"function"}]`

// OraclizeIBin is the compiled bytecode used for deploying new contracts.
const OraclizeIBin = `0x`

// DeployOraclizeI deploys a new Ethereum contract, binding an instance of OraclizeI to it.
func DeployOraclizeI(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OraclizeI, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclizeIABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OraclizeIBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OraclizeI{OraclizeICaller: OraclizeICaller{contract: contract}, OraclizeITransactor: OraclizeITransactor{contract: contract}}, nil
}

// OraclizeI is an auto generated Go binding around an Ethereum contract.
type OraclizeI struct {
	OraclizeICaller     // Read-only binding to the contract
	OraclizeITransactor // Write-only binding to the contract
}

// OraclizeICaller is an auto generated read-only Go binding around an Ethereum contract.
type OraclizeICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeITransactor is an auto generated write-only Go binding around an Ethereum contract.
type OraclizeITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclizeISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OraclizeISession struct {
	Contract     *OraclizeI        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OraclizeICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OraclizeICallerSession struct {
	Contract *OraclizeICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// OraclizeITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OraclizeITransactorSession struct {
	Contract     *OraclizeITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OraclizeIRaw is an auto generated low-level Go binding around an Ethereum contract.
type OraclizeIRaw struct {
	Contract *OraclizeI // Generic contract binding to access the raw methods on
}

// OraclizeICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OraclizeICallerRaw struct {
	Contract *OraclizeICaller // Generic read-only contract binding to access the raw methods on
}

// OraclizeITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OraclizeITransactorRaw struct {
	Contract *OraclizeITransactor // Generic write-only contract binding to access the raw methods on
}

// NewOraclizeI creates a new instance of OraclizeI, bound to a specific deployed contract.
func NewOraclizeI(address common.Address, backend bind.ContractBackend) (*OraclizeI, error) {
	contract, err := bindOraclizeI(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OraclizeI{OraclizeICaller: OraclizeICaller{contract: contract}, OraclizeITransactor: OraclizeITransactor{contract: contract}}, nil
}

// NewOraclizeICaller creates a new read-only instance of OraclizeI, bound to a specific deployed contract.
func NewOraclizeICaller(address common.Address, caller bind.ContractCaller) (*OraclizeICaller, error) {
	contract, err := bindOraclizeI(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &OraclizeICaller{contract: contract}, nil
}

// NewOraclizeITransactor creates a new write-only instance of OraclizeI, bound to a specific deployed contract.
func NewOraclizeITransactor(address common.Address, transactor bind.ContractTransactor) (*OraclizeITransactor, error) {
	contract, err := bindOraclizeI(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &OraclizeITransactor{contract: contract}, nil
}

// bindOraclizeI binds a generic wrapper to an already deployed contract.
func bindOraclizeI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclizeIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OraclizeI *OraclizeIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OraclizeI.Contract.OraclizeICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OraclizeI *OraclizeIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeI.Contract.OraclizeITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OraclizeI *OraclizeIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OraclizeI.Contract.OraclizeITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OraclizeI *OraclizeICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OraclizeI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OraclizeI *OraclizeITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OraclizeI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OraclizeI *OraclizeITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OraclizeI.Contract.contract.Transact(opts, method, params...)
}

// CbAddress is a free data retrieval call binding the contract method 0xc281d19e.
//
// Solidity: function cbAddress() constant returns(address)
func (_OraclizeI *OraclizeICaller) CbAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OraclizeI.contract.Call(opts, out, "cbAddress")
	return *ret0, err
}

// CbAddress is a free data retrieval call binding the contract method 0xc281d19e.
//
// Solidity: function cbAddress() constant returns(address)
func (_OraclizeI *OraclizeISession) CbAddress() (common.Address, error) {
	return _OraclizeI.Contract.CbAddress(&_OraclizeI.CallOpts)
}

// CbAddress is a free data retrieval call binding the contract method 0xc281d19e.
//
// Solidity: function cbAddress() constant returns(address)
func (_OraclizeI *OraclizeICallerSession) CbAddress() (common.Address, error) {
	return _OraclizeI.Contract.CbAddress(&_OraclizeI.CallOpts)
}

// GetPrice is a paid mutator transaction binding the contract method 0x524f3889.
//
// Solidity: function getPrice(_datasource string) returns(_dsprice uint256)
func (_OraclizeI *OraclizeITransactor) GetPrice(opts *bind.TransactOpts, _datasource string) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "getPrice", _datasource)
}

// GetPrice is a paid mutator transaction binding the contract method 0x524f3889.
//
// Solidity: function getPrice(_datasource string) returns(_dsprice uint256)
func (_OraclizeI *OraclizeISession) GetPrice(_datasource string) (*types.Transaction, error) {
	return _OraclizeI.Contract.GetPrice(&_OraclizeI.TransactOpts, _datasource)
}

// GetPrice is a paid mutator transaction binding the contract method 0x524f3889.
//
// Solidity: function getPrice(_datasource string) returns(_dsprice uint256)
func (_OraclizeI *OraclizeITransactorSession) GetPrice(_datasource string) (*types.Transaction, error) {
	return _OraclizeI.Contract.GetPrice(&_OraclizeI.TransactOpts, _datasource)
}

// Query is a paid mutator transaction binding the contract method 0xadf59f99.
//
// Solidity: function query(_timestamp uint256, _datasource string, _arg string) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactor) Query(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg string) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "query", _timestamp, _datasource, _arg)
}

// Query is a paid mutator transaction binding the contract method 0xadf59f99.
//
// Solidity: function query(_timestamp uint256, _datasource string, _arg string) returns(_id bytes32)
func (_OraclizeI *OraclizeISession) Query(_timestamp *big.Int, _datasource string, _arg string) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg)
}

// Query is a paid mutator transaction binding the contract method 0xadf59f99.
//
// Solidity: function query(_timestamp uint256, _datasource string, _arg string) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactorSession) Query(_timestamp *big.Int, _datasource string, _arg string) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg)
}

// Query2 is a paid mutator transaction binding the contract method 0x77228659.
//
// Solidity: function query2(_timestamp uint256, _datasource string, _arg1 string, _arg2 string) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactor) Query2(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg1 string, _arg2 string) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "query2", _timestamp, _datasource, _arg1, _arg2)
}

// Query2 is a paid mutator transaction binding the contract method 0x77228659.
//
// Solidity: function query2(_timestamp uint256, _datasource string, _arg1 string, _arg2 string) returns(_id bytes32)
func (_OraclizeI *OraclizeISession) Query2(_timestamp *big.Int, _datasource string, _arg1 string, _arg2 string) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query2(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg1, _arg2)
}

// Query2 is a paid mutator transaction binding the contract method 0x77228659.
//
// Solidity: function query2(_timestamp uint256, _datasource string, _arg1 string, _arg2 string) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactorSession) Query2(_timestamp *big.Int, _datasource string, _arg1 string, _arg2 string) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query2(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg1, _arg2)
}

// Query2_withGasLimit is a paid mutator transaction binding the contract method 0x85dee34c.
//
// Solidity: function query2_withGasLimit(_timestamp uint256, _datasource string, _arg1 string, _arg2 string, _gaslimit uint256) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactor) Query2_withGasLimit(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg1 string, _arg2 string, _gaslimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "query2_withGasLimit", _timestamp, _datasource, _arg1, _arg2, _gaslimit)
}

// Query2_withGasLimit is a paid mutator transaction binding the contract method 0x85dee34c.
//
// Solidity: function query2_withGasLimit(_timestamp uint256, _datasource string, _arg1 string, _arg2 string, _gaslimit uint256) returns(_id bytes32)
func (_OraclizeI *OraclizeISession) Query2_withGasLimit(_timestamp *big.Int, _datasource string, _arg1 string, _arg2 string, _gaslimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query2_withGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg1, _arg2, _gaslimit)
}

// Query2_withGasLimit is a paid mutator transaction binding the contract method 0x85dee34c.
//
// Solidity: function query2_withGasLimit(_timestamp uint256, _datasource string, _arg1 string, _arg2 string, _gaslimit uint256) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactorSession) Query2_withGasLimit(_timestamp *big.Int, _datasource string, _arg1 string, _arg2 string, _gaslimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query2_withGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg1, _arg2, _gaslimit)
}

// Query_withGasLimit is a paid mutator transaction binding the contract method 0xc51be90f.
//
// Solidity: function query_withGasLimit(_timestamp uint256, _datasource string, _arg string, _gaslimit uint256) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactor) Query_withGasLimit(opts *bind.TransactOpts, _timestamp *big.Int, _datasource string, _arg string, _gaslimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "query_withGasLimit", _timestamp, _datasource, _arg, _gaslimit)
}

// Query_withGasLimit is a paid mutator transaction binding the contract method 0xc51be90f.
//
// Solidity: function query_withGasLimit(_timestamp uint256, _datasource string, _arg string, _gaslimit uint256) returns(_id bytes32)
func (_OraclizeI *OraclizeISession) Query_withGasLimit(_timestamp *big.Int, _datasource string, _arg string, _gaslimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query_withGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg, _gaslimit)
}

// Query_withGasLimit is a paid mutator transaction binding the contract method 0xc51be90f.
//
// Solidity: function query_withGasLimit(_timestamp uint256, _datasource string, _arg string, _gaslimit uint256) returns(_id bytes32)
func (_OraclizeI *OraclizeITransactorSession) Query_withGasLimit(_timestamp *big.Int, _datasource string, _arg string, _gaslimit *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.Query_withGasLimit(&_OraclizeI.TransactOpts, _timestamp, _datasource, _arg, _gaslimit)
}

// SetCustomGasPrice is a paid mutator transaction binding the contract method 0xca6ad1e4.
//
// Solidity: function setCustomGasPrice(_gasPrice uint256) returns()
func (_OraclizeI *OraclizeITransactor) SetCustomGasPrice(opts *bind.TransactOpts, _gasPrice *big.Int) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "setCustomGasPrice", _gasPrice)
}

// SetCustomGasPrice is a paid mutator transaction binding the contract method 0xca6ad1e4.
//
// Solidity: function setCustomGasPrice(_gasPrice uint256) returns()
func (_OraclizeI *OraclizeISession) SetCustomGasPrice(_gasPrice *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.SetCustomGasPrice(&_OraclizeI.TransactOpts, _gasPrice)
}

// SetCustomGasPrice is a paid mutator transaction binding the contract method 0xca6ad1e4.
//
// Solidity: function setCustomGasPrice(_gasPrice uint256) returns()
func (_OraclizeI *OraclizeITransactorSession) SetCustomGasPrice(_gasPrice *big.Int) (*types.Transaction, error) {
	return _OraclizeI.Contract.SetCustomGasPrice(&_OraclizeI.TransactOpts, _gasPrice)
}

// SetProofType is a paid mutator transaction binding the contract method 0x688dcfd7.
//
// Solidity: function setProofType(_proofType bytes1) returns()
func (_OraclizeI *OraclizeITransactor) SetProofType(opts *bind.TransactOpts, _proofType [1]byte) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "setProofType", _proofType)
}

// SetProofType is a paid mutator transaction binding the contract method 0x688dcfd7.
//
// Solidity: function setProofType(_proofType bytes1) returns()
func (_OraclizeI *OraclizeISession) SetProofType(_proofType [1]byte) (*types.Transaction, error) {
	return _OraclizeI.Contract.SetProofType(&_OraclizeI.TransactOpts, _proofType)
}

// SetProofType is a paid mutator transaction binding the contract method 0x688dcfd7.
//
// Solidity: function setProofType(_proofType bytes1) returns()
func (_OraclizeI *OraclizeITransactorSession) SetProofType(_proofType [1]byte) (*types.Transaction, error) {
	return _OraclizeI.Contract.SetProofType(&_OraclizeI.TransactOpts, _proofType)
}

// UseCoupon is a paid mutator transaction binding the contract method 0x60f66701.
//
// Solidity: function useCoupon(_coupon string) returns()
func (_OraclizeI *OraclizeITransactor) UseCoupon(opts *bind.TransactOpts, _coupon string) (*types.Transaction, error) {
	return _OraclizeI.contract.Transact(opts, "useCoupon", _coupon)
}

// UseCoupon is a paid mutator transaction binding the contract method 0x60f66701.
//
// Solidity: function useCoupon(_coupon string) returns()
func (_OraclizeI *OraclizeISession) UseCoupon(_coupon string) (*types.Transaction, error) {
	return _OraclizeI.Contract.UseCoupon(&_OraclizeI.TransactOpts, _coupon)
}

// UseCoupon is a paid mutator transaction binding the contract method 0x60f66701.
//
// Solidity: function useCoupon(_coupon string) returns()
func (_OraclizeI *OraclizeITransactorSession) UseCoupon(_coupon string) (*types.Transaction, error) {
	return _OraclizeI.Contract.UseCoupon(&_OraclizeI.TransactOpts, _coupon)
}

// StandardTokenABI is the input ABI used to generate the binding from.
const StandardTokenABI = `[{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[{"name":"success","type":"bool"}],"type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"success","type":"bool"}],"type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[{"name":"success","type":"bool"}],"type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"remaining","type":"uint256"}],"type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_from","type":"address"},{"indexed":true,"name":"_to","type":"address"},{"indexed":false,"name":"_value","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_owner","type":"address"},{"indexed":true,"name":"_spender","type":"address"},{"indexed":false,"name":"_value","type":"uint256"}],"name":"Approval","type":"event"}]`

// StandardTokenBin is the compiled bytecode used for deploying new contracts.
const StandardTokenBin = `0x60606040526102f3806100126000396000f3606060405236156100565760e060020a6000350463095ea7b3811461005857806318160ddd146100cd57806323b872dd146100d657806370a08231146101c3578063a9059cbb146101f1578063dd62ed3e14610298575b005b6102cc60043560243533600160a060020a03908116600081815260026020908152604080832094871680845294825280832086905580518681529051929493927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060015b92915050565b6101df60005481565b6102cc600435602435604435600160a060020a038316600090815260016020526040812054829010801590610129575060026020908152604080832033600160a060020a03168452909152812054829010155b80156101355750600082115b156102e057600160a060020a03838116600081815260016020908152604080832080548801905588851680845281842080548990039055600283528184203390961684529482529182902080548790039055815186815291519293927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35060016102e4565b600160a060020a03600435166000908152600160205260409020545b60408051918252519081900360200190f35b6102cc60043560243533600160a060020a03166000908152600160205260408120548290108015906102235750600082115b156102eb5733600160a060020a03908116600081815260016020908152604080832080548890039055938716808352918490208054870190558351868152935191937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a35060016100c7565b6101df600435602435600160a060020a038281166000908152600260209081526040808320938516835292905220546100c7565b604080519115158252519081900360200190f35b5060005b9392505050565b5060006100c756`

// DeployStandardToken deploys a new Ethereum contract, binding an instance of StandardToken to it.
func DeployStandardToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StandardToken, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StandardTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}}, nil
}

// StandardToken is an auto generated Go binding around an Ethereum contract.
type StandardToken struct {
	StandardTokenCaller     // Read-only binding to the contract
	StandardTokenTransactor // Write-only binding to the contract
}

// StandardTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardTokenSession struct {
	Contract     *StandardToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandardTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardTokenCallerSession struct {
	Contract *StandardTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StandardTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardTokenTransactorSession struct {
	Contract     *StandardTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StandardTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandardTokenRaw struct {
	Contract *StandardToken // Generic contract binding to access the raw methods on
}

// StandardTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardTokenCallerRaw struct {
	Contract *StandardTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StandardTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardTokenTransactorRaw struct {
	Contract *StandardTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardToken creates a new instance of StandardToken, bound to a specific deployed contract.
func NewStandardToken(address common.Address, backend bind.ContractBackend) (*StandardToken, error) {
	contract, err := bindStandardToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}}, nil
}

// NewStandardTokenCaller creates a new read-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenCaller(address common.Address, caller bind.ContractCaller) (*StandardTokenCaller, error) {
	contract, err := bindStandardToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenCaller{contract: contract}, nil
}

// NewStandardTokenTransactor creates a new write-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardTokenTransactor, error) {
	contract, err := bindStandardToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransactor{contract: contract}, nil
}

// bindStandardToken binds a generic wrapper to an already deployed contract.
func bindStandardToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.StandardTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TestEthUSDABI is the input ABI used to generate the binding from.
const TestEthUSDABI = `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[{"name":"success","type":"bool"}],"type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"success","type":"bool"}],"type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"balances","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[{"name":"amountInCents","type":"uint256"}],"name":"withdraw","outputs":[{"name":"success","type":"bool"}],"type":"function"},{"constant":false,"inputs":[{"name":"myid","type":"bytes32"},{"name":"result","type":"string"},{"name":"proof","type":"bytes"}],"name":"__callback","outputs":[],"type":"function"},{"constant":true,"inputs":[],"name":"exchangeRate","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[],"name":"usingOraclize","outputs":[],"type":"function"},{"constant":false,"inputs":[],"name":"purchase","outputs":[{"name":"success","type":"bool"}],"type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[{"name":"success","type":"bool"}],"type":"function"},{"constant":false,"inputs":[{"name":"delay","type":"uint256"}],"name":"updateExchangeRate","outputs":[],"type":"function"},{"constant":false,"inputs":[{"name":"_amount","type":"uint256"}],"name":"setExchangeRate","outputs":[],"type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"remaining","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[{"name":"_address","type":"address"},{"name":"_amount","type":"uint256"}],"name":"setBalance","outputs":[],"type":"function"},{"inputs":[{"name":"_name","type":"string"},{"name":"_symbol","type":"string"}],"type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Purchase","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Withdraw","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_from","type":"address"},{"indexed":true,"name":"_to","type":"address"},{"indexed":false,"name":"_value","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_owner","type":"address"},{"indexed":true,"name":"_spender","type":"address"},{"indexed":false,"name":"_value","type":"uint256"}],"name":"Approval","type":"event"}]`

// TestEthUSDBin is the compiled bytecode used for deploying new contracts.
const TestEthUSDBin = `0x6060604052604051610bb9380380610bb9833981516080805160a0840160409081526003858701527f55524c00000000000000000000000000000000000000000000000000000000009483019490945283519182018452604482527f6a736f6e2868747470733a2f2f706f6c6f6e6965782e636f6d2f7075626c696360208301527f3f636f6d6d616e643d72657475726e5469636b6572292e555344545f4554482e938201939093527f6c61737400000000000000000000000000000000000000000000000000000000908401528201910181818160066000509080519060200190828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061013a57805160ff19168380011785555b5061016a9291505b808211156101c35760008155600101610126565b8280016001018555821561011e579182015b8281111561011e57825182600050559160200191906001019061014c565b50508060076000509080519060200190828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106101c757805160ff19168380011785555b506101f7929150610126565b5090565b828001600101855582156101b7579182015b828111156101b75782518260005055916020019190600101906101d9565b5050505050506109ae8061020b6000396000f3606060405236156100cf5760e060020a600035046306fdde0381146100d1578063095ea7b31461012f57806318160ddd146101a457806323b872dd146101ad57806327e235e31461029a5780632e1a7d4d146102b257806338bbfa50146103175780633ba0b9a9146103a85780634b1b16bc146103b157806364edfbf0146103b657806370a082311461048257806395d89b41146104b0578063a9059cbb1461050e578063b9e205ae146105b5578063db068e0e14610677578063dd62ed3e14610682578063e30443bc146106b6575b005b6040805160068054602060026001831615610100026000190190921691909104601f81018290048202840182019094528383526106da93908301828280156107875780601f1061075c57610100808354040283529160200191610787565b61074860043560243533600160a060020a03908116600081815260056020908152604080832094871680845294825280832086905580518681529051929493927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060015b92915050565b61049e60035481565b610748600435602435604435600160a060020a038316600090815260046020526040812054829010801590610200575060056020908152604080832033600160a060020a03168452909152812054829010155b801561020c5750600082115b1561078f57600160a060020a03838116600081815260046020908152604080832080548801905588851680845281842080548990039055600583528184203390961684529482529182902080548790039055815186815291519293927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a3506001610793565b61049e60043560086020526000908152604090205481565b61074860043533600160a060020a0316600090815260086020526040812054819083901061079a5760408120805484900390556107a5600254600354600091670de0b6b3a764000091020430600160a060020a031631101561091657506104b061091b565b60408051602060248035600481810135601f81018590048502860185019096528585526100cf9581359591946044949293909201918190840183828082843750506040805160209735808a0135601f81018a90048a0283018a019093528282529698976064979196506024919091019450909250829150840183828082843750949650505050505050610818600090565b61049e60025481565b6100cf565b6107486000600060003411156109055760408051348152905133600160a060020a0316917f2499a5330ab0979cc612135e7883ebc3cd5c9f7a8508f042540c34723348f632919081900360200190a2506040805160025433600160a060020a0316808552600860209081528486208054670de0b6b3a7640000349590950294909404938401905560038054840190558284529351919390927f2499a5330ab0979cc612135e7883ebc3cd5c9f7a8508f042540c34723348f632929081900390910190a26001915061090a565b600160a060020a03600435166000908152600860205260409020545b60408051918252519081900360200190f35b6040805160078054602060026001831615610100026000190190921691909104601f81018290048202840182019094528383526106da93908301828280156107875780601f1061075c57610100808354040283529160200191610787565b61074860043560243533600160a060020a03166000908152600460205260408120548290108015906105405750600082115b1561090e5733600160a060020a03908116600081815260046020908152604080832080548890039055938716808352918490208054870190558351868152935191937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a350600161019e565b6100cf6004355b50604080518082018252600381527f55524c00000000000000000000000000000000000000000000000000000000006020918201528151608081018352604481527f6a736f6e2868747470733a2f2f706f6c6f6e6965782e636f6d2f7075626c6963918101919091527f3f636f6d6d616e643d72657475726e5469636b6572292e555344545f4554482e918101919091527f6c6173740000000000000000000000000000000000000000000000000000000060609190910152565b6004356002556100cf565b61049e600435602435600160a060020a0382811660009081526005602090815260408083209385168352929052205461019e565b600160a060020a0360043516600090815260086020526040902060243590556100cf565b60405180806020018281038252838181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f16801561073a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b604080519115158252519081900360200190f35b820191906000526020600020905b81548152906001019060200180831161076a57829003601f168201915b505050505081565b5060005b9392505050565b600091505b50919050565b604051908404670de0b6b3a764000002915033600160a060020a031690600090839082818181858883f1505060038054919091039055506040805183815290517f2499a5330ab0979cc612135e7883ebc3cd5c9f7a8508f042540c34723348f6329181900360200190a26001915061079f565b600160a060020a031633600160a060020a031614151561083757610002565b6108f38260026040805160208101909152600090819052828180805b83518110156108d557603060f860020a028482815181101561000257016020015160f860020a9081900402600160f860020a031916108015906108c05750603960f860020a028482815181101561000257016020015160f860020a9081900402600160f860020a03191611155b1561091e57811561097e578560001415610976575b60008611156108e857600a86900a909202915b509095945050505050565b60025561090060006105bc565b505050565b600091505b5090565b50600061019e565b506002545b90565b8381815181101561000257016020015160f860020a9081900402600160f860020a0319167f2e00000000000000000000000000000000000000000000000000000000000000141561096e57600191505b600101610853565b600019909501945b600a83029250825060308482815181101561000257016020015160f860020a90819004810204039092019161096e56`

// DeployTestEthUSD deploys a new Ethereum contract, binding an instance of TestEthUSD to it.
func DeployTestEthUSD(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string) (common.Address, *types.Transaction, *TestEthUSD, error) {
	parsed, err := abi.JSON(strings.NewReader(TestEthUSDABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestEthUSDBin), backend, _name, _symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestEthUSD{TestEthUSDCaller: TestEthUSDCaller{contract: contract}, TestEthUSDTransactor: TestEthUSDTransactor{contract: contract}}, nil
}

// TestEthUSD is an auto generated Go binding around an Ethereum contract.
type TestEthUSD struct {
	TestEthUSDCaller     // Read-only binding to the contract
	TestEthUSDTransactor // Write-only binding to the contract
}

// TestEthUSDCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestEthUSDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestEthUSDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestEthUSDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestEthUSDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestEthUSDSession struct {
	Contract     *TestEthUSD       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestEthUSDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestEthUSDCallerSession struct {
	Contract *TestEthUSDCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TestEthUSDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestEthUSDTransactorSession struct {
	Contract     *TestEthUSDTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TestEthUSDRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestEthUSDRaw struct {
	Contract *TestEthUSD // Generic contract binding to access the raw methods on
}

// TestEthUSDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestEthUSDCallerRaw struct {
	Contract *TestEthUSDCaller // Generic read-only contract binding to access the raw methods on
}

// TestEthUSDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestEthUSDTransactorRaw struct {
	Contract *TestEthUSDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestEthUSD creates a new instance of TestEthUSD, bound to a specific deployed contract.
func NewTestEthUSD(address common.Address, backend bind.ContractBackend) (*TestEthUSD, error) {
	contract, err := bindTestEthUSD(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestEthUSD{TestEthUSDCaller: TestEthUSDCaller{contract: contract}, TestEthUSDTransactor: TestEthUSDTransactor{contract: contract}}, nil
}

// NewTestEthUSDCaller creates a new read-only instance of TestEthUSD, bound to a specific deployed contract.
func NewTestEthUSDCaller(address common.Address, caller bind.ContractCaller) (*TestEthUSDCaller, error) {
	contract, err := bindTestEthUSD(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TestEthUSDCaller{contract: contract}, nil
}

// NewTestEthUSDTransactor creates a new write-only instance of TestEthUSD, bound to a specific deployed contract.
func NewTestEthUSDTransactor(address common.Address, transactor bind.ContractTransactor) (*TestEthUSDTransactor, error) {
	contract, err := bindTestEthUSD(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TestEthUSDTransactor{contract: contract}, nil
}

// bindTestEthUSD binds a generic wrapper to an already deployed contract.
func bindTestEthUSD(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestEthUSDABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestEthUSD *TestEthUSDRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TestEthUSD.Contract.TestEthUSDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestEthUSD *TestEthUSDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEthUSD.Contract.TestEthUSDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestEthUSD *TestEthUSDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestEthUSD.Contract.TestEthUSDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestEthUSD *TestEthUSDCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TestEthUSD.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestEthUSD *TestEthUSDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEthUSD.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestEthUSD *TestEthUSDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestEthUSD.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_TestEthUSD *TestEthUSDCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestEthUSD.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_TestEthUSD *TestEthUSDSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _TestEthUSD.Contract.Allowance(&_TestEthUSD.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_TestEthUSD *TestEthUSDCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _TestEthUSD.Contract.Allowance(&_TestEthUSD.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_TestEthUSD *TestEthUSDCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestEthUSD.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_TestEthUSD *TestEthUSDSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _TestEthUSD.Contract.BalanceOf(&_TestEthUSD.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_TestEthUSD *TestEthUSDCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _TestEthUSD.Contract.BalanceOf(&_TestEthUSD.CallOpts, _owner)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_TestEthUSD *TestEthUSDCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestEthUSD.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_TestEthUSD *TestEthUSDSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _TestEthUSD.Contract.Balances(&_TestEthUSD.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_TestEthUSD *TestEthUSDCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _TestEthUSD.Contract.Balances(&_TestEthUSD.CallOpts, arg0)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() constant returns(uint256)
func (_TestEthUSD *TestEthUSDCaller) ExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestEthUSD.contract.Call(opts, out, "exchangeRate")
	return *ret0, err
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() constant returns(uint256)
func (_TestEthUSD *TestEthUSDSession) ExchangeRate() (*big.Int, error) {
	return _TestEthUSD.Contract.ExchangeRate(&_TestEthUSD.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() constant returns(uint256)
func (_TestEthUSD *TestEthUSDCallerSession) ExchangeRate() (*big.Int, error) {
	return _TestEthUSD.Contract.ExchangeRate(&_TestEthUSD.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TestEthUSD *TestEthUSDCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TestEthUSD.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TestEthUSD *TestEthUSDSession) Name() (string, error) {
	return _TestEthUSD.Contract.Name(&_TestEthUSD.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TestEthUSD *TestEthUSDCallerSession) Name() (string, error) {
	return _TestEthUSD.Contract.Name(&_TestEthUSD.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TestEthUSD *TestEthUSDCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TestEthUSD.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TestEthUSD *TestEthUSDSession) Symbol() (string, error) {
	return _TestEthUSD.Contract.Symbol(&_TestEthUSD.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TestEthUSD *TestEthUSDCallerSession) Symbol() (string, error) {
	return _TestEthUSD.Contract.Symbol(&_TestEthUSD.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TestEthUSD *TestEthUSDCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestEthUSD.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TestEthUSD *TestEthUSDSession) TotalSupply() (*big.Int, error) {
	return _TestEthUSD.Contract.TotalSupply(&_TestEthUSD.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TestEthUSD *TestEthUSDCallerSession) TotalSupply() (*big.Int, error) {
	return _TestEthUSD.Contract.TotalSupply(&_TestEthUSD.CallOpts)
}

// __callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(myid bytes32, result string, proof bytes) returns()
func (_TestEthUSD *TestEthUSDTransactor) __callback(opts *bind.TransactOpts, myid [32]byte, result string, proof []byte) (*types.Transaction, error) {
	return _TestEthUSD.contract.Transact(opts, "__callback", myid, result, proof)
}

// __callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(myid bytes32, result string, proof bytes) returns()
func (_TestEthUSD *TestEthUSDSession) __callback(myid [32]byte, result string, proof []byte) (*types.Transaction, error) {
	return _TestEthUSD.Contract.__callback(&_TestEthUSD.TransactOpts, myid, result, proof)
}

// __callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(myid bytes32, result string, proof bytes) returns()
func (_TestEthUSD *TestEthUSDTransactorSession) __callback(myid [32]byte, result string, proof []byte) (*types.Transaction, error) {
	return _TestEthUSD.Contract.__callback(&_TestEthUSD.TransactOpts, myid, result, proof)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_TestEthUSD *TestEthUSDTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TestEthUSD.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_TestEthUSD *TestEthUSDSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TestEthUSD.Contract.Approve(&_TestEthUSD.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_TestEthUSD *TestEthUSDTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TestEthUSD.Contract.Approve(&_TestEthUSD.TransactOpts, _spender, _value)
}

// Purchase is a paid mutator transaction binding the contract method 0x64edfbf0.
//
// Solidity: function purchase() returns(success bool)
func (_TestEthUSD *TestEthUSDTransactor) Purchase(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEthUSD.contract.Transact(opts, "purchase")
}

// Purchase is a paid mutator transaction binding the contract method 0x64edfbf0.
//
// Solidity: function purchase() returns(success bool)
func (_TestEthUSD *TestEthUSDSession) Purchase() (*types.Transaction, error) {
	return _TestEthUSD.Contract.Purchase(&_TestEthUSD.TransactOpts)
}

// Purchase is a paid mutator transaction binding the contract method 0x64edfbf0.
//
// Solidity: function purchase() returns(success bool)
func (_TestEthUSD *TestEthUSDTransactorSession) Purchase() (*types.Transaction, error) {
	return _TestEthUSD.Contract.Purchase(&_TestEthUSD.TransactOpts)
}

// SetBalance is a paid mutator transaction binding the contract method 0xe30443bc.
//
// Solidity: function setBalance(_address address, _amount uint256) returns()
func (_TestEthUSD *TestEthUSDTransactor) SetBalance(opts *bind.TransactOpts, _address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TestEthUSD.contract.Transact(opts, "setBalance", _address, _amount)
}

// SetBalance is a paid mutator transaction binding the contract method 0xe30443bc.
//
// Solidity: function setBalance(_address address, _amount uint256) returns()
func (_TestEthUSD *TestEthUSDSession) SetBalance(_address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TestEthUSD.Contract.SetBalance(&_TestEthUSD.TransactOpts, _address, _amount)
}

// SetBalance is a paid mutator transaction binding the contract method 0xe30443bc.
//
// Solidity: function setBalance(_address address, _amount uint256) returns()
func (_TestEthUSD *TestEthUSDTransactorSession) SetBalance(_address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TestEthUSD.Contract.SetBalance(&_TestEthUSD.TransactOpts, _address, _amount)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xdb068e0e.
//
// Solidity: function setExchangeRate(_amount uint256) returns()
func (_TestEthUSD *TestEthUSDTransactor) SetExchangeRate(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _TestEthUSD.contract.Transact(opts, "setExchangeRate", _amount)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xdb068e0e.
//
// Solidity: function setExchangeRate(_amount uint256) returns()
func (_TestEthUSD *TestEthUSDSession) SetExchangeRate(_amount *big.Int) (*types.Transaction, error) {
	return _TestEthUSD.Contract.SetExchangeRate(&_TestEthUSD.TransactOpts, _amount)
}

// SetExchangeRate is a paid mutator transaction binding the contract method 0xdb068e0e.
//
// Solidity: function setExchangeRate(_amount uint256) returns()
func (_TestEthUSD *TestEthUSDTransactorSession) SetExchangeRate(_amount *big.Int) (*types.Transaction, error) {
	return _TestEthUSD.Contract.SetExchangeRate(&_TestEthUSD.TransactOpts, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_TestEthUSD *TestEthUSDTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TestEthUSD.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_TestEthUSD *TestEthUSDSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TestEthUSD.Contract.Transfer(&_TestEthUSD.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_TestEthUSD *TestEthUSDTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TestEthUSD.Contract.Transfer(&_TestEthUSD.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_TestEthUSD *TestEthUSDTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TestEthUSD.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_TestEthUSD *TestEthUSDSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TestEthUSD.Contract.TransferFrom(&_TestEthUSD.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_TestEthUSD *TestEthUSDTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TestEthUSD.Contract.TransferFrom(&_TestEthUSD.TransactOpts, _from, _to, _value)
}

// UpdateExchangeRate is a paid mutator transaction binding the contract method 0xb9e205ae.
//
// Solidity: function updateExchangeRate(delay uint256) returns()
func (_TestEthUSD *TestEthUSDTransactor) UpdateExchangeRate(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _TestEthUSD.contract.Transact(opts, "updateExchangeRate", delay)
}

// UpdateExchangeRate is a paid mutator transaction binding the contract method 0xb9e205ae.
//
// Solidity: function updateExchangeRate(delay uint256) returns()
func (_TestEthUSD *TestEthUSDSession) UpdateExchangeRate(delay *big.Int) (*types.Transaction, error) {
	return _TestEthUSD.Contract.UpdateExchangeRate(&_TestEthUSD.TransactOpts, delay)
}

// UpdateExchangeRate is a paid mutator transaction binding the contract method 0xb9e205ae.
//
// Solidity: function updateExchangeRate(delay uint256) returns()
func (_TestEthUSD *TestEthUSDTransactorSession) UpdateExchangeRate(delay *big.Int) (*types.Transaction, error) {
	return _TestEthUSD.Contract.UpdateExchangeRate(&_TestEthUSD.TransactOpts, delay)
}

// UsingOraclize is a paid mutator transaction binding the contract method 0x4b1b16bc.
//
// Solidity: function usingOraclize() returns()
func (_TestEthUSD *TestEthUSDTransactor) UsingOraclize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestEthUSD.contract.Transact(opts, "usingOraclize")
}

// UsingOraclize is a paid mutator transaction binding the contract method 0x4b1b16bc.
//
// Solidity: function usingOraclize() returns()
func (_TestEthUSD *TestEthUSDSession) UsingOraclize() (*types.Transaction, error) {
	return _TestEthUSD.Contract.UsingOraclize(&_TestEthUSD.TransactOpts)
}

// UsingOraclize is a paid mutator transaction binding the contract method 0x4b1b16bc.
//
// Solidity: function usingOraclize() returns()
func (_TestEthUSD *TestEthUSDTransactorSession) UsingOraclize() (*types.Transaction, error) {
	return _TestEthUSD.Contract.UsingOraclize(&_TestEthUSD.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amountInCents uint256) returns(success bool)
func (_TestEthUSD *TestEthUSDTransactor) Withdraw(opts *bind.TransactOpts, amountInCents *big.Int) (*types.Transaction, error) {
	return _TestEthUSD.contract.Transact(opts, "withdraw", amountInCents)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amountInCents uint256) returns(success bool)
func (_TestEthUSD *TestEthUSDSession) Withdraw(amountInCents *big.Int) (*types.Transaction, error) {
	return _TestEthUSD.Contract.Withdraw(&_TestEthUSD.TransactOpts, amountInCents)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amountInCents uint256) returns(success bool)
func (_TestEthUSD *TestEthUSDTransactorSession) Withdraw(amountInCents *big.Int) (*types.Transaction, error) {
	return _TestEthUSD.Contract.Withdraw(&_TestEthUSD.TransactOpts, amountInCents)
}

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = `[{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[{"name":"success","type":"bool"}],"type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"success","type":"bool"}],"type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[{"name":"success","type":"bool"}],"type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"remaining","type":"uint256"}],"type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_from","type":"address"},{"indexed":true,"name":"_to","type":"address"},{"indexed":false,"name":"_value","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_owner","type":"address"},{"indexed":true,"name":"_spender","type":"address"},{"indexed":false,"name":"_value","type":"uint256"}],"name":"Approval","type":"event"}]`

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `0x`

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Token *TokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

// UsingOraclizeABI is the input ABI used to generate the binding from.
const UsingOraclizeABI = `[]`

// UsingOraclizeBin is the compiled bytecode used for deploying new contracts.
const UsingOraclizeBin = `0x606060405260068060106000396000f3606060405200`

// DeployUsingOraclize deploys a new Ethereum contract, binding an instance of UsingOraclize to it.
func DeployUsingOraclize(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UsingOraclize, error) {
	parsed, err := abi.JSON(strings.NewReader(UsingOraclizeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UsingOraclizeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UsingOraclize{UsingOraclizeCaller: UsingOraclizeCaller{contract: contract}, UsingOraclizeTransactor: UsingOraclizeTransactor{contract: contract}}, nil
}

// UsingOraclize is an auto generated Go binding around an Ethereum contract.
type UsingOraclize struct {
	UsingOraclizeCaller     // Read-only binding to the contract
	UsingOraclizeTransactor // Write-only binding to the contract
}

// UsingOraclizeCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsingOraclizeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsingOraclizeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsingOraclizeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsingOraclizeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsingOraclizeSession struct {
	Contract     *UsingOraclize    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsingOraclizeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsingOraclizeCallerSession struct {
	Contract *UsingOraclizeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// UsingOraclizeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsingOraclizeTransactorSession struct {
	Contract     *UsingOraclizeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// UsingOraclizeRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsingOraclizeRaw struct {
	Contract *UsingOraclize // Generic contract binding to access the raw methods on
}

// UsingOraclizeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsingOraclizeCallerRaw struct {
	Contract *UsingOraclizeCaller // Generic read-only contract binding to access the raw methods on
}

// UsingOraclizeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsingOraclizeTransactorRaw struct {
	Contract *UsingOraclizeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsingOraclize creates a new instance of UsingOraclize, bound to a specific deployed contract.
func NewUsingOraclize(address common.Address, backend bind.ContractBackend) (*UsingOraclize, error) {
	contract, err := bindUsingOraclize(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UsingOraclize{UsingOraclizeCaller: UsingOraclizeCaller{contract: contract}, UsingOraclizeTransactor: UsingOraclizeTransactor{contract: contract}}, nil
}

// NewUsingOraclizeCaller creates a new read-only instance of UsingOraclize, bound to a specific deployed contract.
func NewUsingOraclizeCaller(address common.Address, caller bind.ContractCaller) (*UsingOraclizeCaller, error) {
	contract, err := bindUsingOraclize(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &UsingOraclizeCaller{contract: contract}, nil
}

// NewUsingOraclizeTransactor creates a new write-only instance of UsingOraclize, bound to a specific deployed contract.
func NewUsingOraclizeTransactor(address common.Address, transactor bind.ContractTransactor) (*UsingOraclizeTransactor, error) {
	contract, err := bindUsingOraclize(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &UsingOraclizeTransactor{contract: contract}, nil
}

// bindUsingOraclize binds a generic wrapper to an already deployed contract.
func bindUsingOraclize(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UsingOraclizeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsingOraclize *UsingOraclizeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UsingOraclize.Contract.UsingOraclizeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsingOraclize *UsingOraclizeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsingOraclize.Contract.UsingOraclizeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsingOraclize *UsingOraclizeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UsingOraclize.Contract.UsingOraclizeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsingOraclize *UsingOraclizeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UsingOraclize.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsingOraclize *UsingOraclizeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsingOraclize.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsingOraclize *UsingOraclizeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UsingOraclize.Contract.contract.Transact(opts, method, params...)
}
