package action

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"github.com/tendermint/tmlibs/common"
	conf2 "github/bsc-task/cli/conf"
	"github/bsc-task/csv"
	"github/bsc-task/log"
	"github/bsc-task/rpc"
	"os"
	"path"
	"strconv"
	"time"
)

type Tx struct {
	FromAddress string `json:"from_address"`
	ToAddress   string `json:"to_address"`
	Amount      string `json:"amount"`
	//IsCollect       int    `json:"is_collect"`
	//Token           string `json:"token"`
	ContractAddress string `json:"contract_address"`
	Fee             int64  `json:"fee"`
}

func NewTx(ContractAddress string, fromAddress string, amount string, fee int64) Tx {
	return Tx{
		FromAddress:     fromAddress,
		Amount:          amount,
		ContractAddress: ContractAddress,
		Fee:             fee,
	}
}

func SendTransaction() {
	csvFilePath := conf2.Config.Transaction.ToAddrFileName
	addrs, err := csv.Read(csvFilePath, 0)
	if err != nil {
		log.Logger.Errorf("无法读取csv文件. 异常信息=%s", err.Error())
	}
	//生成一个批次号
	batchNo := getFileNameWithoutExt(csvFilePath) + "-" + strconv.Itoa(int(time.Now().Unix()))
	execSend(addrs, batchNo)
}

func execSend(toAddress []string, batchNo string) {
	fromAddress := conf2.Config.Transaction.FromAddress
	ContractAddress := conf2.Config.Transaction.ContractAddress
	amount := conf2.Config.Transaction.Amount
	fee := conf2.Config.Transaction.Fee
	if fromAddress == "" {
		log.Logger.Error("配置文件：fromAddress不能为空")
		return
	}
	if amount == "" {
		log.Logger.Error("配置文件：amount不能为空")
		return
	}

	tx := NewTx(ContractAddress, fromAddress, amount, fee)
	log.Logger.WithFields(logrus.Fields{"ContractAddress": ContractAddress, "toAddressCount": len(toAddress), "batchNo": batchNo, "fromAddress": fromAddress, "amount": amount, "fee": fee}).Info("SendTransaction")

	for i, addr := range toAddress {
		tx.ToAddress = addr
		clientRPC := rpc.NewRPC(conf2.Config.BscSignUrl, conf2.Config.BscSignUser, conf2.Config.BscSignPassword)
		response, err := clientRPC.SendRequest("/v1/bsc/transfer", tx)
		if err != nil {
			log.Logger.WithFields(logrus.Fields{"toAddress": addr}).Errorf("调用API出现错误：%s", err.Error())
			continue
		}
		err = responseToLog(getOutputResultFileName(batchNo), addr, response)
		if err != nil {
			log.Logger.WithFields(logrus.Fields{"toAddress": addr}).Errorf("执行结果输出到文件错误：%s", err.Error())
			continue
		}
		log.Logger.WithFields(logrus.Fields{"toAddress": addr, "index": i}).Info("已执行")
	}
	log.Logger.WithFields(logrus.Fields{"batchNo": batchNo}).Info("=======================执行完毕=======================")
}

func responseToLog(outPutFileName string, toAddr string, content []byte) error {
	exists := common.FileExists(outPutFileName)
	if !exists {
		_, err := os.Create(outPutFileName)
		if err != nil {
			return err
		}
	}
	f, err := os.OpenFile(outPutFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.Write(bytes.Join([][]byte{
		[]byte("收款地址 [\""),
		[]byte(toAddr),
		[]byte("\"] 执行结果: "),
		content,
		[]byte("\r\n"),
	}, []byte{}))
	return err
}

func getOutputResultFileName(batchNo string) string {
	return "./" + batchNo + ".log"
}

func getFileNameWithoutExt(fileName string) string {
	fileNameBase := path.Base(fileName)
	fileSuffix := path.Ext(fileName)
	return fileNameBase[0 : len(fileNameBase)-len(fileSuffix)]
}
