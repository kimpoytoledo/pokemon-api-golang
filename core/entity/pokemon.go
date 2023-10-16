package entity

type Pokemon struct {
	Name   string `json:"name"`
	Source string `json:"source,omitempty"`
}
