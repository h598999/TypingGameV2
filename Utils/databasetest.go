package utils

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

func main() {
    // Database connection string
    dsn := "root:JOnas0909@tcp(127.0.0.1:3306)/typinggame_users"

    // Open the connection
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Error opening database:", err)
    }
    defer db.Close()

    // Check the connection
    err = db.Ping()
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }
    fmt.Println("Successfully connected to the database!")

    // Create the Users table
    readTables := "DESCRIBE Users;"

    result, err := db.Exec(readTables)
    if err != nil {
        log.Fatal("Error executing query:", err)
    }
    fmt.Println(result)
}

