package main

// import (
// 	// "fmt"
// 	"github.com/ethereum/go-ethereum/accounts"
// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	// "github.com/ethereum/go-ethereum/crypto"
// 	"github.com/ethereum/go-ethereum/ethclient"
//
// 	// "github.com/stretchr/testify/assert"
// 	"math/big"
// 	"testing"
// 	// "time"
// )

// const testnetKey = `{"address":"94c7a6bc8cdbfbbf9d2b7b12ce3a9c82f41fd74f","privatekey":"2c21018eaf82bfeb57699dc4eee78a714fcd106facdab39964cfc5c81d9620d6","id":"90af7341-e255-4364-b00c-e16b7cfdb095","version":3}`

// func TestSetExchangeRate(t *testing.T) {
// 	key := &accounts.Key{}
// 	key.UnmarshalJSON([]byte(testnetKey))
// 	backend, err := ethclient.Dial("http://localhost:8545")
// 	t.Log("---")
// 	t.Logf("%+v\n", bind.NewKeyedTransactor(key.PrivateKey))
// 	_, tx, token, err := DeployEthUSD(
// 		bind.NewKeyedTransactor(key.PrivateKey),
// 		backend,
// 	)
//
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	t.Logf("tx1:%x", tx.Hash())
// 	t.Logf("token:%v", token)
// 	t.Log("---")
// 	// time.Sleep(20000 * time.Millisecond)
//
// 	opts := bind.TransactOpts{
// 		From:     key.Address,
// 		Signer:   Signer(key.PrivateKey),
// 		Value:    big.NewInt(100000000000000000),
// 		GasLimit: big.NewInt(200000),
// 	}
//
// 	session := EthUSDSession{
// 		Contract:     token,
// 		TransactOpts: opts,
// 	}
//
// 	tx, err = session.Update(big.NewInt(0))
//
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	t.Logf("tx2:%x", tx.Hash())
// }
