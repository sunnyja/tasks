DROP DATABASE IF EXISTS tasks;
CREATE DATABASE tasks;

DROP TABLE IF EXISTS users, tasks, labels, task_label;

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

--задачи
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title TEXT,
    description TEXT,
    status INTEGER DEFAULT 0,
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now()),
    closed_at BIGINT DEFAULT 0,
    updated_at BIGINT DEFAULT 0,
    author_id INTEGER REFERENCES users(id) DEFAULT 0,
    assigned_id INTEGER REFERENCES users(id) DEFAULT 0
);

CREATE TABLE IF NOT EXISTS labels (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS task_label (
    task_id INTEGER REFERENCES tasks(id),
    label_id INTEGER REFERENCES labels(id)
);


--добавление исходных данных в БД
INSERT INTO users(name) VALUES
    ('user1'), 
    ('user2'), 
    ('user3'),
    ('user4'),
    ('user5');

INSERT INTO labels(name) VALUES
    ('срочная'), 
    ('обычная'), 
    ('горящая'),
    ('отложенная'),
    ('оценка');