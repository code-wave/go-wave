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
	user_id bigint NOT NULL,
	title varchar(48) NOT NULL,
	topic varchar(48),
	content text,
	num_of_members integer,
	is_mentor boolean,
	price integer,
	start_date varchar(48),
	end_date varchar(48),
	is_online boolean,
	tech_stack text[],
	created_at timestamp NOT NULL,
	updated_at timestamp,
	PRIMARY KEY(id),
	FOREIGN KEY (user_id) REFERENCES users (id)
);

create table tech_stack (
    id serial NOT NULL,
    tech_name varchar(48) UNIQUE NOT NULL,
    PRIMARY KEY(id)
);

create table study_post_tech_stack (
    study_post_id bigint NOT NULL,
    tech_stack_id bigint NOT NULL,
    FOREIGN KEY (study_post_id) REFERENCES study_post (id) ON DELETE CASCADE,
    FOREIGN KEY (tech_stack_id) REFERENCES tech_stack (id)
);

GRANT ALL PRIVILEGES ON TABLE users to project;
GRANT ALL PRIVILEGES ON TABLE token to project;
GRANT ALL PRIVILEGES ON TABLE study_post to project;
GRANT ALL PRIVILEGES ON TABLE study_post_tech_stack to project;
GRANT ALL PRIVILEGES ON TABLE tech_stack to project;

insert into users (email, password, name, nickname, created_at) values ('test@naver.com', '1234', 'atg', 'nickname', NOW());
commit;