
-- name: CreateProduct :one
INSERT INTO products (name, stock, price, description)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: UpdateProduct :one
UPDATE products
SET name = $1, stock = $2, price = $3, description = $4
WHERE product_code = $5
RETURNING *;

-- name: UpdateProductStock :one
UPDATE products SET stock = $1
WHERE product_code = $2
RETURNING *;

-- name: UpdateProductDescription :one
UPDATE products SET description = $1
WHERE product_code = $2
RETURNING *;

-- name: UpdateProductName :one
UPDATE products SET name = $1
WHERE product_code = $2
RETURNING *;

-- name: UpdateProductPrice :one
UPDATE products SET price = $1
WHERE product_code = $2
RETURNING *;

-- name: DeleteProductByCode :exec
delete from products WHERE product_code = $1;

-- name: GetProduct :one
SELECT * FROM products WHERE $1 = products.name;

-- name: GetUserByCode :one
SELECT * FROM products WHERE $1 = products.product_code;

-- name: CommitOrder :exec

-- $1 user_uuid
-- $2 order_code
--
START TRANSACTION;
    DO $$

    DECLARE
        insufficient_stock BOOLEAN := FALSE;

    -- step 1: test if the amout (of products) cart are within the stock (of products) --
    -- IF TRUE FOR EVERY PRODUCT then insert them in the ordered_products              --
    -- IF NOT raice flag insufficient_stock                                            --
    BEGIN FOR rec IN
        SELECT c.product_code, c.amount, p.stock
        FROM cart c
        JOIN products p ON c.product_code = p.product_code
        WHERE c.user_uuid = $1
    LOOP
        IF rec.amount <= rec.stock THEN

            INSERT INTO ordered_products (order_code, user_uuid, product_code, amount)
            SELECT
                ord.order_code,
                c.product_id,
                c.amount,
                p.price
            FROM cart c, orders ord
            JOIN products p ON c.product_code = p.product_code
            WHERE c.user_uuid = $1 and ord.order_code = $2;

        ELSE
            insufficient_stock := TRUE;
        END IF;
    END LOOP;

    -- step 2: Deduct the ordered amount from the product stock in the products table (since they were bought) --
    UPDATE products
    SET stock = stock - c.amount
    FROM cart c
    WHERE products.product_code = c.product_code
    AND c.user_uuid = $1;

    -- step 3: Remove products from the cart --
    DELETE  FROM cart
    WHERE cart_id = $1;

    -- step 4: if flag insufficient_stock raised throw exception and ROLLBACK the changes --
    IF insufficient_stock THEN
        RAISE EXCEPTION 'Transaction aborted: Insufficient stock for one or more products in the cart.';
        ROLLBACK;
    END IF;

    END $$;

    -- if everything went well the changes will be committed if not then everything will be reverted to the orginal state
COMMIT;
    -- TODO: write some code to return something to inform the client

-- name: DeleteProducts :exec
delete  from products;
