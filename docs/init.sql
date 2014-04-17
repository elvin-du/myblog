drop DATABASE if exists `myblog`;
CREATE DATABASE myblog;
USE myblog;
drop table if exists `admins`;
CREATE TABLE admins(id int(10) PRIMARY KEY AUTO_INCREMENT, name VARCHAR(100) NOT NULL, password VARCHAR(100) NOT NULL);
INSERT INTO admins(name, password) values("root", "63a9f0ea7bb98050796b649e85481845");
drop table if exists `blogs`;
CREATE TABLE blogs(id int(10) PRIMARY KEY AUTO_INCREMENT,title VARCHAR(512), content TEXT NOT NULL, created_date datetime NOT NULL, tag_id int(10) NOT NULL);
INSERT TABLE `blogs`(`title`, `content`,`created_date`, `tag_id`) VALUES('程序员的自我修养', 
'本书主要介绍系统软件的运行机制和原理，涉及在Windows和Linux两个系统平台上，一个应用程序在编译、链接和运行时刻所发生的各种事项，包括：代码指令是如何保存的，库文件如何与应用程序代码静态链接，应用程序如何被装载到内存中并开始运行，动态链接如何实现，C/C++运行库的工作原理，以及操作系统提供的系统服务是如何被调用的。每个技术专题都配备了大量图、表和代码实例，力求将复杂的机制以简洁的形式表达出来。本书最后还提供了一个小巧且跨平台的C/C++运行库MiniCRT，综合展示了与运行库相关的各种技术'，
'2014-4-9 23:22:11', 123456);
drop table if exists `comments`;
CREATE TABLE comments(id int(10) PRIMARY KEY AUTO_INCREMENT, ip VARCHAR(10) NOT NULL, comment TEXT NOT NULL, created_date datetime NOT NULL, blog_id int(10) NOT NULL);
drop table if exists `tags`;
CREATE TABLE tags(id int(10) PRIMARY KEY AUTO_INCREMENT, tag VARCHAR(512) NOT NULL);
