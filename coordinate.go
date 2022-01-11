package main

import (
	"encoding/json"
	"github.com/xuperchain/contract-sdk-go/code"
	"github.com/xuperchain/contract-sdk-go/driver"
)

type ordinate struct {}

const (
	LAWENFORCEMENT = "LawEnforcement"
	BCCERTIFICATE = "BlockChainCertificate"
)

type lawEnforcement struct {
	Uuid			string `json:"uuid"`
	Domain			string `json:"domain"`
	Uploader		string `json:"uploader"`
	UploadTime		string `json:"upload_time"`
	HashAbstract	string `json:"hash_abstract"`
}

func (o *ordinate) Initialize(ctx code.Context) code.Response {
	//初始化执法结果字典
	if _, err := ctx.GetObject([]byte(LAWENFORCEMENT)); err != nil {
		lawMap := map[string][]byte{}
		lawMapByte, _ := json.Marshal(lawMap)
		if err := ctx.PutObject([]byte(LAWENFORCEMENT), lawMapByte); err != nil {
			return code.Error(err)
		}
	}
	//初始化区块链证书字典
	if _, err := ctx.GetObject([]byte(BCCERTIFICATE)); err != nil {
		BCMap := map[string][]byte{}
		BCMapByte, _ := json.Marshal(BCMap)
		if err := ctx.PutObject([]byte(BCCERTIFICATE), BCMapByte); err != nil {
			return code.Error(err)
		}
	}
	return code.OK([]byte("initializing successfully"))
}

func (o *ordinate) UploadLawEnforcement(ctx code.Context) code.Response{
	//验证传入参数格式
	data := lawEnforcement{}
	if err := code.Unmarshal(ctx.Args(), &data); err != nil {
		return code.Error(err)
	}
	//json-->[]byte
	dataByte, _ := json.Marshal(data)
	//获取map
	lawMapByte, err := ctx.GetObject([]byte(LAWENFORCEMENT))
	if err != nil {
		return code.Error(err)
	}
	lawMap := map[string][]byte{}
	_ = json.Unmarshal(lawMapByte, &lawMap)
	//修改map
	lawMap[data.Uuid] = dataByte
	//上传map
	lawMapByte, _ = json.Marshal(lawMap)
	if err := ctx.PutObject([]byte(LAWENFORCEMENT), lawMapByte); err != nil {
		return code.Error(err)
	}
	return code.OK(lawMapByte)
}

func (o *ordinate) GetAllLawEnforcements(ctx code.Context) code.Response {
	//获取map
	lawMapByte, err := ctx.GetObject([]byte(LAWENFORCEMENT))
	if err != nil {
		return code.Error(err)
	}
	lawMap := map[string][]byte{}
	_ = json.Unmarshal(lawMapByte, &lawMap)
	//获取map内所有数据
	rst := make([][]byte, 0, len(lawMap))
	for _, value := range lawMap {
		rst = append(rst, value)
	}
	//转换格式
	rstByte, _ := json.Marshal(rst)
	//回传
	return code.OK(rstByte)
}

func main() {
	//CreateContractAccount()
	//DeployContract()
	//InvokeUploadLawEnforcement("12f4189a58be48828125af7c41806aab", "district_B", "Alice", "2021-12-10 22:57:58", "qwed7b11d1bbb5be258ca4024f477d50ec047ade9cf04853eeb540b2b0b51f73")
	driver.Serve(new(ordinate))
}
