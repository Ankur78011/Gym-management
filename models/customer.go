package models

type CreateCustomer struct {
	Name             string `json:"name" binding:"required"`
	Current_weight   int    `json:"current_weight" binding:"required"`
	Targeted_Weight  int    `json:"target_weight" binding:"required"`
	Personal_trainer string `json:"personal_trainer" binding:"required"`
	Gym_Name         string `json:"gym_name" binding:"required"`
}
type CustomerDetails struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	Gym_name         string `json:"gym_name"`
	Personal_trainer string `json:"personal_trainer"`
	Current_weight   int    `json:"current_weight"`
	Targeted_Weight  int    `json:"targeted_weight"`
}
