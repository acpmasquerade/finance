package types

type Category struct {
    Id          int    `json:"id"          db:"id"`
    Description string `json:"description" db:"description"`
}

type Categories struct {
    Categories []Category `json:"categories"`
}
