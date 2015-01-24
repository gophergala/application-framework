package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"net/http"
)

func addPerson(w http.ResponseWriter, r *http.Request) {
	//this must add at begin of every session code
	c, err := r.Cookie("session")
	if err != nil || c.Value == "" {
		http.Error(w, "Session expired", 401)
		return
	}

	//handle GET/POST methods
	var b string

	if r.Method == "POST" {
		db, _ := sql.Open("sqlite3", "./foo.db")
		defer db.Close()

		q := `insert into person (name,age,address) values ("`
		q += r.FormValue("name") + `",`
		q += r.FormValue("age") + `,"`
		q += r.FormValue("address") + `");`
		//println(q)
		//return
		db.Exec(q)
		defer db.Close()
		http.Redirect(w, r, "/showPersons", http.StatusFound)
	} else {

		//build page content
		b = `<pre>
This is also module that works with forms.

   Add a person
<form method="post" action="">
   Name    <input type="text" size=16 name="name" value="" autofocus>
   Age     <input type="text" size=3 name="age" value="">
   Address <input type="text" size=20 name="address" value="">
           <input type="submit" name="submit" value="Submit"">
</form>`
	}

	//finally show the page
	p := Page{
		Title:  "Check user page",
		Status: c.Value,
		Body:   template.HTML(b),
	}
	t.ExecuteTemplate(w, "index.html", p)
}

func init() {
	http.HandleFunc("/addPerson", addPerson)
}
