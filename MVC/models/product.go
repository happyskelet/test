package models

import (
	"database/sql"
	"github.com/lib/pq"	
	"fmt"
	s "../structures"
)
const (
	DB_USER     = "postgres"
	DB_PASSWORD = "ytgfhmcz,elmcxfcnkbd"
	DB_NAME     = "postgres"
)

func GetProduct(id string) (s.Product){
	connStr := "user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME
	db, err := sql.Open("postgres", connStr)
	defer db.Close();
	if err != nil {fmt.Println(err)}
	rows, err := db.Query("SELECT id,name,description FROM products WHERE id = $1",id);
	var result s.Product;
	for rows.Next(){
		rows.Scan(&result.Id,&result.Name,&result.Description);
	}
	return result;
}


func GetProducts() ([]s.Product){
	connStr := "user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME
	db, err := sql.Open("postgres", connStr)
	defer db.Close();
	if err != nil {fmt.Println(err)}
	rows, err := db.Query("SELECT id,name,description FROM products");
	result:=[]s.Product{};
	for rows.Next(){
		var buffer s.Product;
		rows.Scan(&buffer.Id,&buffer.Name,&buffer.Description);
		result = append(result,buffer);
	}
	
	return result;
}

func AddProduct(data s.Product) (int){
	connStr := "user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME;
	db, err :=sql.Open("postgres",connStr);
	if err!=nil {fmt.Println(err)}
	defer db.Close();
	var InsertId int;
	err = db.QueryRow("INSERT INTO products(name,description) VALUES($1,$2) returning id",data.Name,data.Description).Scan(&InsertId);
	if err!=nil {fmt.Println(err)}
	return InsertId;
}
func UpdateProduct(data s.Product) (bool){
	connStr := "user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME;
	db, err :=sql.Open("postgres",connStr);
	if err!=nil {fmt.Println(err)}
	defer db.Close();
	stmt, err := db.Prepare("update products set name=$1, description=$2 where id=$3")
    if err!=nil {fmt.Println(err)}
    stmt.Exec(data.Name, data.Description, data.Id)
	if err!=nil {fmt.Println(err)}
	return true;
}

func DeleteProduct(id string) (bool){
	connStr := "user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME;
	db, err := sql.Open("postgres",connStr);
	defer db.Close();
	if err != nil {fmt.Println(err)}
	_, err = db.Query("DELETE FROM products WHERE id = $1",id);
	if err != nil {fmt.Println(err)}
	return true;
}