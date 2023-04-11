
CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; 

CREATE TABLE orders(
    id uuid DEFAULT uuid_generate_v4 (),
    primary key (id),
    customer_name varchar(30),
    customer_address varchar(50),
    customer_location jsonb,
    note varchar(45),
    from_time timestamp,
    to_time timestamp,
    branch_id int,
    courier_id int,
    category_id int,
    status_id int,
    comment varchar(45),
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP,
    delivery_day date,
    is_rejected_by_elma boolean,
    instance_id int,
    customer_pinfl int,
    is_confirmed boolean,
    returned boolean,
    card_number int,
    FOREIGN KEY (courier_id) REFERENCES couriers(id),
    FOREIGN KEY (category_id) REFERENCES categories(id),
    FOREIGN KEY (status_id) REFERENCES status(id)
);

INSERT INTO orders values(1,'Usmon','Tashkent Chorsu 21B','location',);

CREATE TABLE status(
    id int primary key,
    name varchar(45)
);
INSERT INTO status values(1,'New');
INSERT INTO status values(2,'In process');
INSERT INTO status values(3,'Delivered');


CREATE TABLE couriers(
    id int primary key,
    name varchar(45),
    phone int,
    region_id int,
    district_id int,
    address varchar(45),
    location varchar(45),
    image varchar(45),
    branch_id int,
    vehicle_model varchar(45),
    vehicle_number varchar(20),
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP,
    FOREIGN KEY (region_id) REFERENCES region(id)
);

-- SELECT * from couriers
-- JOIN region on couriers.region_id = region.id;

SELECT couriers.id, couriers.name, region.name from couriers
JOIN region on couriers.region_id = region.id;

INSERT INTO couriers values(1,'John',8986777,1,2,'Chorsu','24434.34.884.67','image.jpg',2,'Chevrolet','01 999 OFB');
INSERT INTO couriers values(2,'Ali',3246577,2,2,'Bukhara','812633.34.884.67','image.jpg',2,'Audi','80 212 TBC');
INSERT INTO couriers values(3,'Tohir',3289555,3,2,'Jomboy','124734.34.884.67','image.jpg',2,'Chevrolet','01 888 BTM');
INSERT INTO couriers values(4,'Doe',6758908,1,2,'Ko`kcha 21 A','89434.34.884.67','image.jpg',2,'Chevrolet','01 474 OFB');
INSERT INTO couriers values(5,'Temur',7778998,2,2,'Bukhara','812633.34.884.67','image.jpg',2,'Tahoe','80 212 TBM');
INSERT INTO couriers values(6,'Hasan',2326776,3,2,'Jomboy','234734.34.884.67','image.jpg',2,'Hyundai','01 767 VSP');

-- it is region
CREATE TABLE region(
    id int primary key,
    name varchar(45)
);
INSERT INTO region values(1,'Tashkent');
INSERT INTO region values(2,'Bukhara');
INSERT INTO region values(3,'Samarkand');
INSERT INTO region values(4,'Navoi');
INSERT INTO region values(5,'Khorezm');


-- CREATE TABLE courier_timetable(
--     courier_id int,
--     branch_id int,
--     date date,
--     from_time timestamp,
--     to_time timestamp,
--     created_at CURRENT_TIMESTAMP,
--     updated_at CURRENT_TIMESTAMP
--     FOREIGN KEY (courier_id) REFERENCES couriers(id)
-- );

CREATE TABLE categories(
    id int primary key,
    name varchar(45),
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP
);

INSERT INTO categories values(1,'Phone');
INSERT INTO categories values(2,'Computer');
INSERT INTO categories values(3,'Vegetable');
INSERT INTO categories values(4,'Fruit');
INSERT INTO categories values(5,'Grocery');

CREATE TABLE products(
    id int primary key,
    name varchar(45),
    --docs jsonb,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP,
    category_id int,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);
INSERT INTO products (id, name, category_id) values(1,'iPhone 14 Pro',1);
INSERT INTO products (id, name, category_id) values(2,'Pineapple',4);
INSERT INTO products (id, name, category_id) values(3,'Samsung S22',1);
INSERT INTO products (id, name, category_id) values(4,'Watermelon',4);
INSERT INTO products (id, name, category_id) values(5,'Tomato',3);
INSERT INTO products (id, name, category_id) values(6,'Lays',5);
INSERT INTO products (id, name, category_id) values(7,'MacBook Pro 14',2);

SELECT products.id,products.name, categories.name  from products
JOIN categories on products.category_id = categories.id;

CREATE TABLE documents(
    id int primary key,
    name varchar(45),
    created_at CURRENT_TIMESTAMP,
    updated_at CURRENT_TIMESTAMP,
    rule jsonb
);