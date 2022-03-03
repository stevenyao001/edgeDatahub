package pgsql

import (
	"database/sql"
	"fmt"
	"github.com/stevenyao001/edgeCommon/logger"
	"github.com/stevenyao001/edgeCommon/pgsql"
)

type UserPg struct {
	*pgsql.Postgres
	Uid  int    `json:"uid"`
	Name string `json:"name"`
}

func NewUserPgs() *UserPg {
	pgs := &UserPg{
		Postgres: &pgsql.Postgres{
			Db:        nil,
			InsName:   "rootcloud",
			DbName:    "test",
			TableName: "public.user",
		},
	}
	pgs.Conn()

	return pgs
}

func (u *UserPg) Create(name string, gender, birthday, avatarStatus int) (uid int64, err error) {
	sql := fmt.Sprintf("insert into %s (name,gender,birthday,avatar_status) values ($1,$2,$3,$4) RETURNING id", u.TableName)
	err = u.Db.QueryRow(sql, name, gender, birthday, avatarStatus).
		Scan(&uid)
	if err != nil {
		logger.ErrorLog("UserPgs-Create", "sql语句执行报错", "123", err)
	}
	return uid, err
}

func (u *UserPg) Delete(uid int) (rows int64, err error) {
	sql := fmt.Sprintf("delete from  %s where id= $1", u.TableName)
	res, err := u.Db.Exec(sql, uid)
	if err != nil {
		logger.ErrorLog("UserPgs-Create", "sql语句执行报错", "123", err)
		return rows, err
	}

	rows, err = res.RowsAffected()
	if err != nil {
		logger.ErrorLog("UserPgs-Create", "获取影响条数报错", "123", err)
		return rows, err
	}
	return rows, err
}

func (u *UserPg) Update(uid int, name string) (rows int64, err error) {
	sql := fmt.Sprintf("update %s set name=$1 where id= $2", u.TableName)
	res, err := u.Db.Exec(sql, name, uid)
	if err != nil {
		logger.ErrorLog("UserPgs-Update", "sql语句执行报错", "123", err)
		return rows, err
	}

	rows, err = res.RowsAffected()
	if err != nil {
		logger.ErrorLog("UserPgs-Update", "获取影响条数报错", "123", err)
		return rows, err
	}
	return rows, err
}

func (u *UserPg) Userinfo(uid int) (res UserPg, err error) {
	sqls := fmt.Sprintf("select id,name from %s where id=$1", u.TableName)
	err = u.Db.QueryRow(sqls, uid).Scan(&res.Uid, &res.Name)
	if err != nil && err != sql.ErrNoRows {
		logger.ErrorLog("UserPgs-Userinfo", "查询数据报错", "123", err)
	}
	return
}

func (u *UserPg) List(uid int) (res []UserPg, err error) {
	sql := fmt.Sprintf("select id,name from %s where id>$1 order by id asc", u.TableName)
	rows, err := u.Db.Query(sql, uid)
	if err != nil {
		logger.ErrorLog("UserPgs-List", "查询数据报错", "123", err)
		return res, err
	}

	for rows.Next() {
		tmp := UserPg{}
		err := rows.Scan(&tmp.Uid, &tmp.Name)
		if err != nil {
			logger.ErrorLog("UserPgs-List", "数据赋值报错", "123", err)
			return res, err
		}

		res = append(res, tmp)
	}
	return
}
