package models

import "github.com/gocql/gocql"

type UserContancts struct {
			    //User data model struct
	UserName string
	Password string
	Contacts []Contact  //List of user contacts

}

func (Currentuser * UserContancts) DeleteContact(id string,db *gocql.Session) error{
	//delete contact with given contactid from the DB
	err := db.Query("delete from user_data where username = ? and contact_id = ?",Currentuser.UserName , id).Exec()
	return err

}
///////////////////////////////////////////////////////////////////////////////////////////////////////
func (currentuser * UserContancts) DeleteContactNumber(id string , contactid string,db *gocql.Session) error{
	//delete number with given contactid and numberid from the DB
	err := db.Query("delete contact_phonenumbers[?] from user_data where username = ? and contact_id = ?",id ,currentuser.UserName , contactid ).Exec()
	return err

}
///////////////////////////////////////////////////////////////////////////////////////////////////////
func (currentuser * UserContancts) GetUserContacts( db *gocql.Session) (error){
	var c ContactModel
	//Get user data with given user name from DB
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
