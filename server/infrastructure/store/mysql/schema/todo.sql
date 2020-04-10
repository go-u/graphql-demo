CREATE TABLE todo
(
    id         SERIAL PRIMARY KEY ,
    uid        VARCHAR(128) CHARACTER SET ascii NOT NULL,
    text       VARCHAR(200) CHARACTER SET utf8mb4 NOT NULL
);
