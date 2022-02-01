create table ratings
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    criteria_id uuid REFERENCES criterias (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    speech_id uuid REFERENCES speeches (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    user_id  uuid REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE DEFAULT NULL,
    rating     int          null,
    unique (user_id, speech_id, criteria_id)
);

