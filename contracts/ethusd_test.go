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

var (
  key0 *ecdsa.PrivateKey
  key1 *ecdsa.PrivateKey
  addr0 common.Address
  addr1 common.Address
  auth *bind.TransactOpts
  sim *backends.SimulatedBackend
  GAS_LIMIT = big.NewInt(500000)
  NONCE = big.NewInt(2)
)

func TestMain(m *testing.M) {
	key0, _ = crypto.GenerateKey()
	key1, _ = crypto.GenerateKey()
  addr0 = crypto.PubkeyToAddress(key0.PublicKey)
  addr1 = crypto.PubkeyToAddress(key1.PublicKey)
	auth = bind.NewKeyedTransactor(key0)
	sim = backends.NewSimulatedBackend(
		core.GenesisAccount{
			Address: auth.From,
			Balance: big.NewInt(9223372036854775807),
		},
    core.GenesisAccount{
			Address: auth.From,
			Balance: big.NewInt(9223372036854775807),
		})
	flag.Parse()
	os.Exit(m.Run())
}

func Signer(key *ecdsa.PrivateKey) bind.SignerFn {
	return func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		signature, err := crypto.Sign(tx.SigHash().Bytes(), key0)
		if err != nil {
			return nil, err
		}
		return tx.WithSignature(signature)
	}
}

func deploy(initialSupply *big.Int, name string) (*EthUSD) {
	_, _, ethUSD, err := DeployEthUSD(
		auth,
		sim,
		initialSupply,
		name,
	)

	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}

	sim.Commit()

  return ethUSD
}

func TestInitializer(t *testing.T) {
  ethUSD := deploy(big.NewInt(100), "EthUSD")

	name, _ := ethUSD.Name(nil)
	assert.Equal(t, "EthUSD", name)
	balance, _ := ethUSD.BalanceOf(nil, auth.From)
	assert.Equal(t, big.NewInt(100), balance)
}

func TestTransfer(t *testing.T) {
  ethUSD := deploy(big.NewInt(3), "EthUSD")

	_, err := ethUSD.Transfer(&bind.TransactOpts{
		Nonce:    NONCE,
		GasLimit: GAS_LIMIT,
		Signer:   Signer(key0),
	},
		addr1,
		big.NewInt(1),
	)

	if err != nil {
		panic(err)
	}

	sim.Commit()

	senderBalance, _ := ethUSD.BalanceOf(nil, addr0)
	recipientBalance, _ := ethUSD.BalanceOf(nil, addr1)
	assert.Equal(t, big.NewInt(2), senderBalance)
	assert.Equal(t, big.NewInt(1), recipientBalance)
}
