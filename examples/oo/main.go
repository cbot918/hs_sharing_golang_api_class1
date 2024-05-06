package main

import (
	"encoding/json"
	"fmt"
)

// data member
type User struct {
	Name  string
	Email string
}

// constructor
func NewUser(name string, email string) *User {
	return &User{
		Name:  name,
		Email: email,
	}
}

// function member
func (u *User) SetName(newName string) {
	u.Name = newName
}
func (u *User) Print() {
	json, err := json.Marshal(u)
	if err != nil {
		fmt.Println("json marshal error")
		return
	}
	fmt.Println(string(json))
}

func main() {

	user1 := NewUser("yale", "yale@gmail.com")

	user1.Print()

	user1.SetName("node")

	user1.Print()

}
