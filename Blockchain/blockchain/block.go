package blockchain

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}
type BlockChain struct {
	Blocks []*Block
}

/*func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}for dummy blockchain*/
func CreateBlock(data string, PrevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), PrevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newblock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newblock)
}
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
