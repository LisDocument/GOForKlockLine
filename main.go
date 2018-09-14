package main

import (
	"bytes"
	"fmt"
	"hello/models"
	_ "hello/routers"
	"strconv"
)

func main() {
	bc := models.NewBlockChain()

	bc.AddBlock("Send 1")
	bc.AddBlock("Send 2")
	bc.AddBlock("Send 3")
	bc.AddBlock("Send 4")
	bc.AddBlock("Send dadaskhdajdklandmaldnasldnaldnakldnasdaskljdfnakljdaldnasldnad")

	iterator := bc.Iterator()

	iterator.Next()

	for{
		block := iterator.Next()
		if bytes.Equal([]byte{},block.PrevBlockHash){
			break
		}
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := models.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}

	//for _,block := range bc.GetBlocks(){
	//	fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
	//	fmt.Printf("Data: %s\n", block.Data)
	//	fmt.Printf("Hash: %x\n", block.Hash)
	//	pow := models.NewProofOfWork(block)
	//	fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	//	fmt.Println()
	//}

	//beego.Run()
}

