package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"time"

	coreiface "github.com/ipsn/go-ipfs/core/coreapi/interface"
	opt "github.com/ipsn/go-ipfs/core/coreapi/interface/options"
	"github.com/ipsn/go-ipfs/gxlibs/github.com/ipfs/go-ipld-cbor"
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shouldPut := flag.Bool("p", false, "Should we put the randevous?")
	flag.Parse()

	node, err := cbornode.WrapObject(map[string]interface{}{"hi": "hi"}, multihash.SHA2_256, -1)
	if err != nil {
		panic(fmt.Sprintf("error wrapping: %v", err))
	}
	log.Infof("node CID: %v, len: %d", node.Cid().String(), len(node.RawData()))

	dag, err := ipfs.StartIpfs(ctx, "./storage")
	if err != nil {
		panic(fmt.Errorf("error starting ipfs: %v", err))
	}

	if *shouldPut {
		path, err := dag.Put(ctx, bytes.NewReader(node.RawData()), opt.Dag.InputEnc("cbor"))
		if err != nil {
			panic(fmt.Errorf("error adding file: %v", err))
		}
		log.Infof("path: %v", path)
	} else {
		log.Infof("fetching node... %s", node.Cid().String())
		ctx2, cancel2 := context.WithTimeout(ctx, 30*time.Second)
		defer cancel2()
		data, err := dag.Get(ctx2, coreiface.IpldPath(node.Cid()))
		if err != nil {
			panic(fmt.Errorf("error getting: %v", err))
		}
		log.Infof("got the data! %d", len(data.RawData()))
	}

	select {}
}
