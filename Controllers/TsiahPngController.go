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

// RestaurantReq struct
type RestaurantReq struct {
	Name    string `form:"name" json:"name"`
	Price   string `form:"price" json:"price"`
	Purpose string `form:"purpose" json:"purpose"`
}

//TsiahPngRestaurantInsert use mysql handler 新增一筆餐廳資料
func TsiahPngRestaurantInsert(c *gin.Context) {
	var req RestaurantReq

	// err := json.NewDecoder(c.Request.Body).Decode(&req)
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"result": "Error in request",
		})
		return
	}
	// c.Request.FormValue("email")
	if r := models.DBInsertRestaurant(req.Name, req.Price, req.Purpose); r == true {
		c.JSON(http.StatusOK, gin.H{
			"result": "success",
		})
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"result": "fail",
		})
		// c.JSON(http.StatusNoContent)
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
