CREATE SCHEMA coffe_plus;
CREATE TABLE coffe_plus.users (
    id              UUID             PRIMARY KEY,
    version         BIGINT           DEFAULT 1,
    password_hash   VARCHAR(100)    CHECK ( char_length(password_hash) BETWEEN 6 AND 100 ),
    first_name      VARCHAR(100)     CHECK ( char_length(first_name) BETWEEN 3 AND 100),
    last_name       VARCHAR(100)     CHECK ( char_length(last_name) BETWEEN 3 AND 100),
    created_at      TIMESTAMPTZ,
    email           VARCHAR(250) UNIQUE NOT NULL,
    phone_number    VARCHAR(15)      CHECK (
        phone_number ~ '^\+[0-9]+$'
        AND
        char_length (phone_number) BETWEEN 10 AND 13
    ),

    role            VARCHAR(100) DEFAULT 'common'
);

CREATE TABLE coffe_plus.products (
    id           UUID PRIMARY KEY,
    version      BIGINT DEFAULT 1,
    name         VARCHAR(100) NOT NULL,
    description  VARCHAR(1000),
    price        INTEGER NOT NULL,
    is_available BOOLEAN NOT NULL,
    image_url    VARCHAR(1000) NOT NULL
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