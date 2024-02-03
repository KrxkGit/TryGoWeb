package mapper

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"os"
)

type dbConfigInfo struct { /*注意字段需要导出才能解析到json*/
	User     string
	Password string
	Dbname   string
}

var Db *sql.DB /*保存用于导出*/

func init() { /*引入包时自动被调用*/
	fmt.Println("Connecting Database.")
	secret, err := os.Open("./secret.json")
	if err != nil {
		panic("数据库配置文件缺失")
	}
	dbConfig := dbConfigInfo{}
	secretData, _ := io.ReadAll(secret)
	err = json.Unmarshal(secretData, &dbConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	dataSourceName := fmt.Sprintf("%s:%s@%s",
		dbConfig.User, dbConfig.Password, dbConfig.Dbname)
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Fatal(err.Error())
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		panic(err.Error())
	}
	Db = db /*导出变量*/
	fmt.Println("Connected.")
}

func GetOne(id int) (app userInfo, err error) {
	app = userInfo{}
	fmt.Println(id)
	query := `SELECT Id, password, email from user_info where Id=?`
	err = Db.QueryRow(query, id).Scan(&app.Id, &app.password, &app.email)
	return
}

func GetMany(id int) (apps []userInfo, err error) {
	query := `SELECT Id, password, email from user_info where Id<?`
	rows, err := Db.Query(query, id)
	for rows.Next() {
		app := userInfo{}
		err = rows.Scan(&app.Id, &app.password, &app.email)
		if err != nil {
			log.Fatalln(err.Error())
		}
		apps = append(apps, app)
	}
	return
}

func (u userInfo) Update() (err error) {
	query := `UPDATE user_info SET password=?, email=? where Id=?`
	_, err = Db.Exec(query, u.password, u.email, u.Id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return
}

func (u userInfo) Insert() (err error) {
	statement, err := Db.Prepare(`INSERT INTO user_info VALUES (?,?,?)`)
	defer statement.Close()
	_, err = statement.Exec(u.Id, u.password, u.email)
	return
}

func (u userInfo) Delete() (err error) {
	statement, err := Db.Prepare(`DELETE FROM user_info WHERE Id=?`)
	defer statement.Close()
	_, err = statement.Exec(u.Id)
	return
}
