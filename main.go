package main

import (
	"fmt"
	"hello/models"
	_ "hello/routers"
)

func main() {
	bc := models.NewBlockChain()

	bc.AddBlock("Send 1")
	bc.AddBlock("Send 2")

	for _,block := range bc.GetBlocks(){
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}

	//beego.Run()
}

