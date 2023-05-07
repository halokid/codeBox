
CREATE SEQUENCE book_id_seq;

CREATE TABLE book(
                     id bigint PRIMARY KEY DEFAULT nextval('book_id_seq'),
                     title VARCHAR(100) not null,
                     isbn VARCHAR(255),
                     description VARCHAR(255),
                     page int,
                     price numeric
);

ALTER SEQUENCE book_id_seq
    OWNED BY book.id;