package controllers

import (
	"github.com/revel/revel"
	"fmt"
	"strconv"
	"github.com/gocql/gocql"
)

type PhoneNum struct {
	Id int
	ContactId gocql.UUID
	Phonenumber string

}

type ErrorResponse struct {
	error string
}
type LogoutResponse struct {
	error string
	loggedOut bool
}

type ContactModel struct{
	Id gocql.UUID
	FirstName string
	LastName string
	Email string
	PhoneNumbers []string

}
type Contact struct{
	Id gocql.UUID
	FirstName string
	LastName string
	Email string
	PhoneNumbers []string

}

type ContactResponse struct{
	error string
	isAdded bool
	contact Contact

}

type UserContancts struct {
	UserName string
	Password string
	Contacts []Contact

}

type User struct {
	*revel.Controller
	User UserContancts
	Db *gocql.Session
}

func (currentuser * UserContancts) DeleteContact(id string,db *gocql.Session) error{
	err := db.Query("delete from user_data where username = ? and contact_id = ?",currentuser.UserName , id).Exec()
	return err

}
///////////////////////////////////////////////////////////////////////////////////////////////////////
func (currentuser * UserContancts) DeleteContactNumber(id string , contactid string,db *gocql.Session) error{
	fmt.Println(currentuser)
	err := db.Query("delete contact_phonenumbers[?] from user_data where username = ? and contact_id = ?",id ,currentuser.UserName , contactid ).Exec()
	return err

}
///////////////////////////////////////////////////////////////////////////////////////////////////////
func (currentuser * UserContancts) GetUserContacts( db *gocql.Session) (error){
	var c ContactModel
	rows := db.Query("select contact_id,contact_email,contact_fname,contact_lname,contact_phonenumbers from user_data where username= ?" , currentuser.UserName)
	scanner :=rows.Iter().Scanner()
	for scanner.Next(){
		scanner.Scan(&c.Id , &c.Email, &c.FirstName , &c.LastName , &c.PhoneNumbers)
	var contact Contact
		contact.Id = c.Id
		contact.PhoneNumbers = c.PhoneNumbers
		contact.LastName = c.LastName
		contact.FirstName = c.FirstName
		contact.Email = c.Email
		currentuser.Contacts = append(currentuser.Contacts, contact)
	}
	err := rows.Iter().Close()

	return err
}
///////////////////////////////////////////////////////////////////////////////////////////////////////
func (contact * ContactModel) InsertNewContact(username string ,db *gocql.Session) (error){
	//Start Transaction



	err := db.Query("insert into user_data (username ,contact_id , contact_email , contact_fname , contact_lname , contact_phonenumbers ) values(? , ? , ? , ? , ? , ? ) ", username , contact.Id ,  contact.Email , contact.FirstName, contact.LastName, contact.PhoneNumbers  ).Exec()
	if err !=nil {
		return err
	}
	return  nil

}
///////////////////////////////////////////////////////////////////////////////////////////////////////
func (user *User) GetUser() revel.Result {
	response := Responseobject{In : true }
	Username := user.Session["user"]
	if Username == "" {
		return user.RenderJSON(response)

	}
	myuser := UserContancts{UserName:Username}
	err := myuser.GetUserContacts(user.Db)
	if err != nil {
		response.Error="DB error"
		return user.RenderJSON(response)

	}else{
		user.User=myuser
		response.Userdata=myuser
		return user.RenderJSON(response)
	}

}
///////////////////////////////////////////////////////////////////////////////////////////////////////
func (user *User) AddNewContact() revel.Result{

	Error := ErrorResponse{error:""}

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

		Error.error = "Not valid contact data"
		return user.RenderJSON(Error)
	}
	user.Validation.MinSize(email, 7)
	if user.Validation.HasErrors() {

		Error.error = "Email is too short"

		return user.RenderJSON(Error)
	}else {

		uuid ,err := gocql.RandomUUID()
		if err !=nil {
			Error.error=err.Error()
			return user.RenderJSON(Error)
		}
		c := ContactModel{
			FirstName:firstname,
			LastName:lastname,
			Email:email,
			PhoneNumbers:phonenumbers,
			Id: uuid,

		}
		 err = c.InsertNewContact(username,user.Db)
		if err !=nil {
			Error.error=err.Error()
			return user.RenderJSON(Error)
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
		err := myuser.DeleteContact(user.Params.Get("contactid") , user.Db)
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
		myuser := UserContancts{UserName:user.Session["user"]}
		err := myuser.DeleteContactNumber(user.Params.Get("contactnumber") , user.Params.Get("contactid") , user.Db)
		if err != nil {
			return user.RenderJSON("DB Error")
		}
	}

	return user.RenderJSON("")
}
///////////////////////////////////////////////////////////////////////////////////////////////////////
func (user User) Logout() revel.Result{
	user.User.UserName=""
	user.User.Password=""
	user.User.Contacts=[] Contact{}
	user.Session["user"] = ""
	response := LogoutResponse{error :"" , loggedOut:true}

	return user.RenderJSON(response)
}
///////////////////////////////////////////////////////////////////////////////////////////////////////
func startDatabase(user *User) revel.Result{
	var err error
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "address_book"
	user.Db, err= cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	return nil
}
///////////////////////////////////////////////////////////////////////////////////////////////////////
func init(){
	revel.InterceptMethod(startDatabase , revel.BEFORE)
}