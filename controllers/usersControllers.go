package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"simple-social-media/structs"
	"simple-social-media/database"
	"simple-social-media/repository"
	"simple-social-media/helpers"

	"fmt"
)

func Register(ctx *gin.Context) {
	var dataUser structs.Users

	err := ctx.ShouldBindJSON(&dataUser)
	if err != nil {
		panic(err)
	}

	err = repository.Register(database.DbConnection, dataUser)

	if err != nil {
		helpers.GeneralResponse(ctx, http.StatusBadRequest,false, err.Error(), nil, nil)
		return
	}

	helpers.GeneralResponse(ctx, http.StatusOK,true, "Berhasil Terdaftar", nil, nil)
}

func Login(ctx *gin.Context) {
	var dataUser structs.Users

	err := ctx.ShouldBindJSON(&dataUser)
	if err != nil {
		panic(err)
	}

	id, err := repository.Login(database.DbConnection, dataUser)

	if err != nil {
		helpers.GeneralResponse(ctx, http.StatusBadRequest,false, err.Error(), nil, nil)
		return
	}

	token, err := helpers.GenerateToken(id, dataUser.Email, dataUser.Password)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}

	tokenResult := map[string]interface{}{
		"token": token,
	}

	helpers.GeneralResponse(ctx, http.StatusOK,true, "Berhasil Masuk", tokenResult, nil)
}