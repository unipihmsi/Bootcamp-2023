package query

import (
	"database/sql"
	"fmt"
)

type Customers struct {
	Id    uint
	Name  string
	Phone string
	Email string
	Age   uint
}

type DB struct {
	Conn *sql.DB
}

// Implement Create customer
func (db *DB) Create(c *Customers) error {
	stmt, err := db.Conn.Prepare(`INSERT INTO customers (name, phone, email, age) VALUES ($1,$2,$3,$4)`)
	if err != nil {
		return err
	}
	_, errExec := stmt.Exec(c.Name, c.Phone, c.Email, c.Age)
	if err != nil {
		fmt.Println("errExec", errExec)
		return errExec
	}
	return nil
}

// Implement Read customer
func (db *DB) Read() (*[]Customers, error) {
	rows, errExec := db.Conn.Query(`SELECT * FROM customers`)
	if errExec != nil {
		return nil, errExec
	}
	var result []Customers
	for rows.Next() {
		var each Customers

		errScan := rows.Scan(&each.Id, &each.Name, &each.Phone, &each.Email, &each.Age)
		if errScan != nil {
			return nil, errScan
		}
		result = append(result, each)
	}

	return &result, nil
}

// Implement Update customer
func (db *DB) Update(c *Customers) error {
	stmt, err := db.Conn.Prepare(`UPDATE customers SET name = $1, phone = $2, email = $3, age = $4 WHERE id = $5`)
	if err != nil {
		return err
	}
	_, errExec := stmt.Exec(c.Name, c.Phone, c.Email, c.Age, c.Id)
	if err != nil {
		return errExec
	}
	return nil
}

// Implement Delete customer
func (db *DB) Delete(Id int) error {
	_, errExec := db.Conn.Exec(`DELETE FROM customers WHERE id = $1`, Id)
	if errExec != nil {
		return errExec
	}
	return nil
}
