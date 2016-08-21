package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// Generate a new random account and a funded simulator
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)

	sim := backends.NewSimulatedBackend(core.GenesisAccount{Address: auth.From, Balance: big.NewInt(10000000000)})

	// Deploy a token contract on the simulated blockchain
	_, _, token, err := DeployToken(auth, sim, new(big.Int), "My Token", 0, "My Token")
	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}

	sim.Commit()

	name, _ := token.Name(nil)
	fmt.Println(name)
}
