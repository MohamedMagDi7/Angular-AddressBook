	package controllers

import (
	"github.com/revel/revel"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"github.com/gocql/gocql"
	"io/ioutil"
	"encoding/json"
)
type Responseobject struct {
	In bool
	Error string
	Userdata UserContancts
	}
type App struct {
	*revel.Controller
	db *gocql.Session

}

type Userlogins struct{
	username string
	password string
}
func (c App) Index() revel.Result {


	return c.Redirect("/app/index.html")
}


func (userinfo * Userlogins) CheckUsernameExists( db *gocql.Session) error{
	var databasePassword string

	err := db.Query("SELECT password FROM user_logins WHERE username=?", userinfo.username).Scan( &databasePassword)
	fmt.Println(err)
	return err
}

func (userinfo * Userlogins) QueryUser(db *gocql.Session) (string,error){
	var databasePassword string

	err := db.Query("SELECT password FROM user_logins WHERE username=?", userinfo.username).Scan( &databasePassword)
	return databasePassword,err
}
///////////////////////////////////////////////////////////////////////////////////////////////////////
func (userinfo * Userlogins) InsertUser( hashedPassword []byte ,db *gocql.Session) error{
	err :=db.Query("INSERT INTO user_logins(username, password) VALUES(?, ?)",userinfo.username, hashedPassword).Exec()
	return err

}
///////////////////////////////////////////////////////////////////////////////////////////////////////
func (c App) Signin() revel.Result{
	c.Params.Get("user")
		userlogins := map[string]string{}
		content, _ := ioutil.ReadAll(c.Request.Body)
		json.Unmarshal(content, &userlogins)
		response := Responseobject{In: false , Error:"" , Userdata:UserContancts{}}
		userinfo := Userlogins{username:userlogins["username"] , password:userlogins["password"]}
		var databasePassword string
		databasePassword, err := userinfo.QueryUser(c.db)
		if err == gocql.ErrNotFound {
			//no such user
			response.Error ="Username doesn't exist"
			return c.RenderJSON(response)

		} else if  err != nil {

			response.Error ="Internal Server Error please try again"

			return c.RenderJSON(response)
		}

		err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(userlogins["password"]))
		// If wrong password redirect to the login
		if  err != nil {
			//Wrong Password
			response.Error ="wrong password"
			return c.RenderJSON(response)
		} else {
			// If the login succeeded
			c.Session["user"]= userlogins["username"]
			response.In = true
			return c.Redirect("/user")
		}
	}
///////////////////////////////////////////////////////////////////////////////////////////////////////
func (c App) Signup() revel.Result{
	userlogins := map[string]string{}
	content, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(content, &userlogins)
	response := Responseobject{In: false , Error:"" , Userdata:UserContancts{}}
	userinfo := Userlogins{username:userlogins["username"] , password:userlogins["password"]}
	err :=userinfo.CheckUsernameExists(c.db)
	switch {
	case err == nil:
		response.Error = "Please choose a different username"
		return c.RenderJSON( response )

	case err == gocql.ErrNotFound :
			// Username is available
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userlogins["password"]), bcrypt.DefaultCost)
			if err != nil {
				response.Error ="This Password is Not premitted"
				return c.RenderJSON( response )
			}

			err = userinfo.InsertUser(hashedPassword , c.db)
			if err != nil {
				response.Error=err.Error()
				return c.RenderJSON(response)
			}
			c.Session["user"]= userlogins["username"]
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
func startDB(c *App) revel.Result{
	var err error
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "address_book"
	c.db, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	return nil
}

func init(){
	revel.InterceptMethod(startDB , revel.BEFORE)
}