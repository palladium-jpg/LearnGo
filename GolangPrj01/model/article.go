package model

type Article struct {
	ArticleId int64  `xorm:"column:aritcleId" ,json:"articleId"`
	Subject   string `xorm:"column:subject" ,json:"title"`
	Url       string `xorm:"column:url" ,json:"url"`
	ImgUrl    string `xorm:"imgUrl" ,json:"imgUrl"`
	HeadUrl   string `xorm:"headUrl" ,json:"headUrl"`
}

func (Article) TableName() string {
	return "article"
}
