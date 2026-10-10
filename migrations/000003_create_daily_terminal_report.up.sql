CREATE PROCEDURE daily_terminal()
SELECT
    t.id ,
    t.code ,
    t.name AS terminal_name,
    COUNT(tr.id) AS transaction_count,
    SUM(tr.amount) AS total_amount,
    TIME(MAX(tr.date)) AS today_transaction

FROM terminals t
    INNER JOIN transactions tr
ON tr.terminal_id = t.id
    AND tr.date >= CURDATE()
    AND tr.date < DATE_ADD(CURDATE(), INTERVAL 1 DAY)

GROUP BY
    t.id,
    t.code,
    t.name

ORDER BY
    t.id ASC;