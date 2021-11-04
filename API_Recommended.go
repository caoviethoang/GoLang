package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:Caoviethoang@tcp(127.0.0.1:3306)/gotest?parseTime=true"

type User struct {
	// gorm.Model

	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Gender     string    `json:"gender"`
	Latitude   float32   `json:"latitude"`
	Longitude  float32   `json:"longitude"`
	Birthday   time.Time `json:"birthday"`
	LastActive time.Time `json:"lastactive"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&User{})
}

type RecommendBody struct {
	Id        int
	User_Id   int
	Age_Range []int
	Latitude  float64
	Longitude float64
	Distance  int
	Genders   []string
}

func BirthdayBetween(from time.Time, to time.Time) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("birthday BETWEEN ? AND ?", from, to)
	}
}

func WithinDistance(lat float64, long float64, distance int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		// return db.Where("", lat, long, distance)
		//I don't know what to do here
		return db.Where("id IS NOT NULL")
	}
}

func ByGenders(genders []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		return db.Where("gender IN ?", genders)
	}
}

// Select * from users
func Recommended(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var rec RecommendBody
	json.NewDecoder(r.Body).Decode(&rec)

	today := time.Now()
	from := today.AddDate(-rec.Age_Range[1], 0, 0)
	to := today.AddDate(-rec.Age_Range[0], 0, 0)

	fmt.Println(from, today, to)

	var users []User

	DB.Limit(2500).
		Where("id != ?", rec.User_Id).
		Order("last_active DESC").
		Scopes(
			BirthdayBetween(from, to),
			WithinDistance(
				rec.Latitude,
				rec.Longitude,
				rec.Distance),
			ByGenders(rec.Genders)).
		Find(&users)

	json.NewEncoder(w).Encode(users)
}

// Select * from users where id=?
func RecommendedID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)

}

func UpdateUserIgnored(w http.ResponseWriter, r *http.Request) {

}
