package chain

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain(genesis *Block) *Blockchain {
	return &Blockchain{[]*Block{genesis}}
}

func (chain *Blockchain) AddBlock(data string) {
	lastBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := NewBlock(data, lastBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}
