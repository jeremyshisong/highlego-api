package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type GoodsCate struct {
	Id    int    `orm:"column(id);auto" description:"商品类型表id"`
	Cid   int    `orm:"column(cid);null" description:"商品类型id"`
	Cname string `orm:"column(cname);size(20);null" description:"商品类型名称"`
	desc  string `orm:"column(desc);size(20);null" description:"商品类型描述"`
	icon  string `orm:"-;column(icon);size(100);null" description:"商品类型图标链接"`
}

func (t *GoodsCate) TableName() string {
	return "goodsCate"
}

func init() {
	orm.RegisterModel(new(GoodsCate))
}

// AddGoodsCate insert a new GoodsCate into database and returns
// last inserted Id on success.
func AddGoodsCate(m *GoodsCate) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGoodsCateById retrieves GoodsCate by Id. Returns error if
// Id doesn't exist
func GetGoodsCateById(id int) (v *GoodsCate, err error) {
	o := orm.NewOrm()
	v = &GoodsCate{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}


// GetAllGoodsCate retrieves all GoodsCate matches certain condition. Returns empty list if
// no records exist
func GetAllGoodsCate() (ml []GoodsCate, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GoodsCate))
	if _, err := qs.Limit(50, 0).All(&ml); err != nil {
		return nil, errors.New("Error:get all goods cate occur an error")
	}
	return ml, nil
}

// UpdateGoodsCate updates GoodsCate by Id and returns error if
// the record to be updated doesn't exist
func UpdateGoodsCateById(m *GoodsCate) (err error) {
	o := orm.NewOrm()
	v := GoodsCate{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGoodsCate deletes GoodsCate by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGoodsCate(id int) (err error) {
	o := orm.NewOrm()
	v := GoodsCate{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GoodsCate{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
