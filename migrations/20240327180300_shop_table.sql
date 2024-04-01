-- +goose Up
create table products
(
    id   serial primary key,
    name text not null
);

create table orders
(
    id integer primary key
);
create table orders_products
(
    id         serial primary key,
    product_id integer not null,
    order_id   integer not null,
    count      integer not null,
    foreign key (product_id) references products (id) on delete cascade,
    foreign key (order_id) references orders (id) on delete cascade
);
create table shelfs
(
    id         serial primary key,
    product_id integer not null,
    name       text    not null,
    is_main    bool    not null,
    foreign key (product_id) references products (id) on delete cascade
);

insert into products (name)
values ('Ноутбук'),
       ('Телевизор'),
       ('Телефон'),
       ('Системный блок'),
       ('Часы'),
       ('Микрофон');

insert into orders (id)
values (10),
       (11),
       (14),
       (15);

insert into orders_products (product_id, order_id, count)
values (1, 10, 2),
       (1, 14,3),
       (2, 11,3),
       (3, 10,1),
       (4, 14,4),
       (5, 15,1),
       (6, 10,1);

insert into shelfs (product_id, name, is_main)
values (1, 'А', true),
       (2, 'А', true),
       (3, 'Б', true),
       (3, 'З', false),
       (3, 'В', false),
       (4, 'Ж', true),
       (5, 'Ж', true),
       (5, 'А', false),
       (6, 'Ж', true);

CREATE INDEX shelf_name ON shelfs (name);
CREATE INDEX shelf_main ON shelfs (is_main);
CREATE INDEX shelf_product ON shelfs (product_id);

-- +goose Down
drop table orders_products;
drop table shelfs;
drop table products;
drop table orders;
