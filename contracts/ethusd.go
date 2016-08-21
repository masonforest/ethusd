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
const EthUSDABI = `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[],"type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"},{"name":"","type":"address"}],"name":"spentAllowance","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"constant":true,"inputs":[{"name":"","type":"address"},{"name":"","type":"address"}],"name":"allowance","outputs":[{"name":"","type":"uint256"}],"type":"function"},{"inputs":[{"name":"initialSupply","type":"uint256"},{"name":"tokenName","type":"string"}],"type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"}]`

// EthUSDBin is the compiled bytecode used for deploying new contracts.
const EthUSDBin = `0x606060405260405161041a38038061041a833981016040528051608051909101600160a060020a0333166000908152600360209081526040822084905581548351838052601f6002600019610100600186161502019093169290920482018390047f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563908101939192909186019083901060ca57805160ff19168380011785555b5060b89291505b8082111560f7576000815560010160a6565b5050505061031f806100fb6000396000f35b82800160010185558215609f579182015b82811115609f57825182600050559160200191906001019060db565b509056606060405236156100615760e060020a600035046306fdde038114610069578063313ce567146100c757806370a08231146100d357806395d89b41146100eb578063a9059cbb1461014a578063dc3080f214610179578063dd62ed3e1461019e575b6101c2610002565b6040805160008054602060026001831615610100026000190190921691909104601f81018290048202840182019094528383526101c493908301828280156102855780601f1061025a57610100808354040283529160200191610285565b61023260025460ff1681565b61024860043560036020526000908152604090205481565b6101c460018054604080516020601f6002600019610100878916150201909516949094049384018190048102820181019092528281529291908301828280156102855780601f1061025a57610100808354040283529160200191610285565b6101c2600435602435600160a060020a0333166000908152600360205260409020548190101561028d57610002565b6005602090815260043560009081526040808220909252602435815220546102489081565b60046020818152903560009081526040808220909252602435815220546102489081565b005b60405180806020018281038252838181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f1680156102245780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6040805160ff9092168252519081900360200190f35b60408051918252519081900360200190f35b820191906000526020600020905b81548152906001019060200180831161026857829003601f168201915b505050505081565b600160a060020a03821660009081526003602052604090205481810110156102b457610002565b600160a060020a03338116600081815260036020908152604080832080548790039055938616808352918490208054860190558351858152935191937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a3505056`

// DeployEthUSD deploys a new Ethereum contract, binding an instance of EthUSD to it.
func DeployEthUSD(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, tokenName string) (common.Address, *types.Transaction, *EthUSD, error) {
	parsed, err := abi.JSON(strings.NewReader(EthUSDABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthUSDBin), backend, initialSupply, tokenName)
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
	contract, err := bindEthUSD(address, backend.(bind.ContractCaller), backend.(bind.ContractTransactor))
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
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_EthUSD *EthUSDCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EthUSD.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_EthUSD *EthUSDSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _EthUSD.Contract.Allowance(&_EthUSD.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_EthUSD *EthUSDCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _EthUSD.Contract.Allowance(&_EthUSD.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_EthUSD *EthUSDCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EthUSD.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_EthUSD *EthUSDSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _EthUSD.Contract.BalanceOf(&_EthUSD.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_EthUSD *EthUSDCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _EthUSD.Contract.BalanceOf(&_EthUSD.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_EthUSD *EthUSDCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _EthUSD.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_EthUSD *EthUSDSession) Decimals() (uint8, error) {
	return _EthUSD.Contract.Decimals(&_EthUSD.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_EthUSD *EthUSDCallerSession) Decimals() (uint8, error) {
	return _EthUSD.Contract.Decimals(&_EthUSD.CallOpts)
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

// SpentAllowance is a free data retrieval call binding the contract method 0xdc3080f2.
//
// Solidity: function spentAllowance( address,  address) constant returns(uint256)
func (_EthUSD *EthUSDCaller) SpentAllowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EthUSD.contract.Call(opts, out, "spentAllowance", arg0, arg1)
	return *ret0, err
}

// SpentAllowance is a free data retrieval call binding the contract method 0xdc3080f2.
//
// Solidity: function spentAllowance( address,  address) constant returns(uint256)
func (_EthUSD *EthUSDSession) SpentAllowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _EthUSD.Contract.SpentAllowance(&_EthUSD.CallOpts, arg0, arg1)
}

// SpentAllowance is a free data retrieval call binding the contract method 0xdc3080f2.
//
// Solidity: function spentAllowance( address,  address) constant returns(uint256)
func (_EthUSD *EthUSDCallerSession) SpentAllowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _EthUSD.Contract.SpentAllowance(&_EthUSD.CallOpts, arg0, arg1)
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

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns()
func (_EthUSD *EthUSDTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EthUSD.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns()
func (_EthUSD *EthUSDSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EthUSD.Contract.Transfer(&_EthUSD.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns()
func (_EthUSD *EthUSDTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EthUSD.Contract.Transfer(&_EthUSD.TransactOpts, _to, _value)
}
