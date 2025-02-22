
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";  -- Enable UUID extension if not already enabled

--DROP TABLE public.user CASCADE;
--DROP TABLE department CASCADE;



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

INSERT INTO department 
    (name, address) 
	VALUES
    ('Human Resources', '123 Main St, Springfield'),
    ('Finance', '456 Elm St, Springfield'),
    ('IT Support', '789 Oak St, Springfield');

INSERT INTO public.user (name, email, department_id, age, salary)
VALUES
('Alice Johnson', 'alice.johnson@example.com', (SELECT id FROM department WHERE name = 'Human Resources'), 30, 60000),
('Bob Smith', 'bob.smith@example.com', (SELECT id FROM department WHERE name = 'Human Resources'), 28, 55000),
('Charlie Brown', 'charlie.brown@example.com', (SELECT id FROM department WHERE name = 'Human Resources'), 35, 70000),
('Diana Prince', 'diana.prince@example.com', (SELECT id FROM department WHERE name = 'Finance'), 32, 80000),
('Ethan Hunt', 'ethan.hunt@example.com', (SELECT id FROM department WHERE name = 'Finance'), 29, 75000),
('Fiona Gallagher', 'fiona.gallagher@example.com', (SELECT id FROM department WHERE name = 'Finance'), 27, 52000),
('George Costanza', 'george.costanza@example.com', (SELECT id FROM department WHERE name = 'IT Support'), 40, 90000),
('Hannah Baker', 'hannah.baker@example.com', (SELECT id FROM department WHERE name = 'IT Support'), 22, 48000),
('Ian Malcolm', 'ian.malcolm@example.com', (SELECT id FROM department WHERE name = 'IT Support'), 38, 85000);