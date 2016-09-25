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

	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}

	backend.Commit()

	return &TestEthUSDSession{
		Contract:     token,
		TransactOpts: *bind.NewKeyedTransactor(key0),
	}

}

func TestBuy(t *testing.T) {
	token := deploy("TestEthUSD")

	log.Printf("%+v", token)
	// _, err := token.Buy()
	//
	// if err != nil {
	// 	panic(err)
	// }

	backend.Commit()

	name, _ := token.Name()
	backend.Commit()
	assert.Equal(t, "TestEthUSD", name)
}
