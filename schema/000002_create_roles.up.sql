DROP TABLE admins;

ALTER TABLE customers ADD COLUMN role int not null default 1;
