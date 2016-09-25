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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

var (
	senderKey        *ecdsa.PrivateKey
	sender           common.Address
	auth             *bind.TransactOpts
	backend          *backends.SimulatedBackend
	NONCE            = big.NewInt(2)
	STARTING_BALANCE = big.NewInt(10000000000000000)
)

func Signer(key *ecdsa.PrivateKey) bind.SignerFn {
	return func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		signature, err := crypto.Sign(tx.SigHash().Bytes(), senderKey)
		if err != nil {
			return nil, err
		}
		return tx.WithSignature(signature)
	}
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

func deploy(name string) *TestEthUSDSession {
	_, _, token, err := DeployTestEthUSD(
		bind.NewKeyedTransactor(senderKey),
		backend,
		name,
	)

	checkErr(err)

	backend.Commit()

	return &TestEthUSDSession{
		Contract:     token,
		TransactOpts: *bind.NewKeyedTransactor(senderKey),
	}
}

func TestInitialize(t *testing.T) {
	token := deploy("TestEthUSD")

	name, _ := token.Name()
	assert.Equal(t, "TestEthUSD", name)
}

func paidTransactOpts(value *big.Int, key *ecdsa.PrivateKey) *bind.TransactOpts {
	t := bind.NewKeyedTransactor(key)
	t.Value = value
	return t
}

func TestPurchase(t *testing.T) {
	session := deploy("TestEthUSD")
	session.SetExchangeRate(big.NewInt(2))

	_, err := session.Contract.TestEthUSDTransactor.Purchase(
		paidTransactOpts(big.NewInt(1), senderKey),
	)

	checkErr(err)

	backend.Commit()

	balance, _ := session.BalanceOf(sender)
	assert.Equal(t, big.NewInt(2), balance)
}

func TestWithdraw(t *testing.T) {
	session := deploy("TestEthUSD")
	session.SetBalance(sender, big.NewInt(2))
	var COST_OF_TRANSACTION = big.NewInt(1640536)

	_, err := session.Withdraw(big.NewInt(1))

	checkErr(err)

	backend.Commit()

	ethUSDBalance, _ := session.BalanceOf(sender)
	assert.Equal(t, big.NewInt(1), ethUSDBalance)

	weiBalance, _ := backend.BalanceAt(nil, sender, nil)

	expectedWeiBalance := STARTING_BALANCE
	expectedWeiBalance.Sub(expectedWeiBalance, COST_OF_TRANSACTION)
	expectedWeiBalance.Add(expectedWeiBalance, big.NewInt(1))
	assert.Equal(t, expectedWeiBalance, weiBalance)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("%v", err)
	}
}
