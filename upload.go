package main

import (
	"fmt"
	"io/ioutil"
	//	"mime"
	"net/http"
	"os"
	//"log"
	"crypto/rand"
	"crypto/md5"
	"path/filepath"
	
    "github.com/julienschmidt/httprouter"
)

func uploadFileHandler(w http.ResponseWriter, r *http.Request, _  httprouter.Params) {
		if r.Method == "GET" {
				tmpl.ExecuteTemplate(w, "index.html", nil)
		} else {
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			fmt.Printf("Could not parse multipart form: %v\n", err)
			renderError(w, "CANT_PARSE_FORM", http.StatusInternalServerError)
			return
		}
		var fileEndings string
		var fileName string
		var pics []string
		files := r.MultipartForm.File["myFile"]
		for _, fileHeader := range files {
			file, err := fileHeader.Open()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		fmt.Println(r.FormValue("filetitle"))
		fmt.Println(r.FormValue("filetags"))
		fmt.Println(r.FormValue("isprivate"))
		
			defer file.Close()
			// Get and print out file size
			fileSize := fileHeader.Size
			fmt.Println(fileHeader.Filename)
//			fmt.Printf("File size (bytes): %v\n", fileSize)
			// validate file size
			if fileSize > maxUploadSize {
				renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
				return
			}
			fileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				renderError(w, "INVALID_FILE"+http.DetectContentType(fileBytes), http.StatusBadRequest)
				return
			}
			// check file type, detectcontenttype only needs the first 512 bytes
			detectedFileType := http.DetectContentType(fileBytes)
			switch detectedFileType {
			case "video/mp4":
				fileEndings = ".mp4"
				break
			case "video/webm":
				fileEndings = ".webm"
				break
			case "image/gif":
				fileEndings = ".gif"
				break
			default:
				renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
				return
			}
			
			/*var gethash string
			gethash = fmt.Sprintf(fileHeader.Filename + fileEndings + string(fileSize))
			data := []byte(gethash)
			myHash := fmt.Sprintf("%x", md5.Sum(data))
			if myHash == "8dabda60162d4a3c3eb1bf7d43a6ec04" {
				go XHRrespond(w, "File Already Uploaded")
				continue;
			}*/
			
			fileName = randToken(12)
			//		fileEndings, err := mime.ExtensionsByType(detectedFileType)

			if err != nil {
				renderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
				return
			}
			
			newFileName := fileName + fileEndings

			newPath := filepath.Join(uploadPath, newFileName)
//			fmt.Printf("FileType: %s, File: %s\n", detectedFileType, newPath)

			// write file
			newFile, err := os.Create(newPath)
			if err != nil {
				renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
				return
			}
			defer newFile.Close() // idempotent, okay to call twice
			if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
				renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
				return
			}
			
/*
			rdxHset(fileName, "fileName", r.FormValue("filetitle"))
			rdxHset(fileName, "fileTags", r.FormValue("filetags"))
			rdxHset(fileName, "fileVis", r.FormValue("isprivate"))*/
			
			go XHRrespond(w,"/v/"+fileName)
			pics = append(pics, "/files/"+newFileName)
		}
		fmt.Println(pics)
//		http.Redirect(w, r, "/v/"+fileName, http.StatusSeeOther)
//		tmpl.ExecuteTemplate(w, "player.html", pics)
//		tmpl.ExecuteTemplate(w, "returnUploads.html", pics)
	}
}


func ViewPost(w http.ResponseWriter, r *http.Request, postid httprouter.Params) {
		if r.Method == "GET" {
			res := "/files/"+postid.ByName("postid")
			tmpl.ExecuteTemplate(w, "head.html", nil)
			tmpl.ExecuteTemplate(w, "show.html", res)
		}
}



func Fields(w http.ResponseWriter, r *http.Request, postid httprouter.Params) {
	fmt.Println(r.FormValue("filetitle"))
	fmt.Println(r.FormValue("filetags"))
	fmt.Println(r.FormValue("isprivate"))
}



func DelPost(w http.ResponseWriter, r *http.Request, postid httprouter.Params) {
	fmt.Println("./uploads/"+postid.ByName("url")+".mp4")
	os.Remove("./uploads/"+postid.ByName("url")+".mp4")
	go XHRrespond(w,"Deleted")
}


func AddToColl(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//	fmt.Println(r.FormValue("collection"))
//	fmt.Println(r.FormValue("gifname"))
//	fmt.Println(postid.ByName("collection"))
//	fmt.Println(postid.ByName("gifname"))
	go XHRrespond(w,"Not Available Yet")
}



func Search(w http.ResponseWriter, r *http.Request, _  httprouter.Params) {
	fmt.Fprintf(w,"Not Yet Implemented")
}


func XHRrespond(w http.ResponseWriter, message string) {
	fmt.Fprintf(w,message)
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}


func HashIt(strToHash string) string {
	data := []byte(strToHash)
	return fmt.Sprintf("%x", md5.Sum(data))
}



func Ignore(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    //http.ServeFile(w, r, "favicon.ico")
}



