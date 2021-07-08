-- create database
CREATE DATABASE blog;
\c blog

-- creates user
CREATE USER blog_srv WITH PASSWORD 'asdfg';
GRANT CONNECT ON DATABASE blog TO blog_srv;

-- create schema for user tables
CREATE SCHEMA auth;

-- create user tables
CREATE TABLE IF NOT EXISTS auth.users (
    id         BIGSERIAL NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name  VARCHAR(50) NOT NULL,
    email      VARCHAR(50) UNIQUE NOT NULL,
    pwd        VARCHAR(72) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT 'now()',
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS auth.roles (
    id        BIGSERIAL NOT NULL,
    role_name VARCHAR(15) UNIQUE,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS auth.users_roles (
    user_id BIGINT NOT NULL,
    role_id BIGINT NOT NULL,
    PRIMARY KEY (user_id, role_id),
    CONSTRAINT user_roles_fk_user_id FOREIGN KEY (user_id) REFERENCES auth.users(id),
    CONSTRAINT user_roles_fk_role_id FOREIGN KEY (role_id) REFERENCES auth.roles(id)
);

-- create schema for blog tables
CREATE SCHEMA blog;

-- create blog tables
CREATE TABLE IF NOT EXISTS blog.articles (
    id         BIGSERIAL NOT NULL,
    author_id  BIGINT NOT NULL,
    caption    VARCHAR(150) UNIQUE NOT NULL,
    content    TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT 'now()',
    edited_at  TIMESTAMP NOT NULL DEFAULT 'now()',
    deleted_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (id),
    CONSTRAINT articles_fk_author_id FOREIGN KEY (author_id) REFERENCES auth.users(id)
);

CREATE TABLE IF NOT EXISTS blog.tags (
    id       BIGSERIAL NOT NULL,
    tag_name VARCHAR(30) UNIQUE,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS blog.articles_tags (
    article_id BIGINT NOT NULL,
    tag_id     BIGINT NOT NULL,
    PRIMARY KEY (article_id, tag_id),
    CONSTRAINT articles_tags_fk_article_id FOREIGN KEY (article_id) REFERENCES blog.articles(id),
    CONSTRAINT articles_tags_fk_tag_id     FOREIGN KEY (tag_id)     REFERENCES blog.tags(id)
);

-- grant privilegies to the user
GRANT usage ON SCHEMA auth TO blog_srv;
GRANT select, insert, update, delete, trigger ON all tables IN SCHEMA auth TO blog_srv;
GRANT usage, select ON all sequences IN SCHEMA auth TO blog_srv;
GRANT execute ON all functions IN SCHEMA auth TO blog_srv;

GRANT usage ON SCHEMA blog TO blog_srv;
GRANT select, insert, update, delete, trigger ON all tables IN SCHEMA blog TO blog_srv;
GRANT usage, select ON all sequences IN SCHEMA blog TO blog_srv;
GRANT execute ON all functions IN SCHEMA blog TO blog_srv;
