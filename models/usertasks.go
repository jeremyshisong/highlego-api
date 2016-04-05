package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Usertasks struct {
	Id                int       `orm:"column(task_id);auto"`
	Uid               *User     `orm:"column(uid);rel(fk)"`
	DailyShare        string    `orm:"column(daily_share);null"`
	LastCheckInTime   time.Time `orm:"column(last_check_in_time);type(datetime);null"`
	CheckInDays       int       `orm:"column(check_in_days);null"`
	AddrCompleted     string    `orm:"column(addr_completed);null"`
	Registered        string    `orm:"column(registered);null"`
	UserinfoCompleted string    `orm:"column(userinfo_completed);null"`
	TotalCredits      int       `orm:"column(total_credits);null"`
}

func (t *Usertasks) TableName() string {
	return "usertasks"
}

func init() {
	orm.RegisterModel(new(Usertasks))
}

// AddUsertasks insert a new Usertasks into database and returns
// last inserted Id on success.
func AddUsertasks(m *Usertasks) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUsertasksById retrieves Usertasks by Id. Returns error if
// Id doesn't exist
func GetUsertasksById(id int) (v *Usertasks, err error) {
	o := orm.NewOrm()
	v = &Usertasks{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUsertasks retrieves all Usertasks matches certain condition. Returns empty list if
// no records exist
func GetAllUsertasks(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Usertasks))
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

	var l []Usertasks
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

// UpdateUsertasks updates Usertasks by Id and returns error if
// the record to be updated doesn't exist
func UpdateUsertasksById(m *Usertasks) (err error) {
	o := orm.NewOrm()
	v := Usertasks{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUsertasks deletes Usertasks by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUsertasks(id int) (err error) {
	o := orm.NewOrm()
	v := Usertasks{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Usertasks{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
