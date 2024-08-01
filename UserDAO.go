package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type UserDAO struct{
  db *sql.DB
}

func (u *UserDAO) query (query string) ([]User, error){
  var users []User
  var err error
  
  rows, err := u.db.Query(query)
  if err != nil{
    return nil, err
  }
  for rows.Next(){
  defer rows.Close()
    var us User
    if err := rows.Scan(&us.Username, &us.Password); err != nil{
      return nil, err
    }
    users = append(users, us)
  }
  return users, err
}

func (u *UserDAO) exec (query string) (int64, error){
  result, err := u.db.Exec(query)
  if err != nil{
    return 0, err
  }
  id, err := result.LastInsertId()
  if err != nil{
    return 0, err
  }
  return id, nil
}

func (u *UserDAO) queryRow(query string) (User, error){
  var us User
  row := u.db.QueryRow(query)
  if err := row.Scan(&us.Username, &us.Password); err != nil{
    if err == sql.ErrNoRows{
      // fmt.Errorf("No return value %u %e", us, err)
      return us, err
    }
    // fmt.Errorf("%u %e", us, err)
    return us, err
  }
  return us, nil
}

func newUserDAO() (UserDAO, error) {
  cfg := mysql.Config{
    User: "root",
    Passwd: "JOnas0909",
    Net: "tcp",
    Addr: "127.0.0.1",
    DBName: "typinggame_users",
  }
  db, err := sql.Open("mysql", cfg.FormatDSN())
  defer db.Close()
  if err != nil{
    // fmt.Print("Could not connect to database: ", err)
    return UserDAO{}, err
  }
  pingErr := db.Ping()
  if pingErr != nil{
    // fmt.Print(pingErr)
  }
  dao := UserDAO{
    db: db,
  }
  return dao, err
}
