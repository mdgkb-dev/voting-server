create table attachments
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    speech_id uuid REFERENCES speeches (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    meta       json         null,
    filename   varchar(255) null,
    ext        varchar(255) null,
    path       varchar(255) null,
    pivot      varchar(255) null
);


