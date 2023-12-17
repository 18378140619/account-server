package dto

type UserCodeDTO struct {
	Code string `json:"code" binding:"required" message:"code不能为空"`
}
