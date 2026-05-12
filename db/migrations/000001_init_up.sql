CREATE TABLE departments(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    department_id INTEGER REFERENCES departments(id),
    is_active BOOLEAN DEFAULT TRUE,
    created_at DATE DEFAULT CURRENT_DATE
);

CREATE TABLE roles(
    id SERIAL PRIMARY KEY,
    code VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE user_roles(
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    role_id INTEGER REFERENCES roles(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, role_id)
);


insert into departments (name) values ('HR'), ('IT'), ('Security'), ('Finance');
insert into users (username, email, password_hash, first_name, last_name, department_id) values 
('admin', 'admin@example.com', 'hashed_password', 'Admin', 'Admin', 1);
insert into roles (code, name) values ('ADMIN', 'Administrator'), ('EMPLOYEE', 'Regular Employee'), ('APPROVER', 'Approver');
insert into user_roles (user_id, role_id) values (1, 1);

