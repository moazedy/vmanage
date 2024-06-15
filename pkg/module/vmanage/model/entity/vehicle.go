package entity

type Vehicle struct {
	ID    string
	Title string
}

func (v Vehicle) IsEntity() {}
