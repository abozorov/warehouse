package models

type Storage struct {
	Zone       string  `json:"zone"`
	Row        int     `json:"row"`
	AdressCode string  `json:"adress_code" db:"adress_code"`
	Article    string  `json:"article"`
	Name       string  `json:"name"`
	Price      float32 `json:"price"`
	Quantity   int     `json:"quantity"`
}
