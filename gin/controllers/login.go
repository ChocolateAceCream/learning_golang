package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginQuery struct {
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"passwd" json:"password" uri:"pw" xml:"password" binding:"required"`
}

func Login(r *gin.Context) {
	var parsedQuery LoginQuery

	// request with body in JSON format
	//ShouldBindJSON() will bind json from request body into parsedQuery obj
	// if err := r.BindJSON(&parsedQuery); err != nil {
	// 	r.JSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
	// 	return
	// }

	// request with body in form format
	// Bind() will bind form query from request body into parsedQuery obj
	if err := r.Bind(&parsedQuery); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
		return
	}

	fmt.Println("----- username----", parsedQuery.User)
	fmt.Println("----- Password----", parsedQuery.Password)
	if parsedQuery.User != "admin" || parsedQuery.Password != "123qwe" {
		r.JSON(http.StatusBadRequest, gin.H{"errorMsg": "not authorized"})
		return
	}
	r.JSON(http.StatusOK, gin.H{"status": "200"})
}

// GET localhost:3000/v1/info/1?username=admin&passwd=123qwe
func GetInfo(r *gin.Context) {
	var parsedQuery LoginQuery
	id := r.Param("id")
	fmt.Println("----- id----", id)
	if err := r.BindQuery(&parsedQuery); err != nil {
		fmt.Println("----- parsedQuery----", parsedQuery)
		r.JSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
	}
	if parsedQuery.User != "admin" || parsedQuery.Password != "123qwe" {
		r.JSON(http.StatusBadRequest, gin.H{"errorMsg": "not authorized"})
		return
	}
	r.JSON(http.StatusOK, gin.H{"status": "200", "from": "getinfo"})
}
