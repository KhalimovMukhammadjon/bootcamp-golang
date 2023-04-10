\c car;

-- factorial
CREATE FUNCTION factorial(n integer) -- 100
RETURNS integer 
LANGUAGE plpgsql 
AS
$$
DECLARE
  i integer := 1;
  total integer := 1;
BEGIN 
  WHILE i <= n LOOP
    total = total * i;
    i := i + 1;
  END LOOP;
  RETURN total;
END;
$$;
select factorial(10);

-- sum
CREATE FUNCTION counter(n integer) -- 10
RETURNS integer 
LANGUAGE plpgsql 
AS
$$
DECLARE
  i integer := 1;
  total integer := 1;
BEGIN 
  WHILE i <= n LOOP
    total = total * i;
    i := i + 1;
  END LOOP;
  RETURN total;
END;
$$;

select counter(10);



--triger
CREATE OR REPLACE FUNCTION trInsertEmployee() RETURNS TRIGGER AS $$
DECLARE 
  curInserted CURSOR FOR 
    SELECT Id 
    FROM inserted;
  current_row record;
  auditMessage varchar(1000);
BEGIN
  OPEN curInserted;
  LOOP
    FETCH curInserted INTO current_row;
    EXIT WHEN NOT FOUND;
    
    auditMessage := 'New employee with Id = '  current_row.Id  ' is added at ' || CURRENT_TIMESTAMP;
    INSERT INTO Employee_Audit_Test (Message)
    VALUES (auditMessage);
  END LOOP;
  CLOSE curInserted;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

Create table Employee_Audit_Test(
  Audit varchar(45)
);