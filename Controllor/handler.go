package Controllor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"product-trace-server/Service"
)

func HandleCheckUserID(r *gin.Context) {
	_id := r.Query("ID")
	flag := Service.CheckUserID(_id)
	r.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"true": flag,
	})
}
func HandleGetFullToponym(r *gin.Context){
	_id := r.Query("ID")
	str := Service.GetFullToponym(_id)
	r.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"fullName": str,
	})
}
func HandleCreateUnit(r *gin.Context){
	_id := r.Query("ID")
	_name := r.Query("name")
	_description := r.Query("description")
	Service.CreateUnit(_id,_name,_description)
	r.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
func HandleGetUnit(r *gin.Context){
	_id := r.Query("ID")
	unit,flag:=Service.GetUnit(_id)
	if !flag{
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "Unit not fund",
		})
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"id": unit.ID,
		"name": unit.Name,
		"description": unit.Description,
	})
}