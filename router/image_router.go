package router

import (
	"blogX_server/api"
	"blogX_server/api/image_api"
	"blogX_server/middleware"
	"github.com/gin-gonic/gin"
)

func ImageRouter(r *gin.RouterGroup) {
	app := api.App.ImageApi
	r.POST("images", middleware.AuthMiddleware, app.UploadImageView)
	r.POST("images/transfer_deposit", middleware.AuthMiddleware, middleware.BindJsonMiddleware[image_api.TransferDepositRequest], app.TransferDepositView)
	r.GET("images", middleware.AdminMiddleware, app.ImageListView)
	r.DELETE("images", middleware.AdminMiddleware, app.ImageRemoveView)
	return
}
