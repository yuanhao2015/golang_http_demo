package models

import (
	"log"
	"strconv"

	db "../dbs"
)

type Member struct {
	Id        int    `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	UserName  string `json:"user_name" form:"user_name"`
}

func (m *Member) AddMember() (id int64, err error) {
	res, err := db.Conns.Exec("INSERT INTO member(first_name, last_name,user_name) VALUES (?, ?)", m.FirstName, m.LastName, m.UserName)
	if err != nil {
		return
	}
	id, err = res.LastInsertId()
	return
}

func ListMember(page, pageSize int, filters ...interface{}) (lists []Member, count int64, err error) {
	lists = make([]Member, 0)
	where := "WHERE 1=1"
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 3 {
			where = where + " AND " + filters[k].(string) + filters[k+1].(string) + filters[k+2].(string)
		}
	}
	limit := strconv.Itoa((page-1)*pageSize) + "," + strconv.Itoa(pageSize)
	rows, err := db.Conns.Query("SELECT id, first_name, last_name,user_name FROM member " + where + " LIMIT " + limit)
	defer rows.Close()

	if err != nil {
		return
	}
	count = 0
	for rows.Next() {
		var member Member
		rows.Scan(&member.Id, &member.FirstName, &member.LastName, &member.UserName)
		lists = append(lists, member)
		count++
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func OneMember(id int) (m Member, err error) {
	m.Id = 0
	m.FirstName = ""
	m.LastName = ""
	m.UserName = ""
	err = db.Conns.QueryRow("SELECT id, first_name, last_name,user_name FROM member WHERE id=? LIMIT 1", id).Scan(&m.Id, &m.FirstName, &m.LastName, &m.UserName)
	return
}

func (m *Member) UpdateMember(id int) (n int64, err error) {
	res, err := db.Conns.Prepare("UPDATE member SET first_name=?,last_name=? ,user_name=? WHERE id=?")
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}
	rs, err := res.Exec(m.FirstName, m.LastName, m.UserName, m.Id)
	if err != nil {
		log.Fatal(err)
	}
	n, err = rs.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return
}

func DeleteMember(id int) (n int64, err error) {
	n = 0
	rs, err := db.Conns.Exec("DELETE FROM member WHERE id=?", id)
	if err != nil {
		log.Fatalln(err)
		return
	}
	n, err = rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
		return
	}
	return
}
