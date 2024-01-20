package myInterface

//===============================================================================================
//---PERTEMUAN 8---

//---Interface---

type Customer struct {
	Id      uint
	Name    string
	Address string
	Contact
}

// ---Embedded Struct---
type Contact struct {
	Phone, Email string
}

func (c *Customer) GetName() string {
	return c.Name
}

func (c *Customer) GetContact() map[string]string {
	return map[string]string{
		"Phone": c.Contact.Phone,
		"Email": c.Contact.Email,
	}
}

func (c *Customer) GetAddress() string { return c.Address }

// func (c *Customer) GetID() uint {
// 	return c.Id
// }

type Information interface {
	GetName() string
	GetContact() map[string]string
	GetAddress() string
}

type InterfaceKosong interface{}

//type InterfaceKosong any
