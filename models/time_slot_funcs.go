package models

import "gorm.io/gorm"

// Time_slot entity functions

// create a Time_slot
func CreateTime_slot(db *gorm.DB, Time_slot *Time_slot) (err error) {
	err = db.Create(Time_slot).Error
	if err != nil {
		return err
	}
	return nil
}

// get Time_slot
func GetTime_slots(db *gorm.DB, Time_slot *[]Time_slot) (err error) {
	err = db.Find(Time_slot).Error
	if err != nil {
		return err
	}
	return nil
}

// get Time_slot by id
func GetTime_slot(db *gorm.DB, Time_slot *Time_slot, id int) (err error) {
	err = db.Where("id = ?", id).First(Time_slot).Error
	if err != nil {
		return err
	}
	return nil
}

// update Time_slot
func UpdateTime_slot(db *gorm.DB, Time_slot *Time_slot) (err error) {
	db.Save(Time_slot)
	return nil
}

// delete Time_slot
func DeleteTime_slot(db *gorm.DB, Time_slot *Time_slot, id int) (err error) {
	db.Where("id = ?", id).Delete(Time_slot)
	return nil
}
