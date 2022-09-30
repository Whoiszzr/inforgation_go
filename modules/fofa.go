package modules

import (
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/liushuochen/gotable"
	"github.com/liushuochen/gotable/table"
	"io"
	"net/http"
)

type FofaData struct {
	Error   bool       `json:"error"`
	Size    int        `json:"size"`
	Page    int        `json:"page"`
	Mode    string     `json:"mode"`
	Query   string     `json:"query"`
	Results [][]string `json:"results"`
}

func Fofa(ip string, apikey string, mail string) error {
	// 生成查询query
	words := b64.StdEncoding.EncodeToString([]byte("ip=" + ip))
	// 拼接url
	fofaUrl := "https://fofa.info/api/v1/search/all?email=" + mail + "&key=" + apikey +
		"&qbase64=" + words + "&fields=host,title,country_name,province,city,server,protocol,isp"
	fmt.Printf("fofaUrl: %v\n", fofaUrl)
	// 请求
	resp, err := http.Get(fofaUrl)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if err != nil {
		return err
	} else {
		var data FofaData
		_ = json.NewDecoder(resp.Body).Decode(&data)
		if data.Error == true {
			return errors.New("请求失败！请检查网络、APIKEY、EMAIL！")
		} else {
			tableData := handleData(data.Results)
			err, tb := genTable(tableData)
			if err != nil {
				return err
			} else {
				tb
				return nil
			}
		}
	}
}

// 处理数据，返回一个切片，切片单个元素为map，对应键值对
func handleData(fofaData [][]string) []map[string]string {
	var tableData []map[string]string
	for _, result := range fofaData {
		tableData = append(tableData, map[string]string{"host": result[0]})
		tableData = append(tableData, map[string]string{"title": result[1]})
		tableData = append(tableData, map[string]string{"address": result[2] + " " + result[3] + " " + result[4]})
		tableData = append(tableData, map[string]string{"service": result[5]})
		tableData = append(tableData, map[string]string{"protocol": result[6]})
	}
	return tableData
}

func genTable(tableData []map[string]string) (error, *table.Table) {
	tb, err := gotable.Create("host", "标题", "地理位置", "服务名", "协议")
	if err != nil {
		return err, tb
	}
	for _, data := range tableData {
		err = tb.AddRow(data)
		if err != nil {
			return err, tb
		}
	}
	return nil, tb
}
