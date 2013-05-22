package controllers

import(
	"myblog/models"
)

func CheckNamePsw(name,psw string)error{
	medol := models.Model{}
	return medol.CheckNamePsw(name,psw)
}
