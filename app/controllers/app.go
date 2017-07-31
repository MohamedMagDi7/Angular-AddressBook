package controllers

import (
	"github.com/revel/revel"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"github.com/gocql/gocql"
	"Angular-Revel-App/app"
	"Angular-Revel-App/app/models"
)

type App struct {
	*revel.Controller

}


func (c App) Index() revel.Result {


	return c.Redirect("/app/index.html")
}



///////////////////////////////////////////////////////////////////////////////////////////////////////
func (c App) Signin() revel.Result{
	var user models.User
	c.Params.Bind(&user.Logins , "UserLogins")

		response := models.Responseobject{In: false , Error:"" , Userdata:models.User{}}
		var databasePassword string
		databasePassword, err := user.QueryUser(app.DB)
		if err == gocql.ErrNotFound {
			//no such user
			response.Error ="Username doesn't exist"
			return c.RenderJSON(response)

		} else if  err != nil {

			response.Error ="Internal Server Error please try again"

			return c.RenderJSON(response)
		}

		err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(user.Logins.Password))
		// If wrong password redirect to the login
		if  err != nil {
			//Wrong Password
			response.Error ="wrong password"
			return c.RenderJSON(response)
		} else {
			// If the login succeeded
			c.Session["user"]= user.Logins.Username
			response.In = true
			return c.Redirect("/user")
		}
	}
///////////////////////////////////////////////////////////////////////////////////////////////////////
func (c App) Signup() revel.Result{
	var user models.User
	c.Params.Bind(&user.Logins , "userLogins")
	response := models.Responseobject{In: false , Error:"" , Userdata:models.User{}}
	err :=user.CheckUsernameExists(app.DB)
	switch {
	case err == nil:
		response.Error = "Please choose a different username"
		return c.RenderJSON( response )

	case err == gocql.ErrNotFound :
			// Username is available
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Logins.Password), bcrypt.DefaultCost)
			if err != nil {
				response.Error ="This Password is Not premitted"
				return c.RenderJSON( response )
			}

			err = user.InsertUser(hashedPassword , app.DB)
			if err != nil {
				response.Error=err.Error()
				return c.RenderJSON(response)
			}
			c.Session["user"]= user.Logins.Username
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
