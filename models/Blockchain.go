package models

import (
	"github.com/boltdb/bolt"
	"hello/transaction"
	"log"
)

const dbFile  = "block.db"
const blocksBucket  = "block"

/**
区块链
 */
//type Blockchain struct {
//	blocks []*Block
//}
type Blockchain struct {
	db *bolt.DB
	tip []byte
}

func(bc *Blockchain) Db() *bolt.DB{
	return bc.db
}

//把区块添加进区块链挖矿
func(bc *Blockchain) MineBlock(transactions []*transaction.Transaction){
	var lastHash []byte
	//只读方式浏览数据库，获取当前区块链顶端区块的hash，为下以区块做准备
	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("1"))

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	//求出新区块
	newBlock := NewBlock(transactions,lastHash)

	//把新区块加入到数据库区块链中
	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash,newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}
		err = b.Put([]byte("l"),newBlock.Hash)
		bc.tip = newBlock.Hash

		return nil
	})
}

/**
添加新的区块
 */
func (bc *Blockchain) AddBlock(data string){

	//prevBlock := bc.blocks[len(bc.blocks)-1]
	//newBlock := NewBlock(data, prevBlock.Hash)
	//bc.blocks = append(bc.blocks,newBlock)

	var lastHash []byte
	//只读的方式浏览数据库，获取当前区块链顶端区块的哈希，为进入下一区块做准备
	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		//通过链“l”拿到区块链顶端区块hash
		lastHash = b.Get([]byte("l"))

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	//创建下一块区块
	newBlock := NewBlock(data, lastHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		//获取桶
		bucket := tx.Bucket([]byte(blocksBucket))
		//将新建立的区块放入桶中
		err := bucket.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}
		//将l重新赋值为新的区块Hash
		err = bucket.Put([]byte("l"),newBlock.Hash)
		bc.tip = newBlock.Hash

		return nil
	})
}

/**
当需要遍历当前区块链时，创建一个迭代器
 */
func (bc *Blockchain) Iterator() *BlockchainIterator{

	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}

/**
创建链的创世块
 */
func NewGenesisBlock(coinbase *transaction.Transaction) *Block  {
	return NewBlock([]*transaction.Transaction{coinbase},[]byte{})
}

/**
创建链
 */
func NewBlockChain(address string) *Blockchain {

	var tip []byte
	//打开一个数据库文件，如果文件不存在则创建改名字的文件
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil{
		log.Panic(err)
	}
	//读写操作数据库
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		//确认数据库中是否有存在初始桶，存在则跳过，不存在则创建，并建立创世块，并存入
		if b == nil {
			genesisBlock := NewGenesisBlock()
			b,err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}
			err = b.Put(genesisBlock.Hash,genesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			//存放创世块的Hash值
			err = b.Put([]byte("l"),genesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			tip = genesisBlock.Hash

		} else {
			//如果存在桶的话
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	bc := Blockchain{db,tip}
	return &bc
	//return &Blockchain{[]*Block{NewGenesisBlock()}}
}