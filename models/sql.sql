create database myblog;create table myblog.users(userid INT PRIMARY KEY AUTO_INCREMENT,name VARCHAR(50) NOT NULL UNIQUE, password VARCHAR(100) NOT NULL);create table myblog.comments(id INT PRIMARY KEY AUTO_INCREMENT,ip VARCHAR(50) NOT NULL, content VARCHAR(1000) NOT NULL, create_date date NOT NULL, blog_id INT NOT NULL);create table myblog.blogs(id INT PRIMARY KEY AUTO_INCREMENT,content TEXT NOT NULL,title VARCHAR(50) NOT NULL UNIQUE, create_date date NOT NULL,tag_id INT NOT NULL);create table myblog.tags(id INT PRIMARY KEY AUTO_INCREMENT,tag VARCHAR(40) CHARACTER SET UTF8 NOT NULL UNIQUE);


