create table Restaurants(
    id serial not null primary key,
    name varchar
);

create table Courses(
    id serial not null primary key,
    restaurant_id integer not null references Restaurants(id) on delete cascade,
    name varchar
);

create table Foods(
    id serial not null primary key,
    course_id integer not null references Courses(id) on delete cascade,
    name varchar,
    weight int,
    price int,
    calories int,
    info varchar,
    prep int,
    comp varchar
);

create table Users(
    id serial not null primary key,
    login varchar,
    pwd varchar,
    user_type int
);

insert into users (login, pwd, user_type) values ('admin', 'admin', 1), ('user', 'user', 0);
