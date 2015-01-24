###Application Framework
====
**About:**

Application Framework is a full modular web application.
Useful when we have or write many applications and we want to concentrate all of them into single application with autentication.

Modularity is based on some go specific caracteristics and modules are plugable at compile time.

[![last-version-blue](https://cloud.githubusercontent.com/assets/6298396/5602522/8967405e-935b-11e4-8777-de3623ed6ad7.png)] (https://github.com/gophergala/application-framework/archive/master.zip)


**Description:**

Everithing but main.go is a module and have the same structure. You can remove any of mod_*.go file and program compile and run flawless (wow!). You can also add module as you wish. For example if you wish another autenticate module replace only this module.
Basicaly, application is a puzzle of modules.

I used preformated text because is simpler for this job . But monospace fonts is not so prety. So i use Anonymous Pro (see templates/style.html).

Database used is sqlite (see github.com/mattn/go-sqlite3)

**Using:**

Compile program. Se here (https://golang.org/doc/code.html) how.

Run and open http://localhost:8080 in your favorite browser (default user is george without password).

Back button is disabled in browser and is nice to run with Google Chrome in app mode

         google-chrome --app=http://localhost:8080

**Tools used in this project:**

   * compiler http://golang.org
   * ide      https://github.com/visualfc/liteide
   * gopei    https://golang.org/geosoft1/tools for faster development

**How it works:**

We can define this application thus:
   * an autentication mecanism
   * modules link mecanism
   * collection of modules

Basicaly, you have a module template and a tehnique to plugin or plugout into
main application.

init() function make go module plugable. here we put module web handler

         http.HandleFunc("/ModuleName", ModuleName)
	
next we must define handler

         func ModuleName(w http.ResponseWriter, r *http.Request) {
         	//this must add at begin of every session code
         	c, err := r.Cookie("session")
         	if err != nil || c.Value == "" {
         		http.Error(w, "Session expired", 401)
         		return
         	}
         
         	//build page content
         	b := `<pre> Page content`
         	
         	//finally show the page
         	p := Page{
         		Title:    "Module Title",
         		Status:   c.Value,		// e.g connected user
         		Body:     template.HTML(b),
         	}
         	t.ExecuteTemplate(w, "index.html", p)
         }

as you see structure are fixed

   * cookie checker
   * build page content in b variable
   * how page with go template

Note that a simple cookie mecanism are used to implement sessions in modules.
Also, module has access to a global logfile with

         log.Println("message")

Thats all folks about modules. Now, adding a menu line in templates/index.html file like

         <a href="/ModuleName" >ModuleName</a> 

make visible ModuleName to application. Remove this line and coresponding module,recompile application and module are removed. Nothing else to do.

Easy copy/paste module or easy modify old modules to do new modules give a speed in developing applications.

Of course this things can be do better.
