package models

type User struct {
	Name         string `json:"name"`
	PhoneNumber  string `json:"number"`
	EmailAddress string `json:"email"`
	CollegeName  string `json:"college"`
	UserId       string `json:"userid"`
}
