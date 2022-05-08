package dao

import (
	"GolangPrj01/gobal"
	"GolangPrj01/model"
	"fmt"
)

func SelectOneArticle(articleId int64) (*model.Article, error) {
	AID := model.Article{ArticleId: articleId}
	get, err := gobal.SetupDBLink().Get(AID)
	if err != nil {
		return nil, err
	} else {
		if !get {
			return nil, err
		} else {
			return &AID, err
		}
	}
}
func SelectCountAll() (int64, error) {
	//var count int
	i, err := gobal.SetupDBLink().Where("isPublish=?", 1).Count(new(model.Article))
	if err != nil {
		return 0, err
	} else {
		return i, err
	}

}

func SelectAllArticle(page int, pageSize int) ([]*model.Article, error) {
	var articles []*model.Article
	rows, err := gobal.SetupDBLink().Table(model.Article{}.TableName()).Where("isPublish=?", 1).Limit(pageSize).Rows(articles)

	if err != nil {
		fmt.Println("sql is error:")
		fmt.Println(err)
		return nil, err
	}

	//defer rows.Close()
	//var articles []*model.Article
	for rows.Next() {
		//fmt.Println("rows.next:")
		r := &model.Article{}
		if err := rows.Scan(&r.ArticleId, &r.Subject, &r.Url); err != nil {
			fmt.Println("rows.next:")
			fmt.Println(err)
			return nil, err
		}
		articles = append(articles, r)
	}
	return articles, nil

}
