package articles

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Article = struct {
	Title string `Db:"title"`
	Url   string `Db:"url"`
	Text  string `Db:"text"`
}

type ArticleRepository struct {
	Db *sqlx.DB
}

type ConnectOptions struct {
	User     string
	Password string
	Name     string
	Hostname string
	Port     string
}

func NewArticleRepository(opts ConnectOptions) (*ArticleRepository, error) {
	repo := &ArticleRepository{}

	db, err := sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
			opts.User,
			opts.Password,
			opts.Name,
			opts.Hostname,
			opts.Port,
		),
	)
	if err != nil {
		return nil, err
	}

	repo.Db = db

	return repo, nil
}

func (repo *ArticleRepository) Save(article Article) error {
	_, err := repo.Db.NamedExec(`
INSERT INTO 
articles (title, url, text, title_tsv) 
VALUES (:title, :url, :text, to_tsvector(:title)) 
ON CONFLICT DO NOTHING
`, article)

	if err != nil {
		return err
	}

	return nil
}

func (repo *ArticleRepository) GetByTitle(title string) ([]Article, error) {
	var articles []Article

	err := repo.Db.Select(&articles, `
SELECT title, url, text 
FROM articles 
WHERE title_tsv @@ to_tsquery($1)
`, title)

	if err != nil {
		return nil, err
	}

	return articles, nil
}
