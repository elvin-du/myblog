// Author：macs		dumanxiang@gmail.com

package config

import (
	"encoding/json"
	"io/ioutil"
	"myblog/utils"
)

// 项目根目录
var ROOT string
var Config map[string]string

func init() {
	dir, err := utils.ExeDir()
	if err != nil {
		panic(err)
	}

	//不能直接用ROOT来接受utils.ExeDir()的返回值，因为:=会重新生成一个ROOT，
	//这样的话，全局变量就得不到赋值
	ROOT = dir

	//Load配置文件，ReadFile可以解析windows地址格式与linux地址格式的混合
	//比如："C:\golang\src\LBSIM\src/conf/config.json"
	configFile := ROOT + "/conf/config.json"
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	//解析配置文件到Config全局变量
	Config = make(map[string]string)
	err = json.Unmarshal(content, &Config)
	if err != nil {
		panic(err)
	}
}
