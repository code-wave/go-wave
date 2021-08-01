#!bin/bash

set -e

echo '*****************************'
echo "Connecting postgres database engine..."

PGPASSWORD="$POSTGRES_PASSWORD" psql -U $POSTGRES_USER -d $POSTGRES_DB  -W <<-EOSQL
GRANT ALL PRIVILEGES ON DATABASE $POSTGRES_DB TO $POSTGRES_USER;

\c $POSTGRES_DB;

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

create table chat_room (
    id serial NOT NULL,
    room_name varchar(48) UNIQUE NOT NULL,
    client_id bigint NOT NULL,
    host_id bigint NOT NULL,
    study_post_id bigint NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (client_id) REFERENCES users (id),
    FOREIGN KEY (host_id) REFERENCES users (id),
    FOREIGN KEY (study_post_id) REFERENCES study_post (id)
);

create table chat_message (
    chat_room_id bigint NOT NULL,
    chat_room_name varchar(48)  NOT NULL,
    sender_id bigint NOT NULL,
    sender varchar(48) NOT NULL,
    message_type varchar(48) NOT NULL,
    message text NOT NULL,
    created_at timestamp NOT NULL,
    FOREIGN KEY (chat_room_id) REFERENCES chat_room (id),
    FOREIGN KEY (sender_id) REFERENCES users (id)
);

GRANT ALL PRIVILEGES ON TABLE users to $POSTGRES_USER;
GRANT ALL PRIVILEGES ON TABLE token to $POSTGRES_USER;
GRANT ALL PRIVILEGES ON TABLE study_post to $POSTGRES_USER;
GRANT ALL PRIVILEGES ON TABLE study_post_tech_stack to $POSTGRES_USER;
GRANT ALL PRIVILEGES ON TABLE tech_stack to $POSTGRES_USER;
GRANT ALL PRIVILEGES ON TABLE chat_room to $POSTGRES_USER;
GRANT ALL PRIVILEGES ON TABLE chat_message to $POSTGRES_USER;

insert into users (email, password, name, nickname, created_at) values ('test@naver.com', '1234', 'atg', 'nickname', NOW());
insert into users (email, password, name, nickname, created_at) values ('kim@naver.com', '1234', 'fsdf', 'kim', NOW());
insert into users (email, password, name, nickname, created_at) values ('han@naver.com', '1234', 'sdfdsf', 'han', NOW());

insert into tech_stack (tech_name) values ('go');
insert into tech_stack (tech_name) values ('react');
commit;
EOSQL
echo "initializing sql done"
echo '*****************************'