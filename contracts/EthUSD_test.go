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
	senderKey        *ecdsa.PrivateKey
	sender           common.Address
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
	senderKey, _ = crypto.GenerateKey()
	sender = crypto.PubkeyToAddress(senderKey.PublicKey)
	auth = bind.NewKeyedTransactor(senderKey)
	backend = backends.NewSimulatedBackend(
		core.GenesisAccount{
			Address: sender,
			Balance: STARTING_BALANCE,
		})
	flag.Parse()
	os.Exit(m.Run())
}

func deploy(name string, symbol string) *TestEthUSDSession {
	_, _, token, err := DeployTestEthUSD(
		bind.NewKeyedTransactor(senderKey),
		backend,
		name,
		symbol,
	)

	checkErr(err)

	backend.Commit()

	return &TestEthUSDSession{
		Contract:     token,
		TransactOpts: *bind.NewKeyedTransactor(senderKey),
	}
}

func TestInitialize(t *testing.T) {
	token := deploy("EthUSD", "ETHUSD")

	name, _ := token.Name()
	assert.Equal(t, "EthUSD", name)

	symbol, _ := token.Symbol()
	assert.Equal(t, "ETHUSD", symbol)
}

func TestPurchaseWhenExchangeRateLessThanCap(t *testing.T) {
	// 1 Ether
	var AMOUNT_TO_PURCHASE = NewValue(1).toWei()
	// $11 USD per ether
	var EXCHANGE_RATE = big.NewInt(1100)
	// $12 USD per ether
	var EXCHANGE_RATE_CAP = big.NewInt(1200)
	// 11 ETHUSD
	var EXPECTED_AMOUNT = big.NewInt(1100)

	session := deploy("TestEthUSD", "ETHUSD")
	session.SetExchangeRate(EXCHANGE_RATE)
	session.SetExchangeRateCap(EXCHANGE_RATE_CAP)

	_, err := session.Contract.TestEthUSDTransactor.Purchase(
		paidTransactOpts(AMOUNT_TO_PURCHASE, senderKey),
	)

	checkErr(err)

	backend.Commit()

	balance, _ := session.BalanceOf(sender)
	assert.Equal(t, EXPECTED_AMOUNT, balance)
}

func TestPurchaseWhenExchangeRateGreaterThanCap(t *testing.T) {
	// 1 Ether
	var AMOUNT_TO_PURCHASE = NewValue(1).toWei()
	// $13 USD per ether
	var EXCHANGE_RATE = big.NewInt(1300)
	// $12 USD per ether
	var EXCHANGE_RATE_CAP = big.NewInt(1200)
	// 11 ETHUSD
	var EXPECTED_AMOUNT = big.NewInt(1200)

	session := deploy("TestEthUSD", "ETHUSD")
	session.SetExchangeRate(EXCHANGE_RATE)
	session.SetExchangeRateCap(EXCHANGE_RATE_CAP)

	_, err := session.Contract.TestEthUSDTransactor.Purchase(
		paidTransactOpts(AMOUNT_TO_PURCHASE, senderKey),
	)

	checkErr(err)

	backend.Commit()

	balance, _ := session.BalanceOf(sender)
	assert.Equal(t, EXPECTED_AMOUNT, balance)
}

func TestWithdraw(t *testing.T) {
	session := deploy("EthUSD", "ETHUSD")
	session.SetBalance(sender, NewValue(2).toWei())
	var COST_OF_TRANSACTION = big.NewInt(1000000000000072824)
	startingBalance, _ := backend.BalanceAt(nil, sender, nil)

	_, err := session.Withdraw(NewValue(1).toWei())

	checkErr(err)

	backend.Commit()

	ethUSDBalance, _ := session.BalanceOf(sender)
	assert.Equal(t, NewValue(1).toWei(), ethUSDBalance)

	weiBalance, _ := backend.BalanceAt(nil, sender, nil)

	expectedWeiBalance := startingBalance
	expectedWeiBalance.Sub(expectedWeiBalance, COST_OF_TRANSACTION)
	expectedWeiBalance.Add(expectedWeiBalance, NewValue(1).toWei())
	assert.Equal(t, expectedWeiBalance, weiBalance)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("%v", err)
	}
}
