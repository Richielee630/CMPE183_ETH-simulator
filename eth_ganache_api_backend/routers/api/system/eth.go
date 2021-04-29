package system

import (
	"api_backend/models"
	"api_backend/pkg/contract/store"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"strconv"
)

// 生成以太坊地址
func NewAddress() (string,string){
	privateKey,err := crypto.GenerateKey()  // S256 returns an instance of the secp256k1 elliptic.Curve.
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes :=crypto.FromECDSA(privateKey)  // FromECDSA exports a private key into a binary dump.
	key := hexutil.Encode(privateKeyBytes)[2:]
	//私钥生成公钥
	publicKey := privateKey.Public()
	//转换十六进制 剥离0x和前两个字符04，它始终是EC前缀，不是必需的
	publicKeyECDSA, ok:=publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	addr := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return addr,key
}

// 自动充值
func RechargeToMe(addr string) (string,error) {
	privateKey, err := crypto.HexToECDSA(models.RechargeKey)
	if err != nil {
		return "",err
	}
	fromAddress := common.HexToAddress(models.RechargeAddr)
	nonce, err := models.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "",err
	}

	valuef, err := strconv.ParseFloat(models.RechargeAmount, 64)
	if err != nil {
		log.Println("is not a number")
		return "",err
	}
	valueWei, isOk := new(big.Int).SetString(fmt.Sprintf("%.0f", valuef*1000000000000000000), 10)
	if !isOk {
		log.Println("float to bigInt failed!")
		return "",err
	}

	gasLimit := uint64(21000) // in units
	gasPrice, err := models.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return "",err
	}

	chainID := big.NewInt(7777)

	toAddress := common.HexToAddress(addr)
	////4. 构建未签名交易体
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, valueWei, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "",err
	}
	err = models.Client.SendTransaction(context.Background(),signedTx)

	return  signedTx.Hash().Hex(),err
}

func GetAddrBalance(addr string) (string,error) {
	account := common.HexToAddress(addr)
	balance, err := models.Client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return  "",err
	}
	return balance.String(),nil
}

type Tx struct {
	Hash  string `json:"hash"`
	From  string `json:"from"`
	To    string `json:"to"`
	Value string `json:"value"`
	BlockHeight uint64 `json:"block_height"`
	Timestamp uint64 `json:"timestamp"`
	Nonce uint64 `json:"nonce"`
	GasPrice string `json:"gas_price"`
	Gas     uint64 `json:"gas"`

}

func GetTxByHash(hash string)(interface{},error) {
	var  tx  Tx
	res,_,err := models.Client.TransactionByHash(context.Background(),common.HexToHash(hash))
	if err != nil {
		return  nil,err
	}

	reciept,err := models.Client.TransactionReceipt(context.Background(),common.HexToHash(hash))
	if err != nil {
		return  nil,err
	}

	msg,err := res.AsMessage(types.NewEIP155Signer(models.ChainID))
	if err != nil {
		return  nil,err
	}
	blo,err := models.Client.BlockByNumber(context.Background(),reciept.BlockNumber)
	if err != nil {
		return  nil,err
	}
	value,_:=stringTo(msg.Value().String())

	tx = Tx{
		Hash:        hash,
		From:        msg.From().Hex(),
		To:          msg.To().Hex(),
		Value:       value,
		BlockHeight: reciept.BlockNumber.Uint64(),
		Timestamp:   blo.Time(),
		Nonce:       res.Nonce(),
		GasPrice:    res.GasPrice().String(),
		Gas:         res.Gas(),
	}
	return tx,err
}

// 用户转化单位
func stringTo(value string)(string,error){

	valuerf ,err := strconv.ParseFloat(value, 64)
	if err != nil {
		return "",err
	}
	// 保留两位小数
	r := fmt.Sprintf("%.2f", valuerf/1000000000000000000)
	return  r,nil
}

func Trans(from,fromkey,to,value string ) (string,error){

	privateKey, err := crypto.HexToECDSA(fromkey)
	if err != nil {
		return "",err
	}
	fromAddress := common.HexToAddress(from)
	nonce, err := models.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "",err
	}

	valuef, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Println("is not a number")
		return "",err
	}
	valueWei, isOk := new(big.Int).SetString(fmt.Sprintf("%.0f", valuef*1000000000000000000), 10)
	if !isOk {
		log.Println("float to bigInt failed!")
		return "",err
	}

	gasLimit := uint64(21000) // in units
	gasPrice, err := models.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return "",err
	}

	chainID := big.NewInt(7777)

	toAddress := common.HexToAddress(to)
	////4. 构建未签名交易体
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, valueWei, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "",err
	}
	err = models.Client.SendTransaction(context.Background(),signedTx)

	return  signedTx.Hash().Hex(),err

}

func Onchain(typa int, msg string)  (string, error) {
	tx, err := sendStoreContractTrx(typa, msg)
	if err != nil {
		return "", err
	}
	return tx, nil
}

func sendStoreContractTrx(typa int, msg string) (string, error) {

	addr := common.HexToAddress(models.ContractAddr)
	instance, err := Store.NewStore(addr, models.Client)
	if err != nil {
		return "", err
	}

	//1. 由密码导入私钥 详情见 account_key.go
	privateKey, err := crypto.HexToECDSA(models.RechargeKey)
	if err != nil {
		return "", err
	}

	account := common.HexToAddress(models.RechargeAddr)
	//2. 找到随机数
	nonce, err := models.Client.PendingNonceAt(context.Background(), account)
	if err != nil {
		return "", err
	}

	gasLimit := uint64(4700000) // in units
	gasPrice, err := models.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice
	//调用接口
	tx, err := instance.SetItem(auth, big.NewInt(int64(typa)), msg)
	if err != nil {
		return "", err
	}
	log.Println("本次交易hash:", tx.Hash().String())

	return tx.Hash().Hex(), nil
}

func GetChain(typa int) (string,error){
	res, err := getMsg(int64(typa))
	return res, err
}

func getMsg( typa int64) (string, error) {


	var result string

	addr := common.HexToAddress(models.ContractAddr)
	instance, err := Store.NewStore(addr, models.Client)
	if err != nil {
		return "", err
	}

	result, err = instance.NonceToTrxs(nil, big.NewInt(typa))
	if err != nil {
		return "", err
	}
	return result, nil
}
