
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