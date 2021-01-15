// go run F:\develope\go\go_code_path\src\github.com\uipps\goZadmin\study\database\mysql_test.go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	USERNAME = "root"
	PASSWORD = "10y9c2U5"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "laravel"
)

type User struct {
	ID int64 `db:"id"`
	Name sql.NullString  `db:"name"`  //由于在mysql的users表中name没有设置为NOT NULL,所以name可能为null,在查询过程中会返回nil，如果是string类型则无法接收nil,但sql.NullString则可以接收nil值
	Age int `db:"age"`
}

func main()  {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME,PASSWORD,NETWORK,SERVER,PORT,DATABASE)
	DB,err := sql.Open("mysql",dsn)
	if err != nil{
		fmt.Printf("Open mysql failed,err:%v\n",err)
		return
	}
	DB.SetConnMaxLifetime(100*time.Second)  //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)//设置最大连接数
	DB.SetMaxIdleConns(16) //设置闲置连接数

	queryOne(DB)
	queryMulti(DB)
	insertData(DB)
	updateData(DB)
	deleteData(DB)
}

//查询单行
func queryOne(DB *sql.DB) {
	user := new(User)
	row := DB.QueryRow("select * from users where id=?", 2)
	if err := row.Scan(&user.ID, &user.Name, &user.Age); err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Println(*user)
}

//查询多行
func queryMulti(DB *sql.DB) {
	user := new(User)
	rows, err := DB.Query("select * from users where id > ?", 1)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		fmt.Printf("Query failed,err:%v\n", err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("Scan failed,err:%v\n", err)
			return
		}
		fmt.Print(*user)
	}

}


//插入数据
func insertData(DB *sql.DB){
	result,err := DB.Exec("insert INTO users(name,age) values(?,?)","YDZ",23)
	if err != nil{
		fmt.Printf("Insert failed,err:%v\n",err)
		return
	}
	lastInsertID,err := result.LastInsertId()
	if err != nil {
		fmt.Printf("Get lastInsertID failed,err:%v\n",err)
		return
	}
	fmt.Println("LastInsertID:",lastInsertID)
	rowsaffected,err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v\n",err)
		return
	}
	fmt.Println("RowsAffected:",rowsaffected)
}

//更新数据
func updateData(DB *sql.DB){
	result,err := DB.Exec("UPDATE users set age=? where id=?","30",3)
	if err != nil{
		fmt.Printf("Insert failed,err:%v\n",err)
		return
	}
	rowsaffected,err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v\n",err)
		return
	}
	fmt.Println("RowsAffected:",rowsaffected)
}

//删除数据
func deleteData(DB *sql.DB){
	result,err := DB.Exec("delete from users where id=?",1)
	if err != nil{
		fmt.Printf("Insert failed,err:%v\n",err)
		return
	}
	rowsaffected,err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v\n",err)
		return
	}
	fmt.Println("RowsAffected:",rowsaffected)
}