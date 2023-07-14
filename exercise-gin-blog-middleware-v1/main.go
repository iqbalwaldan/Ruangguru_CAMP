package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var Posts = []Post{
	{ID: 1, Title: "Judul Postingan Pertama", Content: "Ini adalah postingan pertama di blog ini.", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 2, Title: "Judul Postingan Kedua", Content: "Ini adalah postingan kedua di blog ini.", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []User{
	{Username: "user1", Password: "pass1"},
	{Username: "user2", Password: "pass2"},
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, err := c.Request.BasicAuth()
		if !err {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		for _, user := range users {
			if username == user.Username && password == user.Password {
				return
			}
		}
		c.AbortWithStatus(http.StatusUnauthorized)
	} // TODO: replace this
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//Set up authentication middleware here // TODO: replace this
	r.Use(authMiddleware())

	r.GET("/posts", func(c *gin.Context) {
		// TODO: answer here
		ids := c.Query("id")
		if ids == "" {
			c.JSON(http.StatusOK, gin.H{"posts": Posts})
			return
		}

		id, err := strconv.Atoi(ids)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
			return
		}

		for _, pos := range Posts {
			if pos.ID == id {
				c.JSON(http.StatusOK, gin.H{"post": pos})
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Postingan tidak ditemukan"})

	})

	r.POST("/posts", func(c *gin.Context) {
		// TODO: answer here
		var pos Post
		err := c.ShouldBindJSON(&pos)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		id := len(Posts) + 1
		now := time.Now()
		pos.ID = id
		pos.CreatedAt = now
		pos.UpdatedAt = now

		Posts = append(Posts, pos)
		c.JSON(http.StatusCreated, gin.H{"post": pos, "message": "Postingan berhasil ditambahkan"})
	})

	return r
}

func main() {
	r := SetupRouter()

	r.Run(":8080")
}
