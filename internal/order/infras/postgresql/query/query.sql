-- name: GetDeleteOrderList :many
select o.order_id,o.user_id,o.order_date,o.amount,o.order_state from "orders".orders o;

-- name: UpdateOrder :exec
UPDATE "orders".orders
SET
    order_state = $2
WHERE order_id = $1;

-- name: GetOrderDetails :many
select
    o.order_id,
    o.product_id,
    o.quantity,
    o.price
from "orders".line_items o
where order_id = $1;
