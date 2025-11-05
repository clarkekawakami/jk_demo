package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Facility entity functions

// create a Facility
func CreateFacility(db *gorm.DB, Facility *Facility) (err error) {
	err = db.Create(Facility).Error
	if err != nil {
		return err
	}
	return nil
}

// get Facility
func GetFacilitys(db *gorm.DB, Facility *[]Facility) (err error) {
	err = db.Find(Facility).Error
	if err != nil {
		return err
	}
	return nil
}

// get Facility by location
func GetFacilitysByLocation(db *gorm.DB, Facility *[]Facility, loc string) (err error) {
	whereClause := "state = ?"
	if len(loc) > 0 && loc[0] >= '0' && loc[0] <= '9' {
		fmt.Println("The first character is a digit.", loc)
		whereClause = "zip = ?"
	} else {
		fmt.Println("The first character is not a digit or the string is empty.", loc)
	}

	err = db.Where(whereClause, loc).Find(Facility).Error
	if err != nil {
		return err
	}
	return nil
}

// get Facility by id
func GetFacility(db *gorm.DB, Facility *Facility, id int) (err error) {
	err = db.Where("id = ?", id).First(Facility).Error
	if err != nil {
		return err
	}
	return nil
}

// update Facility
func UpdateFacility(db *gorm.DB, Facility *Facility) (err error) {
	db.Save(Facility)
	return nil
}

// delete Facility
func DeleteFacility(db *gorm.DB, Facility *Facility, id int) (err error) {
	db.Where("id = ?", id).Delete(Facility)
	return nil
}
