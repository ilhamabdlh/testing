package config

import (
	// "gorm.io/driver/postgresql"
	// "gorm.io/gorm"
	// "github.com/jinzhu/gorm"
	// // "github.com/rs/cors"
	// "github.com/jinzhu/gorm/dialects/postgres"

	"database/sql"
	//"fmt"
	_ "github.com/lib/pq"
	//"time"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "test"
)

func DBInit(){
	
	
	//dsn := "host=localhost port=5432 user=postgres dbname=atop sslmode=disable password=satu2tiga45 TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("user=%s password=%s name=%s", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dsn)

	if err != nil{
		panic("failed to connect database")
	}
	defer db.Close()
	
	db.AutoMigrate(structs.Person{})

	return db
}
