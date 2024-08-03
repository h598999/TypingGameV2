package internal

import (
	"testing"
	"typinggame/models"
)

var dao UserDAO

func setUp(){
  dao = NewUserDAO()
}

func tearDown(){
  dao.db.Close()
}

func TestExec(t *testing.T){
  setUp()
  defer tearDown()

  query := "INSERT INTO Users (Username, Password) VALUES (?, ?)"
  lastinsertid, err := dao.Exec(query, "Test", "Test")
  if err != nil {
    t.Fatalf("Failed to execute query %q: %v", query, err)
  }
  if lastinsertid == 0 {
    t.Fatal("Zero rows lastinsertid")
  }
  query = "DELETE FROM Users WHERE username = ?"
  lastinsertid, err = dao.Exec(query, "Test")
  if err != nil {
    t.Fatalf("Failed to execute query %q: %v", query, err)
  }
  // if lastinsertid == 0 {
  //   t.Fatal("Zero rows lastinsertid")
  // }
}

func TestQuery(t *testing.T){
  setUp()
  defer tearDown()

  query := "SELECT Username, Password FROM Users where Username = ?"
  result, err := dao.Query(query, "Jonas")
  if err != nil{
    t.Fatalf("Failed to execute query %q: %v", query, err)
  }
  if len(result) == 0{
    t.Fatalf("No elements returned")
  }
}

func TestQueryRow(t *testing.T){
  setUp()
  defer tearDown()

  query := "SELECT Username, Password FROM Users where Username = ?"
  result, err := dao.QueryRow(query, "Jonas")
  if err != nil{
    t.Fatalf("Failed to execute query %q: %v", query, err)
  }
  expected_user := models.User{
    Username: "Jonas",
    Password: "123",
  }
  if result != expected_user{
    t.Fatal("Retrieved user does not match expected user")
  }
}

func TestGetTenWords(t *testing.T){
  setUp()
  defer tearDown()

  words, err := dao.getTenWords(0)
  if err != nil{
    t.Fatalf("Error getting ten words")
  }

  if len(words) != 10{
    t.Fatal("Get ten words did not retrieve ten words")
  }
}

