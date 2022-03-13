
create table task (
id serial primary key,
source_id varchar,
sheet_id varchar,
is_active smallint,
created_timestamp timestamp,
last_updated_timestamp timestamp,
constraint task_uk unique(source_id, sheet_id));

create table task_user(
id serial primary key,
task_id int,
user_id varchar
);

create table user_profile (
user_id varchar primary key,
user_name varchar,
user_type varchar
);