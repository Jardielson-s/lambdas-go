package users

import "time"

type Timestamps struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type UserBankX struct {
	Name       string        `json:"name"`
	Email      string        `json:"email"`
	Password   string        `json:"password"`
	Timestamps Timestamps    `json:"timestamps"`
	Account    []BankAccount `json:"account"`
}

type BankAccount struct {
	Account    string     `json:"account"`
	Balance    float32    `json:"balance"`
	Timestamps Timestamps `json:"timestamps"`
}
