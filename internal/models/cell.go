package models

type Cell struct {
	ID int `json:"id"`
	PostCell
}

type PostCell struct {
	Zone       string `json:"zone"`
	Row        int    `json:"row"`
	AdressCode string `json:"adress_code" db:"adress_code"`
}

func PostCellToCell(postC PostCell) Cell {
	return Cell{
		PostCell: postC,
	}
}
