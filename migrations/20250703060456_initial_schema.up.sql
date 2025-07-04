CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    login varchar(50) UNIQUE NOT NULL,
    email varchar(50) UNIQUE NOT NULL,
    password_hash varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS roles (
    id serial PRIMARY KEY,
    name varchar(50) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS projects (
    id serial PRIMARY KEY,
    name varchar NOT NULL,
    s_public bool DEFAULT false,
    created_at timestamp DEFAULT now()
);

CREATE TABLE IF NOT EXISTS boards (
    id serial PRIMARY KEY,
    name varchar(100) NOT NULL,
    created_at timestamp DEFAULT now(),

    project_id_fk int NOT NULL REFERENCES projects(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS labels (
    id serial PRIMARY KEY,
    name varchar(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS tickets (
    id serial PRIMARY KEY,
    name varchar(70) NOT NULL,
    description text,
    created_at timestamp DEFAULT now(),

    user_id_fk int REFERENCES users(id) ON DELETE SET NULL,
    label_id_fk int REFERENCES labels(id) ON DELETE SET NULL,
    board_id_fk int NOT NULL REFERENCES boards(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS tasks (
    id serial PRIMARY KEY,
    name varchar(100) NOT NULL,
    description text,
    created_at timestamp DEFAULT now(),

    user_id_fk int REFERENCES users(id) ON DELETE SET NULL,
    label_id_fk int REFERENCES labels(id) ON DELETE SET NULL,
    ticket_id_fk int NOT NULL REFERENCES tickets(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS project_members (
    user_id_fk int NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    project_id_fk int NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    role_id_fk int NOT NULL REFERENCES roles(id) ON DELETE RESTRICT,

    PRIMARY KEY (user_id_fk, project_id_fk)
)