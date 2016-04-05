package controllers

import (
	"encoding/json"
	"github.com/shawncode/highlego-api/models"
	"strconv"
	"strings"
	"github.com/astaxie/beego"
	"github.com/shawncode/highlego-api/cnst"
	"time"
)

// 用户接口
type UserController struct {
	beego.Controller
}

func (c *UserController) URLMapping() {
	c.Mapping("Post", c.AddUser)
	c.Mapping("GetUserById", c.GetUserById)
	c.Mapping("GetUserByDeviceId", c.GetUserByDeviceId)
	c.Mapping("Put", c.Put)
}


// @Title Login
// @Description 用户通过设备id登陆
// @Param	deviceId		path 	string	true		"user login"
// @Success 200 {object} models.User
// @Failure 403 deviceId is empty
// @router /login/:deviceId [get]
func (c *UserController) Login() {
	deviceId := c.Ctx.Input.Param(":deviceId")
	ret := &cnst.Result{}

	v, err := models.GetUserByDeviceId(deviceId)
	if err != nil {
		ret.Code = cnst.CodeQueryDBFail
		ret.Message = err.Error()
		ret.Value = ""

		c.Data["json"] = ret
	} else {
		if v == nil {
			ip := strings.Split(c.Ctx.Request.RemoteAddr, ":")[0]
			v = &models.User{DeviceId:deviceId, RequestIp:ip}
			models.AddUser(v)
			ret.Code = cnst.CodeRegOK
			ret.Message = cnst.MsgREGOK
			ret.Value = v
			c.Data["json"] = ret
			c.ServeJSON()
		}
		ret.Code = cnst.CodeLoginOK
		ret.Message = cnst.MsgLoginOK
		ret.ServerTime = time.Now().Unix()
		ret.Value = v
		c.Data["json"] = ret
	}
	c.ServeJSON()
}
// Title AddUser
// Description 用户注册
// Param	body		body 	models.User	true		"body for User content"
// Success 201 {int} models.User
// Failure 403 body is empty
// router / [post]
func (c *UserController) AddUser() {
	var v models.User
	ip := strings.Split(c.Ctx.Request.RemoteAddr, ":")[0]
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.RequestIp = ip
		if _, err := models.AddUser(&v); err == nil {
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

// @Title 更新用户信息
// @Description 更新用户信息
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.User	true		"body for User content"
// @Success 200 {object} models.User
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UserController) UpdateUser() {
	idStr := c.Ctx.Input.Param(":id")
	ret := &cnst.Result{}
	id, _ := strconv.Atoi(idStr)
	v := models.User{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateUserById(&v); err == nil {
			ret.Code = cnst.CodeUpdateOK
			ret.Message = cnst.MsgUpdateOK
			ret.Value,err = models.GetUserById(id)
			c.Data["json"] = ret
		} else {
			ret.Code = cnst.CodeUpdateFail
			ret.Message = err.Error()
			ret.Value = ""
			c.Data["json"] = ret
		}
	} else {
		ret.Code = cnst.CodeUpdateFail
		ret.Message = err.Error()
		ret.Value = ""
		c.Data["json"] = ret
	}
	c.ServeJSON()
}

// Title GetUserById
// Description get User by id
// Param	id		path 	string	true		"The key for staticblock"
// Success 200 {object} models.User
// Failure 403 :id is empty
// router /id/:id [get]
func (c *UserController) GetUserById() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUserById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		v.Birthday.Format("20060102")
		c.Data["json"] = v
	}
	c.ServeJSON()
}
// Title GetUserByDeviceId
// Description get User by DeviceId
// Param	deviceId		path 	string	true		"The key for staticblock"
// Success 200 {object} models.User
// Failure 403 :deviceId is empty
// router /deviceId/:deviceId [get]
func (c *UserController) GetUserByDeviceId() {
	idStr := c.Ctx.Input.Param(":deviceId")
	v, err := models.GetUserByDeviceId(idStr)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}






