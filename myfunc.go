package assignment01bca

import (
	"fmt"
	"crypto/sha256"
	"strconv"
	"strings"
	"bufio"
	"os"
)

//defining a structure for a single block of list

type Block struct {
	transaction string
	nonce int
	prev_hash string
	block_hash string

}

//defining a complete blockchain linkedlist structure

type BlockChain struct {
	list []*Block
}

func Create_B() *BlockChain{
	a := new(BlockChain)
	return a
}

//a function which calculates hash using SHA256

func CalculateHash (stringToHash string) string{

	// fmt.Printf("String Received : %s \n \n", stringToHash)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))

}


//creating and adding a new single block in blockchain

func NewBlock(chain *BlockChain,trans string, n int, prev_h string) *Block  {
	b := new(Block)
	b.transaction = trans
	b.nonce = n
	b.prev_hash = prev_h
	b.block_hash = CalculateHash(b.transaction +strconv.Itoa(b.nonce) +b.prev_hash)

	chain.list = append(chain.list, b)


	
	fmt.Printf("\n")
	return b
}


//displaying all the blocks of linkedlist dynamically

func DisplayBlocks(chain *BlockChain) {
	
	
	for index, val := range chain.list {
		fmt.Printf("%s Block :  %d %s\n", strings.Repeat("=", 25), index, strings.Repeat("=", 25))

		fmt.Printf(" Transection: %s \n Nonce: %d \n Previous Block Hash: %s \n Current Block Hash %s \n \n ",val.transaction,val.nonce,val.prev_hash,val.block_hash)
	}
}


//changing/modifying transection information in a block

func ChangeBlock(chain *BlockChain) {
	var ind int
	fmt.Println("Enter the index of block you want to edit transections :")
	fmt.Scan(&ind)
	fmt.Println(ind)

	if(ind<len(chain.list)){
		fmt.Printf("current transection is : %s \n ",chain.list[ind].transaction)
		
	//taking string input new transection to be entered by user
    scanner := bufio.NewScanner(os.Stdin)
    
    
        fmt.Print("Enter modified transection : ")
        
        scanner.Scan()
        
        text := scanner.Text()

		chain.list[ind].transaction=text

		fmt.Println("modified")
	}


}


//verifying that all the blocks are in same condition or corrupted
func VerifyChain(chain *BlockChain) {
	var check = 0

	for _, val := range chain.list {
		
		hash := CalculateHash(val.transaction +strconv.Itoa(val.nonce) +val.prev_hash)

		
		if(hash!=val.block_hash){
			check=1
			break
		}

		
	}

	if check==0 {
		fmt.Printf("The chain is verified with no changes made \n")
	} else {
		fmt.Printf("changes are made in some block of chain \n")
	}

}

