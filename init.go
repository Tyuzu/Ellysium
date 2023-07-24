package main

import (
    "log"
    "os"
    "path/filepath"
	"embed"
	"strconv"

    "github.com/joho/godotenv"
)

//go:embed templates/*.html
var content embed.FS

var Env EnvVars

var PORT = Env.SERVER_PORT

type EnvVars struct {
	WEBSITE_NAME string
	SERVER_PORT string
	MAX_FILE_SIZE string
	CSS_FILE string
}

func init() {

	log.Println("Created By Ellygator on Discord")

  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

Env.WEBSITE_NAME = os.Getenv("WEBSITE_NAME")
Env.SERVER_PORT = os.Getenv("SERVER_PORT")
Env.MAX_FILE_SIZE = os.Getenv("MAX_FILE_SIZE")
Env.CSS_FILE = os.Getenv("CSS_FILE")

haloobaloo := os.Getenv("HALOOBALOO")
if haloobaloo == "" {
  haloobaloo = "Working fine"
}
log.Println(haloobaloo)

var imax int
imax, _ = strconv.Atoi(Env.MAX_FILE_SIZE)
maxUploadSize = int64(imax) * 1024 * 1024 // 8 mb

if Env.SERVER_PORT == "" {
  PORT = "localhost:4000"
}

//  secretKey := os.Getenv("CSS_FILE")
//	log.Println(secretKey)
	
	newpath := filepath.Join(".", "uploads")
	err = os.MkdirAll(newpath, os.ModePerm)
	if err != nil {
		log.Println("\a\a\a\a")
	}
}