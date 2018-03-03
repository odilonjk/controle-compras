package main

type Purchase struct {
	ID    int32   `json:"id,omitempty"`
	Price float32 `json:"price,omitempty"`
	Name  string  `json:"name,omitempty"`
}
