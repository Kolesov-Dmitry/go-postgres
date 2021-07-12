-- get user by email
SELECT id, first_name, last_name, email, pwd, created_at FROM auth.users WHERE email = 'some@email.org';
    
-- get user role by user id
select r.role_name from auth.roles r join auth.users_roles ur on ur.role_id = r.id where ur.user_id = 1

-- get articles sorted by created date with pagination (10 articles on a page)
SELECT 
    a.id, a.caption, a.content, a.edited_at, u.first_name, u.last_name 
FROM 
    blog.articles a
JOIN  
    auth.users u
ON 
    a.author_id = u.id    
WHERE 
    a.deleted_at is NULL AND a.edited_at < to_timestamp('13.07.2021 00:00', 'DD.MM.YYYY H24:MI')
ORDER BY 
    a.edited_at DESC 
LIMIT 10;

-- get article tags by article id
select t.tag_name from blog.tags t join blog.articles_tags art on art.tag_id = t.id where art.article_id = 1
