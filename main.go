package main

import (
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
	"github.com/ethereum/go-ethereum/rpc"
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
			key, ok := new(big.Int).SetString(strings.TrimSuffix(string(content), "\n"), 16)
			if !ok {
				log.Fatal("cannot decode private key", string(content))
			}

			priv := new(ecdsa.PrivateKey)
			priv.PublicKey.Curve = crypto.S256()
			priv.D = key
			priv.PublicKey.X, priv.PublicKey.Y = crypto.S256().ScalarBaseMult(priv.D.Bytes())

			log.Printf("Using account %s to update", crypto.PubkeyToAddress(priv.PublicKey))
			log.Println("Contract:", c.String("contract"))
			log.Println("Provider:", c.String("provider"))

			contractAddress := common.HexToAddress(c.String("contract"))

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
	rpcDial, err := rpc.Dial(provider)
	if err != nil {
		log.Printf("PandaKeeper: connection to  %s failed, reason: %s", provider, err)
		return
	}
	defer rpcDial.Close()

	ethClient := ethclient.NewClient(rpcDial)
	defer ethClient.Close()

	// create caller
	caller, err := NewAggregateUpdaterCaller(address, ethClient)
	if err != nil {
		log.Println("PandaKeeper: NewPoolCaller failed:", err)
		return
	}

	// query next update time
	updateTime, err := caller.GetNextUpdateTime(nil)
	if err != nil {
		log.Println("PandaKeeper: GetNextUpdateTime() failed:", err)
		return
	}

	// still not expired
	if time.Now().Unix() < updateTime.Int64() {
		return
	}

	// create trasactor to update
	poolTransactor, err := NewAggregateUpdaterTransactor(address, ethClient)
	if err != nil {
		log.Println("PandaKeeper: NewAggregateUpdaterTransactor failed:", err)
		return
	}

	auth := bind.NewKeyedTransactor(key)
	tx, err := poolTransactor.Update(auth)
	if err != nil {
		log.Println("PandaKeeper: update transaction failed:", err)
		return
	}

	log.Println("PandaKeeper: update transaction sent:", tx.Hash().String())
}
