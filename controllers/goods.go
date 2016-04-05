package controllers

import (
	"encoding/json"
	"github.com/shawncode/highlego-api/models"
	"strconv"
	"github.com/astaxie/beego"
	"github.com/shawncode/highlego-api/cnst"
)

// 抽奖商品接口
type GoodsController struct {
	beego.Controller
}

func (c *GoodsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Get
// @Description 通过商品id获取商品详情
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Goods
// @Failure 403 :id is empty
// @router /:id [get]
func (c *GoodsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetGoodsById(id)
	if err != nil {
		c.Data["json"] = cnst.Error()
	} else {
		ret := cnst.Succ()
		ret.Value = v
		c.Data["json"] = ret
	}
	c.ServeJSON()
}

// @Title goodsCate
// @Description 获取商品分类接口
// @Success 200 {object} models.GoodsCate
// @Failure 403 :db query error
// @router /cates [get]
func (c *GoodsController) GetALLGoodsCate() {
	ml, err := models.GetAllGoodsCate()
	if err != nil {
		c.Data["json"] = cnst.Error()
	} else {
		ret := cnst.Succ()
		ret.Value = ml
		c.Data["json"] = ret
	}
	c.ServeJSON()
}

// @Title goodsByCate
// @Description 按分类获取商品接口
// @Param	cid		path 	string	true		"The cate id you want to select"
// @Success 200 {object} models.Goods
// @Failure 403 :db query error
// @router /cate/:cid [get]
func (c *GoodsController) GetGoodsByCate() {
	idStr := c.Ctx.Input.Param(":cid")
	cid, _ := strconv.Atoi(idStr)
	ml, err := models.GetGoodsByCate(cid)
	if err != nil {
		c.Data["json"] = cnst.Error()
	} else {
		ret := cnst.Succ()
		ret.Value = ml
		c.Data["json"] = ret
	}
	c.ServeJSON()
}


// Title Update
// Description update the Goods
// Param	id		path 	string	true		"The id you want to update"
// Param	body		body 	models.Goods	true		"body for Goods content"
// Success 200 {object} models.Goods
// Failure 403 :id is not int
// router /:id [put]
func (c *GoodsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Goods{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateGoodsById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Title Delete
// Description delete the Goods
// Param	id		path 	string	true		"The id you want to delete"
// Success 200 {string} delete success!
// Failure 403 id is empty
// router /:id [delete]
func (c *GoodsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteGoods(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Title Post
// Description create Goods
// Param	body		body 	models.Goods	true		"body for Goods content"
// Success 201 {int} models.Goods
// Failure 403 body is empty
// router / [post]
func (c *GoodsController) Post() {
	var v models.Goods
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddGoods(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Get All
// @Description 获取商品列表
// @Param	start	query	string	false	"Start position of result set. Must be an integer"
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Success 200 {object} models.Goods
// @Failure 403
// @router / [get]
func (c *GoodsController) GetAll() {
	var limit int64 = 10
	var offset int64 = 0

	l, err := models.GetAllGoods(offset, limit)
	if err != nil {
		ret := cnst.Error()
		ret.Message = err.Error()
		c.Data["json"] = ret
	} else {
		ret := cnst.Succ()
		ret.Value = l
		c.Data["json"] = ret
	}
	c.ServeJSON()
}
