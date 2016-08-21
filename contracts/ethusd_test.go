//go:generate abigen --sol ethusd.sol --pkg main --out ethusd.go

package main

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

var key *ecdsa.PrivateKey
var auth *bind.TransactOpts
var sim *backends.SimulatedBackend
var GAS_LIMIT = big.NewInt(500000)
var NONCE = big.NewInt(2)

func TestMain(m *testing.M) {
	key, _ = crypto.GenerateKey()
	auth = bind.NewKeyedTransactor(key)
	sim = backends.NewSimulatedBackend(
		core.GenesisAccount{
			Address: auth.From,
			Balance: big.NewInt(9223372036854775807),
		})
	flag.Parse()
	os.Exit(m.Run())
}

func Signer(key *ecdsa.PrivateKey) bind.SignerFn {
	return func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		signature, err := crypto.Sign(tx.SigHash().Bytes(), key)
		if err != nil {
			return nil, err
		}
		return tx.WithSignature(signature)
	}
}

func TestInitializer(t *testing.T) {
	_, _, ethUSD, err := DeployEthUSD(
		auth,
		sim,
		big.NewInt(100),
		"EthUSD",
	)

	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}

	sim.Commit()

	name, _ := ethUSD.Name(nil)
	assert.Equal(t, "EthUSD", name)
	balance, _ := ethUSD.BalanceOf(nil, auth.From)
	assert.Equal(t, big.NewInt(100), balance)
}

func TestTransfer(t *testing.T) {
	_, _, ethUSD, err := DeployEthUSD(
		auth,
		sim,
		big.NewInt(3),
		"EthUSD",
	)

	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}

	sim.Commit()
	recipientKey, _ := crypto.GenerateKey()
	recipient := bind.NewKeyedTransactor(recipientKey).From

	_, err = ethUSD.Transfer(&bind.TransactOpts{
		Nonce:    NONCE,
		GasLimit: GAS_LIMIT,
		Signer:   Signer(key),
	},
		recipient,
		big.NewInt(1),
	)

	if err != nil {
		panic(err)
	}

	sim.Commit()

	senderBalance, _ := ethUSD.BalanceOf(nil, auth.From)
	recipientBalance, _ := ethUSD.BalanceOf(nil, recipient)
	assert.Equal(t, big.NewInt(2), senderBalance)
	assert.Equal(t, big.NewInt(1), recipientBalance)
}
