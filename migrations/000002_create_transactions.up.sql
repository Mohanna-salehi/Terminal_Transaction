CREATE TABLE transactions (
    id int NOT NULL AUTO_INCREMENT,
   terminal_id int NOT NULL,
    amount BIGINT ,
    date DATETIME ,

    PRIMARY KEY (id),

    INDEX idx_transactions_terminal_id (terminal_id),

    CONSTRAINT fk_transactions_terminal
        FOREIGN KEY (terminal_id)
            REFERENCES terminals(id)
            ON UPDATE CASCADE
            ON DELETE RESTRICT
)