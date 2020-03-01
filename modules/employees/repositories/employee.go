package repositories

import (
	"fmt"
	"smart-hr/database"
	"smart-hr/library/logger"
	"smart-hr/modules/employees/models"
)

type EmployeeRepositories struct{}

func (r *EmployeeRepositories) GetAll() (result []models.Employee, err error) {
	query := "select id, employee_id, nik, bpjs_id, join_date, first_name, last_name, place_of_birth, birthday, address, age from employees"
	rows, err := database.DB.Query(query)
	if err != nil {
		logger.Log.Println(err)
		return result, err
	}

	employees := []models.Employee{}
	employee := models.Employee{}

	for rows.Next() {
		errScan := rows.Scan(
			&employee.ID,
			&employee.EmployeeID,
			&employee.NIK,
			&employee.BpjsID,
			&employee.JoinDate,
			&employee.FirstName,
			&employee.LastName,
			&employee.PlaceOfBirth,
			&employee.Birthday,
			&employee.Address,
			&employee.Age,
			// &employee.Sallary,
		)
		if errScan != nil {
			fmt.Println("Failed while scanning users", errScan)
			return result, errScan
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

func (r *EmployeeRepositories) Add(form models.Employee) (result string, err error) {
	fmt.Println(form)
	query := fmt.Sprintf("insert into employees (employee_id, first_name, last_name, join_date, place_of_birth, birthday,")
	query = fmt.Sprintf(" %s address, age, sallary, department_id, level_id, company_id, bpjs_id, nik) ", query)
	query = fmt.Sprintf(" %s values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)", query)
	fmt.Println(form.NIK)
	_, errExec := database.DB.Exec(query, form.EmployeeID, form.FirstName, form.LastName, form.JoinDate, form.PlaceOfBirth, form.Birthday, form.Address, form.Age, form.Sallary, form.DepartmentID, form.LevelID, form.CompanyID, form.BpjsID, form.NIK)
	if errExec != nil {
		logger.Log.Println(errExec)
		return "Error happening while add employee", errExec
	}

	return "New Employee has been added", nil
}

func (r *EmployeeRepositories) GetEmployeesByCompanyID(id int) (employees []models.Employee, err error) {
	query := "select id, employee_id, nik, bpjs_id, join_date, first_name, last_name, place_of_birth, birthday, address, age, sallary from employees where company_id = $1" 
	rows, errQ := database.DB.Query(query, id)

	results := []models.Employee{}
	result := models.Employee{}
	if errQ != nil {
		logger.Log.Println(errQ)
		return employees, errQ
	}
	
	for rows.Next() {
		errS := rows.Scan(
			&result.ID,
			&result.EmployeeID,
			&result.NIK,
			&result.BpjsID,
			&result.JoinDate,
			&result.FirstName,
			&result.LastName,
			&result.PlaceOfBirth,
			&result.Birthday,
			&result.Address,
			&result.Age,
			&result.Sallary,
		)
		if errS != nil {
			return employees, errS
		}

		results = append(results, result)
	}
	
	return results, nil
}
