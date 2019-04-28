package models

import (
	"fmt"
	u "go-rest-api-basic/utils"

	"github.com/jinzhu/gorm"
)

type Profile struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Dob       string `json:"dob"`
	UserId    uint   `json:"user_id"`
}

func (profile *Profile) Validate() (map[string]interface{}, bool) {
	if profile.FirstName == "" {
		return u.Message(false, "Contact name should be on the payload"), false
	}

	if profile.LastName == "" {
		return u.Message(false, "Phone number should be on the payload"), false
	}

	if profile.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	return u.Message(true, "success"), true
}

func (profile *Profile) Create() map[string]interface{} {
	if resp, ok := profile.Validate(); !ok {
		return resp
	}

	GetDB().Create(profile)

	resp := u.Message(true, "success")
	resp["profile"] = profile
	return resp
}

func (profile *Profile) Get(userId uint) []*Profile {
	profiles := make([]*Profile, 0)

	//newProfile := Profile{}
	err := GetDB().Table("profiles").Where("user_id = ?", userId).Find(&profiles).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return profiles
}

func (profile *Profile) Update() *Profile {
	matchedProfile := profile.Get(profile.ID)[0]
	matchedProfile.LastName = profile.LastName
	matchedProfile.FirstName = profile.FirstName
	matchedProfile.Dob = profile.Dob
	err := GetDB().Save(&matchedProfile)
	if err != nil {
		return nil
	}
	return matchedProfile
}
