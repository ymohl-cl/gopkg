package model

// User model to describe user data
/* json raws
{
	"identifier":"identifier-test",
	"pseudo":"pseudo",
	"password":"password",
	"name":"name",
	"lastName":"lastName",
	"age":42,
	"genre":"genre",
	"email":"email"
}
*/
type User struct {
	Identifier string `json:"identifier"`
	Pseudo     string `json:"pseudo" validate:"required"`
	Password   string `json:"password" validate:"required"`
	Name       string `json:"name"`
	LastName   string `json:"lastName"`
	Age        int16  `json:"age"`
	Genre      string `json:"genre"`
	Email      string `json:"email" validate:"required,email"`
}
