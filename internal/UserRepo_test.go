package internal

import (
	"testing"
	m "typinggame/models"
)

var rep UserRepo

func beforeEach(){
  rep = NewUserRepo()
}

func afterEach(){
  rep.dao.db.Close()
}

func TestAddUser(t *testing.T){
  beforeEach()
  defer afterEach()
  newUser := m.User{
    Username: "Test",
    Password: "Test",
  }
  result, err := rep.addUser(&newUser)
  if err != nil{
    t.Fatalf("Error inserting user %v %q", newUser, err)
  }
  if result == 0{
    t.Fatal("Result is 0")
  }

  fetchedUser, err := rep.getUserById(result)
  if err != nil{
    t.Fatal("Error fetching inserted user ", err)
  }
  if (!fetchedUser.Equal(newUser)){
    t.Fatal("Expected user does not match new user")
  }
  _, err = rep.deleteUserById(result)
  if err != nil{
    t.Fatal("Could not delete inserted user")
  }
}

func TestGetUserById(t *testing.T){
  beforeEach()
  defer afterEach()

  newUser := m.User{
    Username: "Test",
    Password: "Test",
  }
  result, err := rep.addUser(&newUser)
  if err != nil{
    t.Fatalf("Error inserting user %v %q", newUser, err)
  }

  fetchedUser, err := rep.getUserById(result)
  if err != nil{
    t.Fatal("Error fetching user", err)
  }
  if (!fetchedUser.Equal(newUser)){
    t.Fatalf("Fetched user did not match expected user")
  }

  _, err = rep.deleteUserById(result)
  if err != nil{
    t.Fatal("Could not delete inserted user")
  }
}

func TestUpdateUserById(t *testing.T){
  beforeEach()
  defer afterEach()
  newUser := m.User{
    Username: "Test",
    Password: "Test",
  }
  result, err := rep.addUser(&newUser)
  if err != nil{
    t.Fatalf("Error inserting user %v %q", newUser, err)
  }

  fetchedUser, err := rep.getUserById(result)
  if err != nil{
    t.Fatal("Error fetching user", err)
  }
  if (!fetchedUser.Equal(newUser)){
    t.Fatalf("Fetched user did not match expected user")
  }
  updatedUser := m.User{
    Username: "Testupdated",
    Password: "Testupdated",
  }
  _, err = rep.updateUserWithId(result, updatedUser)
  if err != nil{
    t.Fatal("Could not update user: ", err)
  }
  fetchedUser, err = rep.getUserById(result)
  if err != nil{
    t.Fatal("Error fetching user", err)
  }
  if (!fetchedUser.Equal(updatedUser)){
    t.Log(fetchedUser.Username)
    t.Fatalf("Fetched user did not match expected user")
  }

  _, err = rep.deleteUserById(result)
  if err != nil{
    t.Fatal("Could not delete inserted user")
  }
}

func DeleteUserById(t *testing.T){
  newUser := m.User{
    Username: "Test",
    Password: "Test",
  }
  result, err := rep.addUser(&newUser)
  if err != nil{
    t.Fatalf("Error inserting user %v %q", newUser, err)
  }

  _, err = rep.deleteUserById(result)
  if err != nil{
    t.Fatal("Could not delete inserted user")
  }

  _, err = rep.getUserById(result)
  if err == nil{
    t.Fatal("Error fetching user", err)
  }
}


