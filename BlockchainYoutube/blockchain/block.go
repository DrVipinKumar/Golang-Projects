package blockchain

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}
type Blockchain struct {
	Blocks []*Block
}

// func (b *Block) GetHash() {
// 	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
// 	hash := sha256.Sum256(info)
// 	b.Hash = hash[:]
// }
func NewBlock(data string, prevhash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevhash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewBlock("Genesis", []byte{})}}
}
func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newblock := NewBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newblock)
}
