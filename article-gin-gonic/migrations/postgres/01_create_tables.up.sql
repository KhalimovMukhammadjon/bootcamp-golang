CREATE TABLE IF NOT EXISTS "author" (
    "id" SERIAL PRIMARY KEY,
    "firstname" varchar(255) NOT NULL,
    "lastname" varchar(255) NOT NULL,
    "created_at" TIMESTAMP DEFAULT(Now()),
    "updated_at"  TIMESTAMP DEFAULT(Now())
);

CREATE TABLE IF NOT EXISTS "article" (
    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL UNIQUE,
    "body" TEXT,
    "author_id" INT,
    "deleted_at" TIMESTAMP DEFAULT NULL,
    "created_at" TIMESTAMP DEFAULT(Now()),
    "updated_at"  TIMESTAMP DEFAULT(Now()),
    CONSTRAINT fk_author FOREIGN KEY(author_id) REFERENCES author(id)
);

