package Controllor

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"product-trace-server/Model"
	"product-trace-server/Service"
	"strconv"
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
	tc.log.Println("CheckUserID:"+_id)
	flag := tc.ts.CheckUserID(_id)
	r.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"result": flag,
	})
}
func (tc *TraceController)HandleGetFullToponym(r *gin.Context){
	_id := r.Query("ID")
	tc.log.Println("GetFullToponym:"+_id)
	str := tc.ts.GetFullToponym(_id)
	r.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"fullName": str,
	})
}

func (tc *TraceController)HandleTransTime(r *gin.Context){
	_timestamp := r.Query("timestamp")
	tc.log.Println("TransTime:"+_timestamp)
	str := tc.ts.TransTime(_timestamp)
	r.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"time": str,
	})
}
func (tc *TraceController)HandleCreateUnit(r *gin.Context){
	var err error
	json := Model.ProductUnit{}
	r.BindJSON(&json)
	_id:=json.ID
	_name := json.Name
	_description := json.Description
	//_id := r.Query("ID")
	//_name := r.Query("name")
	//_description := r.Query("description")
	fmt.Println(_id,_name,_description)
	tc.log.Println("CreateUnit:",_id,_name,_description)
	//var _id_I int
	//_id_I,err= strconv.Atoi(_id)
	err =tc.ts.CreateUnit(_id,_name,_description)
	if err!=nil{
		tc.log.Error("ErrorCreateUnit",_id,_name,_description)
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "ErrorCreateUnit",
		})
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
func (tc *TraceController)HandleGetUnit(r *gin.Context){
	_id := r.Query("ID")
	tc.log.Println("GetUnit:"+_id)
	_id_I,err:= strconv.Atoi(_id)
	if err!=nil{

	}
	unit,flag:=tc.ts.GetUnit(uint(_id_I))
	if !flag{
		tc.log.Error( "Unit not fund",_id)
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
func (tc *TraceController)HandleTransRecord(r *gin.Context){

	id := r.Query("ID")
	state := r.Query("state")
	time :=  r.Query("timestamp")
	tc.log.Println("HandleTransRecord",id,state,time)
	_id,_state,_time:=tc.ts.TransRecord(id,state,time)
	var err error
	var _regionCode int
	var _sCode string
	_regionCode,err=strconv.Atoi(id)
	_sCode=strconv.Itoa(_regionCode>>12)

	if err!=nil{

	}
	r.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"name": _id,
		"state": _state,
		"time": _time,
		"region":tc.ts.GetFullToponym(_sCode),
	})
}