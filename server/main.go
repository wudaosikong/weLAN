package server

import (
	"weLAN/server/config"
	"weLAN/server/controllers"
	"weLAN/server/datasource"
	"weLAN/server/services"
	"fmt"
	"net"
	"time"
	"weLAN/server/cmd"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

func Web() {
	//使用默认配置
	//该实例在生成状态下在"./views"上注册html视图引擎，
	//并从"./locales/*/*"加载区域设置。
	//实例在出现死机时恢复，并记录传入的http请求。
	app := newApp()
	LocalIP = GetIntranetIp()
	//应用App设置
	configation(app)

	//路由设置
	mvcHandle(app)

	config := config.InitConfig()
	addr := ":" + config.Port

	cmd.Open("http://localhost:80")

	app.Run(
		iris.Addr(addr), //在端口8080进行监听
		iris.WithoutServerError(iris.ErrServerClosed), //无服务错误提示
		iris.WithOptimizations,                        //对json数据序列化更快的配置
	)
}

var LocalIP []string

func GetIntranetIp() []string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
	}
	var res []string
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						ip4 := ipnet.IP.String()
						res = append(res, ip4)
					}
				}
			}
		}
	}
	return res
}

// 解决跨域问题
func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}

//构建App
func newApp() *iris.Application {
	app := iris.New()

	//加载跨域配置
	app.Use(Cors)

	//设置日志级别  开发阶段为debug
	app.Logger().SetLevel("debug")
	// 添加两个中间件
	// 1.可以从任何http相关panic中恢复,默认打印包含err内容的warn级别的日志
	// 2.请求日志中间件,请不要将它与框架的日志混淆。这只是针对http请求的。
	app.Use(recover.New())
	app.Use(logger.New(
		logger.Config{
			// 是否记录状态码,默认false
			Status: true,
			// 是否记录远程IP地址,默认false
			IP: true,
			// 是否呈现HTTP谓词,默认false
			Method: true,
			// 是否记录请求路径,默认true
			Path: true,
			// 是否开启查询追加,默认false
			Query: false,
			//是否以列模式显示,默认false
			Columns: false,
		}))

	app.HandleDir("/static", "./server/views")

	//注册视图文件
	app.RegisterView(iris.HTML("./server/views", ".html"))
	app.Get("/", func(context iris.Context) {
		context.View("index.html")
	})
	return app
}

//项目配置
func configation(app *iris.Application) {

	//配置 字符编码
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))

	//错误配置
	//未发现错误
	app.OnErrorCode(iris.StatusNotFound, func(context iris.Context) {
		context.JSON(iris.Map{
			"errmsg": iris.StatusNotFound,
			"msg":    " not found ",
			"data":   iris.Map{},
		})
	})

	app.OnErrorCode(iris.StatusInternalServerError, func(context iris.Context) {
		context.JSON(iris.Map{
			"errmsg": iris.StatusInternalServerError,
			"msg":    " interal error ",
			"data":   iris.Map{},
		})
	})
}

//MVC 架构模式处理
func mvcHandle(app *iris.Application) {

	//启用session
	sessManager := sessions.New(sessions.Config{
		Cookie:  "sessioncookie",
		Expires: 24 * time.Hour,
	})

	engine := datasource.NewMysqlEngine()

	//用户模块功能
	UserService := services.NewUserService(engine)

	User := mvc.New(app.Party("/user"))
	User.Register(
		UserService,
		sessManager.Start,
	)
	User.Handle(new(controllers.UserController))
}
