CREATE TABLE IF NOT EXISTS users
(
    id               serial primary key not null,
    login            varchar(255) unique not null,
    password         varchar(255) not null,
    created_at       timestamp(0) not null default CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS products
(
    id               serial primary key not null,
    name             varchar(255) not null,
    price            float default 0,
    created_at       timestamp(0) not null default CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS orders
(
    id               serial primary key not null,
    user_id          int not null,
    product_id       int not null,
    price            float default 0,
    created_at       timestamp(0) not null default CURRENT_TIMESTAMP,
    CONSTRAINT fk_user  FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE SET NULL,
    CONSTRAINT fk_product  FOREIGN KEY(product_id) REFERENCES products(id) ON DELETE SET NULL
);