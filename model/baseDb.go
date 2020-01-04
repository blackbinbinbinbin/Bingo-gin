package model

type BaseDb interface {
	/**
	 * 初始化操作类
	 */
	Init() interface{}

	/**
	 * 查找数据
	 * @params where map[string]interface{}		eg:{"name": "jinzhu"}
	 * @params limit int
	 * @params offset int
	 * @params fields string  eg: "name, age"
	 * @params orderBy string  eg: "age desc, name"
	 */
	Select(map[string]interface{}, int, int, string, string) interface{}
}