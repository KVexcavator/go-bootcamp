// передача данных структуры между пакетами
package magazine

// здесь вызов полей адреса будет как HomeAddress.City
type Subscriber struct {
	Name        string
	Rate        float64
	Active      bool
	HomeAddress Address
}

// здесь иожно напрямую обращаться к полям адреса(анонимность)
type Employee struct {
	Name   string
	Salary float64
	Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
