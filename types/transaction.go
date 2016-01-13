package types

import (
    "time"
)

type Transaction struct {
    Id              int       `json:"id"               db:"id"`
    Amount          Money     `json:"amount"           db:"amount"`
    Description     string    `json:"description"      db:"description"`
    Incoming        bool      `json:"incoming"         db:"incoming"`
    TransactionDate time.Time `json:"transaction_date" db:"transaction_date"`
    CategoryId      int       `json:"category_id"      db:"_fk_category"`
}

type Transactions struct {
    Transactions []Transaction `json:"transactions"`
}
