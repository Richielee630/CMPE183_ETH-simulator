package models

import (
	"api_backend/pkg/contract/store"
	"api_backend/pkg/setting"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

var (
	nodeUrl                                 string
	MaxPoolSize                             uint64
	ChainID                                 *big.Int
	Client                                  *ethclient.Client
	RechargeAddr,RechargeKey,RechargeAmount,ContractAddr string
)

type mgo struct {
	database   string
	collection string
}

//初始化
func Init() {
	secEth, err := setting.Cfg.GetSection("wsnode")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	nodeUrl = secEth.Key("NODE_URL").String()
	RechargeAddr =  secEth.Key("RECHARGEADDR").String()
	RechargeKey  =  secEth.Key("RECHARGEKEY").String()
	RechargeAmount = secEth.Key("AMOUNT").String()

	Client, err = ethclient.Dial(nodeUrl)
	if err != nil {
		log.Fatal(err)
	}
	ChainID = big.NewInt(7777)

	Deploy()
}

func Deploy(){
	privateKey, err := crypto.HexToECDSA(RechargeKey)
	if err != nil {
		log.Fatal(err)
	}
	fromAddress := common.HexToAddress(RechargeAddr)
	nonce, err := Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(300000) // in
	auth.GasPrice = gasPrice

	address, _, _, err := Store.DeployStore(auth, Client)
	if err != nil {
		log.Fatal(err)
	}
	ContractAddr = address.Hex()
	fmt.Println("合约地址为：",ContractAddr)
	if err != nil {
		panic(err)
	}
}