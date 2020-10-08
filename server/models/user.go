package models

import (
	"time"
)

//定义用户结构体
type User struct {
	//如果field名称为Id，而且类型为int64，并没有定义tag，则会被xorm视为主键，并且拥有自增属性
	UserId     int64     `xorm:"pk autoincr" json:"id"` //主键 自增
	UserName   string    `xorm:"varchar(32)" json:"user_name"`
	CreateTime time.Time `xorm:"DateTime" json:"create_time"`
	Status     int64     `xorm:"default 0" json:"status"`
	Avatar     string    `xorm:"varchar(255)" json:"avatar"`
	MyName     string    `xorm:"varchar(32)" json:"my_name"`
	Password        string    `xorm:"varchar(255)" json:"password"` //用户密码
	LocalIp    string    `xorm:"varchar(15)" json:"local_ip"`
}

/**
 * 从User数据库实体转换为前端请求的resp的json格式
 */
func (this *User) UserToRespDesc() interface{} {
	respDesc := map[string]interface{}{
		"user_name":   this.UserName,
		"id":          this.UserId,
		"create_time": this.CreateTime,
		"status":      this.Status,
		"avatar":      this.Avatar,
		"my_name":     this.MyName,
	}
	return respDesc
}
