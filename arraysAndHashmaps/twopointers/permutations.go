package main

import "fmt"

//this is cursed but it passed all the test cases on leetcode
func createHash(s string, m map[rune]int)int{
	hash:=0
	for _,c:=range s{
		if val,exist:=m[c];exist{
			hash+=int(c)*7+val
		}
	}
	return hash
}

func createS1Map(s1 string)map[rune]int{
	m:=make(map[rune]int)
	for i,c:=range s1{
		m[c]=int(c)*87*i
	}
	return m
}

func checkInclusion(s1 string, s2 string) bool {
	m:=createS1Map(s1)
	s1Hash:=createHash(s1,m)
	if len(s1)==1 && len(s2)==1{
		return s1==s2
	}
	windowSize:=len(s1)
	for i:=1;i<len(s2);i++{
		if i-1+windowSize-1<len(s2){
			s2Hash:=createHash(s2[i-1:i+windowSize-1],m)
			if s1Hash==s2Hash {return true}
		}
	}
	return false
}

//sane one
func checkInclusion2(s1,s2 string)bool{
	if len(s1)>len(s2){
		return false
	}

	count1:=[26]int{}
	count2:=[26]int{}

	for _,s:=range s1{
		count1[int(s-'a')]++
		count2[int(s-'a')]++
	}
	for i,s:=range s2{
		if count1==count2{
			fmt.Println(count1,count2)
			return true
		}
		count2[int(s-'a')]++
		count2[i-2]--
	}
	return false
}

func main() {
	s1 := "ab"
	s2 := "eidboaoo"

	fmt.Println(checkInclusion(s1,s2))
	fmt.Println(checkInclusion2(s1,s2))

}
