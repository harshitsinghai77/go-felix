create table IF NOT EXISTS shortUrl
(
    original_url varchar not null,
    short_url varchar not null,
    created_at time not null,
    expires_at time not null,
    has_expired bool default false
);

create unique index shortUrl_original_index
    on shortUrl (original_url);

create unique index shortUrl_short_index
    on shortUrl (short_url);