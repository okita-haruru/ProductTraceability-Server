package Controllor

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"product-trace-server/Service"
)
type TraceController struct {
	log *logrus.Entry
	ts  *Service.TraceService
}

func NewTraceController(log *logrus.Logger, ts  *Service.TraceService) *TraceController {
	return &TraceController{
		log: log.WithField("controller", "match"),
		ts:  ts,
	}
}

func (tc *TraceController)HandleCheckUserID(r *gin.Context) {
	_id := r.Query("ID")
	flag := tc.ts.CheckUserID(_id)
	r.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"true": flag,
	})
}
func (tc *TraceController)HandleGetFullToponym(r *gin.Context){
	_id := r.Query("ID")
	str := tc.ts.GetFullToponym(_id)
	r.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"fullName": str,
	})
}
func (tc *TraceController)HandleCreateUnit(r *gin.Context){
	_id := r.Query("ID")
	_name := r.Query("name")
	_description := r.Query("description")
	tc.ts.CreateUnit(_id,_name,_description)
	r.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
func (tc *TraceController)HandleGetUnit(r *gin.Context){
	_id := r.Query("ID")
	unit,flag:=tc.ts.GetUnit(_id)
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