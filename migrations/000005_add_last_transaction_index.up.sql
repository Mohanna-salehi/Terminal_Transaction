CREATE INDEX idx_transactions_terminal_date_id
    ON transactions (terminal_id, date DESC, id DESC);

