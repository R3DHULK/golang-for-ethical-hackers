package main

import (
	"fmt"
	"os"
	"github.com/DimitarPetrov/stegify/steg"
)

func main(){
	// 0x45 0x44
	// 0x22 0x33 .....0x45 0x44 0xff
	
	carrierFile := "hulk.jpg"   // File Name Where You Use To Hide
	MaliciousFile := "captain.jpg"    // File Name What You Use To Hide
	encodedFile := "encoded_hulk.jpg"   // Result
	
	err := steg.EncodeByFileNames(carrierFile, MaliciousFile, encodedFile)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}else{
		fmt.Println("***File Has Been Succesfully Encoded***")
	}
}
