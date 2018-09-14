package transaction

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

//挖矿奖励
const subsidy = 50

/**
穿件一个交易的数据结构，交易是由交易ID、交易输入、交易输出组成的
一个交易有多个输入和多个输出，所以这里的类型应该也是切片类型的
 */
type Transaction struct {
	ID 		[]byte
	Vin 	[]TxInput
	Vout 	[]TxOutput
}

/*
1、每一笔交易的输入都会引用之前的一笔或多笔交易输出
2、交易输出保存了输出的值和锁定该输出的信息
3、交易输入保存了引用之前交易输出的交易ID、具体到引用
该交易的第几个输出、能正确解锁引用输出的签名信息
 */

 //交易输出
type TxOutput struct {
	Value 			int //输出的值（可以理解为金额）
	ScriptPubKey 	string //锁定该输出的脚本（未实现地址，不确定具体为某个地址所有）
}

//交易输入
type TxInput struct {
	Txid		[]byte //引用的之前交易的ID
	Vout 		int		//引用之前交易输出的具体是哪个输出（一个交易中输出一般有很多）
	ScriptSig 	string	//能解锁引用输出交易的签名脚本（目前因为未实现地址，暂时无）
}

/*
区块链上存储的交易都是由这些输入输出交易所组成的，
一个输入交易必须引用之前的输出交易，一个输出交易会被之后的输入所引用
问题来了，在最开始的区块链上是先有输入还是先有输出
答案是先有输出，因为是区块链的创世块产生了第一个输出，
这个输出也就是我们常说的挖矿奖励，每一个区块都会有一个这样的输出，
这是奖励给矿工的交易输出，这个输出是凭空产生的
*/

//to 标识此输出建立给谁， 一般是矿工地址，data是交易附带的信息
func NewCoinbaseTX(to,data string) *Transaction{
	if data == "" {
		data = fmt.Sprintf("奖励给 %s",to)
	}
	//此交易中的交易输入，没有交易输入信息
	txin := TxInput{[]byte{},-1,data}
	//交易输出，subsidy为奖励矿工的币的数量
	txout := TxOutput{subsidy,to}
	tx := Transaction{nil, []TxInput{txin},[]TxOutput{txout}}
	tx.SetID()

	return &tx
}

//设置交易Id，交易Id是序列化tx后再hash
func (tx *Transaction) SetID()  {

	var hash [32]byte
	var encoder bytes.Buffer

	enc := gob.NewEncoder(&encoder)
	err := enc.Encode(tx)

	if err != nil{
		log.Panic(err)
	}

	hash = sha256.Sum256(encoder.Bytes())

	tx.ID = hash[:]
}

/**
1.每一个区块至少存储一笔coinbase交易，所以我们在区块的字段中吧Data字段换成交易
2.把所有涉及之前Data字段都要换了，比如NewBlock（），GenesisBlock(),等函数
 */