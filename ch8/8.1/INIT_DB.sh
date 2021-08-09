docker exec 81_postgres_1 psql -U postgres  --command="CREATE DATABASE chapter_eight;"

docker exec 81_postgres_1 psql -U postgres  chapter_eight --command="DROP TABLE IF EXISTS users;CREATE TABLE users(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    username CHAR(25) NOT NULL UNIQUE,
    password CHAR(256) NOT NULL,
    email CHAR(128) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
INSERT INTO users(username , password , email ) VALUES ('test-user' , '!@#$%@#$' , 'test@google.com');"
