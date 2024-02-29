package main

import (
	"fmt"
	"time"
)

type Block struct {
	Timestamp    int64
	Data         string
	PrevBlockPtr *Block
}

type Blockchain struct {
	Head *Block
}

func (bc *Blockchain) NewBlock(data string) *Block {
	newBlock := &Block{
		Timestamp:    time.Now().Unix(),
		Data:         data,
		PrevBlockPtr: bc.Head,
	}

	bc.Head = newBlock
	return newBlock
}

func (bc *Blockchain) DisplayAllBlocks() {
	for block := bc.Head; block != nil; block = block.PrevBlockPtr {
		fmt.Printf("Timestamp: %d, Data: %s\n", block.Timestamp, block.Data)
	}
}

func (bc *Blockchain) ModifyBlock(position int, newData string) bool {
	current := bc.Head
	for i := 0; current != nil && i < position; i++ {
		current = current.PrevBlockPtr
	}

	if current == nil {
		return false
	}

	current.Data = newData
	return true
}

func main() {
	bc := &Blockchain{}

	bc.NewBlock("First Block")
	bc.NewBlock("Second Block")
	bc.NewBlock("Third Block")

	fmt.Println("Displaying all blocks:")
	bc.DisplayAllBlocks()

	success := bc.ModifyBlock(1, "Modified Second Block")
	if success {
		fmt.Println("\nAfter modification:")
		bc.DisplayAllBlocks()
	} else {
		fmt.Println("\nFailed to modify block")
	}
}
