drop table if exists "user" cascade;
create table "user" (
    "id" uuid not null default gen_random_uuid()
    , "name" varchar(200) not null
    , "created_at" timestamp not null default now()
);
alter table "user" add primary key ("id");
comment on table "user" is 'ユーザ';
comment on column "user"."name" is 'ユーザ名';

drop table if exists "event" cascade;
create table "event" (
    "id" uuid not null default gen_random_uuid()
    , "name" varchar(200) not null
    , "date" date not null
    , "created_at" timestamp not null default now()
);
alter table "event" add primary key ("id");
comment on table "event" is 'イベント';
comment on column "event"."name" is 'イベント名';
comment on column "event"."date" is 'イベント開催日';

drop table if exists "event_user" cascade;
create table "event_user" (
    "event_id" uuid not null
    , "user_id" uuid not null
    , "registered_at" timestamp not null default now()
    , "created_at" timestamp not null default now()
);
alter table "event_user" add primary key ("event_id", "user_id");
alter table "event_user" add foreign key ("event_id") references "event" ("id");
alter table "event_user" add foreign key ("user_id") references "user" ("id");
comment on table "event_user" is 'イベント参加者';
comment on column "event_user"."event_id" is 'イベントID';
comment on column "event_user"."user_id" is '参加者ユーザID';
comment on column "event_user"."registered_at" is '参加登録日時';
