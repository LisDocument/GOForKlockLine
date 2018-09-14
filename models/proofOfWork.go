package models

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"hello/util"
	"math"
	"math/big"
)

/**
区块链指针
 */
type ProofOfWork struct {
	block *Block
	target *big.Int
}
/**
区块链获取hash的目标数据
 */
func (pow *ProofOfWork) prepareData(nonce int) []byte{
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			util.IntToHex(pow.block.Timestamp),
			util.IntToHex(int64(targetBits)),
			util.IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
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

/**
POW算法核心
 */
func (pow *ProofOfWork) Run() (int,[]byte)  {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	MaxNonce := math.MaxInt64
	fmt.Printf("Mining the block containing \"%s\"\n",pow.block.Data)

	for nonce < MaxNonce{
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x\t%d",hash,nonce)

		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1{
			break
		}else{
			nonce ++
		}

	}
	fmt.Print("\n\n")
	return nonce,hash[:]
}

/**
有效量机制判别
 */
func (pow *ProofOfWork) Validate() bool{
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)

	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
