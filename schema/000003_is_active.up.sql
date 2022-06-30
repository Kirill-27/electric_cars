ALTER TABLE customers ADD COLUMN is_active bool not null default true;
ALTER TABLE payment_methods ADD COLUMN is_active bool not null default true;
ALTER TABLE bookings ADD COLUMN is_active bool not null default true;
ALTER TABLE stations ADD COLUMN is_active bool not null default true;

