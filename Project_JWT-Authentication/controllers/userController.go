package controllers


import (
	"context"
	"fmt"
	"log"
	"strconv"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	helper "github.com/jayaramsivaramannair/joys_of_GO/Project_JWT-Authentication/helpers"
	"github.com/jayaramsivaramannair/joys_of_GO/Project_JWT-Authentication/models"
	"golang.org/x/crypto/bcrypt"
)


var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validate.New()


func HashPassword()


func VerifyPassword()

func Signup()

func Login()

func GetUsers()

func GetUser () gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return 
		}
	}
}
