package controllers

import (
	models "TsiahPng-Golang/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//TsiahPngGetList use mysql
func TsiahPngGetList(c *gin.Context) {
	if res, count := models.TsiahPngGetList(); count > 0 {
		n := models.Restaurants{Restaurants: res}
		c.JSON(http.StatusOK, n)
	} else {
		n := models.Restaurants{}
		c.JSON(http.StatusOK, n)
	}

}

// //StudentReq struct
// type StudentReq struct {
// 	Name  string `form:"name" json:"name"`
// 	Email string `form:"email" json:"email"`
// }

// //APIInsert use mysql handler 新增一筆學生資料
// func APIInsert(c *gin.Context) {
// 	var studentData StudentReq

// 	// err := json.NewDecoder(c.Request.Body).Decode(&studentData)
// 	err := c.Bind(&studentData)
// 	if err != nil {
// 		c.JSON(http.StatusForbidden, gin.H{
// 			"result": "Error in request",
// 		})
// 		return
// 	}
// 	// c.Request.FormValue("email")
// 	if r := models.DBInsertStudent(studentData.Name, studentData.Email); r == true {
// 		c.JSON(http.StatusOK, gin.H{
// 			"result": "success",
// 		})
// 	} else {
// 		c.JSON(http.StatusForbidden, gin.H{
// 			"result": "fail",
// 		})
// 		// c.JSON(http.StatusNoContent)
// 	}
// }
