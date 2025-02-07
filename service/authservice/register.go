package authservice

import "blog/api/form"

func (authS Service) Register(user form.UserRegister) error {

	err := authS.Storage.Register(user)
	if err != nil {
		return err
	}

	return nil
}

func (authS Service) ValidationRegister(user form.UserRegister) error {
	err := authS.Storage.CheckUsername(user.UserName)
	if err != nil {
		return err
	}

	err = authS.Storage.CheckEmail(user.Email)
	if err != nil {
		return err
	}

	return nil
}
