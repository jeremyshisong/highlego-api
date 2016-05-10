package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Lottery struct {
	Id          int       `orm:"column(lottery_id);auto" description:"夺宝id"`
	Term        int       `orm:"column(term);null" description:"夺宝期号"`
	Title       string    `orm:"column(title);size(20);null" description:"夺宝标题"`
	Goods       *Goods    `orm:"column(gid);rel(fk)" description:"夺宝商品"`
	LeftOrders  int       `orm:"column(leftOrders);null" description:"剩余参与数量"`
	StartTime   time.Time `orm:"column(startTime);type(datetime);null" description:"过期时间"`
	ExpiredTime time.Time `orm:"column(expiredTime);type(datetime);null" description:"过期时间"`
	Status      string    `orm:"column(status);null" description:"夺宝状态:wait=等待开奖,expired=过期,closed=关闭,published=已开奖`
}

func (t *Lottery) TableName() string {
	return "lottery"
}

func init() {
	orm.RegisterModel(new(Lottery))
}

// AddLottery insert a new Lottery into database and returns
// last inserted Id on success.
func AddLottery(m *Lottery) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetLotteryById retrieves Lottery by Id. Returns error if
// Id doesn't exist
func GetLotteryById(id int) (v *Lottery, err error) {
	o := orm.NewOrm()
	v = &Lottery{Id: id}
	if err = o.Read(v); err == nil {
		o.LoadRelated(v, "Goods")
		return v, nil
	}
	return nil, err
}

// GetLotteryBy Goods Id retrieves Lottery by Goods Id. Returns error if
// Id doesn't exist
func GetLotteryByGid(id int) (ml []Lottery, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Lottery))
	qs = qs.Filter("Goods", id)
	qs = qs.Filter("Status", "wait")
	if _, err := qs.Limit(50, 0).All(&ml); err != nil {
		return nil, errors.New("Error:get all expired lotteries occur an error")
	}
	for i, _ := range ml {
		o.LoadRelated(&ml[i], "Goods")
	}
	return ml, nil
}

// GetAllLottery retrieves all Lottery matches certain condition. Returns empty list if
// no records exist
func GetAllLottery(offset int64, limit int64) (ml []Lottery, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Lottery))
	qs = qs.Filter("Status", "wait")
	if _, err := qs.Limit(limit, offset).All(&ml); err == nil {
		for i, _ := range ml {
			o.LoadRelated(&ml[i], "Goods")
		}
		return ml, nil
	}
	return nil, err
}

// UpdateLottery updates Lottery by Id and returns error if
// the record to be updated doesn't exist
func UpdateLotteryById(m *Lottery) (err error) {
	o := orm.NewOrm()
	v := Lottery{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}


// get Lottery which is expired
// the record to be updated doesn't exist
func GetExpired() (ml []Lottery, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Lottery))
	qs = qs.Filter("Status", "wait")
	qs = qs.Filter("ExpiredTime__lte", time.Now())
	if _, err := qs.Limit(50, 0).All(&ml); err != nil {
		return nil, errors.New("Error:get all expired lotteries occur an error")
	}
	return ml, nil
}


// get Lottery which is expired
// the record to be updated doesn't exist
func UpdateExpired(m *Lottery) (id int64, err error) {
	o := orm.NewOrm()
	v := Lottery{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var goods *Goods
		goods = v.Goods
		rate := (float64)(v.LeftOrders / int(goods.Price))
		v.ExpiredTime = time.Now().Add(time.Duration(2 * 60 * 60 * rate) * time.Second)
		o.Update(v)
	}
	return
}

// DeleteLottery deletes Lottery by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLottery(id int) (err error) {
	o := orm.NewOrm()
	v := Lottery{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Lottery{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
