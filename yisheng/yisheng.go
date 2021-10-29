package yisheng

import (
	"context"
	"math/big"
	"time"

	"github.com/anyswap/CrossChain-Bridge/cmd/utils"
	"github.com/anyswap/CrossChain-Bridge/log"
	"github.com/urfave/cli/v2"

	ethclient "github.com/jowenshaw/gethclient"

	"github.com/weijun-sh/gethscan/params"
	"github.com/weijun-sh/gethscan/mongodb"
)

var (
	// ScanSwapCommand scan swaps on eth like blockchain
	YishengCommand = &cli.Command{
		Action:    yisheng,
		Name:      "yisheng",
		Usage:     "scan mongodb",
		ArgsUsage: " ",
		Description: `
scan mongodb
`,
		Flags: []cli.Flag{
			utils.ConfigFileFlag,
			utils.GatewayFlag,
			utils.JobsFlag,
		},
	}
)

type ethSwapScanner struct {
	gateway     string
	chainID *big.Int

	jobCount     uint64

	client *ethclient.Client
	ctx    context.Context

	rpcInterval   time.Duration
	rpcRetryCount int
}

func yisheng(ctx *cli.Context) error {
	utils.SetLogger(ctx)
	params.LoadConfig(utils.GetConfigFilePath(ctx))

	scanner := &ethSwapScanner{
		ctx:           context.Background(),
		rpcInterval:   1 * time.Second,
		rpcRetryCount: 3,
	}
	scanner.gateway = ctx.String(utils.GatewayFlag.Name)

	log.Info("get argument success",
		"gateway", scanner.gateway,
		"jobs", scanner.jobCount,
	)

	scanner.initClient()

	scanner.run()
	return nil
}

func (scanner *ethSwapScanner) initClient() {
	ethcli, err := ethclient.Dial(scanner.gateway)
	if err != nil {
		log.Fatal("ethclient.Dail failed", "gateway", scanner.gateway, "err", err)
	}
	log.Info("ethclient.Dail gateway success", "gateway", scanner.gateway)
	scanner.client = ethcli
	scanner.chainID, err = ethcli.ChainID(scanner.ctx)
	if err != nil {
		log.Fatal("get chainID failed", "err", err)
	}
	log.Info("get chainID success", "chainID", scanner.chainID)
}

func (scanner *ethSwapScanner) run() {
	// getTableFromMongodb
	// sendTx
}

// InitMongodb init mongodb by config
func InitMongodb() {
       log.Info("InitMongodb")
       dbConfig := params.GetMongodbConfig()
       mongodb.MongoServerInit([]string{dbConfig.DBURL}, dbConfig.DBName, dbConfig.UserName, dbConfig.Password)
}

