package domain

type Customer struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func NewCustomer(id string, name string, surname string) Customer {
	return Customer{
		ID:      id,
		Name:    name,
		Surname: surname,
	}
}
