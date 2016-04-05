package controllers

import (
	"encoding/json"
	"github.com/shawncode/highlego-api/models"
	"strconv"
	"github.com/astaxie/beego"
	"github.com/shawncode/highlego-api/cnst"
)

// 抽奖接口
type LotteryController struct {
	beego.Controller
}

func (c *LotteryController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Title Post
// Description create Lottery
// Param	body		body 	models.Lottery	true		"body for Lottery content"
// Success 201 {int} models.Lottery
// Failure 403 body is empty
// router / [post]
func (c *LotteryController) Post() {
	var v models.Lottery
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddLottery(&v); err == nil {
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

// @Title Get
// @Description get Lottery by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Lottery
// @Failure 403 :id is empty
// @router /:id [get]
func (c *LotteryController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetLotteryById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// Title GetExpired
// Description get Get Expired Lotteries
// Success 200 {object} models.Lottery
// router /expired/ [get]
func (c *LotteryController) GetExpired() {
	ml, err := models.GetExpired()
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = ml
	}
	c.ServeJSON()
}

// @Title Get lottery by goods id
// @Description Get Lottery By Goods id
// @Param	gid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Lottery
// @router /goods/:gid [get]
func (c *LotteryController) GetLotteryByGid() {
	idStr := c.Ctx.Input.Param(":gid")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetLotteryByGid(id)
	if err != nil {
		c.Data["json"] = cnst.Error()
	} else {
		ret := cnst.Succ()
		ret.Value = v
		c.Data["json"] = ret
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get Lottery
// @Param	start	query	string	false	"Start position of result set. Must be an integer"
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Success 200 {object} models.Lottery
// @Failure 403
// @router / [get]
func (c *LotteryController) GetAll() {
	var limit int64 = 10
	var offset int64 = 0

	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("start"); err == nil {
		offset = v
	}
	l, err := models.GetAllLottery(offset, limit)
	if err != nil {
		c.Data["json"] = cnst.Error()
	} else {
		ret := cnst.Succ()
		ret.Value = l
		c.Data["json"] = ret
	}
	c.ServeJSON()
}

// Title Update
// Description update the Lottery
// Param	id		path 	string	true		"The id you want to update"
// Param	body		body 	models.Lottery	true		"body for Lottery content"
// Success 200 {object} models.Lottery
// Failure 403 :id is not int
// router /:id [put]
func (c *LotteryController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Lottery{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateLotteryById(&v); err == nil {
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
// Description delete the Lottery
// Param	id		path 	string	true		"The id you want to delete"
// Success 200 {string} delete success!
// Failure 403 id is empty
// router /:id [delete]
func (c *LotteryController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteLottery(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
