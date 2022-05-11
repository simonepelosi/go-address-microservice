package models

// {addresses: [“Address A”, “Address B”, “Address C”, "Address D"]}
type Addresses struct {
	AddressesList []string `json:"addresses"`
}
