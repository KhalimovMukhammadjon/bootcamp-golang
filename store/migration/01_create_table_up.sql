-- CREATE TABLE IF NOT EXISTS movie (
--     id SERIAL PRIMARY KEY,
--     title VARCHAR(50) NOT NULL,
--     description VARCHAR(200) NOT NULL,
--     year date not null,
--     actor_id integer,
--     rating float not null,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
--     foreign key (actor_id) references actor(id)
-- );

-- CREATE TABLE IF NOT EXISTS actor (
--     id SERIAL PRIMARY KEY,
--     firstname VARCHAR(50) NOT NULL,
--     lastname VARCHAR(50) NOT NULL,
--     age int NOT NULL,
--     country VARCHAR(45),
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
-- );

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";  

CREATE TABLE "customers" (
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR NOT NULL, 
    "phone" VARCHAR NOT NULL
);

CREATE TABLE "users" (
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR NOT NULL, 
    "phone_number" VARCHAR NOT NULL
);

CREATE TABLE "couriers" (
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR NOT NULL, 
    "phone_number" VARCHAR NOT NULL
);

CREATE TABLE "categories" (
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR NOT NULL,
);

CREATE TABLE "products" (
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR NOT NULL, 
    "price" NUMERIC NOT NULL,
    "category_id" UUID NOT NULL REFERENCES "categories"("id")
);

CREATE TABLE "orders" (
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "price" NUMERIC NOT NULL,
    "user_id" UUID NOT NULL REFERENCES "users"("id"),
    "customer_id" UUID NOT NULL REFERENCES "customers"("id"),
    "courier_id" UUID NOT NULL REFERENCES "couriers"("id"),
    "product_id" UUID NOT NULL REFERENCES "products"("id"),
    "quantity" NUMERIC NOT NULl DEFAULT 1
);

