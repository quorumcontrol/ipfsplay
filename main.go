package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/ipsn/go-ipfs/gxlibs/github.com/ipfs/go-ipld-cbor"
	"github.com/ipsn/go-ipfs/gxlibs/github.com/multiformats/go-multihash"
	"github.com/ipsn/go-ipfs/pin"

	"github.com/ipsn/go-ipfs/core"
	logging "github.com/ipsn/go-ipfs/gxlibs/github.com/ipfs/go-log"
)

func init() {
	logging.SetLogLevel("core", "debug")
	logging.SetLogLevel("blockstore", "debug")
	logging.SetLogLevel("main", "debug")
}

var log = logging.Logger("main")

func main() {
	shouldPut := flag.Bool("p", false, "Should we put the randevous?")
	flag.Parse()

	cNode, err := cbornode.WrapObject(map[string]interface{}{"hi": "hi"}, multihash.SHA2_256, -1)
	if err != nil {
		panic(fmt.Sprintf("error wrapping: %v", err))
	}
	cNode2, err := cbornode.WrapObject(map[string]interface{}{"hi": "bye"}, multihash.SHA2_256, -1)
	if err != nil {
		panic(fmt.Sprintf("error wrapping: %v", err))
	}

	if *shouldPut {

		log.Infof("putting node")
		node, err := core.NewNode(context.TODO(), &core.BuildCfg{Online: true})
		if err != nil {
			log.Fatalf("Failed to start IPFS node: %v", err)
		}

		log.Infof("adding cnode")
		err = node.DAG.Add(context.TODO(), cNode2)
		if err != nil {
			panic(err)
		}

		err = node.DAG.Add(context.TODO(), cNode)
		if err != nil {
			panic(err)
		}

		node.Pinning.PinWithMode(cNode.Cid(), pin.Recursive)
		err = node.Pinning.Flush()
		if err != nil {
			panic(err)
		}
		node.Blockstore.PinLock().Unlock()

		rnode, _ := node.DAG.Get(context.TODO(), cNode.Cid())
		fmt.Printf("got node: %v", rnode)
	} else {
		ctx, cancelFunc := context.WithTimeout(context.Background(), 60*time.Second)

		node2, err := core.NewNode(ctx, &core.BuildCfg{Online: true})
		if err != nil {
			log.Fatalf("Failed to start IPFS node: %v", err)
		}

		log.Info("sleeping for 15 seconds")
		time.Sleep(15 * time.Second)

		log.Info("finding providers")
		peers, _ := node2.DHT.FindProviders(ctx, cNode.Cid())
		log.Infof("peers found: %d", len(peers))

		log.Infof("getting cnode from 2: %v\n", cNode.Cid().String())

		rNode, err := node2.DAG.Get(context.TODO(), cNode.Cid())

		// d := merkledag.NewDAGService(node2.Blocks)
		// rNode, err := d.Session(context.TODO()).Get(context.TODO(), cNode.Cid())

		if err != nil {
			log.Fatalf("Failed to look up: %v", err)
		}
		cancelFunc()

		log.Infof("rNode: %v", rNode)
	}
	select {}

}
