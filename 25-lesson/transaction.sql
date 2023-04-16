\c car;

CREATE TABLE users_table(
    id int primary key,
    name varchar(45),
    balance int
);

--example 1
BEGIN TRANSACTION;

INSERT INTO users_table(id,name,balance) values (20,'John',1480);
INSERT INTO users_table(id,name,balance) values (21,'Doe',1000);

if (select count(*)from users_table where name = 'John' > 0)
BEGIN
    COMMIT;
END
ELSE
BEGIN
    ROLLBACK;
END

--example 2
INSERT INTO users_table(id,name,balance) values (38,'John',1480);
INSERT INTO users_table(id,name,balance) values (56,'Doe',1000);

BEGIN TRANSACTION;

UPDATE users_table SET balance = balance - 100 WHERE id = 38; 
UPDATE users_table SET balance = balance + 100 WHERE id = 56;

COMMIT;

