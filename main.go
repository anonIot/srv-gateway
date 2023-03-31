package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Publisher interface {
	Publish() error
}

type blogPost struct {
	author string
	title  string
	postId int
}

type Cover struct {
	Id   int
	Name string
}

func (b blogPost) Publish() error {
	fmt.Printf("The title on %s has been published by %s\n", b.title, b.author)
	return nil
}

func test() {

	b := blogPost{"Alex", "understanding structs and interface types", 12345}

	fmt.Println(b.Publish())

	d := &b // pointer receiver for the struct type

	b.author = "Chinedu"

	fmt.Println(d.Publish())

}

func PublishPost(publish Publisher) error {
	return publish.Publish()
}

var db *sql.DB

func main() {

	// app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("SRV Mini Gateway API Service")
	// })

	// app.Listen(":3033")

	// var p Publisher

	// fmt.Println(p)

	// b := blogPost{"Alex", "understanding structs and interface types", 12345}

	// fmt.Println(b)
	// PublishPost(b)

	var err error
	db, err = sql.Open("mysql", "root:tiger@/sd_consumer")
	if err != nil {
		panic(err)
	}
	//covers, err := GetCovers()
	cover, err := GetCover(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cover)
	// for _, cover := range covers {

	// 	fmt.Println(cover)
	// }

}

func GetCover(id int) (*Cover, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	//MSSQL
	// query := "select id,name from customer where id=@id"
	// row := db.QueryRow(query, sql.Named("id", id))

	//MYSQL
	query := "select id,name from customer where id=?"
	row := db.QueryRow(query, id)

	cover := Cover{}
	err = row.Scan(&cover.Id, &cover.Name)
	if err != nil {
		return nil, err
	}
	return &cover, nil

}
func UnitScan() {

}
func AcList() {

}

func acCmd(cmd string) {
	switch cmd {
	case "power":
	case "temp":
	case "speed":
	case "mode":
	case "clean":
	}

}
func acModel(cmd string) bool {

	return true
}

func GetCovers() ([]Cover, error) {

	err := db.Ping()
	if err != nil {

		return nil, err
	}

	query := "select id,name from customer"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	covers := []Cover{}

	for rows.Next() {
		cover := Cover{}
		err = rows.Scan(&cover.Id, &cover.Name)
		if err != nil {
			panic(err)
		}

		covers = append(covers, cover)

	}
	return covers, err

}
