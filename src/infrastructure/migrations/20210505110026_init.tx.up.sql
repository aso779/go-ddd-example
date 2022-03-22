--
-- bks_users
--

CREATE TABLE public.bks_users
(
    id         serial PRIMARY KEY,
    email      character varying(256) UNIQUE NOT NULL,
    password   character varying(256)        NOT NULL,
    created_at timestamp(0)                  NOT NULL,
    updated_at timestamp(0)                  NOT NULL,
    deleted_at timestamp(0)                  NULL
);

--
-- bks_genres
--

CREATE TABLE public.bks_genres
(
    id         serial PRIMARY KEY,
    parent_id  integer                       NOT NULL,
    title      character varying(512) UNIQUE NOT NULL,
    created_at timestamp(0)                  NOT NULL,
    updated_at timestamp(0)                  NOT NULL,
    deleted_at timestamp(0)                  NULL
);

--
-- bks_books
--

CREATE TABLE public.bks_books
(
    id         serial PRIMARY KEY,
    title      character varying(225) NOT NULL,
    created_at timestamp(0)           NOT NULL,
    updated_at timestamp(0)           NOT NULL,
    deleted_at timestamp(0)           NULL
);

INSERT INTO bks_books
VALUES (1, 'book1', '2022-03-21 13:08:54', '2022-03-21 13:08:54', null);
INSERT INTO bks_books
VALUES (2, 'book2', '2022-03-21 13:08:54', '2022-03-21 13:08:54', null);
INSERT INTO bks_books
VALUES (3, 'book3', '2022-03-21 13:08:54', '2022-03-21 13:08:54', null);