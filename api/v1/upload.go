package v1

import (
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")

	fileSize := fileHeader.Size
	fileName := fileHeader.Filename
	url, code := model.UploadFile(file, fileName, fileSize)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"url":     url,
		"message": errmsg.GetErrMsg(code),
	})
}
