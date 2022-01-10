package api

import (
	"fmt"
	"hello/middleware"
	"hello/model"
	"hello/utils/errormsg"
	"hello/utils/img"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListImage(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 15
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, code, total := model.ListImage(pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errormsg.GetErrMsg(code),
	})
}

func PostImg(c *gin.Context) {
	var data model.Imgcontent

	// file, fileHeader, _ := c.Request.FormFile("file")
	// content := c.Request.FormValue("content")
	// fileSize := fileHeader.Size
	// contentType := fileHeader.Header.Get("Content-Type")
	// url, code := model.UpLoadFile(file, contentType, fileSize)

	usernamekey, _ := c.Get("username")
	userinfo, _ := usernamekey.(*middleware.MyClaims)
	_ = c.ShouldBindJSON(&data)
	data.UserID = userinfo.Id
	code = model.PostArt(&data)
	image := img.ResizeImg(data.Imgurl)
	url, _ := model.UpLoadFileImg(image)
	fmt.Println(url)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errormsg.GetErrMsg(code),
	})
}

func GetImgInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetImgInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errormsg.GetErrMsg(code),
	})
}
