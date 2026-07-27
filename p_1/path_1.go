package main

import (
	"fmt"
	"errors"
)

func fuckt(a int) (int, error){
	if(a < 1){
		return 0, errors.New("FUCTORIAL < 1")
	}else if(a == 1){
		return 1, nil
	}else{
		rekurs, err := fuckt(a-1)
		return a * rekurs, err
	}
}

func main(){
	var a int
	b, err := fmt.Scanf("%d", &a)
	if(err != nil){
		fmt.Printf("ERRORS: %v\tMEANING: %d\n", err, b)
		return
	}
	result, err := fuckt(a)
	if(err != nil){
		fmt.Printf("ERRORS: %v\n", err)
		return
	}
	fmt.Printf("RESUTL: %d\n", result)
}