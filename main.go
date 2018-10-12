package main

import (
	"context"
	"fmt"

	logging "github.com/ipsn/go-ipfs/gxlibs/github.com/ipfs/go-log"
	"github.com/quorumcontrol/ipfsplay/ipfs"
)

func init() {
	logging.SetLogLevel("core", "debug")
	logging.SetLogLevel("blockstore", "debug")
	logging.SetLogLevel("main", "debug")
}

var log = logging.Logger("main")

func main() {
	_, err := ipfs.StartIpfs(context.TODO(), "./storage")
	if err != nil {
		panic(fmt.Errorf("error starting ipfs: %v", err))
	}
	select {}
}
