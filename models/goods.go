package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Goods struct {
	Id            int     `orm:"column(gid);auto" description:"商品id"`
	Gname         string  `orm:"column(gname);size(20);null" description:"商品名称"`
	Desc          string  `orm:"column(desc);size(20);null" description:"商品描述"`
	Gpic          string  `orm:"column(goodsPic);size(60);null" description:"商品图片链接"`
	Price         float64 `orm:"column(price);null;digits(10);decimals(2)" description:"商品价格"`
	Punit         int     `orm:"column(priceUnit);null" description:"商品价格单位"`
	regularBuyMax int     `orm:"column(regularBuyMax);null" description:"最大可买数量"`
	Gtype         int     `orm:"column(goodsType);null" description:"商品类型"`
	Gstatus       string  `orm:"column(goodsStatus);null" description:"商品状态:on=上架,off=下架"`
}

func (t *Goods) TableName() string {
	return "goods"
}

func init() {
	orm.RegisterModel(new(Goods))
}

// AddGoods insert a new Goods into database and returns
// last inserted Id on success.
func AddGoods(m *Goods) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGoodsById retrieves Goods by Id. Returns error if
// Id doesn't exist
func GetGoodsById(id int) (v *Goods, err error) {
	o := orm.NewOrm()
	v = &Goods{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllGoods retrieves all Goods matches certain condition. Returns empty list if
// no records exist
func GetAllGoods(
	offset int64, limit int64) (ml []Goods, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Goods))
	qs = qs.Filter("Gstatus", "on")
	// order by:
	var sortFields []string
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&ml); err == nil {
		return ml, nil
	}
	return nil, err
}

// UpdateGoods updates Goods by Id and returns error if
// the record to be updated doesn't exist
func UpdateGoodsById(m *Goods) (err error) {
	o := orm.NewOrm()
	v := Goods{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGoods deletes Goods by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGoods(id int) (err error) {
	o := orm.NewOrm()
	v := Goods{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Goods{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// GetGoodsBy cateId retrieves Goods by cateId. Returns error if
// Cid doesn't exist
func GetGoodsByCate(cid int) (ml []Goods,err error) {
	o := orm.NewOrm()
	c := GoodsCate{Id: cid}
	// ascertain id exists in the database
	if err = o.Read(&c); err == nil {
		qs := o.QueryTable(new(Goods))
		qs =qs.Filter("Gtype",cid)
		if _, err := qs.Limit(50, 0).All(&ml); err != nil {
			return nil, errors.New("Error:get matched cate goods occur an error")
		}
		return ml, nil
	}
	return nil,errors.New("Error:cid error")
}
