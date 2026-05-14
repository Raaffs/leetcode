package main

import (
	"fmt"
	"sync"
)

func main(){
	leaves := []string{"a", "b", "c", "d", "e", "f", "g"}
	root := buildMerkleTree(leaves)
	fmt.Println("Merkle Root:", root)
}

func buildMerkleTree(leaves []string)string{
	var wg sync.WaitGroup
	if len(leaves)==0{
		return ""
	}
	if len(leaves)==1{
		return leaves[0]
	}

	current:=make([]string,len(leaves))

	copy(current,leaves)

	if len(current)%2!=0{
		current = append(current, current[len(current)-1])
	}
	nextLvl:=make([]string,len(current)/2)
	for i:=0;i<len(current);i+=2{
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			nextLvl[i/2]=hash(current[i],current[i+1])
		}()
	}
	wg.Wait()
	return buildMerkleTree(nextLvl)
}

func hash(s1,s2 string)string{
	return s1+s2
}