package common

const (
	// 请求成功
	CodeSuccess = 200

	// 系统错误
	CodeSystemError = 500

	// 数据库报错-语法错误
	CodeDbErrSqlSyntax     = 49001
	// 数据库报错-事务已开启
	CodeDbErrTxHasBegan    = 49002
	// 数据库报错-事务已结束
	CodeDbErrTxDone        = 49003
	// 数据库报错-数据返回多行
	CodeDbErrMultiRows     = 49004
	// 数据库报错-没数据返回
	CodeDbErrNoRows        = 49005
	// 数据库报错-游标已关闭
	CodeDbErrStmtClosed    = 49006
	// 数据库报错-参数错误
	CodeDbErrArgs          = 49007
	// 数据库报错-执行错误
	CodeDbErrNotImplement  = 49008

	// 常规错误
	CodeParamsError	= 40001		//参数错误
)