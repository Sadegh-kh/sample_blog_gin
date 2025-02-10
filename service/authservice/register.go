package authservice

import (
	"blog/api/form"
	"fmt"
)

func (authS Service) Register(user form.UserRegister) error {

	err := authS.Storage.Register(user)
	if err != nil {
		return err
	}

	return nil
}

func (authS Service) ValidationRegister(user form.UserRegister) error {
	exist, err := authS.Storage.CheckUniqUsername(user.UserName)
	if err != nil {
		return err
	} else if !exist {
		return fmt.Errorf("username exist")

	}

	exist, err = authS.Storage.CheckUniqEmail(user.Email)
	if err != nil {
		return err
	} else if !exist {
		return fmt.Errorf("email exist")

	}

	return nil
}
