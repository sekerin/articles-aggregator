CREATE TABLE IF NOT EXISTS articles
(
    id        SERIAL       NOT NULL
        CONSTRAINT articles_pk PRIMARY KEY,
    title     varchar(255) NOT NULL,
    text      text,
    url       varchar(255) NOT NULL,
    title_tsv tsvector
);

CREATE UNIQUE INDEX IF NOT EXISTS articles_url_uindex ON articles (url);

CREATE INDEX IF NOT EXISTS articles_title_gin ON articles USING gin (title_tsv);
