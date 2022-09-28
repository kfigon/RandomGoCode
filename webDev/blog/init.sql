CREATE TABLE if not exists BLOG(
   ID              SERIAL PRIMARY KEY,
   TITLE          VARCHAR(150),
   TEXT          TEXT    NOT NULL
);

insert into BLOG(title, text) VALUES('Hello world', 'bla bla bla lots of text');
insert into BLOG(title, text) VALUES('Second article', 'another article');