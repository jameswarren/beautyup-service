package main

import (
  "github.com/joho/godotenv"
  "log"
  "os"
  // "fmt"
)

func main() {
  // load env vars
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  a := App{}
  a.Initialize(
      os.Getenv("DB_USERNAME"),
      os.Getenv("DB_PASSWORD"),
      os.Getenv("DB_NAME"))

    //fmt.Printf("db-username=%s db-pw=%s db-name=%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"));

  a.Run(":8080")
}
