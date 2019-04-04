package models

// Customer ...
type Customer struct {
	FName  string
	LName  string
	Age    int
	Gender string
	Orders []Order
}

// GetFullName ...
func (c Customer) GetFullName() string {
	return c.FName + " " + c.LName
}
