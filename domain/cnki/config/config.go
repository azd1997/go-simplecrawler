/**********************************************************************
* @Author: Eiger (201820114847@mail.scut.edu.cn)
* @Date: 2020/4/17 1:01
* @Description: The file is for
***********************************************************************/

package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const cfgFile = "./config.json"

var boolMap = map[int]bool{
	0: false,
	1:true,
}

type conf struct {
	DownloadFile   bool
	CrackCode      bool
	DetailPage     bool
	DownloadLink   bool
	StepWaitSecond uint16

	Headers *http.Header
}

type jsonConfig struct {
	DownloadFile int `json:"download_file"`
	CrackCode int `json:"crack_code"`
	DetailPage int `json:"detail_page"`
	DownloadLink int `json:"download_link"`
	StepWaitSecond int `json:"step_wait_second"`
}

func ParseConfig() *conf {
	jcfg := &jsonConfig{}
	content, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		log.Fatalf("read config file: %v\n", err)
	}
	err = json.Unmarshal(content, jcfg)
	if err != nil {
		log.Fatalf("unmarshal config content: %v\n", err)
	}

	cfg := &conf{}
	cfg.CrackCode = boolMap[jcfg.CrackCode]
	cfg.DetailPage = boolMap[jcfg.DetailPage]
	cfg.DownloadFile = boolMap[jcfg.DownloadFile]
	cfg.DownloadLink = boolMap[jcfg.DownloadLink]
	cfg.StepWaitSecond = uint16(jcfg.StepWaitSecond)

	return cfg
}
