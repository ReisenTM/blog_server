package api

import (
	"blogX_server/api/article_api"
	"blogX_server/api/banner_api"
	"blogX_server/api/captcha_api"
	"blogX_server/api/image_api"
	"blogX_server/api/log_api"
	"blogX_server/api/site_api"
	"blogX_server/api/user_api"
)

type Api struct {
	SiteApi    site_api.SiteApi
	LogApi     log_api.LogApi
	ImageApi   image_api.ImageApi
	BannerApi  banner_api.BannerApi
	CaptchaApi captcha_api.CaptchaApi
	UserApi    user_api.UserApi
	ArticleApi article_api.ArticleApi
}

// App 实例化 以供外部调用Api
var App = Api{}
