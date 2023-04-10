\c car

CREATE TABLE person(
    ID int primary key,
    FirstName  varchar(500) NOT NULL,
    LastName varchar(500) NOT NULL,
    Age INT NOT NULL,
    Salary INT NOT NULL,
    Department varchar(500) NOT NULL,
    HireDate varchar(500) NOT NULL,
    Address varchar(500) NOT NULL,
    City varchar(500) NOT NULL,
    Country varchar(500) NOT NULL,
    Phone varchar(500) NOT NULL,
    Email varchar(500) NOT NULL,
    JobTitle varchar(500) NOT NULL,
    Manager varchar(500) NOT NULL,
    Status varchar(500) NOT NULL,
    Notes varchar(500) NOT NULL default 'test',
    Photo varchar(500) NOT NULL default 'jpg',
    CreatedDate TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
    CreatedBy TIMESTAMP NOT NULL default CURRENT_TIMESTAMP
);

insert into person ( ID, FirstName, LastName, Age, Salary, Department, HireDate, Address, City, Country, Phone, Email, JobTitle, Manager, Status) values (1, 'John', 'Doe', 1, 20, 'Support', '10/26/2022', 'Scott', 'Kizil', 'China', '227-406-6723', 'mk22doe0@php.net', 'test', 'Computer Systems Analyst I', 'description');

CREATE INDEX idx ON person (FirstName); -- index

