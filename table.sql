CREATE TABLE Books(
    ID     INT primary key,
	Isbn   INT,
	Title  VARCHAR(30),
	Author_id INT REFERENCES Author(id)
);

CREATE TABLE Author(
	id INT primary key,
    Firstname  VARCHAR(30) ,
	Lastname VARCHAR(30)
);
--FOREIGN KEY
SELECT books.id,books.isbn,books.title,author.firstname,author.lastname
from Books inner join author on author.id =books.Author_id

SELECT * FROM Books;    
SELECT * FROM Books WHERE id='1';
SELECT isbn FROM Books where id=1 ;


INSERT into Books VALUES (1,99,'Murder in Night');
INSERT into Author VALUES (1,'M','K');
UPDATE books SET Author='Rama';


DELETE FROM Books WHERE	ID=1;
