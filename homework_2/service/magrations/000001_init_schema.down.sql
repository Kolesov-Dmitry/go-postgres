-- drop blog tables
DROP TABLE IF EXISTS blog.articles_tags;
DROP TABLE IF EXISTS blog.tags;
DROP TABLE IF EXISTS blog.articles;

-- drop schema for blog tables
DROP SCHEMA IF EXISTS blog;

-- drop user tables
DROP TABLE IF EXISTS auth.users_roles;
DROP TABLE IF EXISTS auth.roles;
DROP TABLE IF EXISTS auth.users;

-- drop schema for user tables
DROP SCHEMA IF EXISTS auth;

-- drop user
DROP USER IF EXISTS blog_srv;