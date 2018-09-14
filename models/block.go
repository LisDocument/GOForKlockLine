package models

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"hello/transaction"
	"log"
	"time"
)

var targetBits = 24
/**
区块
 */
type Block struct {
	Timestamp int64
	//Data []byte
	transactions 	[]*transaction.Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce int
}


//func (b *Block) SetHash(){
//	timestamp := []byte(strconv.FormatInt(b.Timestamp,10))
//	headers := bytes.Join([][]byte{b.PrevBlockHash,b.Data,timestamp},[]byte{})
//	hash := sha256.Sum256(headers)
//
//	b.Hash = hash[:]
//}

func NewBlock(data []*transaction.Transaction, prevBlockHash []byte) *Block{
	block := &Block{time.Now().Unix(),data,prevBlockHash,[]byte{},0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

//区块交易字段，因为每个交易ID是序列化并HASH后的交易数据结构，所以我们只需要把所有的交易ID进行hash可以了
func (b *Block) HashTransactions() []byte  {
	var txHash [32]byte
	var txHashes [][]byte

	for _,tx := range b.transactions{
		txHashes = append(txHashes,tx.ID)
	}

	txHash = sha256.Sum256(bytes.Join(txHashes,[]byte{}))

	return txHash[:]
}


/**
实现序列化
 */
func (b *Block) Serialize() []byte{
	//定义个buff保存序列化后的字符串
	var result = bytes.Buffer{}
	//实例化一个序列化实例，保存到result中
	encoder := gob.NewEncoder(&result)
	//对区划进行实例化
	err := encoder.Encode(b)
	if err != nil{
		log.Panic(err)
	}
	return result.Bytes()
}

/**
反序列化
 */
func DeserializeBlock(d []byte) *Block{
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil{
		log.Panic(err)
	}
	return &block
}

