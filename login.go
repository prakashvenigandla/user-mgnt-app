package main()

import {
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
}

const {
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "23456"
	dbname   = "mydb"
}

func main() {
	//connect to database
	dbinfo := fmt.Sprint("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable")
	host, port, user, password, dbname)
	db, err :=sql.Open("postgres", dbinfo)
	if err := nil {
		log.Fatal(err)
	}
	 defer db.Close()
 
	 //create the gin oruter
	 router := gin.Default()

	 router.use(func(c *gin.context) {
		c.writer.Header().Set("Access-Control-Allow-Origin","*")
	 })

	 //Define the API endpoint for inserting a new person
	 router.Get("/employee", func(c *gin.Context) {

		username := c.Query("username")
		password := c.Query("password")

		var count int
		//Insert the new person into the database
		err := db.QueryRow("SELECT COUNT(*) FROM employee1 WHERE name=$1 password=$2", username, password)

		if err != nil {
			log.Fatal(err)
		}

		if count >0 {
			c.JSON(http.StatusOk, gin.H{"message": "Login success"})
        }else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Invalid username or password"})
		}
	 })
	 //start the server

	 router.Run(":8080")
}