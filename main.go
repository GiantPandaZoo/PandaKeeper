package main

import (
	"context"
	"crypto/ecdsa"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "PandaKeeper",
		Usage:                "./PandaKeeper -h",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "key",
				Value: "./privkey",
				Usage: "raw private key file in hex",
			},
			&cli.StringFlag{
				Name:  "contract",
				Value: "0x20992e494c88e2B08d93944a2d610c441a7c4Aa1",
				Usage: "contract address",
			},
			&cli.StringFlag{
				Name:  "provider",
				Value: "https://rinkeby.infura.io/v3/f081cccbb2744415b20add374caf68c9",
				Usage: "RPC service address",
			},
		},
		Action: func(c *cli.Context) error {
			content, err := ioutil.ReadFile(c.String("key"))
			if err != nil {
				log.Fatal(err)
			}

			priv, err := crypto.HexToECDSA(strings.TrimSuffix(string(content), "\n"))
			if err != nil {
				log.Fatal("cannot decode private key", err)
			}

			publicKey := priv.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				log.Fatal("error casting public key to ECDSA")
			}

			log.Printf("Using account %s to update", crypto.PubkeyToAddress(*publicKeyECDSA))
			log.Println("Contract:", c.String("contract"))
			log.Println("Provider:", c.String("provider"))

			contractAddress := common.HexToAddress(c.String("contract"))

			tryUpdate(c.String("provider"), priv, contractAddress)
			ticker := time.NewTicker(1 * time.Minute)
			for {
				select {
				case <-ticker.C:
					tryUpdate(c.String("provider"), priv, contractAddress)
				}
			}
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func tryUpdate(provider string, key *ecdsa.PrivateKey, address common.Address) {
	// create connection
	client, err := ethclient.Dial(provider)
	if err != nil {
		log.Printf("PandaKeeper: connection to  %s failed, reason: %s", provider, err)
		return
	}
	defer client.Close()

	instance, err := NewAggregateUpdater(address, client)
	if err != nil {
		log.Println("PandaKeeper: NewAggregateUpdater failed:", err)
		return
	}

	// query next update time
	updateTime, err := instance.GetNextUpdateTime(nil)
	if err != nil {
		log.Println("PandaKeeper: GetNextUpdateTime() failed:", err)
		return
	}

	log.Printf("PandaKeeper: Next Update:%s", time.Unix(updateTime.Int64(), 0))

	// still not expired
	if time.Now().Unix() < updateTime.Int64() {
		return
	}

	// query gas price & nonce
	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(key.PublicKey))
	if err != nil {
		log.Println("PandaKeeper: client.PendingNonceAt() failed:", err)
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println("PandaKeeper: client.SuggestGasPrice() failed:", err)
		return
	}

	// create transactor
	auth := bind.NewKeyedTransactor(key)
	auth.GasLimit = uint64(300000)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice

	tx, err := instance.Update(auth)
	if err != nil {
		log.Println("PandaKeeper: update transaction failed:", err)
		return
	}

	log.Println("PandaKeeper: update transaction sent:", tx.Hash().Hex())
}
