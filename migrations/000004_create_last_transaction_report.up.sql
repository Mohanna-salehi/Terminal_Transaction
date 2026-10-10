CREATE PROCEDURE Last_Transaction()
SELECT
    t.id AS terminal_id,
    t.name AS terminal_name,
    tr.id AS transaction_id,
    tr.amount AS amount,
    tr.date AS date
FROM terminals t
    CROSS JOIN LATERAL (
    SELECT
    tran.id,
    tran.amount,
    tran.date
    FROM transactions tran
    WHERE tran.terminal_id = t.id
    ORDER BY tran.date DESC, tran.id DESC
    LIMIT 1
    )
ORDER BY t.id ASC;



-- CREATE PROCEDURE Last_Transaction()
-- SELECT
--     t.id AS terminal_id,
--     t.name AS terminal_name,
--     tr.id AS transaction_id,
--     tr.amount AS amount,
--     tr.date AS date
-- FROM terminals t
--     JOIN LATERAL (
--     SELECT
--     tran.id,
--     tran.amount,
--     tran.date
--     FROM transactions tran
--     WHERE tran.terminal_id = t.id
--     ORDER BY tran.date DESC, tran.id DESC
--     LIMIT 1
--     ) tr ON TRUE
-- ORDER BY t.id ASC;
