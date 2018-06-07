package block

import (
	"log"

	"github.com/boltdb/bolt"
)

// BlockchainIterator struct
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

// Next ...
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})
	if err != nil {
		log.Println("BlockchainIterator db.View error: ", err)
	}

	i.currentHash = block.PrevBlockHash
	return block
}
