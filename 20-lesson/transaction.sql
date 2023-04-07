CREATE TABLE customers(
    id int primary key,
    name varchar(45),
    product_name varchar(45)
);

create table transaction(
    id int,
    amount float,
    currency varchar(45),
    create_at TIMESTAMP,
    customer_id int,
    FOREIGN KEY (customer_id) REFERENCES customers(id)
);

insert into customers values(1,'Muhammadjon','Pineapple');
insert into customers values(2,'Akbar','Banana');
insert into customers values(3,'John','Apple');

ALTER TABLE transaction ADD column create_at TYPE timestamptz;


insert into transaction values(1,32500.00,'UZS',2);
insert into transaction values(2,18500.00,'UZS',3);
insert into transaction values(3,28500.00,'UZS',1);

SELECT *from transaction
JOIN customers on transaction.customer_id = customers.id;
