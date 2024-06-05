
DROP TABLE IF EXISTS BooksList;
CREATE TABLE BooksList (  id INT AUTO_INCREMENT NOT NULL, name VARCHAR(128) NOT NULL, PRIMARY KEY (`id`) );

INSERT INTO BooksList  (name) VALUES  ('Blue Train'),  ('Giant Steps'),  ('Jeru'),  ('Sarah Vaughan');

-- drop user root@localhost;
-- flush privileges;
-- create user root@localhost identified by 'admins_password'


-- CREATE USER 'root'@'localhost' IDENTIFIED BY 'Nta21012001';
CREATE USER 'root'@'%' IDENTIFIED BY 'Nta21012001';
-- GRANT ALL ON *.* TO 'root'@'localhost';
GRANT ALL ON *.* TO 'root'@'%';
FLUSH PRIVILEGES;