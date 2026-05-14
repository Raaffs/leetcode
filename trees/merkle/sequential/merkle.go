package main

import "fmt"

func main() {
	// Example usage
	leaves := []string{"a", "b", "c", "d", "e", "f", "g","h","i"}
	root := buildMerkleTree(leaves)
	fmt.Println("Merkle Root:", root)
	fmt.Println("path for c",merkleProof(4,leaves))
}

func merkleProof(index int, leaves []string) []string {
    path := []string{}
    currentLevel := make([]string, len(leaves))
    copy(currentLevel, leaves)

    currIdx := index

    for len(currentLevel) > 1 {
        if len(currentLevel)%2 != 0 {
            currentLevel = append(currentLevel, currentLevel[len(currentLevel)-1])
        }

        var nextLvl []string
        
        siblingIdx := currIdx ^ 1
        path = append(path, currentLevel[siblingIdx])

        for i := 0; i < len(currentLevel); i += 2 {
            nextLvl = append(nextLvl, currentLevel[i]+currentLevel[i+1])
        }

        currIdx = currIdx / 2
        currentLevel = nextLvl
    }
    return path
}

func buildMerkleTree(leaves []string) string{
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
	var nextLvl []string
	for i:=0;i<len(current);i+=2{
		nextLvl = append(nextLvl, current[i]+current[i+1])
	}
	
	return  buildMerkleTree(nextLvl)
}

