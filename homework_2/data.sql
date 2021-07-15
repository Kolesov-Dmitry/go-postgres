-- select blog database
\c blog

-- insert test users
INSERT INTO auth.users (id, first_name, last_name, email, pwd) VALUES
    (1, 'author', 'user', 'author@email.org',  '123'),
    (2, 'admin',  'user', 'admin@email.org', '456');

-- insert test roles
INSERT INTO auth.roles (id, role_name) VALUES
    (1, 'admin'),
    (2, 'writer');

-- assign user roles
INSERT INTO auth.users_roles (user_id, role_id) VALUES
  (1, 2),
  (2, 1);

-- insert articles
INSERT INTO blog.articles (id, author_id, caption, content) VALUES 
  (1, 1, 'Handle errors in GO', 'Do not discard errors using _ variables. If a function returns an error, examine to make sure whether the function succeeded or not. Better, Handle the error and return it; otherwise, it will rise as a panic error when any unusual situation occurs. Dont use panic errors. Dont use panic for normal error handling. In that case, you can handle the error and multiple return values.'),
  (2, 1, 'Comments in GO', 'Comments documenting declarations should be full sentences, even if that seems a little unnecessary. That way gives them format properly when extracted into godoc documentation. Comments should start with the name of the object described and end in a period. Request represents a call to run a command.');

-- insert tags
INSERT INTO blog.tags (id, tag_name) VALUES
  (1, 'Golang'),
  (2, 'Hot');

-- assign tags to articles
INSERT INTO blog.articles_tags (article_id, tag_id) VALUES
  (1, 1),
  (1, 2),
  (2, 1);
