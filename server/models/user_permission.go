package models

/**
 * 用户的权限明细表（业务结构表）
 */
type UserPermission struct {
	User      *User      `xorm:"extends"` //不需要映射User结构体
	Permission *Permission `xorm:"extends"` //不需要映射permission结构体
}
