package user

//give Input name menandakan kalo buat mapping input dari user
type RegisterUserInput struct {
	Name 		string 	`json:"name" binding:"required"`
	Occupation 	string	`json:"occupation" binding:"required"`
	Email 		string	`json:"email" binding:"required,email"`
	Password 	string	`json:"password" binding:"required"`
}

type LoginInput struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type EmailAvailabilityInput struct{
	Email string `json:"email" binding:"required,email"`
}