package models

import ComponentModel "smart-hr/modules/components/models"

type Employee struct {
	ID 				int 			`json:"id"`
	EmployeeID 		string 			`json:"employee_id"`
	NIK				string			`json:"nik"`
	BpjsID			string			`json:"bpjs_id"`
	JoinDate 		string 			`json:"join_date"`
	FirstName 		string 			`json:"fname"`
	LastName 		string 			`json:"lname"`
	PlaceOfBirth 	string 			`json:"place_of_birth"`
	Birthday 		string 			`json:"birthday"`
	Address 		string 			`json:"address"`
	Age  			int 			`json:"age"`
	Sallary 		int				`json:"sallary"`
	DepartmentID    int     		`json:"department_id"`
	LevelID    		int     		`json:"level_id"`
	CompanyID    	int  			`json:"company_id"`
	Component		[]ComponentModel.Component  `json:"component"`
}