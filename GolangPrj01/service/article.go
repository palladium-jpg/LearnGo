package service

import (
	"GolangPrj01/dao"
	"GolangPrj01/model"
)

func GetOneArticle(id int64) (*model.Article, error) {
	return dao.SelectOneArticle(id)
}

func GetArticleList(page int, pageSize int) ([]*model.Article, error) {
	articles, err := dao.SelectAllArticle(page, pageSize)
	if err != nil {
		return nil, err
	} else {
		return articles, err
	}

}
