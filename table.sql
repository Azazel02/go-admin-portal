CREATE TABLE Books2(
    ID     INT,
	Isbn   INT,
	Title  VARCHAR(30),
	Author VARCHAR(30) 
);

CREATE TABLE Author(
    Firstname  string(30) ,
	Lastname string(30),
);

SELECT * FROM Books;    

SELECT ID=2 FROM Books ;

INSERT into Books VALUES (1,321,'Murder in Night','Ramadesh');

UPDATE books SET Author='Rama';


DELETE FROM Books WHERE	ID=1;
