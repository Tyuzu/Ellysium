package main

import (
	"html/template"
	"log"
	"net/http"
	
    "github.com/julienschmidt/httprouter"
)


const PORT = "localhost:4000"
const maxUploadSize = 8 * 1024 * 1024 // 8 mb
var uploadPath = "./uploads"
var cssPath = "./css"
var userPics = "./userpics"

var tmpl = template.Must(template.ParseGlob("templates/*.html"))



func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/search", Search)
    router.GET("/favicon.ico", Ignore)
    router.GET("/upload", uploadFileHandler)
    router.POST("/upload", uploadFileHandler)
    router.DELETE("/post/:url", DelPost)
    router.GET("/v/:postid", ViewPost)
    router.POST("/addtocoll", AddToColl)
    router.POST("/update", Fields)

//--------FileServer--------//

	static := httprouter.New()
	static.ServeFiles("/files/*filepath", http.Dir(uploadPath))
	static.ServeFiles("/css/*filepath", http.Dir(cssPath))
	static.ServeFiles("/usrimg/*filepath", http.Dir(userPics))
	router.NotFound = static	

//--------Server--------//
	log.Println("Server starting")
    log.Fatal(http.ListenAndServe(PORT, router))
}

func Index(w http.ResponseWriter, r *http.Request, _  httprouter.Params) {
	log.Println("fbhd")
		if r.Method == "GET" {
			tmpl.ExecuteTemplate(w, "index.html", nil)
			return
		}
}