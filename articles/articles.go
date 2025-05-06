package article

import (
	"context"
	"fmt"
	"strings"

	"github.com/himanshu-holmes/rlt-notify/repository"
)

type Article struct {
	articles repository.ArticleRepository
	summaryWordLimit int
}

func NewArticle(articles repository.ArticleRepository, summaryWordLimit int) *Article{
	return &Article{
		articles: articles,
		summaryWordLimit: summaryWordLimit,
	}
}
func (a *Article)ById(ctx context.Context,id int)(repository.SimpleSummaryArticle,error){
   article,err := a.articles.ById(ctx,id)
   if err != nil{
	return repository.SimpleSummaryArticle{},fmt.Errorf("error while retrieving a single article by id: %w",err)
   }

   return repository.SimpleSummaryArticle{
	ID: article.ID,
	Title: article.Title,
	Summary: article.Content,
	More: article.Content,
   },nil
}

func (a *Article) summarize(content string) string {
	words := strings.Split(strings.ReplaceAll(content, "\n", " "), " ")
	if len(words) > a.summaryWordLimit {
		return strings.Join(words[:a.summaryWordLimit], " ") + "..."
	}
	return strings.Join(words, " ")
}