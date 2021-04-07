package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github/bsc-task/log"
	"io/ioutil"
	"net/http"
)

type Client struct {
	rpcUrl      string
	rpcUser     string
	rpcPassword string
}

type TxRes struct {
	Code    int    `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

//初始化一个rpc客户端
func NewRPC(url, user, password string) *Client {
	return &Client{
		rpcUrl:      url,
		rpcUser:     user,
		rpcPassword: password,
	}
}

func (rpc *Client) SendRequest(api string, params interface{}) ([]byte, error) {
	var (
		reqBytes []byte
		err      error
	)
	reqBytes, err = json.Marshal(params)
	if err != nil {
		return nil, err
	}
	reqBuf := bytes.NewBuffer(reqBytes)
	var (
		req *http.Request
	)
	log.Logger.WithFields(logrus.Fields{"API": rpc.rpcUrl + api, "params": string(reqBytes)}).Info("sendRequest")
	if req, err = http.NewRequest(http.MethodPost, rpc.rpcUrl+api, reqBuf); err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	//设置rpc的用户和密码
	//如果为空就不设置
	if rpc.rpcUser != "" && rpc.rpcPassword != "" {
		req.SetBasicAuth(rpc.rpcUser, rpc.rpcPassword)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	//解析resp
	var response TxRes
	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, errors.New(fmt.Sprintf("Parse resp error,Err=【%v】", err))
	}
	data, err := json.Marshal(response)
	return data, nil
}
