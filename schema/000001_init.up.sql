CREATE TABLE customers
(
    id serial not null unique,
    name     text not null,
    email    text not null,
    phone  text not null,
    username text not null,
    password text not null
);

CREATE TABLE admins
(
    id serial not null unique,
    email    text not null,
    username text not null,
    password text not null
);

CREATE TABLE stations
(
    id serial not null unique,
    is_free  boolean not null,
    location_x float not null,
    location_y float not null,
    value_per_min int not null
);

CREATE TABLE payment_methods
(
    id serial not null unique,
    customer_id  int not null,
    details text not null
);

CREATE TABLE bookings
(
    id serial not null unique,
    customer_id  int not null,
    station_id  int not null,
    payment_method_id int not null,
    start_time timestamp without time zone not null,
    end_time timestamp without time zone not null,
    is_paid boolean not null
);