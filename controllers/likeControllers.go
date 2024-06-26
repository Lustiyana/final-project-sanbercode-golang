package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	
	"simple-social-media/structs"
	"simple-social-media/repository"
	"simple-social-media/helpers"
	"simple-social-media/database"
)

func InsertLike(ctx *gin.Context) {
	var like structs.Likes

	tokenWithBearer := ctx.GetHeader("Authorization")

	token, err := helpers.ExtractToken(tokenWithBearer)
	if err != nil {
		panic(err)
	}

	data, err := helpers.VerifyToken(token)
	if err != nil {
		panic(err)
	}

	like.UserID = data.ID

	err = ctx.ShouldBindJSON(&like)
	if err != nil {
		helpers.GeneralResponse(ctx, http.StatusBadRequest, false, "Gagal menyukai", nil, err.Error())
	}

	message, err := repository.InsertLike(database.DbConnection, like)
	if err != nil {
		helpers.GeneralResponse(ctx, http.StatusBadRequest, false, message, nil, err.Error())
		return
	}

	helpers.GeneralResponse(ctx, http.StatusOK, true, message, nil, nil)
}