package Models
//与数据库表格字段相对应
type User struct {
	Id 				string `json:"id";gorm:"primary_key"`
	UserName 		string `json:"username"`
	Name 			string `json:"name"`
	UserPassword 	string `json:"password"`
}
