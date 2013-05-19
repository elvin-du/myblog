create database myblog;
create table myblog.users(userid INT PRIMARY KEY AUTO_INCREMENT,name VARCHAR(50) NOT NULL UNIQUE, password VARCHAR(100) NOT NULL);
create table myblog.comments(id INT PRIMARY KEY AUTO_INCREMENT,userid INT NOT NULL,to_userid INT NOT NULL,comments VARCHAR(1000) NOT NULL,create_date date NOT NULL);
create table myblog.blogs(id INT PRIMARY KEY AUTO_INCREMENT,userid INT NOT NULL,blogs TEXT NOT NULL,create_date date NOT NULL,type_id INT NOT NULL,title VARCHAR(100) NOT NULL UNIQUE);
create table myblog.blog_type(id INT PRIMARY KEY AUTO_INCREMENT,blog_type INT NOT NULL);


