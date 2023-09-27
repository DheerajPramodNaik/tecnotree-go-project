package routes

import (
	"tecnotree-go-project/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, p controllers.UserController) {
	user := r.Group("/api/user")
	user.POST("/register", p.HandleRegister)
	user.POST("/login", p.HandleLogin)
	user.POST("/logout", p.HandleLogout)
}

func ProductRoutes(r *gin.Engine, p controllers.ProductController) {
	user := r.Group("/api/product")
	user.POST("/insert", p.AddProduct)
	user.POST("/getAll", p.SearchProducts)
	user.POST("/search", p.SearchProductById)
}
