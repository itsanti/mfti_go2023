package main

type Contacts struct {
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type UserData struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Surname  string    `json:"surname"`
	City     string    `json:"city,omitempty"`
	Contacts *Contacts `json:"contacts,omitempty"`
}
