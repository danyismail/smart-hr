package repositories

import(
	"smart-hr/database"
	"smart-hr/library/logger"
	PasswordGenerator "smart-hr/library/jwt"
	"smart-hr/modules/companies/models"
)

type CompanyRepositories struct {}

func (cr *CompanyRepositories) AddCompany(company models.Company) (result string, err error) {
	hashPwd := PasswordGenerator.JWT{}
	newHashPassword, errNewPassword := hashPwd.HashPassword(company.Password)
	if errNewPassword != nil {
		logger.Log.Println(errNewPassword)
	}
	query := "insert into companies (name, address, email, username, password, phone, role_id) values ($1, $2, $3, $4, $5, $6, $7)"
	_, err = database.DB.Exec(query, company.Name, company.Address, company.Email, company.Username, newHashPassword, company.Phone, company.RoleID)
	if err != nil {
		logger.Log.Println(err)
		return "Error happening", err
	}

	return "New company has been added", nil
}

func (cr *CompanyRepositories) GetAll() (company []models.Company, err error) {
	query := "select id, name, address, email, username, password, phone, role_id from companies"
	rows, errQuery := database.DB.Query(query)
	if errQuery != nil {
		logger.Log.Println(errQuery)
		return company, errQuery
	}
	var result models.Company
	var arrResult []models.Company
	for rows.Next(){
		errScan := rows.Scan(
			&result.ID,
			&result.Name,
			&result.Address,
			&result.Email,
			&result.Username,
			&result.Password,
			&result.Phone,
			&result.RoleID,
		)

		if errScan != nil {
			logger.Log.Println(errScan)
			return company, errScan
		}

		arrResult = append(arrResult, result)
	}

	return arrResult, nil
}



