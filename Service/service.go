package Service

import (
	"fmt"
	"product-trace-server/Model"
	"product-trace-server/tools"
	"strconv"
)

var weight = [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
var validValue = [11]byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}

func CheckUserID(_ID string)bool{
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
func GetFullToponym(_ID string) string{
	var str string
	var city1 Model.City
	var city2 Model.City
	var city3 Model.City
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
func CreateUnit(_ID string,_name string,_description string){
	unit:=Model.ProductUnit{
		ID:          _ID,
		Name:        _name,
		Description: _description,
	}
	tools.GetDB().Create(unit)
}
func GetUnit(_ID string)(Model.ProductUnit,bool){
	var unit Model.ProductUnit
	var count int64
	fmt.Println(_ID)
	tools.GetDB().Where(&Model.ProductUnit{ID:_ID}).Find(&unit).Count(&count)
	if count==0{
		return unit,false
	}
	return unit,true
}