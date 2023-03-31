START TRANSACTION;
-- DO $$ BEGIN IF NOT EXISTS(
--         SELECT 1
--         FROM pg_namespace
--         WHERE
--             nspname = 'order'
--     ) THEN CREATE SCHEMA "order";
-- END IF;
-- END $$;
CREATE SCHEMA
    IF
    NOT EXISTS "orders";
CREATE EXTENSION
    IF
    NOT EXISTS "uuid-ossp";
CREATE TABLE "orders".orders (
                                 order_id serial PRIMARY KEY,
                                 user_id VARCHAR ( 255 ) NOT NULL,
                                 order_date DATE NOT NULL,
                                 amount NUMERIC ( 8, 2 ) NOT NULL,
                                 order_state VARCHAR ( 2 ) NOT NULL
);
CREATE TABLE "orders".products (
                                   ID serial PRIMARY KEY,
                                   product_name VARCHAR ( 50 ) UNIQUE NOT NULL,
                                   price NUMERIC ( 8, 2 ) NOT NULL,
                                   category VARCHAR ( 50 ) NOT NULL
);
CREATE TABLE "orders".line_items (
                                     ID serial PRIMARY KEY,
                                     order_id INT NOT NULL,
                                     product_id INT NOT NULL,
                                     quantity INT NOT NULL,
                                     price NUMERIC ( 8, 2 ) NOT NULL,
                                     CONSTRAINT fk_line_items_orders_order_temp_id FOREIGN KEY ( order_id ) REFERENCES "orders".orders ( order_id ),
                                     CONSTRAINT fk_line_items_products_id FOREIGN KEY ( product_id ) REFERENCES "orders".products ( ID )
);

CREATE TABLE "orders".users (
                                ID serial PRIMARY KEY,
                                username VARCHAR ( 50 ) UNIQUE NOT NULL,
                                passwd VARCHAR ( 50 ) NOT NULL,
                                email VARCHAR ( 355 ) UNIQUE NOT NULL,
                                created_on TIMESTAMP NOT NULL
);
CREATE UNIQUE INDEX ix_line_items_id ON "orders".line_items ( ID );
CREATE INDEX ix_line_items_order_id ON "orders".line_items ( order_id );
CREATE UNIQUE INDEX ix_orders_id ON "orders".orders ( order_id );
CREATE UNIQUE INDEX ix_product_id ON "orders".products ( ID );
CREATE UNIQUE INDEX ix_users_id ON "orders".users ( ID );
COMMIT;