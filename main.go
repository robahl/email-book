package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func validateEmail(email string) bool {
	emailRegExp, _ := regexp.Compile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return emailRegExp.MatchString(email)
}

const fileName = "emails"

func main() {
	// Create file if it doesn't exsists.
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		_, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Could not create file", err)
			return
		}
	}
	// Read emails
	emails, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Could not read file", err.Error())
		return
	}
	emailList := strings.Fields(string(emails))

	// Add new email if arg
	if len(os.Args) > 1 {
		emailToAdd := os.Args[1]
		if validateEmail(emailToAdd) {

			emailExist := doEmailExists(emailList, emailToAdd)
			if emailExist {
				fmt.Println("Email is already in list")
				fmt.Printf("List: %v", emailList)
				return
			}

			emailList = append(emailList, emailToAdd)
			// Write new email list
			ioutil.WriteFile(fileName, []byte(strings.Join(emailList, "\n")), 0644)
		} else {
			fmt.Println("Email is not valid")
			return
		}
	}

	// Print the list
	fmt.Printf("List: %v", emailList)
}

func doEmailExists(emailList []string, n string) bool {
	for _, e := range emailList {
		if n == e {
			return true
		}
	}
	return false
}
