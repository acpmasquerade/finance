package types

type Debit struct {
    Id          int    `json:"id"           db:"id"`
    Amount      Money  `json:"amount"       db:"amount"`
    Description string `json:"description"  db:"description"`
    OutgoingDay int    `json:"outgoing_day" db:"outgoing_day"`
    CategoryId  int    `json:"category_id"  db:"_fk_category"`
}

type Debits struct {
    Debits []Debit `json:"debits"`
}
