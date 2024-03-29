package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Scanner struct {
	reader *bufio.Reader
}

func (scan *Scanner) prompt(question string) string {
	fmt.Println(question)
	answer, _ := scan.reader.ReadString('\n')
	answer = strings.TrimSpace(answer)
	return answer
}

func createScanner() Scanner {
	return Scanner{reader: bufio.NewReader(os.Stdin)}
}

func main() {

	scanner := createScanner()

	customer := Customer{
		membership: 0,
		User: User{
			name: Name{
				firstName: scanner.prompt("What is your firstName? "), lastname: scanner.prompt("What is your lastname ?")},
			age: Age{age: func() uint16 {
				res := scanner.prompt("How old are you ?")
				val, err := strconv.Atoi(res)
				if err != nil {
					panic("Age is not valid")
				}
				return uint16(val)
			}(),
			},
			auth: Authentication{
				email:    "gigi@hotmail.us",
				password: "84^27812JGdgiha&!!@87yqhed",
			},
		}}

	if !customer.age.hasLegalAge() {
		panic("You don't have the minimum legal age")
	}

	customer.greet()

program:
	for {
		letter := scanner.prompt("[q] Quit, [p] Promote, [d] Downgrade, [g] Greet")
		switch letter {
		case "q":
			fmt.Println("Quitting Program...")
			break program
		case "p":
			customer.promote()
		case "d":
			customer.downgrade()
		case "g":
			customer.greet()
		}
	}

}

type User struct {
	name Name
	age  Age
	auth Authentication
}

type Customer struct {
	membership TypeOfMember
	User
}

func (u *Customer) greet() {
	s := fmt.Sprintf("Hi, my name is %v, I am %v years old, and I am a level %v customer.", u.name.firstName, u.age.age, u.membership)

	fmt.Println(s)
}

type Name struct {
	firstName string
	lastname  string
}

type Age struct {
	age uint16
}

func (age *Age) hasLegalAge() bool {
	const LEGAL_AGE uint16 = 19

	return age.age >= LEGAL_AGE
}

type Authentication struct {
	email    string
	password string
}

func (auth *Authentication) isEmail() bool {
	reg := "/([a-zA-Z\\d.-]+)@(([a-zA-Z\\d]+).([a-z]{2,8}))/gm"
	match, err := regexp.MatchString(reg, auth.email)
	if err != nil {
		panic(err)
	}
	return match
}
func (auth *Authentication) isValidPassword() bool {
	reg := ""
	match, err := regexp.MatchString(reg, auth.password)
	if err != nil {
		panic(err)
	}
	return match
}

type TypeOfMember int

const (
	Basic TypeOfMember = iota
	Gold
	Premium
)

type Membership struct {
	membership TypeOfMember
}

func (customer *Customer) downgrade() {
	switch customer.membership {
	case Premium:
		customer.membership = TypeOfMember(Gold)
		fmt.Printf("%s have been downgraded to Gold. \n", customer.name.firstName)
	case Gold:
		customer.membership = Basic
		fmt.Printf("%s have been promoted to Basic. \n", customer.name.firstName)
	default:
		panic("Cannot downgrade")
	}
}

func (customer *Customer) promote() {
	switch customer.membership {
	case Basic:
		customer.membership = Gold
		fmt.Printf("%s have been promoted to Gold. \n", customer.name.firstName)
	case Gold:
		customer.membership = Premium
		fmt.Printf("%s have been promoted to Premium. \n", customer.name.firstName)
	default:
		panic("Cannot promote")
	}
}
