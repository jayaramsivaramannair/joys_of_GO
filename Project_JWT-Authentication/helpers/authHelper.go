package helper

import (
	"errors"
	"github.com/gin-gonic/gin"
)



func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("Uanthorized to access this resource")
		return err
	}

	return err
}



func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil 

	//Checks if the user is accessing his own data or not
	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorized to access this resource")
		return err
	}

	//If the user is Admin then check again before providing access
	err = CheckUserType(c, userType)
	return err
}