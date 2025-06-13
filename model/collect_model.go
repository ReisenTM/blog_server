package model

// CollectModel 收藏夹
type CollectModel struct {
	Model
	Title        string                    `gorm:"size:32" json:"title"`
	Info         string                    `gorm:"size:256" json:"info"`
	Cover        string                    `gorm:"size:256" json:"cover"`
	ArticleList  []UserArticleCollectModel `gorm:"foreignKey:CollectID" json:"-"`
	ArticleCount uint                      `json:"article_count"`
	IsDefault    bool                      `json:"is_default"` //是否是默认收藏夹
	UserID       uint                      `json:"user_id"`
	UserModel    UserModel                 `gorm:"foreignKey:UserID" json:"user_model"`
}
