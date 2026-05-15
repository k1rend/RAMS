CREATE TABLE departments(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    department_id INTEGER REFERENCES departments(id) NOT NULL,
    manager_id INTEGER REFERENCES users(id),
    job_title VARCHAR(255),
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

CREATE TABLE resources(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    resource TEXT NOT NULL,
    resource_type VARCHAR(50) NOT NULL,
    owner_id INTEGER REFERENCES users(id),
    created_at DATE DEFAULT CURRENT_DATE
);

CREATE TABLE requests(
    id SERIAL PRIMARY KEY,
    applicant_id INTEGER REFERENCES users(id),
    resource_id INTEGER REFERENCES resources(id),
    access_type VARCHAR(50) NOT NULL,
    access_reason TEXT NOT NULL,
    status VARCHAR(50) NOT NULL,
    comments TEXT,
    created_at DATE DEFAULT CURRENT_DATE
); -- TODO: attachment field for supporting documents, e.g., PDF, images, etc.

CREATE TABLE approval_workflow(
    id SERIAL PRIMARY KEY,
    request_id INTEGER REFERENCES requests(id) ON DELETE CASCADE,
    approver_id INTEGER REFERENCES users(id),
    step_order INTEGER NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at DATE DEFAULT CURRENT_DATE
);