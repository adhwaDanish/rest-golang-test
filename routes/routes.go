package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"training-api/controllers"
	"training-api/models"

	"github.com/labstack/echo/v4"
)

func listUsers(c echo.Context) error {

	res, err := models.FetchUsers()
	if err != nil {
		panic("an error occured at line 26 main.go")
	}
	data := map[string]interface{}{
		"users": res.Data,
	}
	fmt.Println(data)
	return c.Render(http.StatusOK, "template", data)
}

func addUsersPage(c echo.Context) error {

	//render addData file
	return c.Render(http.StatusOK, "addData", "")
}
func searchUserPage(c echo.Context) error {

	c.Request().ParseForm()
	//fmt.Print("Button is clicked")
	id, _ := strconv.ParseInt(c.Request().Form.Get("uid"), 10, 64)
	obj, err := models.SearchUser(id)

	if err != nil {
		fmt.Print("ERROR SEARCHING USER ID AND RETURNING OBJ USER")
	}
	//double checking object carried carried
	fmt.Print(obj)

	data := map[string]interface{}{
		"user": obj,
	}
	fmt.Println(data)
	return c.Render(http.StatusOK, "searchUser", data)
}

func userUpdatePage(c echo.Context) error {

	//add functions here
	query := c.Request().URL.Query()
	uid, err := strconv.ParseInt(query.Get("uid"), 10, 64)
	if err != nil {
		fmt.Print("ERROR FETCHING UID FROM TEMPLATE.HTML Err LINE 68 main.go")
	}
	user, err := models.SearchUser(uid)
	if err != nil {
		fmt.Print("ERROR WHEN SEARCHNG USER ID AT LINE 72 main.go ")
	}
	data := map[string]interface{}{
		"user": user,
	}
	//double checking data
	fmt.Print(data)
	return c.Render(http.StatusOK, "putData", data)
}

func userDeletePage(c echo.Context) error {
	query := c.Request().URL.Query()
	uid, err := strconv.ParseInt(query.Get("uid"), 10, 64)
	fmt.Println(uid)

	//call delete function here
	res, err := models.DeleteUser(uid)
	if err != nil {
		fmt.Println("Error delete user at main.go line 88")
	}
	if res <= 1 {
		fmt.Println("Nothing is changed in db as rowsAffected return 0")
	}
	http.Redirect(c.Response(), c.Request(), "/", http.StatusSeeOther)
	return err
}

func updateUser(c echo.Context) error {

	c.Request().ParseForm()
	var user models.Users

	//filling in data from update.html form to user object
	user.UserID, _ = strconv.ParseInt(c.Request().Form.Get("uid"), 10, 64)
	user.FullName = c.Request().Form.Get("fullName")
	user.Email = c.Request().Form.Get("email")
	user.Password = c.Request().Form.Get("password")
	//calling function to update DB
	rowsAffected, err := models.UpdateUser(user)
	if err != nil {
		fmt.Print("Error updating at line 96 main.go")
	}
	if rowsAffected <= 1 {
		fmt.Printf("Nothing is updated as rowsAffected returned : %d", rowsAffected)
	}
	http.Redirect(c.Response(), c.Request(), "/", http.StatusSeeOther)

	return err
}

func addUsers(c echo.Context) error {

	c.Request().ParseForm()

	FullName := c.Request().Form.Get("fullName")
	Email := c.Request().Form.Get("email")
	Password := c.Request().Form.Get("password")
	fmt.Print(FullName + "," + Email + "," + Password)
	res, err := models.RegisterUser(FullName, Email, Password)
	if err != nil {
		fmt.Printf("ERROR WHEN TRYING TO ADD/REGISTER USERS : %d", res.Status)
	}
	http.Redirect(c.Response(), c.Request(), "/", http.StatusSeeOther)
	return err
}

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	//test list all user data
	e.GET("/allusers", controllers.FetchAllUsers)
	//list of APIs routing
	e.GET("/", listUsers)
	e.GET("/add", addUsersPage)
	e.POST("/addUsers", addUsers)
	e.POST("/searchUser", searchUserPage)
	e.GET("/update", userUpdatePage)
	e.POST("/updateUser", updateUser)
	e.GET("/delete", userDeletePage)
	return e
	//e.Logger.Fatal(e.Start(":1323"))

}
