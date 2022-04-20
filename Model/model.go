package Model

type City struct {
	ID           string`gorm:"primaryKey"`
	Name         string `gorm:"index"`
	Parentid     string
	Shortname    string
	Leveltype    uint
	Citycode     uint
	Zipcode      uint
	Lng          float32
	Lat          float32
	Pinyin       string
	Status       uint
}
type ProductUnit struct {
	ID           string `gorm:"primaryKey"`
	Name         string `gorm:"index"`
	Description  string
}
