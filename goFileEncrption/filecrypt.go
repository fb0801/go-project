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
	fmt.Println("CryptoGo")
	fmt.Println("Simple file encrypter for your day-to-day needs.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\tCryptoGo encrypt /path/to/your/file")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("")
	fmt.Println("\t encrypt\tEncrypts a file given a password")
	fmt.Println("\t decrypt\tTries to decrypt a file using a password")
	fmt.Println("\t help\t\tDisplays help text")
	fmt.Println("")

}

func encryptHandle() {
	if len(os.Args) < 3 {
		println("Missing the path to the file. For more information run CryptoGo help")
		os.Exit(0)
	}

	file := os.Args[2]

	if !validateFile(file) {
		panic("File not found")
	}

	password := getPassword()

	fmt.Println("\nEncrypting...")
	filecrypt.Encrypt(file, password)
	fmt.Println("\nFile successfully protected")
}

func decryptHandle(){

}

func getPassword() []byte {
	fmt.Print("Enter password: ")
	password, _ := terminal.ReadPassword(0)
	fmt.Print("\nConfirm password: ")
	password2, _ := terminal.ReadPassword(0)
	if !validatePassword(password, password2) {
		fmt.Print("\nPasswords do not match. Please try again.\n")
		return getPassword()
	}
	return password
}


func validatePassword(){

}

func validateFile(){
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}

	return true
}