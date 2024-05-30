package controllers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	
	"simple-social-media/structs"
	"simple-social-media/repository"
	"simple-social-media/helpers"
	"simple-social-media/database"
)

func InsertComment(ctx *gin.Context) {
	var comment structs.Comments

	tokenWithBearer := ctx.GetHeader("Authorization")

	token, err := helpers.ExtractToken(tokenWithBearer)
	if err != nil {
		panic(err)
	}

	data, err := helpers.VerifyToken(token)
	if err != nil {
		panic(err)
	}

	comment.UserID = data.ID

	err = ctx.ShouldBindJSON(&comment)
	if err != nil {
		helpers.GeneralResponse(ctx, http.StatusBadRequest, false, "Gagal membuat komentar", nil, err.Error())
		return
	}

	err = repository.InsertComment(database.DbConnection, comment)
	if err != nil {
		helpers.GeneralResponse(ctx, http.StatusBadRequest, false, "Gagal membuat komentar", nil, err.Error())
		return
	}

	helpers.GeneralResponse(ctx, http.StatusOK, true, "Berhasil membuat komentar", nil, nil)
}

func UpdateComment(ctx *gin.Context) {
	var comment structs.Comments

	id, _ := strconv.Atoi(ctx.Param("id"))

	err := ctx.ShouldBindJSON(&comment)
	if err != nil {
		helpers.GeneralResponse(ctx, http.StatusBadRequest, false, "Gagal memperbarui komentar", nil, err.Error())
		return
	}

	comment.ID = int64(id)

	err = repository.UpdateComment(database.DbConnection, comment)

	if err != nil {
		helpers.GeneralResponse(ctx, http.StatusBadRequest, false, "Gagal memperbarui komentar", nil, err.Error())
		return
	}

	helpers.GeneralResponse(ctx, http.StatusOK, true, "Berhasil memperbarui komentar", nil, nil)
}

func DeleteComment(ctx *gin.Context) {
	var comment structs.Comments
	id, err := strconv.Atoi(ctx.Param("id"))

	comment.ID = int64(id)

	err = repository.DeleteComment(database.DbConnection, comment)
	if err != nil {
		helpers.GeneralResponse(ctx, http.StatusBadRequest, false, "Gagal menghapus komentar", nil, err.Error())
		return
	}

	helpers.GeneralResponse(ctx, http.StatusOK, true, "Berhasil menghapus komentar", nil, nil)
}
