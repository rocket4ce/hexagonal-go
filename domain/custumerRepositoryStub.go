package domain

type CustumerRepositoryStub struct {
	custumers []Custumer
}

func (s CustumerRepositoryStub) FindAll() ([]Custumer, error) {
	return s.custumers, nil
}

func NewCustumerRepositoryStub() CustumerRepositoryStub {
	custumers := []Custumer{
		{Id: "1", Name: "Dinko", City: "Coquimbo", ZipCode: "123", DateofBirth: "1988-07-17", Status: "1"},
		{Id: "2", Name: "Dinko2", City: "Coquimbo2", ZipCode: "12123", DateofBirth: "1989-07-17", Status: "1"},
	}
	return CustumerRepositoryStub{custumers}
}
