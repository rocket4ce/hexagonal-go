package domain

type Custumer struct {
	Id          string
	Name        string
	City        string
	ZipCode     string
	DateofBirth string
	Status      string
}

type CustumerRepository interface {
	FindAll() ([]Custumer, error)
}
