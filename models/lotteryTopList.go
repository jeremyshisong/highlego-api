package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type LotteryTopList struct {
	Id                    int      `orm:"column(top_list_id);pk"`
	LotteryId             *Lottery `orm:"column(lottery_id);rel(fk)"`
	RichPlayerId          int      `orm:"column(rich_player_id);null"`
	RichPlayerName        string   `orm:"column(rich_player_name);size(45);null"`
	RichPlayerJoinTimes   int      `orm:"column(rich_player_join_times);null"`
	RichPlayerFaviconUrl  string   `orm:"column(rich_player_favicon_url);size(45);null"`
	FirstPlayerId         int      `orm:"column(first_player_id);null"`
	FirstPlayerName       string   `orm:"column(first_player_name);size(45);null"`
	FirstPlayerFaviconUrl string   `orm:"column(first_player_favicon_url);size(45);null"`
	LastPlayerId          int      `orm:"column(last_player_id);null"`
	LastPlayerName        string   `orm:"column(last_player_name);size(45);null"`
	LastPlayerJoinTimes   int      `orm:"column(last_player_join_times);null"`
	LastPlayerFaviconUrl  string   `orm:"column(last_player_favicon_url);size(45);null"`
}

func (t *LotteryTopList) TableName() string {
	return "lottery_top_list"
}

func init() {
	orm.RegisterModel(new(LotteryTopList))
}

// AddLotteryTopList insert a new LotteryTopList into database and returns
// last inserted Id on success.
func AddLotteryTopList(m *LotteryTopList) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetLotteryTopListById retrieves LotteryTopList by Id. Returns error if
// Id doesn't exist
func GetLotteryTopListById(id int) (v *LotteryTopList, err error) {
	o := orm.NewOrm()
	v = &LotteryTopList{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllLotteryTopList retrieves all LotteryTopList matches certain condition. Returns empty list if
// no records exist
func GetAllLotteryTopList(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(LotteryTopList))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []LotteryTopList
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateLotteryTopList updates LotteryTopList by Id and returns error if
// the record to be updated doesn't exist
func UpdateLotteryTopListById(m *LotteryTopList) (err error) {
	o := orm.NewOrm()
	v := LotteryTopList{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteLotteryTopList deletes LotteryTopList by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLotteryTopList(id int) (err error) {
	o := orm.NewOrm()
	v := LotteryTopList{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&LotteryTopList{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
