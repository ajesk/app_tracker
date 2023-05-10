package main

type application struct {
	ID        int    `json:"id"`
	Company   string `json:"company"` // join on company table
	User      int    `json:"user"`    // join on user table
	Title     string `json:"title"`
	MaxSalary int    `json:"max_salary"`
	MinSalary int    `json:"min_salary"`
	Notes     string `json:"notes"`
}
