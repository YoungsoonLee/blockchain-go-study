package block

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, preBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// make first blockchain with genesis block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
