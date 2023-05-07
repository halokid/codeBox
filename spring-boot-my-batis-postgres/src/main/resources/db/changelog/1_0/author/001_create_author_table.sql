
CREATE SEQUENCE author_id_seq;

CREATE TABLE author(
                     id bigint PRIMARY KEY DEFAULT nextval('author_id_seq'),
                     lastname VARCHAR(100) not null,
                     firstname VARCHAR(255)
);

ALTER SEQUENCE author_id_seq
    OWNED BY author.id;