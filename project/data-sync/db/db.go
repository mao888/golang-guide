package db

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MongoClient *qmgo.Client
var MySQLClient *gorm.DB

var MySQLClientCruiser *gorm.DB
var MySQLClientPayment *gorm.DB
var MySQLClientAuth *gorm.DB
var MySQLClientAdmin *gorm.DB
var MySQLClientGM *gorm.DB

func init() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://mongodb-test01.nuclearport.com:27000"})
	if err != nil {
		fmt.Println(err)
	}
	MongoClient = client

	dsn := "root:Redhat@123@(192.168.1.52:3306)/application_console?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	MySQLClient = db

	dsn2 := "root:Redhat@123@(192.168.1.52:3306)/payment?charset=utf8mb4&parseTime=True&loc=Local"
	db2, err := gorm.Open(mysql.Open(dsn2), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	MySQLClientPayment = db2

	dsn3 := "root:Redhat@123@(192.168.1.52:3306)/auth_admin?charset=utf8mb4&parseTime=True&loc=Local"
	db3, err := gorm.Open(mysql.Open(dsn3), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	MySQLClientAuth = db3

	dsn4 := "root:Redhat@123@(192.168.1.52:3306)/admin_console?charset=utf8mb4&parseTime=True&loc=Local"
	db4, err := gorm.Open(mysql.Open(dsn4), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	MySQLClientAdmin = db4

	dsn5 := "root:Redhat@123@(192.168.1.52:3306)/gm-system?charset=utf8mb4&parseTime=True&loc=Local"
	db5, err := gorm.Open(mysql.Open(dsn5), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	MySQLClientGM = db5
}
