package models

type Component struct {
	ID      int    `json:"id"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	ValID   int `json:"val_id"`
	Val string `json:"ammount"`
}

type AllowanceDeduction struct {
	Name  string
	Total int
}

type ValueComponent struct {
	ID int `json:"id"`
	EmployeeID int `json:"employee_id"`
	ComponentID int `json:"component_id"`
	CategoryID int `json:"category_id"`
	KindID int `json:"kind_id"`
	Val int `json:"val"`
}

type Ammount struct {
	ValID int `json:"val_id"`
	Ammount float64 `json:"ammount"`
}

type SallaryParam struct {
	CompanyID int `json:"company_id"`
	Year int `json:"year"`
	Month int `json:"month"`
}
