package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"net/http"
	"os"
)

var t *template.Template

type Page struct {
	Title, Status string
	Body          template.HTML
}

func init() {
	t = template.New("templ")
	t.ParseGlob("templates/*.html")
}

func main() {
	os.Remove("./foo.db")
	db, _ := sql.Open("sqlite3", "./foo.db")

	db.Exec(`
	create table user (id integer not null primary key autoincrement, username text, password text);
	insert into user (username,password) values ("george","");
	insert into user (username,password) values ("john","");
	create table person (id integer not null primary key autoincrement, name text, age int, address text);
	insert into person (name,age,address) values ("George",38,"Sesame Street,Romania");
	insert into person (name,age,address) values ("Gill Bates",55,"Linux Street 10");
	insert into person (name,age,address) values ("Tinus Lorvalds",42,"Windows Bay 1");
	//delete from foo;
	`)
	db.Close()

	http.ListenAndServe(":8080", nil)
}
