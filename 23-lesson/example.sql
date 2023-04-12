CREATE extension if not exists "uuid-ossp";

CREATE TABLE products (
    id uuid primary key DEFAULT uuid_generate_v4(),
    name varchar(100),
    price numeric
);

CREATE TABLE users (
    id uuid primary key DEFAULT uuid_generate_v4(),
    name varchar(100),
    balance numeric
);

CREATE TABLE shop_cart (
    product_id uuid references products(id),
    user_id uuid references users(id),
    created_at timestamp default current_timestamp
);

INSERT INTO products (name,price) values('Apple',18900);
INSERT INTO products (name,price) values('Grape',28800);
INSERT INTO products (name,price) values('Pineapple',48990);
INSERT INTO products (name,price) values('Watermelon',32900);
INSERT INTO products (name,price) values('Strawberry',76900);

INSERT INTO users (name,balance) values('Abdurashid',5600000);
INSERT INTO users (name,balance) values('Ali',2880000);
INSERT INTO users (name,balance) values('John',4899000);
INSERT INTO users (name,balance) values('Doe',1790000);
INSERT INTO users (name,balance) values('Temur',5230000);

INSERT INTO shop_cart values('3acbb3de-e33f-4e4b-b964-7ce8ed451c22','9c7dc930-2e82-4f5f-a4e2-ba522b2e59bb');
INSERT INTO shop_cart values('120372e6-280b-4c70-b2d6-742d427c8098','d915ff72-d747-4d54-89aa-d5429114cbec');
INSERT INTO shop_cart values('120372e6-280b-4c70-b2d6-742d427c8098','a13628a6-7c97-4507-b3e6-65cbb94c7eb4');
INSERT INTO shop_cart values('120372e6-280b-4c70-b2d6-742d427c8098','6b563c36-82f3-4e5d-9692-9f977773246a');
INSERT INTO shop_cart values('120372e6-280b-4c70-b2d6-742d427c8098','9c7dc930-2e82-4f5f-a4e2-ba522b2e59bb');


-- SELECT DISTINCT users.name, ARRAY_AGG(products.name)from shop_cart
-- JOIN users on shop_cart.user_id = users.id
-- JOIN products on shop_cart.product_id = products.id GROUP BY users.name; -- use

SELECT users.name,users.balance, products.name, products.price from shop_cart
JOIN users on shop_cart.user_id = users.id
JOIN products on shop_cart.product_id = products.id;


select * from shop_cart
where shop_cart.user_id = '9c7dc930-2e82-4f5f-a4e2-ba522b2e59bb' and shop_cart.product_id = '120372e6-280b-4c70-b2d6-742d427c8098';

delete from where orders.user_id = '' and orders.product_id = '';

-- totat amount
create or replace procedure payment6(user_id uuid, product_id uuid) language plpgsql
as
$$
    declare
        prices numeric;
        balances numeric;
    begin
        select users.balance from users into balances where users.id = user_id;
        select products.price from products into prices where products.id = product_id;

        if balances > prices then update users set balance = balance - (select products.price from products where products.id = product_id) where users.id = user_id; commit;
        else  rollback;
        end if;
    end;
$$;
