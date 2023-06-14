package play010_gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

func Play() {
	router := gin.Default()

	router.POST("/users", createUser)
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUser)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)

	router.Run(":80")
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 在数据库或其他存储中创建用户
	users = append(users, user)

	c.JSON(http.StatusCreated, user)
}

func getUsers(c *gin.Context) {
	var users []User
	users = append(users, User{
		ID:       "abc",
		Username: "name",
		Email:    "1@2",
	})
	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	id := c.Param("id")

	// 在数据库或其他存储中查找特定用户
	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func updateUser(c *gin.Context) {
	id := c.Param("id")

	// 在数据库或其他存储中更新特定用户
	for i, user := range users {
		if user.ID == id {
			var updatedUser User
			if err := c.ShouldBindJSON(&updatedUser); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// 更新用户信息
			users[i] = updatedUser

			c.JSON(http.StatusOK, updatedUser)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")

	// 在数据库或其他存储中删除特定用户
	for i, user := range users {
		if user.ID == id {
			// 从切片中删除用户
			users = append(users[:i], users[i+1:]...)

			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
