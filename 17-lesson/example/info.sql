\c car;

CREATE TABLE cars(
    id int primary key,
    name varchar(45),
    price integer,
    color varchar(45)
);

INSERT INTO cars (id,name,price,color) values (1,'BMW',195000,'Black');
INSERT INTO cars (id,name,price,color) values (2,'Ferrari',378000,'Red');
INSERT INTO cars (id,name,price,color) values (3,'Tesla',178000,'White');
INSERT INTO cars (id,name,price,color) values (4,'Tesla',156000,'Green');

UPDATE cars SET price = 240000 WHERE id = 1;

DELETE FROM cars WHERE name = 'Tesla';

SELECT * FROM cars;



