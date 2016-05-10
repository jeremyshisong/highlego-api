package controllers

import (
	"github.com/astaxie/beego"
	"github.com/shawncode/highlego-api/cnst"
	"github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/charge"
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"github.com/shawncode/highlego-api/models"
	"time"
	"encoding/json"
	"strings"
)

// 支付凭证接口
type ChargeController struct {
	beego.Controller
}


func (c *ChargeController) URLMapping() {
	c.Mapping("GetCharge", c.GetCharge)
	c.Mapping("PostHooks", c.PostHooks)
}

// @Title Get
// @Description 获取支付凭证
// @Param	orderNo		query 	string	false		"订单号"
// @Param	amount		query 	string	false		"支付金额"
// @Param	channel		query 	string	false		"支付渠道,如:alipay"
// @Success 200
// @Failure 403 :orderNo is empty
// @router / [get]
func (c *ChargeController) GetCharge() {
	orderNo := c.Ctx.Input.Query("orderNo")
	amountStr := c.Ctx.Input.Query("amount")
	channel := c.Ctx.Input.Query("channel")
	amount, _ := strconv.ParseUint(amountStr, 0, 64)
	pingpp.Key = "sk_test_5mXfL4en1iP00Cujz5r1KuzP"
	params := &pingpp.ChargeParams{
		Order_no:  orderNo,
		App:       pingpp.App{Id: "app_4G4e1GH8OqfH08CW"},
		Amount:    amount,
		Channel:   channel,
		Currency:  "cny",
		Client_ip: strings.Split(c.Ctx.Request.RemoteAddr, ":")[0],
		Subject:   "Your Subject",
		Body:      "Your Body",
	}
	//返回的第一个参数是 charge 对象，你需要将其转换成 json 给客户端，或者客户端接收后转换。
	ch, err := charge.New(params)
	if err != nil {
		ret := cnst.Error()
		ret.Message="get charge error"
		c.Data["json"] = params
	} else {
		ret := cnst.Succ()
		ret.Value = ch
		c.Data["json"] = ret
	}
	c.ServeJSON()
}

// @Title Post
// @Description 获取支付回调
// @Param	body body 	pingpp.webhook	true		"body for webhook content"
// @Success 200
// Failure 403 body is empty
// @router / [post]
func (c *ChargeController) PostHooks() {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Ctx.Request.Body)
	//signature := c.Ctx.Request.Header.Get("x-pingplusplus-signature")
	webhook, err := pingpp.ParseWebhooks(buf.Bytes())
	fmt.Println(webhook.Type)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(c.Ctx.ResponseWriter, "fail")
		return
	}

	if webhook.Type == "charge.succeeded" {
		// TODO your code for charge
		ret_order := webhook.Data.Object["order_no"]
		ret_channel := webhook.Data.Object["channel"]
		ret_pay_time := webhook.Data.Object["time_paid"]
		ret_paid := webhook.Data.Object["paid"]
		c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
		order_id, _ := strconv.Atoi(ret_order.(string))
		v, err := models.GetLotteryOrderById(order_id)
		if err == nil {
			v.PayType = ret_channel.(string)
			time_stamp, _ := json.Number.Int64(ret_pay_time.(json.Number))
			v.PayTime = time.Unix(time_stamp, 0)
			v.Status = "wait"
			if !ret_paid.(bool) {
				v.Status = "fail"
			}
			if err := models.UpdateLotteryOrderById(v); err == nil {
				c.Data["json"] = v
			} else {
				c.Data["json"] = err.Error()
			}
		} else {
			c.Data["json"] = err.Error()
		}
		c.ServeJSON()
	} else if webhook.Type == "refund.succeeded" {
		// TODO your code for refund
		c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	} else {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	}
}
