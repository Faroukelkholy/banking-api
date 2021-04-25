-- object: bank | type: SCHEMA --
CREATE SCHEMA bank;
set search_path = bank;
select * from pg_extension;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE  "account_name" AS ENUM('current', 'savings');
CREATE TABLE customers (id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,name character varying(50) NOT NULL);
CREATE TABLE accounts (id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,name "account_name" DEFAULT 'current', balance float DEFAULT 0.0, customer_id uuid REFERENCES customers (id));
CREATE TABLE transactions(id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY, sender_id uuid Not NULL REFERENCES accounts (id), receiver_id uuid Not NULL REFERENCES accounts (id), amount float Not NULL, created_at TIMESTAMP  NOT NULL DEFAULT now());
