drop table if exists "user" cascade;
create table "user" (
    "id" uuid not null default gen_random_uuid()
    , "name" varchar(100) not null
    , "created_at" timestamptz not null default now()
);
alter table "user" add primary key ("id");
comment on table "user" is 'User';
comment on column "user"."name" is 'User name';

drop table if exists "event" cascade;
create table "event" (
    "id" uuid not null default gen_random_uuid()
    , "name" varchar(100) not null
    , "description" varchar(400) not null
    , "date" date not null
    , "created_at" timestamptz not null default now()
);
alter table "event" add primary key ("id");
comment on table "event" is 'Event information';
comment on column "event"."name" is 'Event name';
comment on column "event"."date" is 'Event date';

drop table if exists "event_user" cascade;
create table "event_user" (
    "event_id" uuid not null
    , "user_id" uuid not null
    , "registered_at" timestamptz not null
    , "created_at" timestamptz not null default now()
);
alter table "event_user" add primary key ("event_id", "user_id");
alter table "event_user" add foreign key ("event_id") references "event" ("id");
alter table "event_user" add foreign key ("user_id") references "user" ("id");
comment on table "event_user" is 'Event participants';
comment on column "event_user"."registered_at" is 'A user registered at';
