package models

import (
	"bytes"
	"crypto/sha256"
	"math/big"
	"strconv"
	"time"
)

var targetBits = 24
/**
区块
 */
type Block struct {
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash          []byte
}
/**
区块链
 */
type Blockchain struct {
	blocks []*Block
}

type ProofOfWork struct {
	block *Block
	target *big.Int
}
/**
指向目标的指针
 */
func NewProofOfWork(b *Block) *ProofOfWork  {
	target := big.NewInt(1)
	//右移xxx位
	target.Lsh(target,uint(256-targetBits))

	pow := &ProofOfWork{b,target}

	return pow
}

func (bc *Blockchain) GetBlocks() []*Block  {
	return bc.blocks
}

func (b *Block) SetHash(){
	timestamp := []byte(strconv.FormatInt(b.Timestamp,10))
	headers := bytes.Join([][]byte{b.PrevBlockHash,b.Data,timestamp},[]byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block{
	block := &Block{time.Now().Unix(),[]byte(data),prevBlockHash,[]byte{}}
	block.SetHash()
	return block
}

func (bc *Blockchain) AddBlock(data string){

	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks,newBlock)

}

func NewGenesisBlock() *Block  {
	return NewBlock("Genesis Block",[]byte{})
}

func NewBlockChain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}