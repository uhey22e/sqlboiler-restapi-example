truncate table "user" cascade;
insert into "user" ("id", "name") values
('00000001-0001-0000-0000-000000000000', 'ユーザ01'),
('00000001-0002-0000-0000-000000000000', 'ユーザ02'),
('00000001-0003-0000-0000-000000000000', 'ユーザ03'),
('00000002-0001-0000-0000-000000000000', 'ユーザ04'),
('00000002-0002-0000-0000-000000000000', 'ユーザ05'),
('00000002-0003-0000-0000-000000000000', 'ユーザ06')
;

truncate table "event" cascade;
insert into "event" ("id", "name", "date") values
('00000000-0000-0000-0000-000000000001', 'イベント01', '2015-10-21'),
('00000000-0000-0000-0000-000000000002', 'イベント02', '2015-10-22'),
('00000000-0000-0000-0000-000000000003', 'イベント03', '2015-11-01'),
('00000000-0000-0000-0000-000000000004', 'イベント04', '2015-11-02'),
('00000000-0000-0000-0000-000000000005', 'イベント05', '2015-12-01')
;

truncate table "event_user" cascade;
insert into "event_user" ("event_id", "user_id", "registered_at") values
('00000000-0000-0000-0000-000000000001', '00000001-0002-0000-0000-000000000000', '2015-10-14 00:00:00'),
('00000000-0000-0000-0000-000000000001', '00000001-0003-0000-0000-000000000000', '2015-10-14 01:00:00')
;
