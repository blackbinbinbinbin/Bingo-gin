package controller

import (
	"github.com/gin-gonic/gin"
	"Bingo-gin/controller/validator/index"
	"Bingo-gin/common"
	"Bingo-gin/model"
	"strconv"
)

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