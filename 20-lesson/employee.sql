create table employee (
    id int primary key not null,
    name varchar(50) not null,
    age int not null,
    position varchar(50) not null,
    country varchar(50) not null,
    salary int not null
);

create table factory(
    id int primary key not null,
    name varchar(50) not null,
    employee_id int,
    location varchar(50) default 'uzb',
    FOREIGN KEY (employee_id) REFERENCES employee(id),
    account_amount int not null
);

insert into employee values(1,'Akbar',24,'Senior','uzb',720);
insert into employee values(7,'Temur',21,'Junior','rus',450);
insert into employee values(8,'Husan',21,'Middle','rus',550);

insert into factory values(1,'Store',2,'Tashkent',500);
insert into factory values(2,'Store',1,'Samarkand',600);
insert into factory values(3,'Store',1,'Bukhara',700);

SELECT *from factory
JOIN employee on factory. employee_id = employee.id;

SELECT name, sum(salary) from employee GROUP BY name; -- sort by name and it introduces all salary

SELECT age, count(age) from employee GROUP BY age;

SELECT position, count(position) from employee GROUP BY position;


SELECT *from factory where location = 'Bukhara';
SELECT *from factory where age > 20;


SELECT *from factory
JOIN employee on factory.employee_id = employee.id;

SELECT * from factory
JOIN employee on factory.employee_id = employee.id where salary > 550;

SELECT *from employee where salary > 600;








