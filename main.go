package main

import (
	"fmt"
	"time"
)

type Block struct {
	timestamp time.Time
	transaction []string
	prevhash	[]byte
	Hash        []byte
}


func main() {
	abc := []string{"A sent 10 coins to B and C"}
	xyz := Blocks(abc, []byte{})
	fmt.Println("This is our First Block")
	Print(xyz)

	def := []string{"D sent 25 coins to E and F"}
	uvw := Blocks(def, xyz.Hash)
	fmt.Println("This is our First Block")
	Print(uvw)
}

func Blocks(transactions []string , prevhash []byte) *Block {
	currentTime:= time.Now()
	return &Block{
		timestamp: currentTime,
		transactions: transactions,
		prevhash: prevhash,
		Hash: NewHash(currentTime, transactions,prevhash),
	}
}

func NewHash(time time.Time, transactions[]string , prevhash[]byte) []byte{
	input := append(prevhash, time.String()...)
	for transactions:= range transactions {
		input = append(input, string(rune(transaction))...)
	}
	hash := sha256.sum256(input)
	return hash[:]
}

func Print(block *Block){
	fmt.Printf("\ttime: %s\n", block.timestamp.String())
	fmt.Printf("\tprevhash: %x\n", block.prevhash)
	fmt.Printf("\thash: %x\n", block.Hash)
}
func Transaction(block *Block) {
	fmt.Println("\tTransactions:")
	for i, transaction := range block.transactions {
		fmt.Printf("\t\t&v: %q\n", i, transaction)
	}
}