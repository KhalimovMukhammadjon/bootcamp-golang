\c car

CREATE TABLE students(
    id int,
    name varchar(45),
    grade int
);  

INSERT INTO students(id,name,grade) values(1,'Akbar',5);
INSERT INTO students(id,name,grade) values(2,'Ali',4);
INSERT INTO students(id,name,grade) values(3,'John',3);
INSERT INTO students(id,name,grade) values(4,'Doe',5);
INSERT INTO students(id,name,grade) values(5,'Ali',5);
INSERT INTO students(id,name,grade) values(8,'Akbar',4);


SELECT grade, count(*) from students GROUP BY grade; -- 

--SELECT grade, count(*) from students GROUP BY grade ORDER BY grade;

SELECT * from students ORDER BY grade DESC; -- Z-A

SELECT grade, sum(grade) from students GROUP BY grade; -- sum

SELECT name, sum(grade) from students GROUP BY name; -- by name

-- SELECT name, COUNT(*) FROM students GROUP BY grade ORDER BY name;

INSERT INTO students(id,name,grade) values(9,'Husan',80);
INSERT INTO students(id,name,grade) values(10,'Akbar',75);

SELECT name, SELECT MAX(grade) FROM students;

--SELECT MAX(price) FROM car;

--  create function
create function result(a int, b int) returns int
language plpgsql
AS $$
DECLARE 
    result integer;
    begin
    result = a*b;
    return result;
    end;
    $$;

select result(3,4);
SELECT AVG(grade) from students;

-- 
CREATE function tr() returns float
LANGUAGE plpgsql
AS 
$$
    BEGIN
    RETURN AVG(grade) FROM students;
    END;
$$;
SELECT tr();



