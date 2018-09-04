package models

import (
	"database/sql"
	"log"
	"os"
	"time"
)

//Student struct
type Student struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"Name,omitempty"`
	Email       string `json:"Email,omitempty"`
	CreatedTime string `json:"CreatedTime,omitempty"`
}

// Students is a slice of Student
type Students struct {
	Students []Student `json:"students"`
}

func dialSQL() (*sql.DB, error) {
	ConnectionStr := ""
	if os.Getenv("PROJECT_ID") == "tsiahpng" {
		//google cloud sql
		ConnectionStr = os.Getenv("CLOUD_CONNECTION")
	} else {
		//本機
		// ConnectionStr = "Junxiang:rmp4vu;6@tcp(127.0.0.1:3306)/TsiahPng_db"
		ConnectionStr = "root:@tcp(127.0.0.1:3306)/TsiahPng_db"

	}

	return sql.Open("mysql", ConnectionStr)
}
func dbGetConn() *sql.DB {

	db, err := dialSQL()
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect: ", err.Error())
	}

	return db
}

// DBGetStudents MySQL 取得所有學生資料
func DBGetStudents() (data []Student, count int) {
	db := dbGetConn()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM `School`")

	defer rows.Close()

	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var student Student
		rows.Scan(&student.ID, &student.Name, &student.Email, &student.CreatedTime)
		data = append(data, student)
		count++
	}
	if err = rows.Err(); err != nil {
		log.Fatalln(err)
	}
	//getTime()
	return data, count
}

// DBInsertStudent MySQL 新增一筆學生資料
func DBInsertStudent(name string, email string) (r bool) {

	// name := c.Request.FormValue("name")
	// email := c.Request.FormValue("email")
	//name = "A"
	//email = "B"

	//開啟db
	db := dbGetConn()
	defer db.Close()

	//傳統執行query //每次執行内部都会去连接池获取一个新的连接，效率低
	/////////////////////////////////////////
	// query := fmt.Sprintf("INSERT INTO `junxiang_db`.`school` (`name`,`email`,`CreatedTime`) VALUES('%s','%s','%s')", name, email, getTime())
	// result, err := db.Query(query)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// defer result.Close()

	/////////////////////////////////////////

	//https://www.jianshu.com/p/340eb943be2e
	//https://studygolang.com/articles/3022
	//方式4  Begin函数内部会去获取连接  多筆新增時效率高
	/////////////////////////////////////////
	// query := fmt.Sprintf("INSERT INTO `School` (`name`,`email`,`CreatedTime`) VALUES('%s','%s','%s')", name, email, getTime())

	//Begin函数内部会去获取连接
	tx, err := db.Begin()
	errCheck(err)

	//每次循环用的都是tx内部的连接，没有新建连接，效率高
	// rs, err := tx.Exec(query)
	rs, err := tx.Exec("INSERT INTO `School` (`name`,`email`,`CreatedTime`) VALUES(?, ?, ?)", name, email, getTime())

	errCheck(err)

	//最后释放tx内部的连接
	tx.Commit()

	rowCount, err := rs.RowsAffected()
	errCheck(err)

	log.Printf("inserted %d rows", rowCount)
	/////////////////////////////////////////

	r = true
	return r

}

//Restaurant Restaurant
type Restaurant struct {
	ID      int    `json:"id,omitempty"`
	Img     string `json:"img,omitempty"`
	Name    string `json:"name,omitempty"`
	Price   string `json:"price,omitempty"`
	Purpose string `json:"purpose,omitempty"`
}

// Restaurants is a slice of Restaurant
type Restaurants struct {
	Restaurants []Restaurant `json:"Restaurants"`
}

// TsiahPngGetList MySQL 取得所有餐廳資料
func TsiahPngGetList() (data []Restaurant, count int) {
	db := dbGetConn()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM `restaurants_list`")

	defer rows.Close()

	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var restaurant Restaurant
		rows.Scan(&restaurant.ID, &restaurant.Img, &restaurant.Name, &restaurant.Price, &restaurant.Purpose)
		data = append(data, restaurant)
		count++
	}
	if err = rows.Err(); err != nil {
		log.Fatalln(err)
	}
	//getTime()
	return data, count
}

// DBInsertRestaurant MySQL 新增一筆餐廳資料
func DBInsertRestaurant(name string, price string, purpose string) (r bool) {

	//開啟db
	db := dbGetConn()
	defer db.Close()

	//Begin函数内部会去获取连接
	tx, err := db.Begin()
	errCheck(err)

	//每次循环用的都是tx内部的连接，没有新建连接，效率高
	// rs, err := tx.Exec(query)
	rs, err := tx.Exec("INSERT INTO `restaurants_list` (`img`,`name`,`price`,`purpose`) VALUES(?, ?, ?, ?)", "no-image.jpg", name, price, purpose)

	errCheck(err)

	//最后释放tx内部的连接
	tx.Commit()

	rowCount, err := rs.RowsAffected()
	errCheck(err)

	log.Printf("inserted %d rows", rowCount)
	/////////////////////////////////////////

	r = true
	return r

}

//取得時間並format
func getTime() (r string) {
	//p := fmt.Println
	t := time.Now()
	//_, z := t.Zone()

	//p(z)
	//p(t.Format("2006-01-02 15:04:05"))
	r = t.Format("2006-01-02 15:04:05")
	return r
}

func errCheck(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
