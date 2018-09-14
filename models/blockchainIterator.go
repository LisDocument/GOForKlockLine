package models

import (
	"github.com/boltdb/bolt"
	"log"
)

type BlockchainIterator struct {
	currentHash []byte
	db *bolt.DB
}

/**
迭代器迭代 ？？？非根节点向后迭代。末节点朝根节点迭代
 */
func (i *BlockchainIterator) Next() *Block  {

	var block *Block

	//只读方式打开区块链数据库
	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		//获取数据库中当前区块被序列化的区块
		encoderBlock := b.Get(i.currentHash)
		//反序列化操作
		block = DeserializeBlock(encoderBlock)

		return nil
	})
	if err != nil{
		log.Panic(err)
	}

	//吧迭代器的当前区块hash设置为上一区块的hash，实现迭代的作用
	i.currentHash = block.PrevBlockHash

	return block
}
