package main
import (
	"fmt"
)
func main(){
	sentence:="see. this works? perfectly. fine"
	formated:=""
	for i :=0;i<len(sentence);i++{
		if i==0{
			formated+=string(sentence[i]-32)
			continue
		}
		if (sentence[i]=='.' || sentence[i]=='?') && i+2<len(sentence){
			formated+=string(sentence[i])
			formated+=string(' ')
			formated+=string(sentence[i+2]-32)
			i+=2
		}else{
			formated+=string(sentence[i])
		}
	}
	fmt.Println(formated);
}
