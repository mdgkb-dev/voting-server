-- auto-generated definition
create table events
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name       varchar(255) null,
    date_from  time(6)      null,
    date_to    time(6)      null
--     created_at      null,
--     updated_at datetime     null
);

