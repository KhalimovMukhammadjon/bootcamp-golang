\c order;

CREATE TABLE IF NOT EXISTS orders(
    id int primary key,
    name varchar(45) NOT NULL,
    couriers_id int NOT NULL,
    FOREIGN KEY (couriers_id) REFERENCES couriers(id)
);

CREATE TABLE IF NOT EXISTS couriers(
    id int primary key,
    name varchar(45) NOT NULL,
    branch_id int NOT NULL,
    address varchar(45) NOT NULL,
    car_id int NOT NULL,
    FOREIGN KEY (branch_id) REFERENCES branch(id),
    FOREIGN KEY (car_id) REFERENCES car_models(id)
);

CREATE TABLE IF NOT EXISTS branch(
    id int primary key,
    name varchar(45) NOT NULL,
    location varchar(45) NOT NULL
);

CREATE TABLE IF NOT EXISTS car_models(
    id int primary key,
    name varchar(45) NOT NULL,
    car_number varchar(45) NOT NULL
);

-- add new car to car's table
INSERT INTO car_models (id,name,car_number) values (1,'BMW','W022BM');
INSERT INTO car_models (id,name,car_number) values (2,'Ferrari','F712BB');
INSERT INTO car_models (id,name,car_number) values (3,'Porshce','P911ST');
INSERT INTO car_models (id,name,car_number) values (4,'Bugatti','P888ST');
INSERT INTO car_models (id,name,car_number) values (5,'Tahoe','B777SS');
INSERT INTO car_models (id,name,car_number) values (6,'Porshce','P911ST');


INSERT INTO branch (id,name,location) values(1,'Chorsu','Tashkent Chorsu');
INSERT INTO branch (id,name,location) values(2,'Ko`kcha','Tashkent Ko`kcha');
INSERT INTO branch (id,name,location) values(3,'Drujba','Tashkent Drujba');

-- add new courier to courier's table
INSERT INTO couriers (id,name,branch_id,address,car_id)values(2,'John',1,'Tashkent',1);

-- join
SELECT car_models.name, branch.name FROM couriers JOIN branch ON branch.id = couriers.branch_id
JOIN car_models ON car_models.id = couriers.car_id;

INSERT INTO couriers (id,name,branch_id,address,car_id)values(2,'John',2,'Tashkent',2);


-- add new order to order's table
INSERT INTO orders (id,name,couriers_id) values (1,'for debit card',1);






