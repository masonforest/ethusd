package main

//go:generate abigen --solc=solc-0.3.6 --sol TestEthUSD.sol --pkg main --out TestEthUSD.go

import (
	"crypto/ecdsa"
	"flag"
	"fmt"
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
	var WITHDRAWAL_EXCHANGE_RATE = big.NewInt(1200)
	// 12 ETHUSD
	var WITHDRAWAL_AMOUNT = big.NewInt(1200)
	var COST_OF_TRANSACTION = big.NewInt(941802)

	var weiPurchased = big.NewInt(0)
	weiPurchased.Mul(PURCHASE_AMOUNT, PURCHASE_EXCHANGE_RATE)

	var weiWithdrawn = big.NewInt(0)
	weiWithdrawn.Mul(WITHDRAWAL_AMOUNT, NewValue(1).toWei())
	weiWithdrawn.Mul(weiWithdrawn, PURCHASE_EXCHANGE_RATE)
	weiWithdrawn.Div(weiWithdrawn, WITHDRAWAL_EXCHANGE_RATE)

	fmt.Printf("P: %s\n", weiPurchased)
	fmt.Printf("W: %s\n", weiWithdrawn)
	startingBalance, _ := backend.BalanceAt(nil, account1, nil)
	var expectedBalance = big.NewInt(0)
	expectedBalance.Sub(startingBalance, COST_OF_TRANSACTION)
	expectedBalance.Add(expectedBalance, weiPurchased)
	expectedBalance.Sub(expectedBalance, weiWithdrawn)

	session := deploy("TestEthUSD", "ETHUSD")
	session.SetExchangeRate(PURCHASE_EXCHANGE_RATE)

	_, err := session.Contract.TestEthUSDTransactor.Purchase(
		paidTransactOpts(PURCHASE_AMOUNT, account1Key),
	)
	backend.Commit()

	session.SetExchangeRate(WITHDRAWAL_EXCHANGE_RATE)
	weiBalance, _ := backend.BalanceAt(nil, account1, nil)
	fmt.Printf("Before: %s\n", weiBalance)
	_, err = session.Withdraw(WITHDRAWAL_AMOUNT)

	checkErr(err)

	backend.Commit()
	weiBalance, _ = backend.BalanceAt(nil, account1, nil)
	fmt.Printf("After: %s\n", weiBalance)

	weiBalance, _ = backend.BalanceAt(nil, account1, nil)
	assert.Equal(t, expectedBalance, weiBalance)
}

/*
	When the contract has less value than the totalSupply of
	tokens issued the contract should pay
	out a rate equal to the total supply divided by the
	amount of value left in the contract.
*/

func checkErr(err error) {
	if err != nil {
		log.Fatalf("%v", err)
	}
}
