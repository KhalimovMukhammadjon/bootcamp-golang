CREATE TABLE course(
    id int primary key not null,
    name varchar(45),
    category_id int,
    price int,
    currency_id int,
    language_id int,
    FOREIGN KEY (category_id) REFERENCES category(id),
    FOREIGN KEY (currency_id) REFERENCES currency(id),
    FOREIGN KEY (language_id) REFERENCES language(id)
);

CREATE TABLE category(
    id int primary key not null,
    name varchar(45)
);

CREATE TABLE currency(
    id int primary key,
    name varchar(45)
);

CREATE TABLE language(
    id int primary key,
    name varchar(45)
);

INSERT INTO language (id,name) values(1,'Uzbek');
INSERT INTO language (id,name) values(2,'English');
INSERT INTO language (id,name) values(3,'Russian');