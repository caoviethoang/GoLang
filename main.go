package main

import (
	"database/sql"
	"fmt"
	"time"

	// "gorm.io/gorm"
	// "gorm.io/driver/sqlite"
	"github.com/beego/beego/v2/client/orm"
	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	err := orm.RegisterDataBase("default", "mysql", "root:Caoviethoang@/gotest?charset=utf8")
	if err != nil {
		glog.Fatal("Failed to register database %v", err)
	}

	// name := "default"

	// // Drop table and re-create.
	// force := true

	// // Print log.
	// verbose := true

	// // Error.
	// err1 := orm.RunSyncdb(name, force, verbose)
	// if err1 != nil {
	// 	glog.Fatal("Failed to run sync, error: %v ", err1)
	// }
}

// type User struct {
// 	Usermame string `json:"username"`
// }

type User struct {
	UserId     uint32
	Username   string    `fake:"{username}_{number:1000}"`
	Gender     string    `fake:"{gender}"`
	Latitude   float64   `fake:"{latitude}"`
	Longitude  float64   `fake:"{longitude}"`
	Birthday   time.Time `fake:"{year}-{month}-{day}" format:"1921-01-02"`
	LastActive time.Time `fake:"{date}" format:"2006-01-02"`
}

func Insert(u User) {
	db, err := sql.Open("mysql", "root:Caoviethoang@tcp(127.0.0.1:3306)/gotest")

	insertForm, err := db.Prepare("INSERT INTO users(username,gender,latitude,longitude,birthday,last_active) VALUES(?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	result, error := insertForm.Exec(u.Username, u.Gender, u.Latitude, u.Longitude, gofakeit.Date().Format("2006-01-02"), gofakeit.DateRange(
		time.Date(2021, 10, 1, 0, 0, 0, 0, time.Now().Location()),
		time.Date(2021, 11, 1, 0, 0, 0, 0, time.Now().Location())).Format("2006-01-02 15:04:05"))
	// result, error := insertForm.Exec(u.Username, u.Gender, u.Latitude, u.Longitude, "1921-01-01", "2021-11-01 09:42:47")
	if err != error {
		panic(error.Error())
	}

	fmt.Println(result)

	defer db.Close()

}

func InsertFakeData(n int) {
	for i := 0; i < n; i++ {
		var u User
		gofakeit.Struct(&u)

		fmt.Println(u)
		Insert(u)

	}
}

func main() {

	InsertFakeData(200000)
	// fmt.Println(gofakeit.Number(100000, 10000000))

	// fmt.Println("Go MySQL")

	// db, err := sql.Open("mysql", "root:Caoviethoang@tcp(127.0.0.1:3306)/gotest")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer db.Close()

	// insert, err := db.Query("INSERT INTO users VALUES('Hoang')")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer insert.Close()
	// fmt.Println("Succesfully inserted into MySQL database")

	// results, err := db.Query("SELECT username from users")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// for results.Next() {
	// 	var username User

	// 	err = results.Scan(&username.Usermame)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	fmt.Println(username.Usermame)
	// }

}
