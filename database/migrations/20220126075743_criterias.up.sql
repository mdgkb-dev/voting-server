CREATE TABLE criterias
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    event_id uuid REFERENCES events (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
name varchar,
stars integer
);
