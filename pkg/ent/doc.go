// ent 文档：https://entgo.io/docs

// 1.初始化实体对象，会在命令运行目录下创建 ent 目录
// go run entgo.io/ent/cmd/ent init User

// 2.显示声明表字段类型：/ent/schema/User.go
/**
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("age").
			Positive(),
		field.String("name").
			Default("unknown"),
	}
}
 */

// 3. 生成数据库操作代码
// ent generate ./ent/schema

// 4. 查看所有声明结构体
// go run entgo.io/ent/cmd/ent describe ./ent/schema

package ent

