package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"errors"
	"fmt"
)

type LotteryOrder struct {
	Id             int       `orm:"column(rid);pk" description:"夺宝订单id"`
	Gid            int       `orm:"column(gid);null" description:"夺宝商品id"`
	Tid        	   int       `orm:"column(tid);type(date);null" description:"夺宝标题"`
	Status         string    `orm:"column(status);null" description:"夺宝状态:wait=等待开奖,published=已开奖,win=中奖,fail=未中奖"`
	PublishTime    time.Time `orm:"column(publishTime);null" description:"夺宝公布时间"`
	Uid            int       `orm:"column(uid);null" description:"夺宝订单用户id"`
	LotteryNo      string    `orm:"column(lotteryNo);null" description:"夺宝号码"`
	PayType		   string	 `orm:"column(payType);null" description:"夺宝订单支付方式"`
	PayTime		   time.Time `orm:"column(payTime);null" description:"夺宝订单支付时间"`
}


func (t *LotteryOrder) TableName() string {
	return "lotteryOrder"
}

func init() {
	orm.RegisterModel(new(LotteryOrder))
}

// AddLotteryOrder insert a new LotteryOrder into database and returns
// last inserted Id on success.
func AddLotteryOrder(m *LotteryOrder, u *User) (id int64, err error) {
	o := orm.NewOrm()
	o.Begin()
	o.Update(u)
	o.Insert(m);
	if errC := o.Commit(); errC != nil {
		errR := o.Rollback()
		err = errR
	}
	return
}

// GetLotteryOrderById retrieves LotteryOrder by Id. Returns error if
// Id doesn't exist
func GetLotteryOrderById(id int) (v *LotteryOrder, err error) {
	o := orm.NewOrm()
	v = &LotteryOrder{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetLotteryOrderById retrieves LotteryOrder by Uid. Returns error if
// Id doesn't exist
func GetLotteryOrderByUId(id int,status string) (ml []LotteryOrder, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(LotteryOrder))
	qs = qs.Filter("Uid", id)
	qs = qs.Filter("Status", status)
	if _, err := qs.Limit(50, 0).All(&ml); err != nil {
		return nil, errors.New("Error:get all published LotteryOrders occur an error")
	}
	return ml, nil
}

// GetLotteryOrderById retrieves LotteryOrder by Uid. Returns error if
// Id doesn't exist
func GetAllLotteryOrderUId(id int) (ml []LotteryOrder, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(LotteryOrder))
	qs = qs.Filter("Uid", id)
	if _, err := qs.Limit(50, 0).All(&ml); err != nil {
		return nil, errors.New("Error:get all published LotteryOrders occur an error")
	}
	return ml, nil
}



	// GetAllLotteryOrder retrieves all LotteryOrder matches certain condition. Returns empty list if
// no records exist
func GetAllLotteryOrder(offset int64, limit int64) (ml []LotteryOrder, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(LotteryOrder))
	if _, err := qs.Limit(limit, offset).All(&ml); err == nil {
		return ml, nil
	}
	return nil, err
}

// UpdateLotteryOrder updates LotteryOrder by Id and returns error if
// the record to be updated doesn't exist
func UpdateLotteryOrderById(m *LotteryOrder) (err error) {
	o := orm.NewOrm()
	v := LotteryOrder{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteLotteryOrder deletes LotteryOrder by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLotteryOrder(id int) (err error) {
	o := orm.NewOrm()
	v := LotteryOrder{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&LotteryOrder{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
