package v1

import (
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加分类
func AddCate(c *gin.Context) {
	var data model.Category
	c.ShouldBindJSON(&data)
	code = model.CheckCate(data.Name)
	if code == errmsg.SUCCSE {
		model.CreateCate(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个分类下的文章
func GetCateInfo(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	data, code = model.GetCateInfo(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询分类列表
func GetCate(c *gin.Context) {
	// strconv.Atoi 将str转换为int
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, total := model.GetCate(pageSize, pageNum)
	code = errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑分类
func EditCate(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckCate(data.Name)
	if code == errmsg.SUCCSE {
		model.EditCate(id, &data)
	} else if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除分类
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteCate(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
