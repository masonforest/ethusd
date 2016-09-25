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
	key0             *ecdsa.PrivateKey
	key1             *ecdsa.PrivateKey
	addr0            common.Address
	addr1            common.Address
	auth             *bind.TransactOpts
	backend          *backends.SimulatedBackend
	NONCE            = big.NewInt(2)
	STARTING_BALANCE = big.NewInt(10000000000000000)
)

func Signer(key *ecdsa.PrivateKey) bind.SignerFn {
	return func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		signature, err := crypto.Sign(tx.SigHash().Bytes(), key0)
		if err != nil {
			return nil, err
		}
		return tx.WithSignature(signature)
	}
}

func TestMain(m *testing.M) {
	key0, _ = crypto.GenerateKey()
	key1, _ = crypto.GenerateKey()
	addr0 = crypto.PubkeyToAddress(key0.PublicKey)
	addr1 = crypto.PubkeyToAddress(key1.PublicKey)
	auth = bind.NewKeyedTransactor(key0)
	backend = backends.NewSimulatedBackend(
		core.GenesisAccount{
			Address: addr0,
			Balance: STARTING_BALANCE,
		},
		core.GenesisAccount{
			Address: addr1,
			Balance: STARTING_BALANCE,
		})
	flag.Parse()
	os.Exit(m.Run())
}

func deploy(name string) *TestEthUSDSession {
	_, _, token, err := DeployTestEthUSD(
		bind.NewKeyedTransactor(key0),
		backend,
		name,
	)

	checkErr(err)

	backend.Commit()

	return &TestEthUSDSession{
		Contract:     token,
		TransactOpts: *bind.NewKeyedTransactor(key0),
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
		paidTransactOpts(big.NewInt(1), key0),
	)

	checkErr(err)

	backend.Commit()

	balance, _ := session.BalanceOf(addr0)
	assert.Equal(t, big.NewInt(2), balance)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("%v", err)
	}
}
