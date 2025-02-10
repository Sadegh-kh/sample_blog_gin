package postgresql

import (
	"blog/api/form"
	"database/sql"
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

func (s Storage) CheckUniqUsername(userName string) (bool, error) {
	row := s.DB.QueryRow("SELECT * FROM user WHERE user.username=" + userName)
	err := row.Scan()
	if err == sql.ErrNoRows {
		return true, nil
	}

	return false, err
}

func (s Storage) CheckUniqEmail(email string) (bool, error) {
	row := s.DB.QueryRow("SELECT * FROM user WHERE user.email=" + email)
	err := row.Scan()
	if err == sql.ErrNoRows {
		return true, nil
	}

	return false, err
}
