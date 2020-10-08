package main

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const fileName = "emails"

func validateEmail(email string) bool {
	emailRegExp, _ := regexp.Compile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return emailRegExp.MatchString(email)
}

func doEmailExists(emailList []string, n string) bool {
	for _, e := range emailList {
		if n == e {
			return true
		}
	}
	return false
}

func printList(emailList []string) {
	fmt.Println("E-mail list")
	for i, email := range emailList {
		fmt.Println(i+1, ")", email)
	}
}

func printHighlightedLIst(emailList []string, selectedEmail string, textColor string) {
	colors := map[string]color.Attribute{
		"red":   color.FgHiRed,
		"blue":  color.FgHiBlue,
		"green": color.FgHiGreen,
	}

	c := colors[textColor]

	hightLight := color.New(c)

	fmt.Println("E-mail list")
	for i, email := range emailList {
		if email == selectedEmail {
			hightLight.Println(i+1, ")", email)
		} else {
			fmt.Println(i+1, ")", email)
		}
	}
}

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
				printHighlightedLIst(emailList, emailToAdd, "blue")
				return
			}

			emailList = append(emailList, emailToAdd)
			// Write new email list
			ioutil.WriteFile(fileName, []byte(strings.Join(emailList, "\n")), 0644)
			printHighlightedLIst(emailList, emailToAdd, "green")
			return
		} else {
			fmt.Println("Email is not valid")
			return
		}
	}

	// Print the list
	printList(emailList)
}
