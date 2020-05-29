### Установка переменных окружения
```shell script
cp .env.dist .env
```

### Запуск
Запуск сервисов
```shell script
docker-compose up
```

### Миграции
```shell script
docker run \
--rm -v $(pwd)/migrations:/migrations \
--network articles-aggregator_default \
migrate/migrate \
-database "postgres://postgres:password@postgres:5432/postgres?sslmode=disable" \
-path /migrations up
```

### Запуск тестов
```shell script
make test
```

### Использование
Запустить парсер
 - `endpoint` - Страница со списком статей
 - `urls` - XPath для выборки из страницы со списком статей ссылки на страницы с детальтными статьями
 - `title` - Xpath для выборки названия статьи из страницы с детальным описанием
 - `text` - XPath для выборки текста статьи из страницы с детальным описанием 
```shell script
docker run \
--rm \
-v $(pwd):/go/src/github.com/sekerin/artticles-aggregator \
--network articles-aggregator_default \
-e PARSER_ADDRESS=parser:3030 \
golang:1.13 \
go run github.com/sekerin/artticles-aggregator/cmd parser \
--endpoint https://habr.com/ru/ \
--urls="//*/a[@class='post__title_link']/@href" \
--title="//*/span[contains(@class, 'post__title-text')]" \
--text="//*/div[@id='post-content-body']"
```

Запуск поиска по названию
- `search` - строка для поиска
```shell script
docker run \
--rm \
-v $(pwd):/go/src/github.com/sekerin/artticles-aggregator \
--network articles-aggregator_default \
-e SEARCH_ADDRESS=search:3030 \
golang:1.13 \
go run github.com/sekerin/artticles-aggregator/cmd search --search <Text for search>
```
