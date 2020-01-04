![](bingo-gin.jpg)
   ___ _                               _       
  / __(_)_ __   __ _  ___         __ _(_)_ __  
 /__\// | '_ \ / _` |/ _ \ _____ / _` | | '_ \ 
/ \/  \ | | | | (_| | (_) |_____| (_| | | | | |
\_____/_|_| |_|\__, |\___/       \__, |_|_| |_|
               |___/             |___/        
               
 ![Travis](https://img.shields.io/badge/build-passing-brightgreen.svg) 
 ![Travis](https://img.shields.io/badge/language-golang-blue.svg) 
 ![Travis](https://img.shields.io/badge/godoc-reference-blue.svg) 
 [![Travis](https://img.shields.io/badge/vendor-gin-important.svg)](https://github.com/gin-gonic/gin)
 ![Travis](https://img.shields.io/badge/version-v1.0.0-yellowgreen.svg) 
 ![Travis](https://img.shields.io/badge/platform-linux|mac|windows-inactive.svg)

# Bingo-gin,an easy web framework for Go

Bingo-gin is a web framework based on the Gin Web Framework. It was created specifically for APIs. Bingo-gin highly encapsulates the Gin Web Framework and still maintains the original features of the Gin Web Framework, allowing developers to write multi-process, asynchronous, and highly available applications with minimal learning cost and effort.

- Base on Gin Web Framework
- Database Gorm ORM support
- Mysql Clients
- RESTful supported
- High performance router
- Fast and flexible parameter validator
- Powerful log component
- Universal connection pools
- Remote Console support


        
### Build
```
go build ./...
```

### Config
```
// all config setting in config/config.ini
// you can get a config value like this
cfg := common.NewConfig()
configValue := cfg.GetValue("", "CONFIG_NAME")
```


### Router
##### Create a new router group
```
// router/initRouter.go
func InitRouter(router *gin.Engine) {
	//设置路由中间件
	router.Use(exception.SetUp(), logger.SetUp())

	BaseRouter := &Router{Root: router}
	//首页Api
	BaseRouter.Index()

	return
}
```
##### Set router to controller
```
// router/index.go
func (r *Router) Index() {
	// 设置路由分组
	r.IndexApi = r.Root.Group("")
	// 设置controller
	controller := controller.IndexController{}


	r.IndexApi.GET("/index", controller.IndexAction)
}
```

### Controller
##### Create a controller like this
```
type IndexController struct {
	BaseController
}

func (ctl *IndexController) IndexAction(c *gin.Context) {
	ctl.BaseController.context = c
	/**
	参数验证
	 */
	validator := index.IndexValidator{}
	if err := c.ShouldBind(&validator); err != nil {
		ctl.ErrorResponse(common.CodeParamsError, "参数错误", "")
		return
	}

	/**
	参数
	 */
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	//name := c.DefaultQuery("name", "0")

	/**
	查询数据
	 */
	userTable := model.UserTableInstance()
	ids := make([]int, 1)
	ids[0] = id
	userList := userTable.GetUsersById(ids)

	/**
	返回响应
	 */
	returnData := make(map[string]interface{})
	returnData["userList"] = userList
	ctl.SuccessResponse("请求成功", returnData)
}
```


### Params Validate
Param validate in controller seem like gin
```
// controller/validator
type IndexValidator struct {
	Name	string    `form:"name" binding:"required"`
	Id  	string    `form:"id" binding:"required"`
}
```

### Database Orm
##### Create a table orm
```
// model/userTable.go

type UserTable struct {
	MysqlDb
}

var instanceUserTable *UserTable
var once sync.Once

/**
 * 单例模式
 */
func UserTableInstance() *UserTable{
	once.Do(func(){
		instanceUserTable = &UserTable{}
		instanceUserTable.New()
	})
	return instanceUserTable
}

func (DB *UserTable) New() *UserTable {
	DB.Dababase = "bingo"
	DB.Table = "user"
	DB.DebugMode = true
	if DBCon == nil || DBCon.DB().Ping() == nil {
		DB.Init()
	}
	return DB
}
```

##### Use table orm
```
userTable := model.UserTableInstance()
```

##### Database CURD 
```
// model/mysqlDb.go

// 根据条件查找数据
//	example:
//  	var rowsEntity []entity.TableEntity
//		DB.Select(&rowsEntity, where, limit, 0, "*", "id")
func (DB *MysqlDb) Select (
	entity interface{},
	where map[string]interface{},
	limit int, offset int, fields string, orderBy string) interface{} {

	// 处理 where 参数
	objDb := DB.dbconn.Table(DB.Table)
	for fieldName, val := range where {
		if reflect.TypeOf(val).Kind() == reflect.Slice || reflect.TypeOf(val).Kind() == reflect.Array {
			objDb = objDb.Where(fieldName+" IN (?)", val)
		} else {
			objDb = objDb.Where(fieldName+" = ?", val)
		}
	}

	rows := objDb.Limit(limit).Offset(offset).Order(orderBy)
	if result := rows.Find(entity); result.Error != nil {
		dbErr := common.NewError(common.CodeDbErrSqlSyntax, fmt.Sprintf("数据库报错-%s", result.Error.Error()))
		panic(dbErr)
	}

	return entity
}


// 执行查询sql语句
//	example:
//  	var rowEntity entity.TableEntity
//	    DB.QuerySql(&rowEntity,"select * from `table` where id in (?)", 2)
func (DB *MysqlDb) QuerySql(
	entity interface{},
	sql string, values ...interface{}) interface{} {

	objDb := DB.dbconn.Table(DB.Table)
	result := objDb.Raw(sql, values).Find(entity)
	if result.Error != nil {
		dbErr := common.NewError(common.CodeDbErrSqlSyntax, fmt.Sprintf("数据库报错-%s", result.Error.Error()))
		panic(dbErr)
	}
	return entity
}


// 执行原生的sql语句
//	example:
//	    DB.ExecSql("DROP TABLE table", nil)
//	    DB.ExecSql("UPDATE orders SET shipped_at=? WHERE id IN (?)", time.Now(), []int64{11,22,33})
func (DB *MysqlDb) ExecSql(sql string, values ...interface{}) {

	objDb := DB.dbconn.Table(DB.Table)
	result := objDb.Exec(sql, values)
	if result.Error != nil {
		dbErr := common.NewError(common.CodeDbErrSqlSyntax, fmt.Sprintf("数据库报错-%s", result.Error.Error()))
		panic(dbErr)
	}
}


// 插入数据
//	example:
//		var table = entity.TableEntity{fieldName: value}
//		DB.Insert(&table)
func (DB *MysqlDb) Insert(entity interface{}) interface{} {

	objDb := DB.dbconn.Table(DB.Table)
	result := objDb.Create(entity)
	if result.Error != nil {
		dbErr := common.NewError(common.CodeDbErrSqlSyntax, fmt.Sprintf("数据库报错-%s", result.Error.Error()))
		panic(dbErr)
	}
	return entity
}


// 根据实体来修改某一行数据
//	example:
//		var updateData = map[string]interface{}{"fieldName": "value"}
//		DB.UpdateOneRow(&tableEntity, updateData)
func (DB *MysqlDb) UpdateOneRow(entity interface{}, updateData map[string]interface{}) {

	objDb := DB.dbconn.Table(DB.Table)
	result := objDb.Model(entity).Updates(updateData)
	if result.Error != nil {
		dbErr := common.NewError(common.CodeDbErrSqlSyntax, fmt.Sprintf("数据库报错-%s", result.Error.Error()))
		panic(dbErr)
	}
}


// 根据条件查询修改多行数据
//	example:
//		var where = map[string]interface{}{"id": []int{1, 2}}
//		var updateData = map[string]interface{}{"fieldName": "value"}
//		DB.UpdateBatchRows(where, updateData)
func (DB *MysqlDb) UpdateBatchRows(where map[string]interface{}, updateData map[string]interface{}) {

	objDb := DB.dbconn.Table(DB.Table)
	for fieldName, val := range where {
		if reflect.TypeOf(val).Kind() == reflect.Slice || reflect.TypeOf(val).Kind() == reflect.Array {
			objDb = objDb.Where(fieldName+" IN (?)", val)
		} else {
			objDb = objDb.Where(fieldName+" = ?", val)
		}
	}

	result := objDb.Updates(updateData)
	if result.Error != nil {
		dbErr := common.NewError(common.CodeDbErrSqlSyntax, fmt.Sprintf("数据库报错-%s", result.Error.Error()))
		panic(dbErr)
	}
}


// 根据实体entity删除某行数据
//	example:
//		DB.DeleteOneRow(&tableEntity)
func (DB *MysqlDb) DeleteOneRow(entity interface{}) {

	objDb := DB.dbconn.Table(DB.Table)
	result := objDb.Delete(&entity)
	if result.Error != nil {
		dbErr := common.NewError(common.CodeDbErrSqlSyntax, fmt.Sprintf("数据库报错-%s", result.Error.Error()))
		panic(dbErr)
	}
}


// 根据搜索条件删除多行
//	example:
//		var where = map[string]interface{}{"fieldName": "value"}
//		DB.DeleteBatchRows(where)
func (DB *MysqlDb) DeleteBatchRows(where map[string]interface{}) {

	objDb := DB.dbconn.Table(DB.Table)
	for fieldName, val := range where {
		if reflect.TypeOf(val).Kind() == reflect.Slice || reflect.TypeOf(val).Kind() == reflect.Array {
			objDb = objDb.Where(fieldName+" IN (?)", val)
		} else {
			objDb = objDb.Where(fieldName+" = ?", val)
		}
	}
	result := objDb.Delete(nil)
	if result.Error != nil {
		dbErr := common.NewError(common.CodeDbErrSqlSyntax, fmt.Sprintf("数据库报错-%s", result.Error.Error()))
		panic(dbErr)
	}
}
```


### Log
Gin web framework output all msg to console will log into log/console.log.
And framework will clear the log seven days ago
```
/**
设置日志格式，和输出重定向到文件，并且清理过期日志
 */
log.SetFlags(log.LstdFlags | log.Llongfile |log.LUTC)
consoleLogFileName := cfg.GetValue("console.log", "APP_CONSOLE_LOG_FILE")
now := time.Now()
logFile := fmt.Sprintf("%s/%s.%s.log", common.ProjectLogPath, consoleLogFileName, now.Format("20060102"))
logWriter, e := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0755)
if e != nil {
    log.Println("打开日志文件错误")
} else {
    gin.DefaultWriter = io.Writer(logWriter)
    log.SetOutput(gin.DefaultWriter)
}
checkLastDate := time.Now().AddDate(0, 0, -7)
checkLastDateInt, _ := strconv.Atoi(checkLastDate.Format("20060102"))
filepath.Walk(common.ProjectLogPath, func (path string, info os.FileInfo, err error) error {
    if pos := strings.Index(path, consoleLogFileName); pos != -1 {
        logFileSliceString := strings.Split(path, ".")
        for _, str := range logFileSliceString {
            date , _ := strconv.Atoi(str)
            if date != 0 && date < checkLastDateInt {
                os.Remove(path)
            }
        }
    }
    return nil
})
```

If request the api, framework will log the access into log/access.log.
```
// 日志格式
logger.WithFields(logrus.Fields{
    "http_status_code"  : statusCode,
    "latency_time" : latencyTime,
    "client_ip"    : clientIP,
    "req_method"   : reqMethod,
    "req_uri"      : reqUri,
    "req_post_data": reqPostData,
    "response_time": endTime,
    "response_code": responseCode,
    "response_msg" : responseMsg,
    "response_data": responseData,
    "response_body": responseBody,
}).Info()

// Output:
// {"client_ip":"127.0.0.1","http_status_code":200,"latency_time":0,"level":"info","msg":"","req_method":"GET","req_post_data":"","req_uri":"/index","response_body":"","response_code":0,"response_data":null,"response_msg":"","response_time":"2019-12-31T16:07:48.941439+08:00","time":"2019-12-31 16:07:48"}
```


### Exception
In request api you can panic anywhere.
```
// router/middleware/exception/exception.go

func SetUp() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				// 处理异常
				log.Println("【Error】系统错误：")
				log.Println(fmt.Sprintf("ErrorMsg：%s", err))
				log.Println(fmt.Sprintf("RequestURL：%s  %s%s", c.Request.Method, c.Request.Host, c.Request.RequestURI))
				log.Println(fmt.Sprintf("RequestUA：%s", c.Request.UserAgent()))
				log.Println("DebugStack：")
				log.Println(string(debug.Stack()))

				utilGin := response.Gin{Ctx: c}
				errMsg := "系统异常"
				cfg := common.NewConfig()
				appMode := cfg.GetValue("", "APP_ENV_MODE")
				// 如果打开测试开关，将详细信息返回至响应中
				if appMode == "debug" {
					errMsg = errMsg + fmt.Sprintf("：%s", err)
				}

				code := common.CodeSystemError
				switch t:=err.(type) {
				case *common.Error:
					code = t.Code
				}
				utilGin.Response(code, errMsg, nil)
			}
		}()
		c.Next()
	}
}
```