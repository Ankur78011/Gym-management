package models

type Report struct {
	Customer_id int    `json:"customer_id"`
	Month       string `json:"month"`
	Year        int    `json:"year"`
	Weight      int    `json:"weight"`
}

type NewReport struct {
	Weight int `json:"weight" binding:"required"`
}
