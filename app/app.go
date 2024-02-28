package app

import (
	"blockchain-core/config"
	"blockchain-core/global"
	"blockchain-core/repository"
	"blockchain-core/service"
	. "blockchain-core/types"
	"errors"
	"strings"

	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/inconshreveable/log15"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	config     *config.Config
	service    *service.Service
	repository *repository.Repository

	log log15.Logger
}

func NewApp(config *config.Config) {
	a := &App{
		config: config,
		log:    log15.New("module", "app"),
	}

	var err error

	if a.repository, err = repository.NewRepository(config); err != nil {
		panic(err)
	}

	a.service = service.NewService(a.repository, 1)
	a.log.Info("Module started", "time", time.Now().Unix())

	sc := bufio.NewScanner(os.Stdin)
	useCase()

	for {
		from := global.FROM()

		if from != "" {
			a.log.Info("Current Connected Wallet", "from", from)
			fmt.Println()
		}

		sc.Scan()
		input := strings.Split(sc.Text(), " ")
		if err = a.inputValueAssessment(input); err != nil {
			a.log.Error("Failed to call cli", "err", err, "input", input)
			fmt.Println()
		}
	}

}

func (a *App) inputValueAssessment(input []string) error {
	msg := errors.New("check Use Case")

	if len(input) == 0 {
		return msg
	} else {
		from := global.FROM()

		switch input[0] {
		case CreateWallet:
			fmt.Println("CreateWallet ----------------")
			if wallet := a.service.MakeWallet(); wallet == nil {
				panic("failed to create")
			} else {
				a.log.Info("Success To Create Wallet", "pk", wallet.PrivateKey, "pu", wallet.PublicKey)
			}

		case ConnectWallet:

			if from != "" {
				a.log.Debug("Already Connected", "from", from)
				fmt.Println()
			} else {
				if wallet, err := a.service.GetWallet(input[1]); err != nil {
					if err == mongo.ErrNoDocuments {
						a.log.Debug("Failed To Find Wallet pk is Nil", "pk", input[1])
					} else {
						a.log.Crit("Failed To Find Wallet", "pk", input[1], "err", err)
					}
				} else {
					global.SetFrom(wallet.PublicKey)
					fmt.Println()
					a.log.Info("Success To Connect Wallet", "from", wallet.PublicKey)
					fmt.Println()
				}
			}

		case ChangeWallet:
			if from == "" {
				a.log.Debug("Connect Wallet First")
			} else {
				if wallet, err := a.service.GetWallet(input[1]); err != nil {
					if err == mongo.ErrNoDocuments {
						a.log.Debug("Failed To Find Wallet pk is Nil", "pk", input[1])
					} else {
						a.log.Crit("Failed To Find Wallet", "pk", input[1], "err", err)
					}
				} else {
					global.SetFrom(wallet.PublicKey)
					fmt.Println()
					a.log.Info("Success To Connect Wallet", "from", wallet.PublicKey)
				}
			}

		case TransferCoin:
			fmt.Println("TransferCoin in Switch")
		case MintCoin:
			fmt.Println("MintCoin in Switch")
		default:
			return msg
		}
	}

	return nil
}

func useCase() {
	fmt.Println()
	fmt.Println("This is Akaps Module For BlockChain Core With Mongo")
	fmt.Println()
	fmt.Println("Use Case")

	fmt.Println("1. ", CreateWallet)
	fmt.Println("2. ", ConnectWallet, " <PK>")
	fmt.Println("3. ", ChangeWallet, " <PK>")
	fmt.Println("4. ", TransferCoin, " <To> <Amount>")
	fmt.Println("5. ", MintCoin, "<To> <Amount>")
	fmt.Println()
}
