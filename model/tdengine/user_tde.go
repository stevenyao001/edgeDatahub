package tdengine

import (
	"fmt"
	"github.com/stevenyao001/edgeCommon/logger"
	"github.com/stevenyao001/edgeCommon/tdengine"
	"time"
)

type UserTd struct {
	*tdengine.TdEngine
	Uid  int       `json:"uid"`
	Name string    `json:"name"`
	Ts   int       `json:"ts"`
	Ts2  time.Time `json:"ts2"`
}

func NewUserTd() *UserTd {
	tde := &UserTd{
		TdEngine: &tdengine.TdEngine{
			Db:        nil,
			InsName:   "rootcloud",
			DbName:    "test",
			TableName: "user1",
		},
	}

	tde.Conn()

	return tde
}

func (u *UserTd) Insert(uid int, name string) (ts int64, err error) {
	ts = time.Now().UnixNano()
	sqls := fmt.Sprintf("insert into `%s` (uid,name,ts) values (%d,'%s',%d)", u.TableName, uid, name, ts)
	_, err = u.Db.Exec(sqls)
	if err != nil {
		logger.ErrorLog("UserTde-Insert", "执行sql报错", "", err)
		return
	}

	return
}

func (u *UserTd) Find() (res []UserTd, err error) {
	sqls := fmt.Sprintf("select uid,name,ts from %s limit 10", u.TableName)

	rows, err := u.Db.Query(sqls)
	if err != nil {
		logger.ErrorLog("UserTde-Find", "执行sql报错", "", err)
		return
	}
	defer rows.Close()

	res = make([]UserTd, 0, 0)
	for rows.Next() {
		tmp := UserTd{}
		err = rows.Scan(&tmp.Uid, &tmp.Name, &tmp.Ts2)
		if err != nil {
			logger.ErrorLog("UserTde-Find", "赋值报错", "", err)
			return
		}
		res = append(res, tmp)
	}

	return
}
