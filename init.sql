create table beer
(
    id       int     not null
            constraint beer_pk
            primary key,
    name     varchar not null,
    brewery  varchar,
    country  varchar,
    price    float4,
    currency varchar not null
);

create unique index beer_id_uindex
    on beer (id);