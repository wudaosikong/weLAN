package models

/**
 * 商店服务类别结构体
 */
type Service struct {
	Id          int     `xorm:"pk autoincr" json:"id"`           // 主键自增
	Name        string  `xorm:"varchar(32)" json:"name"`         //名称
	Description string  `xorm:"varchar(255)" json:"description"` //服务描述
}
