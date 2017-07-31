package models

import (
	"github.com/gocql/gocql"
	"fmt"
)
type Responseobject struct {
	In bool
	Error string
	Userdata UserContancts
}

type UserLogins struct{
	Username string
	Password string
}

func (userinfo * UserLogins) CheckUsernameExists( db *gocql.Session) error{
	var databasePassword string

	err := db.Query("SELECT password FROM user_logins WHERE username=?", userinfo.Username).Scan( &databasePassword)
	fmt.Println(err)
	return err
}

func (userinfo * UserLogins) QueryUser(db *gocql.Session) (string,error){
	var databasePassword string

	err := db.Query("SELECT password FROM user_logins WHERE username=?", userinfo.Username).Scan( &databasePassword)
	return databasePassword,err
}
///////////////////////////////////////////////////////////////////////////////////////////////////////
func (userinfo * UserLogins) InsertUser( hashedPassword []byte ,db *gocql.Session) error{
	err :=db.Query("INSERT INTO user_logins(username, password) VALUES(?, ?)",userinfo.Username, hashedPassword).Exec()
	return err

}
