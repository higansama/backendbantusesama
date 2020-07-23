package router

import (
	"backend_skripsi/controller"

	"github.com/gin-gonic/gin"
)

func RouterHome(r *gin.Engine) {
	rLogin := r.Group("/")
	// rLogin.Use(middleware.AuthUser)
	rLogin.POST("/create/campaign", controller.CreateCampaign)
	rLogin.POST("/follow/campaign", controller.FollowCampaign)
	rLogin.POST("/like/campaign", controller.LikeCampaign)
	rLogin.POST("/remember/campaign", controller.RememberCampaign)
	rLogin.POST("/action/campaign", controller.ActionCampaign)
	rLogin.POST("/share/campaign", controller.ShareCampaign)
	rLogin.GET("/followed/campaign", controller.GetFollowCampaign)
	rLogin.GET("/liked/campaign", controller.GetLikeCampaign)
	rLogin.GET("/remembered/campaign", controller.GetRememberCampaign)
	rLogin.GET("/actions/campaign", controller.GetActionCampaign)
	rLogin.GET("/shared/campaign", controller.GetShareCampaign)
	rLogin.POST("/tanggapi/campaign", controller.TambahKomentar)
	rLogin.GET("/detail/campaign", controller.DetailCampaign)

	// rLogin.POST("/process", controller.LoginProses)
	// r.GET("/logout", controller.Logout)
}
