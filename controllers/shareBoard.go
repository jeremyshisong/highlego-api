package controllers

import (
	"encoding/json"
	"errors"
	"github.com/shawncode/highlego-api/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// 晒单接口[待实现]
type ShareBoardController struct {
	beego.Controller
}

func (c *ShareBoardController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Title Post
// Description create ShareBoard
// Param	body		body 	models.ShareBoard	true		"body for ShareBoard content"
// Success 201 {int} models.ShareBoard
// Failure 403 body is empty
// router / [post]
func (c *ShareBoardController) Post() {
	var v models.ShareBoard
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddShareBoard(&v); err == nil {
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

// Title Get
// Description get ShareBoard by id
// Param	id		path 	string	true		"The key for staticblock"
// Success 200 {object} models.ShareBoard
// Failure 403 :id is empty
// router /:id [get]
func (c *ShareBoardController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetShareBoardById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// Title Get All
// Description get ShareBoard
// Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// Param	offset	query	string	false	"Start position of result set. Must be an integer"
// Success 200 {object} models.ShareBoard
// Failure 403
// router / [get]
func (c *ShareBoardController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllShareBoard(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Title Update
// Description update the ShareBoard
// Param	id		path 	string	true		"The id you want to update"
// Param	body		body 	models.ShareBoard	true		"body for ShareBoard content"
// Success 200 {object} models.ShareBoard
// Failure 403 :id is not int
// router /:id [put]
func (c *ShareBoardController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.ShareBoard{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateShareBoardById(&v); err == nil {
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
// Description delete the ShareBoard
// Param	id		path 	string	true		"The id you want to delete"
// Success 200 {string} delete success!
// Failure 403 id is empty
// router /:id [delete]
func (c *ShareBoardController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteShareBoard(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
