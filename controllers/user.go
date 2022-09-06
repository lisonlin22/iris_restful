package controllers

import (
	"gdopo.com/database"
	"gdopo.com/models"
	"gdopo.com/public"
	"github.com/kataras/iris/v12"
)

// 接口
func User(p iris.Party) {
	p.Get("/", list)
	p.Post("/", create)
	p.Put("/{id:uint64}", update)
	p.Get("/{id:uint64}", retrieve)
	p.Delete("/{id:uint64}", delete)
}

// 信息
func retrieve(c iris.Context) {
	respone := public.Result{}
	user := models.User{}
	queryset := database.DB.Find(&user, c.Params().Get("id"))
	if queryset.Error != nil {
		respone.Code = 1
		respone.Msg = "未找到"
		c.JSON(respone.Format())
		return
	}
	respone.Data = user
	c.JSON(respone.Format())
}

// 创建
func create(c iris.Context) {
	params := models.User{}
	respone := public.Result{}

	// 读取表单
	if err := c.ReadJSON(&params); err != nil {
		respone.Code = 1
		respone.Msg = "参数不正确"
		c.JSON(respone.Format())
		return
	}
	// 判断是否创建成功
	if queryset := database.DB.Save(&params).Last(&params).Error; queryset != nil {
		respone.Code = 1
		respone.Msg = "创建失败"
		c.JSON(respone.Format())
		return
	}

	// 创建成功后返回最后一条数据
	database.DB.Last(&params)
	respone.Data = params
	c.JSON(respone.Format())
}

// 列表
func list(c iris.Context) {
	respone := public.Result{}
	user := []models.User{}
	queryset := database.DB.Find(&user)
	if queryset.Error != nil {
		respone.Code = 1
		respone.Msg = "未找到"
		c.JSON(respone.Format())
		return
	}
	respone.Data = user
	c.JSON(respone.Format())
}

// 删除
func delete(c iris.Context) {
	respone := public.Result{}
	user := models.User{}
	queryset := database.DB.Delete(&user, c.Params().Get("id"))
	if queryset.Error != nil {
		respone.Code = 1
		respone.Msg = "删除失败"
		c.JSON(respone.Format())
		return
	}
	c.JSON(respone.Format())
}

// 修改
func update(c iris.Context) {
	params := models.User{}
	respone := public.Result{}
	id := c.Params().Get("id")
	// 读取表单
	if err := c.ReadJSON(&params); err != nil {
		respone.Code = 1
		respone.Msg = "参数不正确"
		c.JSON(respone.Format())
		return
	}
	// 判断是否创建成功
	if queryset := database.DB.Model(&params).Where("id = ?", id).Update(params).Error; queryset != nil {
		respone.Code = 1
		respone.Msg = "更新失败"
		c.JSON(respone.Format())
		return
	}

	respone.Data = params
	c.JSON(respone.Format())
}
