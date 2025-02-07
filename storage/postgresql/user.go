package postgresql

import (
	"blog/api/form"
	"log"
)

func (s Storage) Register(user form.UserRegister) error {
	stetm, err := s.DB.Prepare("INSERT INTO user(username,email,password) VALUE($1,$2,$3)")
	if err != nil {
		return err
	}
	defer stetm.Close()

	_, err = stetm.Exec(user.UserName, user.Email, user.Password)
	if err != nil {
		return err
	}

	log.Println("Register user successfully")
	return nil
}
