package repositories

import(
	"smart-hr/database"
	// CompanyModel "hr/modules/companies/models"
	"smart-hr/modules/companies/models"
	UserModel "smart-hr/modules/users/models"
	"fmt"
	"smart-hr/library/logger"
	"smart-hr/library/jwt"
)

type UserRepositories struct{}

type loginData struct{
	ID 			int		`json:"id"`
	Username 	string  `json:"username"`
	Name 		string  `json:"name"`
	Email 		string  `json:"email"`
	RoleID 		int     `json:"roleID"`
	IsLogin     bool    `json:"is_login"`
}

func(r *UserRepositories) GetAll()(result []UserModel.User, err error){
	query := "select username, u.email, r.name, c.office_name from users u join role r on u.role_id = r.id join company c on c.id = u.company_id  "
	rows, err := database.DB.Query(query)
	if err != nil {
		logger.Log.Println(err)
		return result, err
	}

	users := []UserModel.User{}
	user := UserModel.User{}


	for rows.Next(){
		errScan := rows.Scan(
			&user.Username,
			&user.Email,
			&user.Role,
			&user.Company,
		)
		if errScan != nil {
			fmt.Println("Failed while scanning users", errScan) 
			return result, err
		}

		users = append(users, user)
	}
	
	return users, nil
}

func(r *UserRepositories) Login(form UserModel.LoginForm)(user loginData , err error){
	query := "select id, username, name, email, role_id ,password from companies where username = $1"
	var hPassword string
	userData := loginData{}
	row := database.DB.QueryRow(query, form.Username)
	err = row.Scan(&userData.ID, &userData.Username, &userData.Name, &userData.Email, &userData.RoleID, &hPassword)
	fmt.Println(hPassword)
	if err != nil {
		logger.Log.Println(err)
		return userData, err
	}

	//cocokan password bcrypt
	B:= jwt.JWT{}
	errMatch := B.ComparePassword(form.Password, hPassword)
	if errMatch != nil {
		logger.Log.Println("[Wrong Password Bcrypt doesnt match]", errMatch)
		return userData, errMatch
	} else {
		queryLogin := "update is_login set login_status = 1 where user_id = $1"
		_, errUpdateLogin := database.DB.Exec(queryLogin, userData.ID)
		if err != nil {
			logger.Log.Println(errUpdateLogin)
			return userData , errUpdateLogin
		} else {
			userData.IsLogin = true
			return userData, nil
		}
	}
}

func (r *UserRepositories) Logout(userID int) (result string, err error) {
	queryLogout := "update is_login set login_status = 0 where user_id = $1"
	fmt.Println("masuk habis query")
	fmt.Println(userID)

	_, errQ := database.DB.Exec(queryLogout, userID)
	if errQ != nil {
		logger.Log.Println(errQ)
	}

	return "Logout successfully", nil
}

func(r *UserRepositories) Add(form models.Company)(result string, err error){
	bC := jwt.JWT{}
	hashPwd, err := bC.HashPassword(form.Password)

	query := fmt.Sprintf("insert into companies (name, address, email, username, password, phone, role_id) values ($1 ,$2, $3, $4, $5, $6, $7);")
	_, err = database.DB.Exec(query, form.Name, form.Address, form.Email, form.Username, hashPwd, form.Phone, form.RoleID)

	if err != nil {
		logger.Log.Println(err)
		return "Error, please check your input", err
	}

	return "New client has been created", nil

}



