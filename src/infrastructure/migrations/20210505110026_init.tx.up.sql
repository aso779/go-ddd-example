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
VALUES (1, null, 'All', '2022-03-21 13:08:54', '2022-03-21 13:08:54', null),
       (2, 1, 'Novels', '2022-03-21 13:08:54', '2022-03-21 13:08:54', null),
       (3, 1, 'Drama', '2022-03-21 13:08:54', '2022-03-21 13:08:54', null),
       (4, 1, 'Comedy', '2022-03-21 13:08:54', '2022-03-21 13:08:54', null)
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
VALUES (1, 1, 'book1', '', 4300, 'USD', '2022-03-21 13:08:54', '2022-03-21 13:08:54', null),
       (2, 1, 'book2', '', 5200, 'USD', '2022-03-21 13:08:54', '2022-03-21 13:08:54', null),
       (3, 2, 'book3', '', 1700, 'EUR', '2022-03-21 13:08:54', '2022-03-21 13:08:54', null)
;