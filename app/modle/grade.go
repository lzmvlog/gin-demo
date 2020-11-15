package modle

type Grade struct {
	Id        int    `orm:"id,primary" json:"id"`
	Gradename string `orm:"gradename"  json:"gradename"`
}
