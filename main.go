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

func addEmail(emailList []string, emailToAdd string) {
	if validateEmail(emailToAdd) {
		emailExist := doEmailExists(emailList, emailToAdd)
		if emailExist {
			fmt.Println("Email is already in list")
			printHighlightedLIst(emailList, emailToAdd, "blue")
			return
		}

		emailList = append(emailList, emailToAdd)
		ioutil.WriteFile(fileName, []byte(strings.Join(emailList, "\n")), 0644)
		printHighlightedLIst(emailList, emailToAdd, "green")
	} else {
		fmt.Println("Email is not valid")
	}
}

func deleteEmail(emailList []string, email string) {
	index := -1
	for i, e := range emailList {
		if e == email {
			index = i
			break
		}
	}

	if index > -1 {
		tmpList := make([]string, len(emailList))
		copy(tmpList, emailList)

		newList := append(emailList[:index], emailList[index+1:]...)
		ioutil.WriteFile(fileName, []byte(strings.Join(newList, "\n")), 0644)
		printHighlightedLIst(tmpList, email, "red")
	} else {
		fmt.Println("Email not in list")
		printList(emailList)
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
	nrOfArgs := len(os.Args)

	// Add new email if arg
	if len(os.Args) > 1 {
		switch arg := os.Args[1]; arg {
		case "-d":
			if nrOfArgs > 2 {
				emailToRemove := os.Args[2]
				deleteEmail(emailList, emailToRemove)
			} else {
				fmt.Println("You need to provide an email")
			}
		default:
			emailToAdd := os.Args[1]
			addEmail(emailList, emailToAdd)
		}
	} else {
		printList(emailList)
	}
}
