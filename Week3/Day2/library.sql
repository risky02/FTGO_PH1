CREATE DATABASE library;

CREATE TABLE IF NOT EXISTS authors (
    author_id INT AUTO_INCREMENT NOT NULL,
    author_name VARCHAR(50) UNIQUE,
    country VARCHAR(30),
    PRIMARY KEY (author_id)
);

CREATE TABLE IF NOT EXISTS books (
    book_id INT AUTO_INCREMENT NOT NULL,
    title VARCHAR(255) NOT NULL UNIQUE, 
    author_id INT NOT NULL,
    publication_year YEAR(4),
    available_quantity INT(3) NOT NULL,
    PRIMARY KEY (book_id),
    FOREIGN KEY (author_id) REFERENCES authors (author_id)
);

CREATE TABLE IF NOT EXISTS book_loans (
    loan_id INT AUTO_INCREMENT NOT NULL,
    book_id INT NOT NULL,
    borrower_name VARCHAR(100) NOT NULL,
    loan_date DATE NOT NULL, 
    return_date DATE NOT NULL,
    PRIMARY KEY (loan_id),
    FOREIGN KEY (book_id) REFERENCES books (book_id)
);

--  membuat index di table books field title dan author_id
CREATE INDEX idx_title_author_id ON books (title, author_id);

--  membuat index di table book_loans field loan_date
CREATE INDEX idx_loan_date ON book_loans (loan_date);

--  membuat index di table book_loans field borrower_name dan return_date
CREATE INDEX idx_borrower_name_return_date ON book_loans (borrower_name, return_date);


------------------------ Input Data ------------------------------------------------------
INSERT INTO `authors` (`author_name`, `country`)
VALUES
	('Abdul Malik Karim Amrullah', 'Indonesia'),
    ('Andrea Hirata', 'Indonesia'),
    ('Arthur Conan Doyle', 'England'),
    ('Eka Kurniawan', 'Indonesia'),
    ('Habiburrahman El Shirazy', 'Indonesia'),
	('Raditya Dika', 'Indonesia');

INSERT INTO `books` (`title`, `author_id`, `publication_year`,  `available_quantity`)
VALUES
	('A Study in Scarlet', '3', '1901', 30),
    ('Sang Pemimpi', '2', '2006', 30),
    ('Ayah', '2', '2015', 30),
    ('Tenggelamnya Kapal Van der Wijck', '1', '1938', 20),
    ('Ketika Cinta Bertasbih', '5', '2007', 30),
    ('Cinta Brontosaurus', '6', '2006', 40),
    ('Cantik itu Luka', '4', '2002', 30);

INSERT INTO `book_loans` (`book_id`,	`borrower_name`, `loan_date`, `return_date`)
VALUES
	(5, 'Mandzukic', '20231001', '20231008'),
    (3, 'Didier Drogba', '20231001','20231008'),
    (3, 'Samir Nasri', '20231002','20231009'),
    (2, 'Mesut Ozil', '20231002','20231009'),
    (1, 'Michael Ballack', '20231002','20231009'),
    (4, 'Andriy Shevchenko', '20231003','20231010');

