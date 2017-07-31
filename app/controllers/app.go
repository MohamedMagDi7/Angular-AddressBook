package controllers

import (
	"github.com/revel/revel"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"github.com/gocql/gocql"
	."Angular-Revel-App/app/models"
	"Angular-Revel-App/app"
)

type App struct {
	*revel.Controller

}


func (c App) Index() revel.Result {


	return c.Redirect("/app/index.html")
}



///////////////////////////////////////////////////////////////////////////////////////////////////////
func (c App) Signin() revel.Result{
	var userLogins UserLogins
	c.Params.Bind(&userLogins , "UserLogins")
	fmt.Println(userLogins)
		response := Responseobject{In: false , Error:"" , Userdata:UserContancts{}}
		var databasePassword string
		databasePassword, err := userLogins.QueryUser(app.DB)
		if err == gocql.ErrNotFound {
			//no such user
			response.Error ="Username doesn't exist"
			return c.RenderJSON(response)

		} else if  err != nil {

			response.Error ="Internal Server Error please try again"

			return c.RenderJSON(response)
		}

		err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(userLogins.Password))
		// If wrong password redirect to the login
		if  err != nil {
			//Wrong Password
			response.Error ="wrong password"
			return c.RenderJSON(response)
		} else {
			// If the login succeeded
			c.Session["user"]= userLogins.Username
			response.In = true
			return c.Redirect("/user")
		}
	}
///////////////////////////////////////////////////////////////////////////////////////////////////////
func (c App) Signup() revel.Result{
	var userLogins UserLogins
	c.Params.Bind(&userLogins , "userLogins")
	response := Responseobject{In: false , Error:"" , Userdata:UserContancts{}}
	err :=userLogins.CheckUsernameExists(app.DB)
	switch {
	case err == nil:
		response.Error = "Please choose a different username"
		return c.RenderJSON( response )

	case err == gocql.ErrNotFound :
			// Username is available
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userLogins.Password), bcrypt.DefaultCost)
			if err != nil {
				response.Error ="This Password is Not premitted"
				return c.RenderJSON( response )
			}

			err = userLogins.InsertUser(hashedPassword , app.DB)
			if err != nil {
				response.Error=err.Error()
				return c.RenderJSON(response)
			}
			c.Session["user"]= userLogins.Username
			response.In=true
			return c.Redirect("/user")

	case err != nil:
			//Database Error
			response.Error ="Internal Server Error please try again"
			return c.RenderJSON(response)

	default:
			return c.RenderJSON(response)
	}
}
///////////////////////////////////////////////////////////////////////////////////////////////////////
