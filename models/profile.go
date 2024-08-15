package model

import "bmacharia/jwt-go-rbac/database"

/*****************************************************************/

type Profile struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"uniqueIndex"`
	Image     string `gorm:"size:255" json:"image,omitempty"`
	Phone     string `gorm:"size:20;unique" json:"phonenumber,omitempty"`
	LastName  string `gorm:"size:255" json:"lastname,omitempty"`
	FirstName string `gorm:"size:255" json:"firstname,omitempty"`
}

/*****************************************************************/

func (profile *Profile) Save() (*Profile, error) {
	err := database.Database.DB.Create(&profile).Error
	if err != nil {
		return &Profile{}, err
	}
	return profile, nil
}

/*****************************************************************/

func GetProfile(Profile *Profile, id int) (err error) {
	err = database.Database.DB.Where("user_id=?", id).First(&Profile).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/

func GetProfiles(Profile *[]Profile) (err error) {
	err = database.Database.DB.Find(Profile).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/

func UpdateProfile(Profile *Profile) (err error) {
	err = database.Database.DB.Omit("ID").Updates(Profile).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/

func DeleteProfile(Profile *Profile) (err error) {
	err = database.Database.DB.Delete(Profile).Error
	if err != nil {
		return err
	}
	return nil
}

/*****************************************************************/
