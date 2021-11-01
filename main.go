package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

var Config ConfigFile
var PublicIp string

func main() {
	if runtime.GOOS != "linux" {
		Quit("only linux environment is supported")
	}

	configFile, err := os.Open(path.Join(ConfigPath, ConfigFilename))

	if err != nil {
		Quit("config file not found")
	}

	defer configFile.Close()

	bytes, err := ioutil.ReadAll(configFile)

	if err != nil {
		Quit("could not read config file")
	}

	err = json.Unmarshal(bytes, &Config)

	if err != nil {
		Quit("could not unmarshal config file")
	}

	ip, err := GetIp()

	if err != nil {
		Quit("could not get public ip")
	}

	PublicIp = ip

	for i := 0; i < len(Config.DnsList); i++ {
		Config.DnsList[i].Handle()
	}
}