package main

import (
    _ "github.com/go-sql-driver/mysql"
    _ "database/sql"
    "github.com/jmoiron/sqlx"

    "finance/types"

    "encoding/json"
    "fmt"
)

func main() {
    db, err := sqlx.Open("mysql", "root:@192.168.0.20/finance")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    debits := []types.Debit{}
    debits = append(debits, types.Debit{})
    db.Select(&debits, `
        SELECT *
        FROM debits
    `)

    json, err := json.Marshal(debits)
    if err != nil {
        panic(err.Error())
        return
    }

    fmt.Println(string(json))
}
