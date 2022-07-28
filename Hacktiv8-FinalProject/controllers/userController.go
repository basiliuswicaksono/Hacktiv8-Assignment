package controllers

import (
	"encoding/json"
	"finalProject/helpers"
	"finalProject/models"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

type UserLogin struct {
	Email    string `valid:"required~Email is required, email~Invalid format email"`
	Password string `valid:"required~Password is required,minstringlength(6)~Password has to have minimum length of 6 characters"`
}

type UserUpdateValidation struct {
	Username string `valid:"required~Username is required"`
	Email    string `valid:"required~Email is required,email~Invalid format email"`
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{db: db}
}

func (u *UserController) Register(ctx *gin.Context) {
	var user models.User
	var err error

	err = ctx.ShouldBindJSON(&user)
	if err != nil {
		badRequestJsonResponse(ctx, err.Error())
		return
	}

	_, errCreate := govalidator.ValidateStruct(&user)
	if errCreate != nil {
		badRequestJsonResponse(ctx, errCreate.Error())
		return
	}

	err = u.db.Create(&user).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusCreated, gin.H{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
		"age":      user.Age,
	})
}

func (u *UserController) Login(ctx *gin.Context) {
	var userLogin UserLogin

	err := ctx.ShouldBindJSON(&userLogin)
	if err != nil {
		badRequestJsonResponse(ctx, err.Error())
		return
	}

	_, errLogin := govalidator.ValidateStruct(&userLogin)
	if errLogin != nil {
		badRequestJsonResponse(ctx, errLogin.Error())
		return
	}

	var userDB models.User

	err = u.db.First(&userDB, "email=?", userLogin.Email).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			// noDataJsonResponse(ctx, err.Error())
			unauthorizeJsonResponse(ctx, "username / password is not match")
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	isValid := helpers.ComparePassword(userDB.Password, userLogin.Password)
	if !isValid {
		unauthorizeJsonResponse(ctx, "username / password is not match")
		return
	}

	token, err := helpers.GenerateToken(userDB.ID, userDB.Email)
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
	}

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"token": token,
	})
}

func (u *UserController) UpdateUser(ctx *gin.Context) {
	// id, _ := ctx.Get("id")
	email, _ := ctx.Get("email")

	var (
		userUpdate UserUpdateValidation
		err        error
		userDB     models.User
	)

	err = ctx.ShouldBindJSON(&userUpdate)
	if err != nil {
		badRequestJsonResponse(ctx, err.Error())
		return
	}

	_, errUpdate := govalidator.ValidateStruct(&userUpdate)
	if errUpdate != nil {
		badRequestJsonResponse(ctx, errUpdate.Error())
		return
	}

	err = u.db.First(&userDB, "email=?", email).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	err = u.db.Model(&userDB).Updates(models.User{Username: userUpdate.Username, Email: userUpdate.Email}).Error
	if err != nil {
		var newErr map[string]interface{}
		byteErr, _ := json.Marshal(err)
		json.Unmarshal((byteErr), &newErr)

		if newErr["Code"] == "23505" {
			badRequestJsonResponse(ctx, newErr["Detail"])
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"id":         userDB.ID,
		"email":      userDB.Email,
		"username":   userDB.Username,
		"age":        userDB.Age,
		"updated_at": userDB.UpdatedAt,
	})
}

func (u *UserController) DeleteUser(ctx *gin.Context) {
	email, _ := ctx.Get("email")

	var (
		err    error
		userDB models.User
	)

	err = u.db.First(&userDB, "email=?", email).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			noDataJsonResponse(ctx, err.Error())
			return
		}
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	err = u.db.Delete(&userDB).Error
	if err != nil {
		internalServerJsonResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
