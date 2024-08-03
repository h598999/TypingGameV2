package internal

import (
	"database/sql"
	"fmt"
  m "typinggame/models"
	"github.com/go-sql-driver/mysql"
)

type UserDAO struct{
  db *sql.DB
}

type Word struct{
  id int64
  word string
}


func (u* UserDAO) getTenWords(startid int64) ([]string, error){
  var words []string
  var err error

  query := fmt.Sprintf("SELECT * FROM words WHERE Id >= %v AND Id <= %v", startid, startid+10)
  stmt, err := u.db.Prepare(query)
  if err != nil{
    fmt.Print("Error preparing query: ", err)
    return words, err
  }
  defer stmt.Close()

  rows, err := u.db.Query(query)
  if err != nil{
    return nil, err
  }
  for rows.Next(){
  defer rows.Close()
  var word Word
    if err := rows.Scan(&word.id, &word.word); err != nil{
      return nil, err
    }
    words = append(words, word.word)
  }
  return words ,nil

}

func (u *UserDAO)Query(query string, args ...interface{}) ([]m.User, error){
  var users []m.User
  var err error

  stmt, err := u.db.Prepare(query)
  if err != nil{
    fmt.Print("Error preparing query: ", err)
    return nil, err
  }
  defer stmt.Close()
  
  rows, err := stmt.Query(args...)
  if err != nil{
    return nil, err
  }
  defer rows.Close()

  for rows.Next(){
    var us m.User
    // var id int64
    if err := rows.Scan(&us.Username, &us.Password); err != nil{
      return nil, err
    }
    users = append(users, us)
  }
  return users, nil 
}

func (u *UserDAO)Exec(query string, args ...interface{}) (int64, error){
  tx, err := u.db.Begin()
  if err != nil{
    fmt.Print("Error starting transaction")
    return 0, err
  }

  //Preparing the query
  stmt, err := tx.Prepare(query)
  if err != nil{
    fmt.Print("Error preparing query: ", err)
    tx.Rollback()
    return 0, err
  }
  defer stmt.Close()

  // Execute the query
  result, err := stmt.Exec(args...)
  if err != nil{
    fmt.Print("Error executing query")
    tx.Rollback()
    return 0, err
  }
  
  err = tx.Commit()
  if err != nil{
    fmt.Print("Error commiting transaction: ", err)
    return 0, err
  }

  // Get the id of the executed query
  affected, err := result.LastInsertId()
  if err != nil{
    return 0, err
  }
  return affected, nil
}

func (u *UserDAO) QueryRow(query string, args ...interface{}) (m.User, error){
  var us m.User

  tx, err := u.db.Begin()
  if err != nil{
    fmt.Print("Error staring transaction ", err)
    return us, err
  }

  stmt, err := tx.Prepare(query)
  if err != nil{
    fmt.Print("Error preparing query: ", err)
    tx.Rollback()
    return us, err
  }
  defer stmt.Close()

  row := stmt.QueryRow(args...)
  if err := row.Scan(&us.Username, &us.Password); err != nil{
    if err == sql.ErrNoRows{
      // fmt.Errorf("No return value %u %e", us, err)
      tx.Rollback()
      return us, err
    }
    // fmt.Errorf("%u %e", us, err)
    tx.Rollback()
    return us, err
  }
  err = tx.Commit()
  if err != nil{
    fmt.Print("Error commiting transaction: ", err)
    return us, err
  }
  return us, nil
}

func (u *UserDAO) TestConn(){
  pingErr := u.db.Ping()
  if pingErr != nil{
    fmt.Printf("Err: %v \n", pingErr)
    return 
  }
  fmt.Print("Connected\n")
}

func NewUserDAO() (UserDAO) {
  cfg := mysql.Config{
    User: "root",
    Passwd: "JOnas0909",
    Net: "tcp",
    Addr: "127.0.0.1",
    DBName: "typinggame_users",
  }
  db, err := sql.Open("mysql", cfg.FormatDSN())
  if err != nil{
    // fmt.Print("Could not connect to database: ", err)
    return UserDAO{} 
  }
  pingErr := db.Ping()
  if pingErr != nil{
    // fmt.Print(pingErr)
  }
  dao := UserDAO{
    db: db,
  }
  return dao 
}
