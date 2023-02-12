package main

import (
	"fmt"
	"github.com/DimitarPetrov/stegify/steg"
)

func main(){
	// 0x45 0x44
	// 0x22 0x33 .....0x45 0x44 0xff
	
	encodedFile := "encoded_hulk.jpg" // Encoded File Name
	ResultFile := "captain-america.jpg"  // Desired Result File Name
	
	err := steg.DecodeByFileNames(encodedFile, ResultFile)
	if err != nil{
		fmt.Errorf("[-] Can't Decode The File")
	}else{
		fmt.Println("[*] File Has Been Succesfully Decoded")
	}
}

