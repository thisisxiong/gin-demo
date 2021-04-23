package Models

import "gin-demo/Databases"

type User struct {
	//gorm.Model
	Id int
	Username string
	Password string
	CreatedAt int `gorm:"column:created_time"` //通过tag设置列名为‘created_time’
	UpdatedTime int //列名为·updated_time·
}

// 基本模型的定义
/*type Model struct {
	ID        uint `gorm:"primary_key"` //设置主键
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}*/

/**
  设置表名 默认为结构体的复数形式
 */
func (User) TableName() string {
	return "user"
}

func (this *User) Insert() (id int, err error) {
	result := Databases.Db.Create(&this)
	id = this.Id
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
