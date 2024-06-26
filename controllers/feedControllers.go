package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"simple-social-media/structs"
	"simple-social-media/database"
	"simple-social-media/repository"
	"simple-social-media/helpers"
)

func GetAllFeed(ctx *gin.Context) {
	feeds, err := repository.GetAllFeed(database.DbConnection)
	
	if err != nil {
		helpers.GeneralResponse(ctx, http.StatusBadRequest, false, "Gagal menampilkan feed", nil, err)
		return
	}
	
	helpers.GeneralResponse(ctx, http.StatusOK, true, "Berhasil menampilkan feed", feeds, nil)
}


func InsertFeed(ctx *gin.Context) {
	var feed structs.Feeds
	tokenWithBearer := ctx.GetHeader("Authorization")

	token, err := helpers.ExtractToken(tokenWithBearer)
	if err != nil {
		panic(err)
	}

	data, err := helpers.VerifyToken(token)
	if err != nil {
		panic(err)
	}

	feed.UserID = data.ID

	err = ctx.ShouldBindJSON(&feed)
	if err != nil {
		helpers.GeneralResponse(ctx, http.StatusBadRequest, false, "Gagal membuat feed", nil, err.Error())
	}

	err = repository.InsertFeed(database.DbConnection, feed)
	if err != nil {
		helpers.GeneralResponse(ctx, http.StatusBadRequest, false, "Gagal membuat feed", nil, err.Error())
		return
	}

	helpers.GeneralResponse(ctx, http.StatusOK, true, "Berhasil membuat feed", nil, nil)
}

func UpdateFeed(ctx *gin.Context) {
	var feed structs.Feeds

	id, _ := strconv.Atoi(ctx.Param("id"))
	
	err := ctx.ShouldBindJSON(&feed)
	if err != nil {
		helpers.GeneralResponse(ctx, http.StatusBadRequest, false, "Gagal memperbarui feed", nil, err.Error())
		return
	}

	feed.ID = int64(id)

	err = repository.UpdateFeed(database.DbConnection, feed)

	if err != nil {
		helpers.GeneralResponse(ctx, http.StatusBadRequest, false, "Gagal mengedit feed", nil, err.Error())
		return
	}

	helpers.GeneralResponse(ctx, http.StatusOK, true, "Berhasil mengedit feed", nil, nil)
}

func DeleteFeed(ctx *gin.Context) {
	var feed structs.Feeds

	id, err := strconv.Atoi(ctx.Param("id"))

	feed.ID = int64(id)

	err = repository.DeleteFeed(database.DbConnection, feed)
	if err != nil {
		panic(err)
	}

	helpers.GeneralResponse(ctx, http.StatusOK, true, "Berhasil menghapus feed", nil, nil)
}

func GetDetailFeed(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	
	if err != nil {
		helpers.GeneralResponse(ctx, http.StatusBadRequest, false, "Gagal menampilkan detail feed", nil, err.Error())
		return
	}

	result, err := repository.GetDetailFeed(database.DbConnection, id)

	if err != nil {
		helpers.GeneralResponse(ctx, http.StatusBadRequest, false, "Gagal menampilkan detail feed", nil, err.Error())
		return
	}

	helpers.GeneralResponse(ctx, http.StatusOK, true, "Berhasil menampilkan detail feed", result, nil)
}