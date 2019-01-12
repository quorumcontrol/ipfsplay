package ipfs

import (
	"context"
	"fmt"

	"github.com/ipsn/go-ipfs/core"
	"github.com/ipsn/go-ipfs/core/coreapi"
	coreiface "github.com/ipsn/go-ipfs/core/coreapi/interface"
)

const (
	nBitsForKeypairDefault = 2048
)

func StartIpfs(ctx context.Context, repoPath string) (coreiface.DagAPI, error) {
	// if !fsrepo.IsInitialized(repoPath) {
	// 	conf, err := config.Init(os.Stdout, nBitsForKeypairDefault)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("error initializing file: %v", err)
	// 	}

	// 	transformer, ok := config.Profiles["server"]
	// 	if !ok {
	// 		return nil, fmt.Errorf("invalid configuration profile: %s", "server")
	// 	}

	// 	if err := transformer.Transform(conf); err != nil {
	// 		return nil, fmt.Errorf("error transforming: %v", err)
	// 	}
	// 	fmt.Printf("conf: %v", conf)

	// 	if err := fsrepo.Init(repoPath, conf); err != nil {
	// 		return nil, fmt.Errorf("error initializng fsrepo: %v", err)
	// 	}
	// }

	// repo, err := fsrepo.Open(repoPath)
	// if err != nil {
	// 	return nil, fmt.Errorf("error opening repo: %v", err)
	// }

	ncfg := &core.BuildCfg{
		// Repo:      repo,
		Online:    true,
		Permanent: true,
		Routing:   core.DHTOption,
	}

	node, err := core.NewNode(ctx, ncfg)
	if err != nil {
		return nil, fmt.Errorf("error new node: %v", err)
	}

	api, err := coreapi.NewCoreAPI(node)
	if err != nil {
		return nil, fmt.Errorf("error creating node: %v", err)
	}
	return api.Dag(), nil
}
