
package main

import(
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db, err :=sql.Open("postgres", "user=postgres password=Strong@Passw0rd dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin","*")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "GET,POST,PUT,DELETE,OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "Content-Type,Authorization")
         
		if c.Request.Method == "OPTIONS" {
	    c.AbortWithStatus(http.StatusNoContent)
		return
		}
	})

	router.POST("/user", func(c *gin.Context) {

		var user struct {
			Name   string `json:"name"`
			Phnno  string `json:"phnno"`
			Email  string `json:"email"`
			City   string `json:"city"`
			State  string `json:"state"`
		}

		if err :=c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H("error":"Invalid request"))
			return
		}

		result, err :=Exec("Insert into empl(name,phnno,email,city,state) VALUES ($1,$2,$3,$4,$5)",
	user.Name, user.Phnno, user.Email, user.city, user.State)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H("error", "Failed to get rows affected"))
		return
	}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return
		}

		c.JSON(http.StatusOk, gin.H{"message": fmt.Sprintf("%d rows affected", rowsAffected)})
	})
	router.Run(":8081")

}