package Service

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"product-trace-server/Model"
	"product-trace-server/tools"
	"strconv"
	"time"
)

var weight = [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
var validValue = [11]byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}

type TraceService struct {
	log *logrus.Entry
}
func NewTraceService(log *logrus.Logger) *TraceService {
	return &TraceService{
		log: log.WithField("services", "Match"),
	}
}

func (ts *TraceService)CheckUserID(_ID string)bool{
	if len(_ID)!=18{
		return false
	}
		id18 := []byte(_ID)
	nSum := 0
	for i := 0; i < len(id18)-1; i++ {
		n, _ := strconv.Atoi(string(id18[i]))
		nSum += n * weight[i]
	}
	//mod得出18位身份证校验码
	mod := nSum % 11
	if validValue[mod] == id18[17] {
		return true
	}
	return false
}
func (ts *TraceService)GetFullToponym(_ID string) string{
	var str string
	var city1 Model.City
	var city2 Model.City
	var city3 Model.City
	ts.log.Println("GetFullToponym:"+_ID)
	tools.GetDB().Find(&city1, "id=?", _ID)
	if _ID[len(_ID)-2:len(_ID)] != "00" {
		tools.GetDB().Find(&city2, "id=?", city1.Parentid)
		tools.GetDB().Find(&city3, "id=?", city2.Parentid)
		str=city3.Name+city2.Name+city1.Name
	}else if  _ID[len(_ID)-2:len(_ID)] == "000"{
		str=city1.Name
	}else {
		tools.GetDB().Find(&city2, "id=?", city1.Parentid)
		str=city2.Name+city1.Name
	}
	return str
}
func (ts *TraceService)CreateUnit(_ID uint,_name string,_description string)error{
	ts.log.Println("CreateUnit %s %s %s",_ID,_name,_description)
	unit:=Model.ProductUnit{
		ID:          _ID,
		Name:        _name,
		Description: _description,
	}
	result:=tools.GetDB().Create(unit)
	if result.Error!=nil{
		ts.log.Error("CreateUnitError:")
		return result.Error
	}
	return nil
}
func (ts *TraceService)GetUnit(_ID uint)(Model.ProductUnit,bool){
	var unit Model.ProductUnit
	var count int64
	fmt.Println(_ID)
	tools.GetDB().Where(&Model.ProductUnit{ID:_ID}).Find(&unit).Count(&count)
	if count==0{
		ts.log.Error("not find ID")
		return unit,false
	}
	return unit,true
}
func (ts *TraceService)TransRecord(_unit string,_state string,_time string)(string,string,string){
	uunit,err:=strconv.Atoi(_unit)
	if err!=nil{

	}
	unit,flag:=ts.GetUnit(uint(uunit))
	if !flag{
		return "","",""
	}
	state:=getState(_state)
	timeI,err:=strconv.Atoi(_time)
	fmt.Println(timeI)
	if err!=nil {

	}
	timeLayout := "2006-01-02 15:04:05"

	return unit.Name,state,time.Unix(int64(timeI), 0).Format(timeLayout)
}

func (ts *TraceService)TransTime(_time string)(string){

	timeI,err:=strconv.Atoi(_time)
	fmt.Println(timeI)
	if err!=nil {

	}
	timeLayout := "2006-01-02 15:04:05"

	return time.Unix(int64(timeI), 0).Format(timeLayout)
}

func getState(_state string)string{
	switch _state {
	case "1":
		return "生产"
	case "2":
		return "进入"
	case "3":
		return "离开"
	case "4":
		return "进口"
	}
	return "Error"
}