package models

import "github.com/gocql/gocql"


type ErrorResponse struct {
	//response struct to send to the front end when error happens
	Error string
}

type LogoutResponse struct {
	//response struct to send to the front end when logout is requested
	Error string
	LoggedOut bool
}


type ContactResponse struct{
	//response struct to send back when adding new contact is requested with status and added contact
	error string
	isAdded bool
	contact Contact

}

type Contact struct{
	//contact data to send back to the user
	Id gocql.UUID
	FirstName string
	LastName string
	Email string
	PhoneNumbers []string

}
type ContactModel struct{
	//Model contact Struct with contact data and methods
	//notice that we cannot use struct with methods as response that's because it would not be encoded successfully into json object
	Id gocql.UUID
	FirstName string
	LastName string
	Email string
	PhoneNumbers []string

}
func (contact * ContactModel) InsertNewContact(username string ,db *gocql.Session) (error){
	//insert new contact data to the database
	err := db.Query("insert into user_data (username ,contact_id , contact_email , contact_fname , contact_lname , contact_phonenumbers ) values(? , ? , ? , ? , ? , ? ) ", username , contact.Id ,  contact.Email , contact.FirstName, contact.LastName, contact.PhoneNumbers  ).Exec()
	if err !=nil {
		//if error happens
		return err
	}
	return  nil

}
