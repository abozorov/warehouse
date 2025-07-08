package models

type Counterparty struct {
	ID      int    `json:"id"`
	PostCounterparty
}

type PostCounterparty struct {
	Name    string `json:"name"`
	Contact string `json:"contact"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

func PostCounterpartyToCounterparty(postC PostCounterparty) Counterparty {
	return Counterparty{
		PostCounterparty: postC,
	}
}