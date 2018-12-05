package main

import (
	"time"
	"github.com/jinzhu/gorm"
	"fmt"
)

type User struct {
	UserID        int64  `gorm:"primary_key;column:user_id"`
	Name      string     		`gorm:"unique_index:idx_name;column:name"`// if column is not defined, it will change camel case to snake case automatically
	Dependents	  []Dependent    `gorm:"foreignkey:user_id"`
	CreatedAt time.Time
}

type Dependent struct {
	DependentID		int64	`gorm:"primary_key; column:dependent_id"`
	Name 			string	`gorm:"column:name"`
	Relation 		string	`gorm:"column:relation"`
	UserID		int64		`gorm:"column:user_id"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
}



func TestOneToMany(db *gorm.DB) {

	db.LogMode(true)
	db.AutoMigrate(&User{}, &Dependent{})

	//Learned: When uuid is not assigned a value. save command will save all associated records without select and update. but if the value is assgined. It will do select and update
	user1 := User{ Name: "test1", Dependents: []Dependent {
		{ Name: "test11", Relation: "spouse"},
		{ Name: "test12", Relation: "son"}}}

	user2 := 	User{UserID:2, Name: "test2", Dependents: []Dependent {
		{DependentID:21, Name: "test21", Relation: "spouse"},
		{DependentID:22, Name: "test22", Relation: "daughter"}}}

	// user1 data doesn't have primary key. The create is equivalent to the following sql
	// it will do once insert for all records
	db.Create(&user1)

	// user2 has primary key, the create is equivalent to the following sql for the associated entities
	// it will insert the user record, but it will do update, select and insert for all associated entities
	db.Create(&user2)

	users := []User{}


	// preload will to the following sql
	// the following code will do select * from users where name = 'test1'
	//         and select * from dependents where user_id = 1
	db.Where("Name = ? ", "test1").Preload("Dependents").Find(&users)
	fmt.Println("Users", users)

	// the following code will do
	// select * from users join dependents on dependents.user_id = users.user_id where dependents.name = 'test22'
	// select * from dependents where user_id = 2
	users2 := []User{}
	err := db.Joins("JOIN dependents on dependents.user_id = users.user_id").Where("dependents.name = ?", "test22").Preload("Dependents").Find(&users2).Error

	fmt.Println(users2)
	if nil != err {
		fmt.Println("Error in search users with dependent's name.")
	}


	// update change user2's dependent test22 relation from daughter to so
	changeDependent := Dependent{}
	err = db.Where("DependentID = ?", "test22").First(&changeDependent).Error

	err = db.Model(&changeDependent).Update("Relation","son", "UpdatedBy", "change").Error
	if nil != err {
		fmt.Println("Failed to update dependents")
	}
	db.Where("Name = ? ", "test1").Preload("Dependents").Find(&users)
	fmt.Println("Users", users)


	// delete the data with user1
	deleteUser := User{UserID:1}
	db.Delete(&deleteUser)

	var allUsers []User
	db.Preload("Dependents").Find(&allUsers)
	fmt.Println(allUsers)

	// user1 is deleted but the dependents are still there
	var allDependents []Dependent
	db.Find(&allDependents)
	fmt.Println(allDependents)


}