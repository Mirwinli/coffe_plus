CREATE SCHEMA coffe_plus;
CREATE TABLE coffe_plus.users (
    id              UUID             PRIMARY KEY,
    version         BIGINT           DEFAULT 1,
    password_hash   TEXT             CHECK ( char_length(password_hash) >= 3),
    first_name      VARCHAR(100)     CHECK ( char_length(first_name) BETWEEN 3 AND 100),
    last_name       VARCHAR(100)     CHECK ( char_length(last_name) BETWEEN 3 AND 100),
    created_at      TIMESTAMPTZ,
    email           VARCHAR(250) UNIQUE NOT NULL,
    phone_number    VARCHAR(15)      CHECK (
        phone_number ~ '^\+?[0-9]{10,13}$'
        AND
        char_length (phone_number) BETWEEN 10 AND 13
    ),

    role            VARCHAR(100) DEFAULT 'common'
);

CREATE TABLE coffe_plus.category (
    id UUID PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    name VARCHAR(100) UNIQUE NOT NULL CHECK ( char_length(name) BETWEEN 3 AND 100)
);

CREATE TABLE coffe_plus.products (
    id           UUID PRIMARY KEY,
    version      BIGINT DEFAULT 1,
    name         VARCHAR(100) UNIQUE NOT NULL CHECK ( char_length(name) BETWEEN 3 AND 100),
    description  VARCHAR(1000)         CHECK ( char_length(description) BETWEEN 3 AND 1000),
    price        NUMERIC(12, 2) NOT NULL,
    is_available BOOLEAN NOT NULL,
    category_id   UUID NOT NULL REFERENCES coffe_plus.category(id),
    public_id    TEXT NOT NULL,
    image_url    TEXT NOT NULL
);

CREATE TABLE coffe_plus.orders (
    id          UUID                  PRIMARY KEY,
    user_id     UUID         NOT NULL REFERENCES coffe_plus.users(id),
    status      VARCHAR(150) NOT NULL,
    total_price DECIMAL       NOT NULL,
    created_at  TIMESTAMPTZ  NOT NULL
);

CREATE TABLE coffe_plus.order_items (
    id              UUID             PRIMARY KEY,
    order_id        UUID    NOT NULL REFERENCES coffe_plus.orders(id),
    product_id      UUID    NOT NULL REFERENCES coffe_plus.products(id),
    quantity        BIGINT  NOT NULL,
    price_at_time   DECIMAL NOT NULL
);

INSERT INTO coffe_plus.users (id,password_hash, first_name, last_name, created_at, email, phone_number,role)
VALUES (gen_random_uuid(),'$2a$12$PtJeLujI1V7zaULvLqteG.uhPzLrhX8cNblBAgghlJz5XcVPVz7hK','Назар','Кушнірюк',NOW(),'mirwinli.tech@gmail.com','0974526184','admin');