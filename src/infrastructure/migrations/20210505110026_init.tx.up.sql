--
-- bks_genres
--

CREATE TABLE public.bks_genres
(
    id         serial PRIMARY KEY,
    parent_id  integer                       NULL,
    title      character varying(512) UNIQUE NOT NULL,
    created_at timestamp(0)                  NOT NULL,
    updated_at timestamp(0)                  NOT NULL,
    deleted_at timestamp(0)                  NULL
);

INSERT INTO bks_genres
VALUES (1, null, 'All', '2022-04-01 13:00:00', '2022-04-01 13:00:00', null),
       (2, 1, 'DDD', '2022-04-01 13:00:00', '2022-04-01 13:00:00', null),
       (3, 1, 'Patterns', '2022-04-01 13:00:00', '2022-04-01 13:00:00', null),
       (4, 1, 'Testing', '2022-04-01 13:00:00', '2022-04-01 13:00:00', null),
       (5, 1, 'Microservices', '2022-04-01 13:00:00', '2022-04-01 13:00:00', null)
;

ALTER TABLE public.bks_genres
    ADD CONSTRAINT rel_parent_id_id FOREIGN KEY (parent_id) REFERENCES public.bks_genres (id);

--
-- bks_books
--

CREATE TABLE public.bks_books
(
    id             serial PRIMARY KEY,
    genre_id       integer                 NOT NULL,
    title          character varying(225)  NOT NULL,
    description    character varying(1024) NULL,
    price_amount   integer                 NOT NULL,
    price_currency character varying(3)    NOT NULL,
    created_at     timestamp(0)            NOT NULL,
    updated_at     timestamp(0)            NOT NULL,
    deleted_at     timestamp(0)            NULL
);

ALTER TABLE public.bks_books
    ADD CONSTRAINT rel_genre_id_id FOREIGN KEY (genre_id) REFERENCES public.bks_genres (id);

INSERT INTO bks_books
VALUES (1, 1, 'Domain-Driven Design: Tackling Complexity in the Heart of Software', '', 5650, 'USD',
        '2022-04-01 13:00:00', '2022-04-01 13:00:00', null),
       (2, 1, 'Implementing Domain-Driven Design', '', 3130, 'USD', '2022-04-01 13:00:00', '2022-04-01 13:00:00', null),
       (3, 2, 'Patterns of Enterprise Application Architecture', '', 5217, 'USD', '2022-04-01 13:00:00',
        '2022-04-01 13:00:00', null),
       (4, 2, 'Design Patterns: Elements of Reusable Object-Oriented Software', '', 2990, 'USD', '2022-04-01 13:00:00',
        '2022-04-01 13:00:00', null),
       (5, 3, 'Extreme Programming Explained: Embrace Change', '', 3578, 'USD', '2022-04-01 13:00:00',
        '2022-04-01 13:00:00', null),
       (6, 4,
        'Practical Event-Driven Microservices Architecture: Building Sustainable and Highly Scalable Event-Driven Microservices',
        '', 5155, 'USD', '2022-04-01 13:00:00',
        '2022-04-01 13:00:00', null),
       (7, 3, 'Refactoring: Improving the Design of Existing Code', '', 5155, 'USD', '2022-04-01 13:00:00',
        '2022-04-01 13:00:00', null);

--
-- bks_authors
--

CREATE TABLE public.bks_authors
(
    id         serial PRIMARY KEY,
    name       character varying(225) NOT NULL,
    created_at timestamp(0)           NOT NULL,
    updated_at timestamp(0)           NOT NULL,
    deleted_at timestamp(0)           NULL
);

INSERT INTO bks_authors
VALUES (1, 'Eric Evans', '2022-04-01 13:00:00', '2022-04-01 13:00:00', null),
       (2, 'Vernon Vaughn', '2022-04-01 13:00:00', '2022-04-01 13:00:00', null),
       (3, 'Fowler Martin', '2022-04-01 13:00:00', '2022-04-01 13:00:00', null),
       (4, 'Gamma Erich', '2022-04-01 13:00:00', '2022-04-01 13:00:00', null),
       (5, 'Helm Richard', '2022-04-01 13:00:00', '2022-04-01 13:00:00', null),
       (6, 'Johnson Ralph', '2022-04-01 13:00:00', '2022-04-01 13:00:00', null),
       (7, 'Vlissides John', '2022-04-01 13:00:00', '2022-04-01 13:00:00', null),
       (8, 'Kent Beck', '2022-04-01 13:00:00', '2022-04-01 13:00:00', null),
       (9, 'Cynthia Andres', '2022-04-01 13:00:00', '2022-04-01 13:00:00', null),
       (10, 'Hugo Filipe Oliveira Rocha', '2022-04-01 13:00:00', '2022-04-01 13:00:00', null);

--
-- bks_books_authors
--

CREATE TABLE public.bks_books_authors
(
    id        serial PRIMARY KEY,
    book_id   integer NOT NULL,
    author_id integer NOT NULL
);

ALTER TABLE public.bks_books_authors
    ADD CONSTRAINT rel_book_id_id FOREIGN KEY (book_id) REFERENCES public.bks_books (id);

ALTER TABLE public.bks_books_authors
    ADD CONSTRAINT rel_author_id_id FOREIGN KEY (author_id) REFERENCES public.bks_authors (id);

INSERT INTO bks_books_authors
VALUES (1, 1, 1),
       (2, 2, 2),
       (3, 3, 3),
       (4, 4, 4),
       (5, 4, 5),
       (6, 4, 6),
       (7, 4, 7),
       (8, 5, 8),
       (9, 5, 9),
       (10, 6, 10),
       (11, 7, 8),
       (12, 7, 3);
