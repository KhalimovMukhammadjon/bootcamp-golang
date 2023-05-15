-- FUNCTION // create function in postgres sql
CREATE function sum_of_numbers (n integer)
RETURNS integer LANGUAGE plpgsql AS

$$
DECLARE 
    i integer := 1; --var i = 1 
    total integer := 0;
BEGIN
    WHILE i <= n LOOP -- {
        total := total + i;
        i := i +1;
    END LOOP;   -- }
    RETURN total;
END;
$$;


-- FIZZBUZZ
--version 1
DO
$$
BEGIN
    FOR n IN 1..100 LOOP
        IF n % 15 = 0 THEN
            RAISE NOTICE 'FizzBuzz';
        ELSIF n % 3 = 0 THEN
            RAISE NOTICE 'Fizz';
        ELSIF n % 5 = 0 THEN
            RAISE NOTICE 'Buzz';
        ELSE
            RAISE NOTICE '%', n;
        END IF;
    END LOOP;
END $$;


-- version 2
CREATE function fizzbuzz (n integer)
RETURNS varchar LANGUAGE plpgsql AS
$$
DECLARE
    output varchar := '';
BEGIN
    IF n % 15 = 0 THEN
        output := 'FizzBuzz';
    ELSIF n % 3 = 0 THEN
        output := 'Fizz';
    ELSIF n % 5 = 0 THEN
        output := 'Buzz';
    END IF;
    RETURN output;
END
$$;

--TRIGGER
CREATE OR REPLACE FUNCTION calculate_or
