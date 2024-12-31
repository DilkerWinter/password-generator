package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"password-generator/database"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
)

func passwordToDatabase(password string, passwordName string) error {
	db, err := database.Database()
	if err != nil {
		return fmt.Errorf("Could not connect to database : %v", err)
	}
	defer db.Close()

	query := "INSERT INTO passwords (password_name, password) VALUES ($1, $2)"
	_, err = db.Exec(query, passwordName, password)
	if err != nil{
		return fmt.Errorf("could not insert password: %v", err)
	}

	return nil
}

func addPassword() {
	var length int
	var err error

	for {
		fmt.Printf("\n\nEnter the desired length for the password (number only): ")
		var lengthInput string
		fmt.Scan(&lengthInput)

		length, err = strconv.Atoi(lengthInput)
		if err != nil || length <= 0 {
			fmt.Println("Invalid length. Please enter a valid number greater than 0.")
		} else {
			break
		}
	}

	var includeUpperCaseLetters bool
	for {
		fmt.Printf("Should the password include uppercase letters? (y/n): ")
		var upperCaseLetters string
		fmt.Scan(&upperCaseLetters)

		upperCaseLetters = strings.ToLower(upperCaseLetters)

		if upperCaseLetters == "y" || upperCaseLetters == "yes" {
			includeUpperCaseLetters = true
			break
		} else if upperCaseLetters == "n" || upperCaseLetters == "no" {
			includeUpperCaseLetters = false
			break
		} else {
			fmt.Println("Invalid input, please enter 'y', 'yes', 'n' or 'no'.")
		}
	}

	var includeSpecialChars bool
	for {
		fmt.Printf("Should the password include special characters? (y/n): ")
		var specialCharsInput string
		fmt.Scan(&specialCharsInput)

		specialCharsInput = strings.ToLower(specialCharsInput)

		if specialCharsInput == "y" || specialCharsInput == "yes" {
			includeSpecialChars = true
			break
		} else if specialCharsInput == "n" || specialCharsInput == "no" {
			includeSpecialChars = false
			break
		} else {
			fmt.Println("Invalid input, please enter 'y', 'yes', 'n' or 'no'.")
		}
	}

	var passwordName string
	for {
		fmt.Printf("What is the purpose or name for this password? ")
		fmt.Scan(&passwordName)

		if len(passwordName) > 0 {
			break
		} else {
			fmt.Println("Password name cannot be empty. Please enter a valid name.")
		}
	}

	password := passwordGenerator(length, includeSpecialChars, includeUpperCaseLetters, passwordName)
	fmt.Printf("Generated Password: %s\n", password)

	var copyToClipboard bool
	for {
		fmt.Printf("Wish to copy the password to clipboard? (y/n): ")
		var clipboardInput string
		fmt.Scan(&clipboardInput)

		clipboardInput = strings.ToLower(clipboardInput)

		if clipboardInput == "y" || clipboardInput == "yes" {
			copyToClipboard = true
			break
		} else if clipboardInput == "n" || clipboardInput == "no" {
			copyToClipboard = false
			break
		} else {
			fmt.Println("Invalid input, please enter 'y', 'yes', 'n' or 'no'.\n")
		}
	}

	if copyToClipboard {
		err := clipboard.WriteAll(password)
		if err != nil {
			fmt.Printf("Error copying to clipboard:\n", err)
			return
		}
		fmt.Printf("Text copied to clipboard\n")
	}

	var storePassword bool
	for {
		fmt.Printf("Want to store that Password? (y/n): ")
		var storePasswordInput string
		fmt.Scan(&storePasswordInput)

		storePasswordInput = strings.ToLower(storePasswordInput)

		if storePasswordInput == "y" || storePasswordInput == "yes" {
			storePassword = true
			break
		} else if storePasswordInput == "n" || storePasswordInput == "no" {
			storePassword = false
			break
		} else {
			fmt.Println("Invalid input, please enter 'y', 'yes', 'n' or 'no'.\n")
		}
	}

	if storePassword {
		err := passwordToDatabase(password, passwordName)
		if err != nil {
			fmt.Printf("Error storing password in database: %v\n", err)
		} else {
			fmt.Println("Password stored in database successfully.")
		}
	}

	fmt.Println("\nPress any key to continue...")
	_, _ = bufio.NewReader(os.Stdin).ReadByte()
}

func passwordGenerator(length int,includeSpecialChar bool,includeUpperCaseLetters bool, passwordName string) string {
	
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	numbers := "0123456789"
	specialChars := "!@#$%^&*()-_=+[]{}|;:,.<>?/~"

	characters := lowercase + numbers
	if includeSpecialChar {
		characters += specialChars
	}
	if includeUpperCaseLetters {
		characters += uppercase
	}

	var password []rune
	for i := 0; i < length; i++ {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		password = append(password, rune(characters[randomIndex.Int64()]))
	} 

	return string(password)

}

func clearTerminal() {
	cmd := exec.Command("clear")
	if os.PathSeparator == '\\' {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {

	for {
		clearTerminal()

		fmt.Printf("----Password Generator---- \n\n")
		fmt.Printf("[1] Generate New Password\n")
		fmt.Printf("[2] Search Password\n")
		fmt.Printf("[3] Exit\n")
		fmt.Print("Choose an option: ")

		reader := bufio.NewReader(os.Stdin)
		optionStr, _, err := reader.ReadRune()

		if err != nil {
			fmt.Println("Error reading input")
			return
		}

		switch optionStr {
		case '1':
			addPassword()
		case '2':
			fmt.Println("\nYou selected: Search Password")
		case '3':
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please choose a valid option.")
		}
	}
}
