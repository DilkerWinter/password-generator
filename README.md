# Password Generator

## 📖 Desciption


A simple and customizable Password Generator created in GoLang for use in the terminal. This tool allows you to generate strong passwords with different options, such as including uppercase letters and special characters, and it also provides functionality for storing and retrieving passwords from a PostgreSQL database.

## 💻 Requirements

- **GoLang**
- **PostgreSQL**

### Instalation

1 - Clone the repository:

```git clone git@github.com:DilkerWinter/password-generator.git ```

2 - Change directory to the project

```cd password-generator ```


3 - Set up PostgreSQL database

```sql CREATE DATABASE passwordgenerator; ```

4 - Create the table

```sql
 CREATE TABLE passwords (
  id SERIAL PRIMARY KEY,
  password_name VARCHAR(255),
  password TEXT
);
```

5 - Configure the database on the app creating a database/database.go file

```go 
package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func Database() (*sql.DB, error) {
	connStr := "postgres://username:password@localhost:5432/passwordgenerator?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
```

6 - Build the project

```go build -o passwordgenerator main.go ```

7 - Now you can run the app

```./passwordgenerator```

if you want to create a command line to use the app

```sudo mv passwordgenerator /usr/local/bin/```

and verify if you have in your ~/.bashrc or ~/.zshrc the following line

```export PATH="/usr/local/bin:$PATH"```


