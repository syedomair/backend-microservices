
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";  -- Enable UUID extension if not already enabled

DROP TABLE public.user CASCADE;
DROP TABLE department CASCADE;



CREATE TABLE IF NOT EXISTS department (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),  -- Generate a new UUID by default
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS public.user (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),  -- Generate a new UUID by default
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    department_id UUID REFERENCES department(id) ON DELETE SET NULL,  -- Foreign key constraint
    age INT CHECK (age >= 0),  -- Ensures age is non-negative
    salary DOUBLE PRECISION CHECK (salary >= 0)  -- Ensures salary is non-negative
);
