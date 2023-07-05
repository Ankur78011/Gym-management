package models

type Gym struct {
	Id       int    `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Owner_id int    `json:"owner_id" binding:"required"`
}
