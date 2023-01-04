package models

import (
	"fmt"
	"net/http"
	"training-api/db"
)

type Users struct {
	UserID   int64  `json:"userID"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// this function is to list all users
func FetchUsers() (Response, error) {
	var obj Users
	var arrobj []Users
	var res Response

	//create connection with DB
	fmt.Println("Checking error creating connection DB(userModel line 22 ): ")
	conn := db.CreateCon()
	//sql statement to query data can be changed here to perform using ORM
	sqlStatement := "SELECT * FROM Users"
	//returning 2 values rows return value(if any) and err if encountered any
	rows, err := conn.Query(sqlStatement)
	defer rows.Close()
	//check if there are any errors encountered
	if err != nil {
		fmt.Println("Checking error(userModel line 32 ): ")
		return res, err
	}
	//looping for data until the end of the table
	for rows.Next() {
		//rows.Scan(&obj.userID, &obj.fullName, &obj.email, &obj.password)
		err = rows.Scan(&obj.UserID, &obj.FullName, &obj.Email, &obj.Password)
		//checking the data position if encounter any error
		if err != nil {
			fmt.Println("Checking error(userModel line 40 ): ")
			return res, err
		}
		//append the object to slices
		arrobj = append(arrobj, obj)
		//repeat the loop until the data is all retrieved
	}
	//returning the response struct with datas and no err encountered
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj
	return res, nil
}

// adding data to he database
func RegisterUser(fullName string, email string, password string) (Response, error) {

	var res Response
	conn := db.CreateCon()

	sqlStatement := "INSERT INTO Users (fullName, email, password) VALUES (?,?,?)"
	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(fullName, email, password)
	if err != nil {
		return res, err
	}

	lastInsertedId, err1 := result.LastInsertId()
	if err1 != nil {
		return res, err
	}
	fmt.Print(lastInsertedId)
	//checking using postman
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func UpdateUser(user Users) (int64, error) {
	//var res Response
	conn := db.CreateCon()

	sqlStatement := "UPDATE Users SET fullName = ?, email= ?, password= ? WHERE userID = ?"
	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		fmt.Println("Error Preparing sql statement for update lin 94 at user.models.go")
	}
	result, err := stmt.Exec(user.FullName, user.Email, user.Password, user.UserID)
	if err != nil {
		fmt.Print("ERROR UPDATING DATA TO DATABASE LINE 98 user.model.go")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Print("error at line 102 users.model.go")
	}
	return rowsAffected, err
}

func DeleteUser(uid int64) (int64, error) {

	conn := db.CreateCon()
	sqlStatement := "DELETE FROM Users WHERE `userID` = ?"
	stmt, err := conn.Prepare(sqlStatement)
	if err != nil {
		fmt.Println("Error prepare statement at line 113 user.models.go")
	}
	result, err := stmt.Exec(uid)
	if err != nil {
		fmt.Println("Error executing query with id :", uid)
	}
	row, _ := result.RowsAffected()
	return row, err
}

func SearchUser(uid int64) (Users, error) {
	//var res Response
	var obj Users
	fmt.Printf("ENTERED SEARCH USER MODEL WITH UID : %d", uid)
	conn := db.CreateCon()
	sqlStatement := "SELECT * FROM Users WHERE userID = ?"
	err := conn.QueryRow(sqlStatement, uid).Scan(&obj.UserID, &obj.FullName, &obj.Email, &obj.Password)

	fmt.Println(obj.Email)
	if err != nil {
		fmt.Print("ERROR PREPARE STATEMENT LINE 95 user.models.go")
		return obj, err
	}

	// res.Status = http.StatusOK
	// res.Message = "Success"
	// res.Data = obj
	return obj, nil
}
