CREATE TABLE terminals (
    id int NOT NULL AUTO_INCREMENT,
    code VARCHAR(50) NOT NULL,
    name VARCHAR(100) ,
    PRIMARY KEY (id),
    UNIQUE KEY idx_terminals_code (code)
)