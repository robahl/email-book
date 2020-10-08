package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
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

func printHighlightedList(emailList []string, selectedEmail string, textColor string) {
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

var remove bool

func init() {
	flag.BoolVar(&remove, "remove", false, "indicates if it is a 'remove address' command")
}

func main() {
	flag.Parse()

	if err := initStorageFile(); err != nil {
		fmt.Printf("Failed to initialize storage: %v", err)
		return
	}

	// Read emails
	emails, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Could not read file", err.Error())
		return
	}

	emailList := strings.Fields(string(emails))

	// Add new email if arg
	if len(flag.Args()) > 0 {
		email := flag.Arg(0)
		if remove {
			removeEmail(emailList, email)
		} else {
			addEmail(emailList, email)
		}
		return
	}
}

func initStorageFile() error {
	// Create file if it doesn't exsists.
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		_, err := os.Create(fileName)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
	}
	return nil
}

func removeEmail(emailList []string, email string) {
	if !doEmailExists(emailList, email) {
		fmt.Println("Email is not yet in list")
	} else {
		fmt.Printf("Will remove %q fromt the list\n", email)
	}
	emailList = removeFromList(emailList, email)
	printList(emailList)
}

func removeFromList(emailList []string, email string) []string {
	out := make([]string, 0, len(emailList))
	for _, el := range emailList {
		if el != email {
			out = append(out, el)
		}
	}
	return out
}

func addEmail(emailList []string, email string) {
	if validateEmail(email) {
		emailExist := doEmailExists(emailList, email)
		if emailExist {
			fmt.Println("Email is already in list")
			printHighlightedList(emailList, email, "blue")
			return
		}

		emailList = append(emailList, email)
		// Write new email list
		ioutil.WriteFile(fileName, []byte(strings.Join(emailList, "\n")), 0644)
		printHighlightedList(emailList, email, "green")
		return
	}

	fmt.Println("Email is not valid")
	return
}
