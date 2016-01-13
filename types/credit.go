package types

type Credit struct {
    Id          int    `json:"id"           db:"id"`
    Amount      Money  `json:"amount"       db:"amount"`
    Description string `json:"description"  db:"description"`
    IncomingDay int    `json:"incoming_day" db:"incoming_day"`
    CategoryId  int    `json:"category_id"  db:"_fk_category"`
}

type Credits struct {
    Credits []Credit`json:"credits"`
}
