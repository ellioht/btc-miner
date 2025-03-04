package cpu

import (
	"context"
	"github.com/ellioht/btc-miner/common"
	"github.com/ellioht/btc-miner/core"
	"github.com/ellioht/btc-miner/core/script"
	"github.com/ellioht/btc-miner/log"
	"github.com/ellioht/btc-miner/merkle"
	"github.com/ellioht/btc-miner/rpc"
	"sync"
	"sync/atomic"
	"time"
)

type Miner struct {
	ctx context.Context
	sync.Mutex
	started       atomic.Bool
	mining        atomic.Bool
	collecting    atomic.Bool
	cfg           *MiningConfig
	currentHeight uint64
	jonChan       chan rpc.GetBlockTemplateResult
	stopChan      chan struct{}
	currentNonce  uint32
}

func NewCpuMiner(ctx context.Context, cfg *MiningConfig) *Miner {
	return &Miner{
		ctx:      ctx,
		cfg:      cfg,
		jonChan:  make(chan rpc.GetBlockTemplateResult),
		stopChan: make(chan struct{}),
	}
}

func (m *Miner) Start() {
	m.Lock()
	defer m.Unlock()

	m.started.Store(true)

	go m.collectJob()
	go m.mine()

	m.mining.Store(true)
}

func (m *Miner) Stop() {
	m.Lock()
	defer m.Unlock()

	m.started.Store(false)
	m.mining.Store(false)
	m.collecting.Store(false)
	close(m.stopChan)
}

func (m *Miner) IsStarted() bool {
	return m.started.Load()
}

func (m *Miner) IsMining() bool {
	return m.mining.Load()
}

func (m *Miner) mine() {
	log.Info("Mining started")

	miningTicker := time.NewTicker(10 * time.Second)
	defer miningTicker.Stop()

	for {
		select {
		case <-m.ctx.Done():
			log.Info("Mining stopping")
			return
		case <-m.stopChan:
			log.Info("Mining stopped")
			return
		case <-miningTicker.C:
			if !m.mining.Load() {
				continue
			}
			log.Info("Mining block...", "nonce", m.currentNonce)
		case job := <-m.jonChan:
			m.mining.Store(true)
			log.Info("Mining new block...", "height", job.Height, "transactions", len(job.Transactions))

			txIds := make([]common.Hash, len(job.Transactions)+1)

			coinbaseMessageBytes := []byte("Hello, world!")

			coinbaseScript, err := script.EncodeCoinbaseScript(int64(job.Height), coinbaseMessageBytes)
			if err != nil {
				log.Warn("Error encoding coinbase script", "error", err)
				continue
			}

			coinbaseTx, err := core.MakeCoinbaseTx(coinbaseScript, script.PubKeyScript(common.StringToAddress("")), job.CoinbaseValue)
			if err != nil {
				log.Warn("Error creating coinbase transaction", "error", err)
				continue
			}

			txIds[0] = coinbaseTx.TxId

			for i, tx := range job.Transactions {
				txIds[i+1] = tx.TxID
			}

			tree := merkle.NewTreeFromHashes(txIds)
			calculatedMerkleRoot := tree.ComputeMerkleRoot()

			version := job.Version
			prevBlockHash := job.PreviousBlockHash
			timestamp := uint32(time.Now().Unix())

			bits, err := common.CompactToUint32(job.Bits)
			if err != nil {
				log.Warn("Error converting bits to uint32", "error", err)
				continue
			}

			go func() {
				nonce := uint32(1)
				for {
					header := core.NewHeader(version, prevBlockHash, calculatedMerkleRoot, timestamp, bits, nonce)
					if IsSolved(header) {
						log.Info("Solved block")
						close(m.stopChan)
						break
					}

					nonce++
					m.currentNonce = nonce
				}
			}()
		}
	}
}

func (m *Miner) collectJob() {
	log.Info("Job collector started")
	m.collecting.Store(true)

	jobTicker := time.NewTicker(time.Duration(m.cfg.JobFetchInterval) * time.Second)
	defer jobTicker.Stop()

	for {
		select {
		case <-m.ctx.Done():
			return
		case <-m.stopChan:
			log.Info("Job collector shut down.")
			return
		case <-jobTicker.C:
			// TODO: remove
			if m.currentHeight > 0 {
				continue
			}

			api := rpc.NewApiClient(m.cfg.RpcUrl)

			bt, err := api.GetBlockTemplate(rpc.CreateGetBlockTemplateParams())
			if err != nil {
				log.Warn("Error getting block template", "error", err)
				continue
			}

			log.Info("Job collected", "height", bt.Height, "transactions", len(bt.Transactions))

			height := uint64(bt.Height)

			if height > m.currentHeight {
				log.Info("New block template received", "height", height)
				m.currentHeight = height
				m.jonChan <- *bt
			}
		}
	}
}
