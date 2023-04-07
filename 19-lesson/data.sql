\c car;


CREATE TABLE IF NOT EXISTS store_users(
    id int primary key,
    name varchar(45) NOT NULL
);

CREATE TABLE IF NOT EXISTS store_product(
    id int primary key,
    name varchar(45) NOT NULL,
    price numeric,
    count int
);

CREATE TABLE IF NOT EXISTS store_orders(
    id int primary key,
    user_id int NOT NULL,
    product_id int NOT NULL,
    count int,
    FOREIGN KEY (user_id) REFERENCES store_users(id),
    FOREIGN KEY (product_id) REFERENCES store_product(id)
);

INSERT INTO store_users (id,name) values (1,'John');
INSERT INTO store_users (id,name) values (2,'Ali');
INSERT INTO store_users (id,name) values (3,'Akbar');

INSERT INTO store_users (balance) values ('123000') where id = 1;

INSERT INTO store_product (id,name,price,count) values (1,'Apple',17800,5);
INSERT INTO store_product (id,name,price,count) values (2,'Pineapple',17800,4);
INSERT INTO store_product (id,name,price,count) values (3,'Banana',17800,6);

INSERT INTO store_orders (id, user_id,product_id,count) values (1,2,2,6);
INSERT INTO store_orders (id, user_id,product_id,count) values (2,1,3,4);


SELECT *from store_orders 
join store_users on store_orders.user_id = store_users.id
join store_product on store_orders.product_id = store_product.id;

-- add new column to users
ALTER TABLE store_users ADD balance int;

-- edit column data type
ALTER TABLE store_users ALTER balance type int; 

UPDATE store_users SET balance = 320 WHERE id = 2;
