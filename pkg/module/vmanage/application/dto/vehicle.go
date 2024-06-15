package dto

type Vehicle struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func (v Vehicle) IsDto() {}
