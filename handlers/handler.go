package handler

import (
	"bufio"
	"fmt"
	"os"
	"password-generator/services"
	"password-generator/utils"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
)


func MainMenu () {
	for {
		utils.ClearTerminal()

		fmt.Printf("----- Password Generator ----- \n\n")
		fmt.Printf("[1] - Generate New Password\n")
		fmt.Printf("[2] - Search Password\n")
		fmt.Printf("[3] - Delete Password\n")
		fmt.Printf("\n[0] - Exit\n")
		fmt.Print("\nChoose an option: ")

		reader := bufio.NewReader(os.Stdin)
		optionStr, _, err := reader.ReadRune()

		if err != nil {
			fmt.Println("Error reading input")
			return
		}

		switch optionStr {
		case '1':
			AddPasswordMenu()
		case '2':
			SearchPasswordMenu()
		case '3':
			DeletePasswordMenu()
		case '0':
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please choose a valid option.")
		}
	}
}

func AddPasswordMenu() {
	var length int
	var err error
	utils.ClearTerminal()
	fmt.Println("----- Create Password -----")
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

	password := services.PasswordGenerator(length, includeSpecialChars, includeUpperCaseLetters, passwordName)
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
		err := services.PasswordToDatabase(password, passwordName)
		if err != nil {
			fmt.Printf("Error storing password in database: %v\n", err)
		} else {
			fmt.Println("Password stored in database successfully.")
		}
	}

	fmt.Println("\nPress any key to continue...")
	_, _ = bufio.NewReader(os.Stdin).ReadByte()
}

func SearchPasswordMenu() {
	utils.ClearTerminal()
	fmt.Printf("----- Search Password ----- \n\n")
	fmt.Printf("[1] - Search Passwords By Name\n")
	fmt.Printf("[2] - Search All Passwords\n")
	fmt.Printf("\n[0] - Go Back\n")
	fmt.Print("\nChoose an option: ")

	reader := bufio.NewReader(os.Stdin)
	optionStr, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println("Error reading input")
		return
	}

	switch optionStr {
	case '1':
		SearchPasswordByNameMenu()
	case '2':
		SearchAllPasswordsMenu()
	case '0':
		return
	default:
		fmt.Println("Invalid option. Please choose a valid option.")
	}

}


func SearchPasswordByNameMenu () {
	fmt.Printf("Insert the name of the password: ")
	var passwordNameInput string
	fmt.Scan(&passwordNameInput)

	password := services.SearchPasswordByNameDatabase(passwordNameInput)
	fmt.Println("Senha:", password)
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

	fmt.Println("Press any key to continue...")
	_, _ = bufio.NewReader(os.Stdin).ReadByte()

}

func SearchAllPasswordsMenu() {
	passwords, err := services.SearchAllPasswordsDatabase()
	if err != nil {
		fmt.Printf("Error retrieving passwords: %v\n", err)
		return
	}

	for {
		utils.ClearTerminal()

		fmt.Printf("----- List of Passwords -----\n")
		for i, password := range passwords {
			fmt.Printf("[%d] %s\n", i+1, password.Name)
		}

		fmt.Println("\n[0] Go Back")
		fmt.Printf("\nChoose a password to copy (enter the number): ")
		var choice int
		_, err := fmt.Scan(&choice)

		if err != nil || choice < 0 || choice > len(passwords) {
			fmt.Println("Invalid choice, please try again.")
			continue
		}

		if choice == 0 {
			fmt.Println("Exiting...")
			break
		}

		passwordToCopy := passwords[choice-1].Password

		fmt.Printf("Do you want to copy the password for '%s' to clipboard? (y/n): ", passwords[choice-1].Name)
		var clipboardInput string
		fmt.Scan(&clipboardInput)
		clipboardInput = strings.ToLower(clipboardInput)

		if clipboardInput == "y" || clipboardInput == "yes" {
			err := clipboard.WriteAll(passwordToCopy)
			if err != nil {
				fmt.Printf("Error copying to clipboard: %v\n", err)
			} else {
				fmt.Printf("Password for '%s' copied to clipboard!\n", passwords[choice-1].Name)
			}
		} else {
			fmt.Println("Password not copied.")
		}

		fmt.Println("Press any key to continue...")
		_, _ = bufio.NewReader(os.Stdin).ReadByte()
	}
}

func DeletePasswordMenu() {
	passwords, err := services.SearchAllPasswordsDatabase()
	if err != nil {
		fmt.Printf("Error retrieving passwords: %v\n", err)
		return
	}

	for {
		utils.ClearTerminal()

		fmt.Printf("----- Delete Password -----\n")
		for i, password := range passwords {
			fmt.Printf("[%d] %s\n", i+1, password.Name)
		}

		fmt.Println("\n[0] Go Back")
		fmt.Printf("\nChoose a password to delete (enter the number): ")
		var choice int
		_, err := fmt.Scan(&choice)

		if err != nil || choice < 0 || choice > len(passwords) {
			fmt.Println("Invalid choice, please try again.")
			continue
		}

		if choice == 0 {
			fmt.Println("Exiting...")
			break
		}

		passwordToDelete := passwords[choice-1].Name

		fmt.Printf("Are you sure you want to delete '%s'? (y/n): ", passwordToDelete)
		var confirmationInput string
		fmt.Scan(&confirmationInput)
		confirmationInput = strings.ToLower(confirmationInput)

		if confirmationInput == "y" || confirmationInput == "yes" {
			err := services.DeletePasswordByNameDatabase(passwordToDelete)
			if err != nil {
				fmt.Printf("Error deleting password: %v\n", err)
			} else {
				fmt.Printf("Password '%s' deleted successfully!\n", passwordToDelete)
			}

			passwords, err = services.SearchAllPasswordsDatabase()
			if err != nil {
				fmt.Printf("Error refreshing passwords: %v\n", err)
				return
			}
		} else {
			fmt.Println("Password not deleted.")
		}

		fmt.Println("Press any key to continue...")
		_, _ = bufio.NewReader(os.Stdin).ReadByte()
	}
}
