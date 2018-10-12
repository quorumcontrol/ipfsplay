package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/ipfs/go-ipld-cbor"

	logging "github.com/ipsn/go-ipfs/gxlibs/github.com/ipfs/go-log"
	multihash "github.com/ipsn/go-ipfs/gxlibs/github.com/multiformats/go-multihash"
	"github.com/quorumcontrol/ipfsplay/ipfs"
)

func init() {
	logging.SetLogLevel("core", "debug")
	logging.SetLogLevel("blockstore", "debug")
	logging.SetLogLevel("main", "debug")
}

var log = logging.Logger("main")

func main() {

	cNode, err := cbornode.WrapObject(map[string]interface{}{"hi": "hi"}, multihash.SHA2_256, -1)
	if err != nil {
		panic(fmt.Sprintf("error wrapping: %v", err))
	}
	// cNode2, err := cbornode.WrapObject(map[string]interface{}{"hi": "bye"}, multihash.SHA2_256, -1)
	// if err != nil {
	// 	panic(fmt.Sprintf("error wrapping: %v", err))
	// }

	ipfsAPI, err := ipfs.StartIpfs(context.TODO(), "./storage")
	if err != nil {
		panic(fmt.Errorf("error starting ipfs: %v", err))
	}

	id, err := ipfsAPI.Add(bytes.NewReader(cNode.RawData()))
	if err != nil {
		panic(fmt.Errorf("error adding file: %v", err))
	}
	log.Infof("id: %s", id)

	select {}
}
