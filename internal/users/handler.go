package users

import "github.com/gin-gonic/gin"

type UsersHandler struct {

}

func (h *UsersHandler) getUsers(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "Hello, World!",
    })
}

func (h *UsersHandler) Register(r *gin.Engine) {
    r.GET("/users", h.getUsers)
}
