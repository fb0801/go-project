package main

import 

func main(){
	if len(os.Args) < 2{
		printHelp()
		os.Exit(0)
	}
	function := os.Args[1]

	switch function{
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("Run CryptoGo encrypt to encrypt a file, and CryptoGo decrypt to decrypt a file.")
		os.Exit(1)
	}
}

func printHelp(){

}