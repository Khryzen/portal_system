-- database name: portal_system;
use portal_system;

-- users
create table users (
	  id int not null auto_increment,
    username varchar(30),
    user_password varchar(80),
    first_name varchar(30),
    last_name varchar(30),
    email varchar(50),
    contact_number varchar(20),
    last_login datetime,
    primary key (id)
);

select * from users;