/*用户表*/
create table users
(
	id int primary key auto_increment,
	username varchar(32) not null,
	password varchar(32) not null,
	sex varchar(2) not null check(sex='男' or sex='女'),
	email varchar(64),
	regTime datetime not null,
	userType varchar(12) not null check(userType='系统管理员' or userType='普通用户'),
	picUrl varchar(255)
);
INSERT INTO users(username, password, sex, email, regTime, userType, picUrl)
values
('admin', '123', '男', 'admin@email.com', now(), '系统管理员', 'static/files/golang2.jpg');


/*bbs类型*/
create table bbsKinds
(
	id int primary key auto_increment,
	kindName varchar(32) not null,
	kindintro text,
	picPath varchar(255),
	addTime datetime not null,
	addAuthor varchar(32) not null 
);

/*bbs内容*/
create table bbs
(
	id int primary key auto_increment,
	bbsTitle varchar(32) not null,
	bbsContent text,
	bbsKind int not null,
	addTime datetime not null,
	addAuthor int not null,
	constraint bbs_bbsType_rf foreign key(bbsKind) references bbsKinds(id),
	constraint bbs_bbsAuthor_rf foreign key(addAuthor) references users(id)
);

/*回复*/
create table response
(
	id int primary key auto_increment,
	bbsId int not null,
	content text not null,
	addTime datetime not null,
	addAuthor varchar(32) not null,
	constraint response_bbsId_rf foreign key(bbsId) references bbs(id)
);




