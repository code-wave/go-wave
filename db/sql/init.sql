-- CREATE USER project WITH PASSWORD 'password';
-- CREATE DATABASE projectdb;
GRANT ALL PRIVILEGES ON DATABASE projectdb TO project;

\c projectdb;

create table users (
    id serial NOT NULL,
    email varchar(48) NOT NULL UNIQUE,
    password text NOT NULL,
    name varchar(20) NOT NULL,
    nickname varchar(20) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    PRIMARY KEY (id)
);

create table token (
    id serial NOT NULL,
    email varchar(48) NOT NULL UNIQUE,
    refresh_token text UNIQUE,
    expires_at timestamp,
    PRIMARY KEY (id)
);

create table study_post (
	id serial NOT NULL,
	user_id serial NOT NULL,
	title varchar(48) NOT NULL,
	topic varchar(48),
	content text,
	num_of_members integer,
	is_mento boolean,
	price integer,
	start_date varchar(48),
	end_date varchar(48),
	is_online boolean,
	created_at timestamp NOT NULL,
	updated_at timestamp,
	PRIMARY KEY(id)
);

GRANT ALL PRIVILEGES ON TABLE users to project;
GRANT ALL PRIVILEGES ON TABLE token to project;
GRANT ALL PRIVILEGES ON TABLE study_post to project;

alter table study_post add constraint study_post foreign key (user_id) references users(id);
insert into users (email, password, name, nickname, created_at) values ('test@naver.com', '1234', 'atg', 'nickname', NOW());
commit;