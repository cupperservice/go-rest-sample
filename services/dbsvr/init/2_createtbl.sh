#!/bin/bash

psql -U root -d dic << "EOL"
create table word (
    id                  serial,
    content             text not null,
    create_date         timestamp with time zone,
    update_date         timestamp with time zone,
    reference_date      timestamp with time zone,
    primary key (id)
);

create table example (
    id                  serial,
    word_id             int references word (id),
    content             text not null,
    create_date         timestamp with time zone,
    update_date         timestamp with time zone,
    primary key (id)
);

create unique index on word (content);
create index on word (reference_date);
EOL
