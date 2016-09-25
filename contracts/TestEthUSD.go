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
const EthUSDABI = `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"type":"function"},{"constant":false,"inputs":[{"name":"myid","type":"bytes32"},{"name":"result","type":"string"},{"name":"proof","type":"bytes"}],"name":"__callback","outputs":[],"type":"function"},{"constant":true,"inputs":[],"name":"exchangeRate","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[],"name":"purchase","outputs":[{"name":"success","type":"bool"}],"type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[{"name":"delay","type":"uint256"}],"name":"updateExchangeRate","outputs":[],"type":"function"},{"inputs":[],"type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Purchase","type":"event"}]`

// EthUSDBin is the compiled bytecode used for deploying new contracts.
const EthUSDBin = `0x60606040526100627f110000000000000000000000000000000000000000000000000000000000000060008054600160a060020a031614156101355761013360005b60006000610249731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed5b3b90565b610125600061023c81604060405190810160405280600381526020017f55524c0000000000000000000000000000000000000000000000000000000000815260200150606060405190810160405280603f81526020017f68747470733a2f2f706f6c6f6e6965782e636f6d2f7075626c69633f636f6d6d81526020017f616e643d72657475726e5469636b6572292e555344545f4554482e6c61737400815260200150600080548190600160a060020a031681141561037e5761037c6000610041565b610970806106266000396000f35b505b600060009054906101000a9004600160a060020a0316600160a060020a03166338cc4831604051817c01000000000000000000000000000000000000000000000000000000000281526004018090506020604051808303816000876161da5a03f11561000257505060408051805160018054600160a060020a031916909117908190557f688dcfd70000000000000000000000000000000000000000000000000000000082527fff00000000000000000000000000000000000000000000000000000000000000851660048301529151600160a060020a0392909216925063688dcfd7916024808301926000929190829003018183876161da5a03f1156100025750505050565b5050565b5060005b919050565b111561027d575060008054600160a060020a031916731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed1790556001610244565b600061029c739efbea6358bed926b293d2ce63a730d6d98d43dd61005e565b11156102d3575060008054739efbea6358bed926b293d2ce63a730d6d98d43dd600160a060020a0319919091161790556001610244565b60006102f27320e12a1f859b3feae5fb2a0a32c18f5a65555bbf61005e565b11156103295750600080547320e12a1f859b3feae5fb2a0a32c18f5a65555bbf600160a060020a0319919091161790556001610244565b6000610348739a1d6e5c6c8d081ac45c6af98b74a42442afba6061005e565b1115610240575060008054600160a060020a031916739a1d6e5c6c8d081ac45c6af98b74a42442afba601790556001610244565b505b600060009054906101000a9004600160a060020a0316600160a060020a03166338cc4831604051817c01000000000000000000000000000000000000000000000000000000000281526004018090506020604051808303816000876161da5a03f115610002575050604051805160018054600160a060020a031916909117908190557f524f388900000000000000000000000000000000000000000000000000000000825260206004838101828152895160248601528951600160a060020a0394909416955063524f3889948a949193849360449290920192868201929091829185918391869160009190601f850104600302600f01f150905090810190601f16801561049f5780820380516001836020036101000a031916815260200191505b50925050506020604051808303816000876161da5a03f11561000257505060405151915050670de0b6b3a764000062030d403a02018111156104e857600091505b509392505050565b600160009054906101000a9004600160a060020a0316600160a060020a031663adf59f9982878787604051857c01000000000000000000000000000000000000000000000000000000000281526004018084815260200180602001806020018381038352858181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f16801561059f5780820380516001836020036101000a031916815260200191505b508381038252848181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f1680156105f85780820380516001836020036101000a031916815260200191505b509550505050505060206040518083038185886185025a03f1156100025750506040515193506104e091505056606060405236156100565760e060020a600035046306fdde03811461005857806338bbfa50146100b35780633ba0b9a91461017a57806364edfbf01461018357806370a08231146101bd578063b9e205ae146101e2575b005b6040805160028054602060018216156101000260001901909116829004601f81018290048202840182019094528383526102a693908301828280156103655780601f1061033a57610100808354040283529160200191610365565b60408051602060248035600481810135601f81018590048502860185019096528585526100569581359591946044949293909201918190840183828082843750506040805160209735808a0135601f81018a90048a0283018a01909352828252969897606497919650602491909101945090925082915084018382808284375094965050505050505061036d60008054600160a060020a031681141561045d5761045b60005b60006000610835731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed5b3b90565b61031460035481565b610326600060003411156104505760035433600160a060020a03168252600460205260409091208054349290920290910190556001610454565b610314600435600160a060020a0381166000908152600460205260409020545b919050565b61005660043561045781604060405190810160405280600381526020017f55524c0000000000000000000000000000000000000000000000000000000000815260200150606060405190810160405280603f81526020017f68747470733a2f2f706f6c6f6e6965782e636f6d2f7075626c69633f636f6d6d81526020017f616e643d72657475726e5469636b6572292e555344545f4554482e6c61737400815260200150600080548190600160a060020a03168114156105c0576105be6000610159565b60405180806020018281038252838181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f1680156103065780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60408051918252519081900360200190f35b604080519115158252519081900360200190f35b820191906000526020600020905b81548152906001019060200180831161034857829003601f168201915b505050505081565b600160a060020a031633600160a060020a031614151561038c57610002565b6104488260026040805160208101909152600090819052828180805b835181101561042a57603060f860020a028482815181101561000257016020015160f860020a9081900402600160f860020a031916108015906104155750603960f860020a028482815181101561000257016020015160f860020a9081900402600160f860020a03191611155b1561052d57811561058e578560001415610585575b600086111561043d57600a86900a909202915b509095945050505050565b600355505050565b5060005b90565b5050565b505b600060009054906101000a9004600160a060020a0316600160a060020a03166338cc48316040518160e060020a0281526004018090506020604051808303816000876161da5a03f11561000257505060408051805160018054600160a060020a031916909117908190557fc281d19e0000000000000000000000000000000000000000000000000000000082529151600160a060020a0392909216925063c281d19e91600482810192602092919082900301816000876161da5a03f1156100025750506040515191506104549050565b8381815181101561000257016020015160f860020a9081900402600160f860020a0319167f2e00000000000000000000000000000000000000000000000000000000000000141561057d57600191505b6001016103a8565b60001995909501945b600a83029250825060308482815181101561000257016020015160f860020a90819004810204039092019161057d565b505b600060009054906101000a9004600160a060020a0316600160a060020a03166338cc48316040518160e060020a0281526004018090506020604051808303816000876161da5a03f115610002575050604051805160018054600160a060020a031916909117908190557f524f388900000000000000000000000000000000000000000000000000000000825260206004838101828152895160248601528951600160a060020a0394909416955063524f3889948a9491938493604490920192868201929091829185918391869160009190601f850104600302600f01f150905090810190601f1680156106c75780820380516001836020036101000a031916815260200191505b50925050506020604051808303816000876161da5a03f11561000257505060405151915050670de0b6b3a764000062030d403a020181111561071057600091505b509392505050565b600160009054906101000a9004600160a060020a0316600160a060020a031663adf59f99828787876040518560e060020a0281526004018084815260200180602001806020018381038352858181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f1680156107ae5780820380516001836020036101000a031916815260200191505b508381038252848181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f1680156108075780820380516001836020036101000a031916815260200191505b509550505050505060206040518083038185886185025a03f115610002575050604051519350610708915050565b1115610869575060008054600160a060020a031916731d3b2638a7cc9f2cb3d298a3da7a90b67e5506ed17905560016101dd565b6000610888739efbea6358bed926b293d2ce63a730d6d98d43dd610176565b11156108bf575060008054739efbea6358bed926b293d2ce63a730d6d98d43dd600160a060020a03199190911617905560016101dd565b60006108de7320e12a1f859b3feae5fb2a0a32c18f5a65555bbf610176565b11156109155750600080547320e12a1f859b3feae5fb2a0a32c18f5a65555bbf600160a060020a03199190911617905560016101dd565b6000610934739a1d6e5c6c8d081ac45c6af98b74a42442afba60610176565b1115610968575060008054600160a060020a031916739a1d6e5c6c8d081ac45c6af98b74a42442afba6017905560016101dd565b5060006101dd56`

// DeployEthUSD deploys a new Ethereum contract, binding an instance of EthUSD to it.
func DeployEthUSD(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthUSD, error) {
	parsed, err := abi.JSON(strings.NewReader(EthUSDABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthUSDBin), backend)
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

// TestEthUSDABI is the input ABI used to generate the binding from.
const TestEthUSDABI = `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"type":"function"},{"constant":false,"inputs":[{"name":"myid","type":"bytes32"},{"name":"result","type":"string"},{"name":"proof","type":"bytes"}],"name":"__callback","outputs":[],"type":"function"},{"constant":true,"inputs":[],"name":"exchangeRate","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[],"name":"usingOraclize","outputs":[],"type":"function"},{"constant":false,"inputs":[],"name":"purchase","outputs":[{"name":"success","type":"bool"}],"type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"type":"function"},{"constant":false,"inputs":[{"name":"delay","type":"uint256"}],"name":"updateExchangeRate","outputs":[],"type":"function"},{"constant":false,"inputs":[{"name":"_amount","type":"uint256"}],"name":"setExchangeRate","outputs":[],"type":"function"},{"inputs":[{"name":"_name","type":"string"}],"type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Purchase","type":"event"}]`

// TestEthUSDBin is the compiled bytecode used for deploying new contracts.
const TestEthUSDBin = `0x60606040526040516106103803806106108339815160a0820160409081526003838501527f55524c000000000000000000000000000000000000000000000000000000000060809093019290925281518084018352603f81527f68747470733a2f2f706f6c6f6e6965782e636f6d2f7075626c69633f636f6d6d60208201527f616e643d72657475726e5469636b6572292e555344545f4554482e6c617374009201919091520160028054825160008390527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace602060018416156101000260001901909316849004601f90810184900482019386019083901061012557805160ff19168380011785555b506101559291505b808211156101665760008155600101610111565b82800160010185558215610109579182015b82811115610109578251826000505591602001919060010190610137565b5050506104a68061016a6000396000f35b5090566060604052361561006c5760e060020a600035046306fdde03811461006e57806338bbfa50146100c95780633ba0b9a91461015a5780634b1b16bc1461016357806364edfbf01461016857806370a08231146101a5578063b9e205ae146101c9578063db068e0e1461025a575b005b6040805160028054602060018216156101000260001901909116829004601f810182900482028401820190945283835261026593908301828280156103245780601f106102f957610100808354040283529160200191610324565b60408051602060248035600481810135601f810185900485028601850190965285855261006c9581359591946044949293909201918190840183828082843750506040805160209735808a0135601f81018a90048a0283018a01909352828252969897606497919650602491909101945090925082915084018382808284375094965050505050505061032c600090565b6102d360035481565b61006c565b6102e56000600034111561040f575060035433600160a060020a031660009081526004602052604090208054349290920290910190556001610413565b6102d3600435600160a060020a038116600090815260046020526040902054919050565b604080518082018252600381527f55524c00000000000000000000000000000000000000000000000000000000006020918201528151606081018352603f81527f68747470733a2f2f706f6c6f6e6965782e636f6d2f7075626c69633f636f6d6d918101919091527f616e643d72657475726e5469636b6572292e555344545f4554482e6c6173740091015261006c565b60043560035561006c565b60405180806020018281038252838181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f1680156102c55780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60408051918252519081900360200190f35b604080519115158252519081900360200190f35b820191906000526020600020905b81548152906001019060200180831161030757829003601f168201915b505050505081565b600160a060020a031633600160a060020a031614151561034b57610002565b6104078260026040805160208101909152600090819052828180805b83518110156103e957603060f860020a028482815181101561000257016020015160f860020a9081900402600160f860020a031916108015906103d45750603960f860020a028482815181101561000257016020015160f860020a9081900402600160f860020a03191611155b1561041657811561047657856000141561046e575b60008611156103fc57600a86900a909202915b509095945050505050565b600355505050565b5060005b90565b8381815181101561000257016020015160f860020a9081900402600160f860020a0319167f2e00000000000000000000000000000000000000000000000000000000000000141561046657600191505b600101610367565b600019909501945b600a83029250825060308482815181101561000257016020015160f860020a90819004810204039092019161046656`

// DeployTestEthUSD deploys a new Ethereum contract, binding an instance of TestEthUSD to it.
func DeployTestEthUSD(auth *bind.TransactOpts, backend bind.ContractBackend, _name string) (common.Address, *types.Transaction, *TestEthUSD, error) {
	parsed, err := abi.JSON(strings.NewReader(TestEthUSDABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestEthUSDBin), backend, _name)
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
