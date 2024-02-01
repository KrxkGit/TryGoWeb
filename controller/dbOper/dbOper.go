package dbOper

import (
	"encoding/json"
	"fmt"
	"github.com/KrxkGit/TryGoWeb/mapper"
	"log"
	"net/http"
	"strconv"
)

func RegisterDbOperRoute() {
	http.HandleFunc("/db", handleDbOperation)
}

func handleDbOperation(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case "GET":
		handleGet(w, r)
		break
	case "POST":
		handlePost(w, r)
		break
	case "PUT":
		handlePut(w, r)
		break
	case "DELETE":
		handleDelete(w, r)
		break
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, _ := strconv.Atoi(query["id"][0])
	log.Println(id)

	app, err := mapper.GetOne(id)
	if err != nil {
		log.Println(err.Error())
	}

	res, _ := json.Marshal(app)
	w.Write(res)

	apps, err := mapper.GetMany(id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(apps)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err.Error())
	}
	form := r.PostForm
	if err != nil {
		log.Fatalln(err.Error())
	}
	id, _ := strconv.Atoi(form["id"][0])
	app, _ := mapper.GetOne(id)

	app = app.Set(id, form["password"][0], form["email"][0])
	fmt.Println(app)
	app.Update()

	app2, _ := mapper.GetOne(id)

	data, _ := json.Marshal(app2)
	fmt.Fprintln(w, string(data))
	//log.Println(id)
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	form := r.PostForm
	app := mapper.Create()
	id, err := strconv.Atoi(form["id"][0])
	if err != nil {
		log.Fatalln(err.Error())
	}
	password := form["password"][0]
	email := form["email"][0]
	app = app.Set(id, password, email)

	app.Insert()
	w.WriteHeader(http.StatusOK)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//fmt.Println(r.PostForm)
	id, err := strconv.Atoi(r.URL.Query()["id"][0])
	fmt.Println(id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	app := mapper.Create()
	app = app.Set(id, "", "")
	app.Delete()
	w.WriteHeader(http.StatusOK)
}
