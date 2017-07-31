package models

import (
	"github.com/gocql/gocql"
	"fmt"
)

type Responseobject struct {
	In bool
	Error string
	Userdata User
}

type ErrorResponse struct {
	//response struct to send to the front end when error happens
	Error string
}

type LogoutResponse struct {
	//response struct to send to the front end when logout is requested
	Error string
	LoggedOut bool
}

type LoginData struct{
	Username string
	Password string
}
type User struct {
	Logins LoginData
	Contacts []Contact

}

func (user * User) CheckUsernameExists(db* gocql.Session) error{
	var databasePassword string

	err := db.Query("SELECT password FROM user_logins WHERE username=?", user.Logins.Username).Scan( &databasePassword)
	fmt.Println(err)
	return err
}

func (user * User) QueryUser(db * gocql.Session) (string,error){
	var databasePassword string

	err := db.Query("SELECT password FROM user_logins WHERE username=?", user.Logins.Username).Scan( &databasePassword)
	return databasePassword,err
}
///////////////////////////////////////////////////////////////////////////////////////////////////////
func (user * User) InsertUser(hashedPassword []byte , db * gocql.Session) error{
	err :=db.Query("INSERT INTO user_logins(username, password) VALUES(?, ?)", user.Logins.Username, hashedPassword).Exec()
	return err

}

func (user * User) GetUserContacts( db *gocql.Session) error{
	var newcontact Contact
	rows := db.Query("select contact_id,contact_email,contact_fname,contact_lname,contact_phonenumbers from user_data where username= ?" , user.Logins.Username)
	scanner :=rows.Iter().Scanner()
	for scanner.Next(){
		scanner.Scan(&newcontact.Id , &newcontact.Email, &newcontact.FirstName , &newcontact.LastName , &newcontact.PhoneNumbers)
		user.Contacts = append(user.Contacts, newcontact)
	}
	err := rows.Iter().Close()
	return err
}

func (user * User) AddtoContacts(contact  Contact) {
	user.Contacts = append(user.Contacts, contact)
	return
}