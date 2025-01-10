
-- name: CreateCart :one
INSERT INTO cart (User_UUID, product_code, amount)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: UpdateCart :one
UPDATE cart
SET User_UUID = $1, product_code = $2,amount = $3
WHERE ID = $4
RETURNING *;

-- name: UpdateCartUser :one
UPDATE cart SET User_UUID = $1
WHERE ID = $2
RETURNING *;

-- name: UpdateCartProduct :one
UPDATE cart SET product_code = $1
WHERE ID = $2
RETURNING *;

-- name: UpdateCartAmount :one
UPDATE cart SET amount = $1
WHERE ID = $2
RETURNING *;

-- name: DeleteUserCart :exec
delete from cart WHERE User_UUID = $1;

-- name: DeleteCartProduct :exec
delete from cart WHERE product_code = $1 and User_UUID = $2;

-- name: GetUserCart :one
SELECT * FROM cart WHERE $1 = cart.User_UUID;

-- name: GetCartProduct :one
SELECT * FROM cart WHERE $1 = cart.User_UUID and $2 = cart.product_code;

-- name: DeleteAllCarts :exec
delete  from cart;
