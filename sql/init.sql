create database if not exists web_URL_Shortener;

use web_URL_Shortener;

create table if not exists URL_conversion(
       id integer AUTO_INCREMENT not null,
       short_url varchar(255),
       long_url text not null,
       expiration bigint,
       create_at timestamp,
       primary key(id)
);