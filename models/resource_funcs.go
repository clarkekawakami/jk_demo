package models

import "gorm.io/gorm"

// Resource entity functions

// create a Resource
func CreateResource(db *gorm.DB, Resource *Resource) (err error) {
	err = db.Create(Resource).Error
	if err != nil {
		return err
	}
	return nil
}

// get Resource
func GetResources(db *gorm.DB, Resource *[]Resource) (err error) {
	err = db.Find(Resource).Error
	if err != nil {
		return err
	}
	return nil
}

// get Resource by id
func GetResource(db *gorm.DB, Resource *Resource, id int) (err error) {
	err = db.Where("id = ?", id).First(Resource).Error
	if err != nil {
		return err
	}
	return nil
}

// update Resource
func UpdateResource(db *gorm.DB, Resource *Resource) (err error) {
	db.Save(Resource)
	return nil
}

// delete Resource
func DeleteResource(db *gorm.DB, Resource *Resource, id int) (err error) {
	db.Where("id = ?", id).Delete(Resource)
	return nil
}
