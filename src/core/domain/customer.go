package domain

type Customer struct {
	ID      string `json:"id"`
	Name    string `json:"name" binding:"required" example:"Discovery" maxLength:"255"`
	Surname string `json:"surname" binding:"required" example:"Discovery" maxLength:"255"`
}

func NewCustomer(id string, name string, surname string) Customer {
	return Customer{
		ID:      id,
		Name:    name,
		Surname: surname,
	}
}
