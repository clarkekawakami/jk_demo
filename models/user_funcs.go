package models

import "gorm.io/gorm"

// User entity functions

// create a User
func CreateUser(db *gorm.DB, User *User) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

// get User
func GetUsers(db *gorm.DB, User *[]User) (err error) {
	err = db.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

// get User by id
func GetUser(db *gorm.DB, User *User, id int) (err error) {
	err = db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

// update User
func UpdateUser(db *gorm.DB, User *User) (err error) {
	db.Save(User)
	return nil
}

// delete User
func DeleteUser(db *gorm.DB, User *User, id int) (err error) {
	db.Where("id = ?", id).Delete(User)
	return nil
}
