CREATE TABLE users
(
    id serial PRIMARY KEY,
    user_name text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);

CREATE TABLE calendars
(
    id serial PRIMARY KEY,
    calendar_name text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);

CREATE TABLE entries
(
    id serial PRIMARY KEY,
    calendar_id integer,
    user_id integer,
    entried_date smallint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    UNIQUE(calendar_id, user_id, entried_date),
    FOREIGN KEY (calendar_id) REFERENCES calendars(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX ON entries (calendar_id);
CREATE INDEX ON entries (user_id);