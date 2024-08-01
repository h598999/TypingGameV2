package main 

import(
  "fmt"
)

type User struct{
  id int
  Username string
  Password string
}

func (u User) Greet(){
  fmt.Printf("Hello my name is %s\n", u.Username)
}

func newUser(Username string, Password string) *User{
  return &User{
    Username:  Username,
    Password: Password,
  }

}
