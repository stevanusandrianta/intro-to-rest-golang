create table books (id serial, title varchar, author varchar, year varchar);

insert into books (title, author, year) values('Title 1', 'Author 1', '2021');
insert into books (title, author, year) values('Title 2', 'Author 2', '2022');
insert into books (title, author, year) values('Title 3', 'Author 3', '2023');
insert into books (title, author, year) values('Title 4', 'Author 4', '2024');
insert into books (title, author, year) values('Title 5', 'Author 5', '2025');

select * from books;