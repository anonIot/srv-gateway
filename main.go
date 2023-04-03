package main

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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

// var db *sql.DB
var db *sqlx.DB

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
	db, err = sqlx.Open("mysql", "root:tiger@/sd_consumer")
	if err != nil {
		panic(err)
	}

	// err = DeleteCover(1)
	// if err != nil {
	// 	fmt.Println(err)
	// 	//panic(err)
	// 	return
	// }

	//covers, err := GetCoversX()
	covers, err := getCoverX(4)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(covers)
	// for _, cover := range covers {

	// 	fmt.Println(cover)
	// }

}

func GetCoversX() ([]Cover, error) {
	query := "select id,name from customer"
	covers := []Cover{}
	err := db.Select(&covers, query)
	if err != nil {
		return nil, err
	}

	return covers, nil

}

func getCoverX(id int) (*Cover, error) {
	query := "select id,name from customer where id=?"
	cover := Cover{}
	err := db.Get(&cover, query, id)
	if err != nil {
		return nil, err
	}
	return &cover, nil

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

func AddCover(cover Cover) error {
	query := "insert into customer(name) values (?)"
	result, err := db.Exec(query, cover.Name)

	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected <= 0 {
		return errors.New("can not insert")
	}

	return nil
}

func UpdateCover(cover Cover) error {
	query := "update customer set name=? where id=?"
	result, err := db.Exec(query, cover.Name, cover.Id)

	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected <= 0 {
		return errors.New("can not update")
	}

	return nil
}
func DeleteCover(id int) error {
	query := "delete from customer where id=?"
	result, err := db.Exec(query, id)

	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected <= 0 {
		return errors.New("can not delete")
	}

	return nil
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
