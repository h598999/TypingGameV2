package internal

import (
	"fmt"
  m "typinggame/models"
)

type UserRepo struct{
  dao UserDAO
}

// Create
func (ur *UserRepo)addUser(u* m.User) (int64, error){
  query := "INSERT INTO Users (Username, password) VALUES (?, ?)"
  result, err := ur.dao.Exec(query, u.Username, u.Password)
  if err != nil{
    fmt.Print(err.Error())
    return 0, err
  }
  if result == 0{
    return 0, err
  }
  return result, nil
}

// Read
func (ur *UserRepo)getUserById(id int64) ( m.User, error ){
  query := fmt.Sprintf("SELECT Username, Password FROM Users where id = %v", id)
  us, err := ur.dao.QueryRow(query)
  if err != nil{
    fmt.Print(err.Error())
  }
  return us, nil
}

func (ur *UserRepo) getUserByUsername(Username string) (m.User, error){
  query := ("SELECT * FROM Users where Username = ?")
  us, err := ur.dao.QueryRow(query, Username)
  if err != nil{
    fmt.Print(err.Error())
    return us, err
  }
  return us, nil 
}

// Update
func (ur *UserRepo) updateUserWithId(id int64, us m.User) (int64, error){
  query := "UPDATE Users SET Username = ?, Password = ? WHERE id = ?"
  affected, err := ur.dao.Exec(query, us.Username, us.Password, id)
  if err != nil{
    fmt.Print(err.Error())
    return 0, err
  }
  return affected, nil
}

// Delete
func (ur *UserRepo) deleteUserById(id int64) (int64, error){
  query := fmt.Sprintf("DELETE FROM Users WHERE id = %v", id)
  affected, err := ur.dao.Exec(query)
  if err != nil{
    fmt.Print(err.Error())
    return 0, err
  }
  return affected, nil
}

func NewUserRepo() (UserRepo){
  return UserRepo{
    dao: NewUserDAO(),
  }
}
