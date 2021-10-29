package params

import (
	"encoding/json"
	"errors"

	"github.com/BurntSushi/toml"
	"github.com/anyswap/CrossChain-Bridge/common"
	"github.com/anyswap/CrossChain-Bridge/log"
)

var (
	configFile string
	yishengConfig = &Config{}
	mongodbConfig = &MongoDBConfig{}
	gatewayConfig = &GatewayConfig{}
	accountConfig = &AccountConfig{}
)

// LoadConfig load config
func LoadConfig(filePath string) *Config {
	log.Println("LoadConfig Config file is", filePath)
	if !common.FileExist(filePath) {
		log.Fatalf("LoadConfig error: config file '%v' not exist", filePath)
	}

	config := &Config{}
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		log.Fatalf("LoadConfig error (toml DecodeFile): %v", err)
	}

	var bs []byte
	if log.JSONFormat {
		bs, _ = json.Marshal(config)
	} else {
		bs, _ = json.MarshalIndent(config, "", "  ")
	}
	log.Println("LoadConfig finished.", string(bs))

       mongodbConfig = config.MongoDB
	gatewayConfig = config.Gateway
	accountConfig = config.Account

       if err := accountConfig.CheckConfig(); err != nil {
		log.Fatalf("LoadConfig Check config failed. %v", err)
	}

	configFile = filePath // init config file path
	return yishengConfig
}

// CheckConfig check scan config
func (a *AccountConfig) CheckConfig() (err error) {
	if !common.FileExist(a.KeyFrom) {
		log.Fatalf("LoadConfig error: config file '%v' not exist", a.KeyFrom)
	}
	if !common.FileExist(a.PasswdFrom) {
		log.Fatalf("LoadConfig error: config file '%v' not exist", a.PasswdFrom)
	}
	if a.AddressTo == "" || !common.IsHexAddress(a.AddressTo) {
		return errors.New("wrong 'AddressTo' " + a.AddressTo)
	}
	return nil
}


type Config struct {
       MongoDB *MongoDBConfig
       Gateway *GatewayConfig
	Account *AccountConfig
}

// MongoDBConfig mongodb config
type MongoDBConfig struct {
       DBURL      string
       DBName     string
       UserName   string `json:"-"`
       Password   string `json:"-"`
}

type GatewayConfig struct {
	URL string
}

type AccountConfig struct {
	KeyFrom string
	PasswdFrom string
	AddressTo string
}

// GetMongodbConfig get mongodb config
func GetMongodbConfig() *MongoDBConfig {
       return mongodbConfig
}

// GetMongodbConfig get mongodb config
func GetGatewayConfig() *GatewayConfig {
       return gatewayConfig
}

// GetMongodbConfig get mongodb config
func GetAccountConfig() *AccountConfig {
       return accountConfig
}

