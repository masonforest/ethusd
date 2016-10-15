package main

//go:generate abigen --solc=solc-0.3.6 --sol TestEthUSD.sol --pkg main --out TestEthUSD.go

import (
	"crypto/ecdsa"
	"flag"
	"log"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

type Value big.Int

func (v *Value) toWei() *big.Int {
	return big.NewInt(0).Mul(((*big.Int)(v)), common.Ether)
}

func NewValue(value int64) *Value {
	return (*Value)(big.NewInt(value))
}

var (
	account1Key      *ecdsa.PrivateKey
	account1         common.Address
	account2Key      *ecdsa.PrivateKey
	account2         common.Address
	auth             *bind.TransactOpts
	backend          *backends.SimulatedBackend
	NONCE            = big.NewInt(2)
	STARTING_BALANCE = NewValue(4).toWei()
)

func paidTransactOpts(value *big.Int, key *ecdsa.PrivateKey) *bind.TransactOpts {
	t := bind.NewKeyedTransactor(key)
	t.Value = value
	return t
}

func TestMain(m *testing.M) {
	account1Key, _ = crypto.GenerateKey()
	account1 = crypto.PubkeyToAddress(account1Key.PublicKey)
	auth = bind.NewKeyedTransactor(account1Key)
	backend = backends.NewSimulatedBackend(
		core.GenesisAccount{
			Address: account1,
			Balance: STARTING_BALANCE,
		},
		core.GenesisAccount{
			Address: account2,
			Balance: STARTING_BALANCE,
		},
	)
	flag.Parse()
	os.Exit(m.Run())
}

func deploy(name string, symbol string) *TestEthUSDSession {
	_, _, token, err := DeployTestEthUSD(
		bind.NewKeyedTransactor(account1Key),
		backend,
		name,
		symbol,
	)

	checkErr(err)

	backend.Commit()

	return &TestEthUSDSession{
		Contract:     token,
		TransactOpts: *bind.NewKeyedTransactor(account1Key),
	}
}

func TestInitialize(t *testing.T) {
	token := deploy("EthUSD", "ETHUSD")

	name, _ := token.Name()
	assert.Equal(t, "EthUSD", name)

	symbol, _ := token.Symbol()
	assert.Equal(t, "ETHUSD", symbol)
}

func TestPurchase(t *testing.T) {
	// 1 Ether
	var PURCHASE_AMOUNT = NewValue(1).toWei()
	// $12 USD per ether
	var EXCHANGE_RATE = big.NewInt(1200)
	// 12 ETHUSD
	var EXPECTED_AMOUNT = big.NewInt(1200)

	session := deploy("TestEthUSD", "ETHUSD")
	session.SetExchangeRate(EXCHANGE_RATE)

	_, err := session.Contract.TestEthUSDTransactor.Purchase(
		paidTransactOpts(PURCHASE_AMOUNT, account1Key),
	)

	checkErr(err)

	backend.Commit()

	balance, _ := session.BalanceOf(account1)
	assert.Equal(t, EXPECTED_AMOUNT, balance)
	totalSupply, _ := session.TotalSupply()
	assert.Equal(t, EXPECTED_AMOUNT, totalSupply)
}

/*
	When the contract has greater than or equal value
	to totalSupply of tokens issued the contract should pay
	out the exchange rate.
*/

func TestWithdrawWhenContractIsFull(t *testing.T) {
	var PURCHASE_AMOUNT = NewValue(1).toWei()
	// $12 USD per ether
	var PURCHASE_EXCHANGE_RATE = big.NewInt(1200)
	// $12 USD per ether
	var WITHDRAWL_EXCHANGE_RATE = big.NewInt(1200)
	// 12 ETHUSD
	var WITHDRAWL_AMOUNT = big.NewInt(1200)
	var COST_OF_TRANSACTION = big.NewInt(926340)

	startingBalance, _ := backend.BalanceAt(nil, account1, nil)
	session := deploy("TestEthUSD", "ETHUSD")
	session.SetExchangeRate(PURCHASE_EXCHANGE_RATE)

	_, err := session.Contract.TestEthUSDTransactor.Purchase(
		paidTransactOpts(PURCHASE_AMOUNT, account1Key),
	)

	_, err = session.Withdraw(WITHDRAWL_AMOUNT)
	session.SetExchangeRate(WITHDRAWL_EXCHANGE_RATE)

	checkErr(err)

	backend.Commit()
	weiBalance, _ := backend.BalanceAt(nil, account1, nil)
	expectedWeiBalance := startingBalance.Sub(startingBalance, COST_OF_TRANSACTION)
	assert.Equal(t, expectedWeiBalance, weiBalance)
	// // 12 ETHUSD
	// var PURCHASE_AMOUNT = big.NewInt(1200)
	// var PURCHASE_EXCHANGE_RATE = big.NewInt(1200)
	// // var WITHDRAWL_AMOUNT = big.NewInt(1200)
	// // var WITHDRAWL_EXCHANGE_RATE = big.NewInt(1300)
	// var EXPECTED_AMOUNT = NewValue(1).toWei()
	// var COST_OF_TRANSACTION = big.NewInt(0)
	//
	// startingBalance, _ := backend.BalanceAt(nil, account1, nil)
	// fmt.Printf("1: %s\n", startingBalance)
	// session := deploy("EthUSD", "ETHUSD")
	// startingBalance, _ = backend.BalanceAt(nil, account1, nil)
	// fmt.Printf("2: %s\n", startingBalance)
	// session.SetExchangeRate(PURCHASE_EXCHANGE_RATE)
	// _, err := session.Contract.TestEthUSDTransactor.Purchase(
	// 	paidTransactOpts(PURCHASE_AMOUNT, account1Key),
	// )
	// // session.SetExchangeRate(WITHDRAWL_EXCHANGE_RATE)
	//
	// // _, err = session.Withdraw(WITHDRAWL_AMOUNT)
	//
	// checkErr(err)
	//
	// backend.Commit()
	//
	// balance, _ := session.BalanceOf(account1)
	// fmt.Printf("balance: %s\n", balance)
	//
	// weiBalance, _ := backend.BalanceAt(nil, account1, nil)
	// fmt.Printf("3: %s\n", weiBalance)
	// expectedWeiBalance := startingBalance
	// expectedWeiBalance.Sub(expectedWeiBalance, COST_OF_TRANSACTION)
	// expectedWeiBalance.Add(expectedWeiBalance, EXPECTED_AMOUNT)
	// assert.Equal(t, expectedWeiBalance, weiBalance)
}

/*
	When the contract has less value than the totalSupply of
	tokens issued the contract should pay
	out a rate equal to the total supply divided by the
	amount of value left in the contract.
*/

// func TestWithdrawWhenContractIsLessThanFull(t *testing.T) {
// 	session := deploy("EthUSD", "ETHUSD")
// 	session.SetBalance(account1, NewValue(2).toWei())
// 	var COST_OF_TRANSACTION = big.NewInt(1000000000000077851)
// 	startingBalance, _ := backend.BalanceAt(nil, account1, nil)
//
// 	_, err := session.Withdraw(NewValue(1).toWei())
//
// 	checkErr(err)
//
// 	backend.Commit()
//
// 	ethUSDBalance, _ := session.BalanceOf(account1)
// 	assert.Equal(t, NewValue(1).toWei(), ethUSDBalance)
//
// 	weiBalance, _ := backend.BalanceAt(nil, account1, nil)
// 	expectedWeiBalance := startingBalance
// 	expectedWeiBalance.Sub(expectedWeiBalance, COST_OF_TRANSACTION)
// 	expectedWeiBalance.Add(expectedWeiBalance, NewValue(1).toWei())
// 	assert.Equal(t, expectedWeiBalance, weiBalance)
// }

func checkErr(err error) {
	if err != nil {
		log.Fatalf("%v", err)
	}
}
