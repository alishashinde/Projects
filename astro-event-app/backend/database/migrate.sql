
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100),
    email VARCHAR(100) UNIQUE,
    password TEXT,
    date_joined TIMESTAMP,
    location VARCHAR(100)
);

CREATE TABLE astronomical_events (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    type VARCHAR(100),
    description TEXT,
    date_time TIMESTAMP,
    visibility VARCHAR(255)
);

CREATE TABLE observations (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    event_id INTEGER REFERENCES astronomical_events(id),
    date TIMESTAMP,
    location VARCHAR(100),
    description TEXT
);

CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    observation_id INTEGER REFERENCES observations(id),
    user_id INTEGER REFERENCES users(id),
    content TEXT,
    created_at TIMESTAMP
);

CREATE TABLE notifications (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    event_id INTEGER REFERENCES astronomical_events(id),
    notification_date TIMESTAMP
);
