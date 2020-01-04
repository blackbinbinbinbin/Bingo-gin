package index

type IndexValidator struct {
	Name	string    `form:"name" binding:"required"`
	Id  	string    `form:"id" binding:"required"`
}