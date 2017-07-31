package controllers

import (
	"github.com/revel/revel"
	"fmt"
	"strconv"
	"github.com/gocql/gocql"
	. "Angular-Revel-App/app/models"
	"Angular-Revel-App/app"
)


///////////////////////////////////////////////////////////////////////////////////////////////////////
type User struct {
	*revel.Controller
     //User data
}

func (user *User) GetUser() revel.Result {
	response := Responseobject{In : true }
	Username := user.Session["user"]
	if Username == "" {
		return user.RenderJSON(response)

	}
	myuser := UserContancts{UserName:Username}
	err := myuser.GetUserContacts(app.DB)
	if err != nil {
		response.Error="DB error"
		return user.RenderJSON(response)

	}else{	
		response.Userdata=myuser
		return user.RenderJSON(response)
	}

}
///////////////////////////////////////////////////////////////////////////////////////////////////////
func (user *User) AddNewContact() revel.Result{

	errorResponse := ErrorResponse{Error:""}

	firstname:= user.Params.Get("firstname");
	lastname := user.Params.Get("lastname");
	email    := user.Params.Get("email");
	username := user.Session["user"]
	fmt.Println(firstname);
	var phonenumbers [] string

	i := 1
	for user.Params.Get("phone" + strconv.Itoa(i)) != "" {
		str := user.Params.Get("phone" + strconv.Itoa(i))
		phonenumbers = append(phonenumbers,str)
		i++
	}

	user.Validation.Required(firstname)
	user.Validation.Required(lastname)
	user.Validation.Required(email)
	user.Validation.MaxSize(firstname,50)
	user.Validation.MaxSize(lastname ,50)
	user.Validation.MaxSize(email ,50)
	if user.Validation.HasErrors() {

		errorResponse.Error = "Not valid contact data"
		return user.RenderJSON(errorResponse)
	}
	user.Validation.MinSize(email, 7)
	if user.Validation.HasErrors() {

		errorResponse.Error = "Email is too short"

		return user.RenderJSON(errorResponse)
	}else {

		uuid ,err := gocql.RandomUUID()
		if err !=nil {
			errorResponse.Error=err.Error()
			return user.RenderJSON(errorResponse)
		}
		c := ContactModel{
			FirstName:firstname,
			LastName:lastname,
			Email:email,
			PhoneNumbers:phonenumbers,
			Id: uuid,

		}
		 err = c.InsertNewContact(username,app.DB)
		if err !=nil {
			errorResponse.Error=err.Error()
			return user.RenderJSON(errorResponse)
		}

		return user.RenderJSON(c)
	}
}
///////////////////////////////////////////////////////////////////////////////////////////////////////
func (user User) Delete() revel.Result {
	user.Validation.Clear()
	user.Validation.Required(user.Params.Get("contactid"))
	if user.Validation.HasErrors() {
		return user.RenderJSON("Internal Server Error")
	} else {
		myuser := UserContancts{UserName:user.Session["user"]}
		err := myuser.DeleteContact(user.Params.Get("contactid") , app.DB)
		if err != nil {
			return user.RenderJSON("DB Error")
		}
	}
			return user.RenderJSON("")

}
///////////////////////////////////////////////////////////////////////////////////////////////////////
func (user User) DeleteNum() revel.Result{
	user.Validation.Clear()
	user.Validation.Required(user.Params.Get("contactid"))
	if user.Validation.HasErrors() {
		return user.RenderJSON("Internal Server Error")
	} else {
		myuser := &UserContancts{UserName:user.Session["user"]}
		err := myuser.DeleteContactNumber(user.Params.Get("contactnumber") , user.Params.Get("contactid") , app.DB)
		if err != nil {
			return user.RenderJSON("DB Error")
		}
	}

	return user.RenderJSON("")
}
///////////////////////////////////////////////////////////////////////////////////////////////////////
func (user User) Logout() revel.Result{
	user.Session["user"] = ""
	response := LogoutResponse{Error :"" , LoggedOut:true}

	return user.RenderJSON(response)
}
///////////////////////////////////////////////////////////////////////////////////////////////////////