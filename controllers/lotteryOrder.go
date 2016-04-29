package controllers

import (
	"encoding/json"
	"github.com/shawncode/highlego-api/models"
	"strconv"
	"github.com/astaxie/beego"
	"github.com/shawncode/highlego-api/cnst"
)

// 夺宝订单接口
type LotteryOrderController struct {
	beego.Controller
}

type Params struct {
	Gid     int       `required:"true" description:"夺宝商品id"`
	Tid     int       `required:"true" description:"夺宝标题"`
	Uid     int       `required:"true" description:"夺宝订单用户id"`
	PayType string     `required:"true" description:"夺宝订单支付方式"`
}

func (c *LotteryOrderController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description 新增一条夺宝订单
// @Param	body	body  controllers.Params 	true		"body for Object content"
// @Success 201 {object}  models.LotteryOrder
// @Success 402 {object}  controllers.Params
// @Failure 403 controllers.Params is null
// @router / [post]
func (c *LotteryOrderController) Post() {
	var params Params
	ret := cnst.Error()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &params); err != nil {
		ret.Message = err.Error()
		c.Data["json"] = ret
		c.ServeJSON()
		return
	}

	var v models.LotteryOrder
	v.Gid = params.Gid
	v.Tid = params.Tid
	v.PayType = params.PayType
	v.Uid = params.Uid


	u, err := models.GetUserById(v.Uid);
	if err != nil {
		ret.Message = err.Error()
		c.Data["json"] = ret
		c.ServeJSON()
		return
	}

	g, err := models.GetGoodsById(v.Gid)
	if err != nil {
		ret.Message = err.Error()
		c.Data["json"] = ret
		c.ServeJSON()
		return
	}

	if u.Coins < g.Punit {
		ret := cnst.Error()
		ret.Code = cnst.CodeNoEnoughCoins
		ret.Message = "not enough coins"
		c.Data["json"] = ret
		c.ServeJSON()
		return
	}


	ml, err := models.GetLotteryByGid(g.Id);
	if err != nil {
		ret.Message = err.Error()
		c.Data["json"] = ret
		c.ServeJSON()
		return
	} else if len(ml) == 0 {
		ret.Message = "not invidal Lottery for orders"
		c.Data["json"] = ret
		c.ServeJSON()
		return
	}

	l := ml[0]

	v.Tid = l.Id
	v.PublishTime = l.ExpiredTime
	v.Gid = g.Id
	v.Uid = u.Id

	v.LotteryNo = "121221212"

	if _, err := models.AddLotteryOrder(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		ret := cnst.Succ()
		ret.Value = v
		c.Data["json"] = ret
	} else {
		c.Data["json"] = cnst.Error()
	}
	c.ServeJSON()
}


// @Title Get
// @Description get LotteryOrder by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.LotteryOrder
// @Failure 403 :id is empty
// @router /:id [get]
func (c *LotteryOrderController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetLotteryOrderById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}
// @Title Get Unpublished by uid
// @Description 用户待开奖订单
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.LotteryOrder
// @Failure 403 :uid is empty
// @router /wait/user/:uid [get]
func (c *LotteryOrderController) GetWaitByUser() {
	idStr := c.Ctx.Input.Param(":uid")
	id, _ := strconv.Atoi(idStr)
	ret := cnst.Succ()
	v, err := models.GetLotteryOrderByUId(id,"wait")
	if err != nil {
		ret := cnst.Error()
		ret.Message = err.Error()
		c.Data["json"] = ret
	} else {
		ret.Value = v
		c.Data["json"] = ret
	}
	c.ServeJSON()
}

// @Title Get
// @Description 用户全部订单
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.LotteryOrder
// @Failure 403 :uid is empty
// @router /all/user/:uid [get]
func (c *LotteryOrderController) GetPublishedByUser() {
	idStr := c.Ctx.Input.Param(":uid")
	id, _ := strconv.Atoi(idStr)
	ret := cnst.Succ()
	v, err := models.GetAllLotteryOrderUId(id)
	if err != nil {
		ret := cnst.Error()
		ret.Message = err.Error()
		c.Data["json"] = ret
	} else {
		ret.Value = v
		c.Data["json"] = ret
	}
	c.ServeJSON()
}

// @Title Get
// @Description 用户中奖订单
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.LotteryOrder
// @Failure 403 :uid is empty
// @router /win/user/:uid [get]
func (c *LotteryOrderController) GetWinByUser() {
	idStr := c.Ctx.Input.Param(":uid")
	id, _ := strconv.Atoi(idStr)
	ret := cnst.Succ()
	v, err := models.GetLotteryOrderByUId(id,"win")
	if err != nil {
		ret := cnst.Error()
		ret.Message = err.Error()
		c.Data["json"] = ret
	} else {
		ret.Value = v
		c.Data["json"] = ret
	}
	c.ServeJSON()
}

// @Title Get
// @Description 用户未中奖订单
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.LotteryOrder
// @Failure 403 :uid is empty
// @router /fail/user/:uid [get]
func (c *LotteryOrderController) GetFailByUser() {
	idStr := c.Ctx.Input.Param(":uid")
	id, _ := strconv.Atoi(idStr)
	ret := cnst.Succ()
	v, err := models.GetLotteryOrderByUId(id,"fail")
	if err != nil {
		ret := cnst.Error()
		ret.Message = err.Error()
		c.Data["json"] = ret
	} else {
		ret.Value = v
		c.Data["json"] = ret
	}
	c.ServeJSON()
}

// Title Get All
// Description get LotteryOrder
// Param	start	query	string	false	"Start position of result set. Must be an integer"
// Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// Success 200 {object} models.LotteryOrder
// Failure 403
// router / [get]
func (c *LotteryOrderController) GetAll() {
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

	l, err := models.GetAllLotteryOrder(offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Title Update
// Description update the LotteryOrder
// Param	id		path 	string	true		"The id you want to update"
// Param	body		body 	models.LotteryOrder	true		"body for LotteryOrder content"
// Success 200 {object} models.LotteryOrder
// Failure 403 :id is not int
// router /:id [put]
func (c *LotteryOrderController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.LotteryOrder{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateLotteryOrderById(&v); err == nil {
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
// Description delete the LotteryOrder
// Param	id		path 	string	true		"The id you want to delete"
// Success 200 {string} delete success!
// Failure 403 id is empty
// router /:id [delete]
func (c *LotteryOrderController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteLotteryOrder(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
