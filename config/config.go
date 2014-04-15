// Author：oliver		dumanxiang@gmail.com

package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"myblog/utils"
)

var (
	Config = make(map[string]string)
	ROOT   = "."
)

func init() {
	ROOT, err := utils.ExeDir()
	if err != nil {
		log.Fatal(err)
	}

	//Load配置文件，ReadFile可以解析windows地址格式与linux地址格式的混合
	//比如："C:\golang\src\LBSIM\src/conf/config.json"
	configFile := ROOT + "/config.json"
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &Config)
	if err != nil {
		log.Fatal(err)
	}
}
