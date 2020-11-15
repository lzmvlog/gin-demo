package modle

type Student struct {
	Id      string `orm:"id,primary" json:"id"`
	Name    string `orm:"name"       json:"name"`
	Age     int    `orm:"age"        json:"age"`
	Gradeid int    `orm:"gradeid"    json:"gradeid" xorm:"index"`
}
