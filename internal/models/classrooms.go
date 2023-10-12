package models

type Classroom struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Desc string `json:"description"`
}
