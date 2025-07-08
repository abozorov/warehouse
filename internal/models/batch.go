package models

import "time"

type Batch struct {
	ID          int       `json:"id"`   //
	Date        time.Time `json:"Date"` //
	Contact     string    `json:"contact"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	ProductName string    `json:"product_name" db:"product_name"`
	Price       float32   `json:"price"`
	Zone        string    `json:"zone"`
	Row         int       `json:"row"`
	Username    string    `json:"username"` //
	Role        string    `json:"role"`
	FullName    string    `json:"full_name" db:"full_name"`
	Active      bool      `json:"-"`
	PostBatch
}

type PostBatch struct {
	Type             string `json:"type"`                                     //
	CounterpartyName string `json:"counterparty_name" db:"counterparty_name"` //
	Article          string `json:"article"`                                  //
	Quantity         int    `json:"quantity"`                                 //
	AdressCode       string `json:"adress_code" db:"adress_code"`             //
}

func PostBatchToBatch(postB PostBatch) Batch {
	return Batch{
		PostBatch: postB,
	}
}
