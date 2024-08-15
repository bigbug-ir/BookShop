package model

/*****************************************************************/

type Password struct {
	PrevPassword      string `json:"prev_passwrod"`
	NewPassWord       string `json:"new_password"`
	RepeatNewPassWord string `json:"repeat_new_password"`
}

/*****************************************************************/

func CheckPassword(Password *Password) (err error) {
	if Password.NewPassWord == Password.PrevPassword {
		return err
	}
	if Password.NewPassWord != Password.RepeatNewPassWord {
		return err
	}
	return nil
}

/*****************************************************************/
