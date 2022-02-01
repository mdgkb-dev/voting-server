CREATE TABLE speeches
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    event_id uuid REFERENCES events (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    title       varchar(255) null,
    description varchar(255) null,
    author      varchar(255) null,
--     date_from   time(6)      null,
--     date_to     time(6)      null,
--     created_at  datetime     null,
--     updated_at  datetime     null,
    meta        json         null,
--     user_id     int          null,
--     category_id int unsigned null,
    message     text     null
);
