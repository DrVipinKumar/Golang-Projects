package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

const Difficulty = 16

func (pow *ProofOfWork) Validate() bool {
	var initHash big.Int
	var hash [32]byte
	data := pow.InitData(pow.Block.Nonce)
	hash = sha256.Sum256(data)
	initHash.SetBytes(hash[:])
	return initHash.Cmp(pow.Target) == -1
}
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	pow := &ProofOfWork{b, target}
	return pow
}
func ToBytes(num int64) []byte {
	var buff = new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Fatal(err)
	}
	return buff.Bytes()
}
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.Block.Data,
		pow.Block.PrevHash,
		ToBytes(int64(nonce)),
		ToBytes(int64(Difficulty))},
		[]byte{})
	return data
}
func (pow *ProofOfWork) Run() (int, []byte) {
	var initHash big.Int
	var hash [32]byte
	var nonce = 0
	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		initHash.SetBytes(hash[:])
		if initHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()
	return nonce, hash[:]
}
