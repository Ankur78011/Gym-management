package models

type Owner struct {
	Id    int    `json:"Id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type CreateOwner struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
