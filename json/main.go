package main

import (
	"encoding/csv"
	"strconv"

	//"bytes"
	//"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	//"os"

)

type Oid struct {
	OOid string `json:"$oid"`
}

type Xxjson struct {
	Ooid Oid `json:"_id"`
	Id string `json:"Id"`
	NickName string `json:"NickName"`
	Gender string `json:"Gender"`
	Location string `json:"Location"`
	BriefIntroduction string `json:"BriefIntroduction"`
	Url string `json:"Url"`
	Viplevel int `json:"Viplevel"`
	Fans string `json:"Fans"`
	Follows string `json:"Follows"`
	Tweets string `json:"Tweets"`
}


func UMSreadFile(filename string)(Xxjson, error){
	var xj Xxjson
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("ReadFile error")
		panic(err.Error())
	}
	defer jsonFile.Close()
	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("err reading json data: ", err)
		panic(err.Error())
	}
	if err := json.Unmarshal(bytes, &xj); err != nil {
		fmt.Println("Unmarshal error")
		panic(err.Error())
	}

	return xj, nil
}

func DCDreadFile(filename string) ([]Xxjson, error) {
	var xj []Xxjson
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening JSON file", err)
		panic(err.Error())
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	for {
		var xjj Xxjson
		err := decoder.Decode(&xjj)
		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Println("Error decoding JSON:", err)
			panic(err.Error())
		}
		xj = append(xj, xjj)
	}
	return xj, nil
}

func main() {
	//读取json文件
	jsdata, err := DCDreadFile("weibo-users.json")
	if err != nil {
		fmt.Println("main readFile error")
		panic(err.Error())
	}
	fmt.Println(jsdata)
	//创建csv文件
	f, err := os.Create("test2.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")//写入utf-8 bom
	w := csv.NewWriter(f)
	for _, xxjson := range jsdata{
		line := [][]string{{xxjson.Ooid.OOid}, {xxjson.Id}, {xxjson.NickName},
		{xxjson.Gender}, {xxjson.Location}, {xxjson.BriefIntroduction},
		{xxjson.Url}, {strconv.Itoa(xxjson.Viplevel)}, {xxjson.Fans}, {xxjson.Follows}, {xxjson.Tweets}}
		err := w.WriteAll(line)
		if err != nil {
			panic(err.Error())
		}
	}
	w.Flush()
}
