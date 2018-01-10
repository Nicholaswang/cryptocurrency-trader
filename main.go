package main 
import (
	"fmt"
	"log"
	"os"
	"flag"
	"github.com/Nicholaswang/cryptocurrency-trader/api"
	"github.com/Nicholaswang/cryptocurrency-trader/util"
	"github.com/Nicholaswang/cron"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func getBalance() {
	balance := api.Balances()
	fmt.Printf("Balances: %s", balance)
}

var (
    configFile = flag.String("config", "", "Configuration file")
)

func main() {
	flag.Parse()

	if *configFile == "" {
		fmt.Fprintln(os.Stderr, "Missing configuration file")
		flag.PrintDefaults()
		os.Exit(1)
	}
	
	cfg, err := util.LoadConfigFile(*configFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error occurred while loading config file:", err)
	}
	api.KEY = cfg.GateIOConfigs[0].Key
	api.SECRET = cfg.GateIOConfigs[0].Secret

	cr := cron.New()
	cr.Start()
	defer cr.Stop()
	//cr.AddFunc("0 1/* * * ?", getPrice)
	cr.AddFunc("@every 5s", getBalance)
	/*
	var ret string = GetPairs()
	fmt.Printf("Pairs: %s", ret)
	info := Marketinfo()
    fmt.Printf("Marketinfo: %s", info)
    tickers_info := Tickers()
    fmt.Printf("Tickers info: %s", tickers_info)
	*/
	select {}
}
