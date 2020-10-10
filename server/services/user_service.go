package services

import (
	"fmt"
	"weLAN/server/models"

	"xorm.io/xorm"
)

type UserService interface {
	//通过用户用户名+密码 获取用户实体 如果查询到，返回用户实体，并返回true
	//否则 返回 nil ，false
	GetByUserNameAndPassword(username, password, localip string) (models.User, bool)
	GetByUserName(username string) bool
	//获取用户总数
	GetUserCount() (int64, error)
	AddUser(userRegister models.User) bool
	UserList() interface{}
}

func NewUserService(db *xorm.Engine) UserService {
	return &userSevice{
		engine: db,
	}
}

/**
 * 用户的服务实现结构体
 */
type userSevice struct {
	engine *xorm.Engine
}

/**
 * 查询用户总数
 */
func (ac *userSevice) GetUserCount() (int64, error) {
	count, err := ac.engine.Count(new(models.User))

	if err != nil {
		panic(err.Error())
	}
	return count, nil
}

/**
 * 通过用户名和密码查询用户
 */
func (ac *userSevice) GetByUserNameAndPassword(username, password, localip string) (models.User, bool) {
	var user models.User

	ac.engine.Where(" user_name = ? and password = ? ", username, password).Get(&user)

	if user.UserId != 0 {
		userInsert := models.User{
			LocalIp: localip,
		}
		rowNum, err := ac.engine.ID(user.UserId).Update(&userInsert)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(rowNum) //rowNum 受影响的记录条数
		fmt.Println()
	}

	return user, user.UserId != 0
}

func (ac *userSevice) GetByUserName(username string) bool {
	var user models.User

	ac.engine.Where(" user_name = ? ", username).Get(&user)
	fmt.Println(user)
	fmt.Println()
	return user.UserId != 0
}

func (ac *userSevice) GetInfo(username string) bool {
	var user models.User

	ac.engine.Where(" user_name = ? ", username).Get(&user)
	fmt.Println(user)
	return user.UserId != 0
}

func (ac *userSevice) AddUser(userRegister models.User) bool {
	userInsert := models.User{
		UserName: userRegister.UserName,
		Password: userRegister.Password,
		MyName:   userRegister.MyName,
		LocalIp:  userRegister.LocalIp,
	}
	rowNum, err := ac.engine.Insert(&userInsert)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rowNum) //rowNum 受影响的记录条数
	fmt.Println()
	return true
}

func (ac *userSevice) UserList() interface{} {
	var users []string

	ac.engine.Table("user").Cols("my_name").Find(&users)

	// usersJS, err := json.Marshal(users)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(usersJS))

	respDesc := map[string]interface{}{
		"users": users,
	}
	return respDesc
}
