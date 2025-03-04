package main

import (
	"context"
	"github.com/ellioht/btc-miner/cpu"
)

func main() {
	ctx := context.Background()

	cfg := &cpu.MiningConfig{
		RpcUrl:           "",
		JobFetchInterval: 10,
	}

	miner := cpu.NewCpuMiner(ctx, cfg)
	miner.Start()

	<-ctx.Done()
}
