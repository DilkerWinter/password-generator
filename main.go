package main

import (
	"bufio"
	"crypto/rand"
	"database/sql"
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

func searchPasswordByNameDatabase(passwordName string) string {
	db, err := database.Database()
	if err != nil {
		return fmt.Sprintf("Could not connect to database: %v", err)
	}
	defer db.Close()

	query := "SELECT password FROM passwords WHERE password_name = $1"
	row := db.QueryRow(query, passwordName)

	var password string
	err = row.Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Sprintf("No password found with name: %v", passwordName)
		}
		return fmt.Sprintf("Could not retrieve password: %v", err)
	}

	return password
}



func searchPasswordByName () {
	fmt.Printf("Insert the name of the password: ")
	var passwordNameInput string
	fmt.Scan(&passwordNameInput)

	password := searchPasswordByNameDatabase(passwordNameInput)
	fmt.Println("Senha: %s\n", password)
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

func searchAllPasswordsDatabase() ([]struct {
    Name     string
    Password string
}, error) {
	db, err := database.Database()
	if err != nil {
		return nil, fmt.Errorf("Could not connect to database: %v", err)
	}
	defer db.Close()

	query := "SELECT password_name, password FROM passwords"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve passwords: %v", err)
	}
	defer rows.Close()

	var passwords []struct {
		Name     string
		Password string
	}

	for rows.Next() {
		var passwordName, password string
		if err := rows.Scan(&passwordName, &password); err != nil {
			return nil, fmt.Errorf("could not scan row: %v", err)
		}
		passwords = append(passwords, struct {
			Name     string
			Password string
		}{
			Name:     passwordName,
			Password: password,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return passwords, nil
}


func searchAllPasswords() {
	passwords, err := searchAllPasswordsDatabase()
	if err != nil {
		fmt.Printf("Error retrieving passwords: %v\n", err)
		return
	}

	for {
		clearTerminal()

		fmt.Printf("\n---- List of Passwords ----\n")
		for i, password := range passwords {
			fmt.Printf("[%d] %s: %s\n", i+1, password.Name, password.Password)
		}

		fmt.Println("\n[0] Exit")
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



func searchPasswordMenu() {
	clearTerminal()
	fmt.Printf("----Password Generator---- \n\n")
	fmt.Printf("[1] - Search Passwords By Name\n")
	fmt.Printf("[2] - Search All Passwords\n")
	fmt.Printf("[0] - Go Back\n")
	fmt.Print("Choose an option: ")

	reader := bufio.NewReader(os.Stdin)
	optionStr, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println("Error reading input")
		return
	}

	switch optionStr {
	case '1':
		searchPasswordByName()
	case '2':
		searchAllPasswords()
	case '0':
		return
	default:
		fmt.Println("Invalid option. Please choose a valid option.")
	}

}

func main() {

	for {
		clearTerminal()

		fmt.Printf("----Password Generator---- \n\n")
		fmt.Printf("[1] - Generate New Password\n")
		fmt.Printf("[2] - Search Password\n")
		fmt.Printf("[3] - Exit\n")
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
			searchPasswordMenu()
		case '3':
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please choose a valid option.")
		}
	}
}
