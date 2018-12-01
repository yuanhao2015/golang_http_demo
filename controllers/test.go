package controllers

import (
	"../models"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type memberlist struct {
	Status  int             `json:"status"`
	Message []models.Member `json:"message"`
}

type membermsg struct {
	Status  int         `json:"status"`
	Content string      `json:"content"`
	Data    interface{} `json:"data"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	onlineUser := map[string]string{"title": "标题", "content": "内容"}
	t, err := template.ParseFiles("public/index.html")
	if err != nil {
		fmt.Println("111111")
		fmt.Println(err)
	}
	err = t.Execute(w, onlineUser)
	if err != nil {
		fmt.Println("222222")
		fmt.Println(err)
	}
}

func MemberList(w http.ResponseWriter, r *http.Request) {
	//var page, pageSize int
	filters := make([]interface{}, 0)
	filters = append(filters, "id", "<>", "0")
	page, _ := strconv.Atoi(r.FormValue("page"))
	pageSize, _ := strconv.Atoi(r.FormValue("page_size"))
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	list, _, err := models.ListMember(page, pageSize, filters...)
	fmt.Printf("%v", list)
	MembersList := &memberlist{Status: 0, Message: list}
	js, err := json.Marshal(MembersList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func MemberAdd(w http.ResponseWriter, r *http.Request) {
	m := new(models.Member)
	m.FirstName = r.FormValue("first_name")
	m.LastName = r.FormValue("last_name")
	m.UserName = r.FormValue("user_name")
	w.Header().Set("Content-Type", "application/json")
	if id, err := m.AddMember(); err != nil {
		Membersmsg := &membermsg{Status: -1, Content: "Failed", Data: nil}
		js, err := json.Marshal(Membersmsg)
		w.Write(js)
		fmt.Println(err)
	} else {
		m.Id = int(id)
		Membersmsg := &membermsg{Status: 0, Content: "Success", Data: m}
		js, _ := json.Marshal(Membersmsg)
		w.Write(js)
	}
}
func MemberDelete(w http.ResponseWriter, r *http.Request) {
	mid, _ := strconv.Atoi(r.FormValue("id"))
	w.Header().Set("Content-Type", "application/json")
	if n, err := models.DeleteMember(mid); err != nil {
		Membersmsg := &membermsg{Status: -1, Content: "Failed", Data: nil}
		js, err := json.Marshal(Membersmsg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(js)
		fmt.Println(err)
	} else {
		if n != 1 {
			Membersmsg := &membermsg{Status: 0, Content: "找不到对应id", Data: n}
			js, _ := json.Marshal(Membersmsg)
			w.WriteHeader(http.StatusOK)
			w.Write(js)
			return
		}
		Membersmsg := &membermsg{Status: 0, Content: "Success", Data: n}
		js, _ := json.Marshal(Membersmsg)
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	}
}

func MemberGet(w http.ResponseWriter, r *http.Request) {
	mid, _ := strconv.Atoi(r.FormValue("id"))
	mem, err := models.OneMember(mid)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(mem, err)
	if err != nil {
		Membersmsg := &membermsg{Status: -1, Content: "Failed", Data: nil}
		js, err := json.Marshal(Membersmsg)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(js)
		fmt.Println(err)
		//log.Fatal(err)
	} else {
		Membersmsg := &membermsg{Status: 0, Content: "Success", Data: mem}
		js, _ := json.Marshal(Membersmsg)
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	}

}
