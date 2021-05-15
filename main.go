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
		Usage:                "The Keeper in Giant Panda Zoo To Perform Routine Tasks",
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
			&cli.IntFlag{
				Name:  "gasLimit",
				Value: 5000000,
				Usage: "gas limit for updating",
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
			contractAddress := common.HexToAddress(c.String("contract"))
			gasLimit := c.Int("gasLimit")
			provider := c.String("provider")
			log.Println("Contract:", contractAddress)
			log.Println("Provider:", provider)
			log.Println("GasLimit:", gasLimit)

			tryUpdate(provider, priv, contractAddress, gasLimit, common.Big0)
			ticker := time.NewTicker(1 * time.Minute)

			lastUpdateTime := common.Big0

			for {
				select {
				case <-ticker.C:
					// fail-retry
					succ, updateTime := tryUpdate(provider, priv, contractAddress, gasLimit, lastUpdateTime)
					if succ {
						lastUpdateTime = updateTime
					}
				}
			}
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func tryUpdate(provider string, key *ecdsa.PrivateKey, address common.Address, gasLimit int, lastUpdateTime *big.Int) (bool, *big.Int) {
	// create connection
	client, err := ethclient.Dial(provider)
	if err != nil {
		log.Printf("PandaKeeper: connection to  %s failed, reason: %s", provider, err)
		return false, common.Big0
	}
	defer client.Close()

	instance, err := NewAggregateUpdater(address, client)
	if err != nil {
		log.Println("PandaKeeper: NewAggregateUpdater failed:", err)
		return false, common.Big0
	}

	// query next update time
	updateTime, err := instance.GetNextUpdateTime(nil)
	if err != nil {
		log.Println("PandaKeeper: GetNextUpdateTime() failed:", err)
		return false, common.Big0
	}

	log.Printf("PandaKeeper: Next Update:%s", time.Unix(updateTime.Int64(), 0))

	// still not expired
	if time.Now().Unix() < updateTime.Int64() {
		return false, common.Big0
	}

	// if updateTime is within 1 minute without change
	// assume the transaction has not yet confirmed
	if lastUpdateTime.Cmp(updateTime) == 0 && time.Since(time.Unix(updateTime.Int64(), 0)) < time.Minute {
		log.Println("PandaKeeper: still not confirmed within 1 minute:", lastUpdateTime, updateTime)
		return false, common.Big0
	}

	// query gas price & nonce
	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(key.PublicKey))
	if err != nil {
		log.Println("PandaKeeper: client.PendingNonceAt() failed:", err)
		return false, common.Big0
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println("PandaKeeper: client.SuggestGasPrice() failed:", err)
		return false, common.Big0
	}

	// create transactor
	auth := bind.NewKeyedTransactor(key)
	auth.GasLimit = uint64(gasLimit)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice

	tx, err := instance.Update(auth)
	if err != nil {
		log.Println("PandaKeeper: update transaction failed:", err)
		return false, common.Big0
	}

	log.Println("PandaKeeper: update transaction sent:", tx.Hash().Hex())
	return true, updateTime
}
