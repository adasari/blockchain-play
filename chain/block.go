package chain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Data     []byte
	PrevHash []byte
	Hash     []byte
	Nonce    int
}

func NewBlock(data string, prevHash []byte) *Block {
	b := &Block{Data: []byte(data), PrevHash: prevHash}
	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return b
}

func (b *Block) SetHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func Genesis() *Block {
	return NewBlock("Genesis", []byte{})
}
