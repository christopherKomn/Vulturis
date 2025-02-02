// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: products.sql

package database

import (
	"context"
	"database/sql"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (name, stock, description)
VALUES (
    $1,
    $2,
    $3
)
RETURNING product_code, name, stock, description
`

type CreateProductParams struct {
	Name        string
	Stock       int32
	Description sql.NullString
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct, arg.Name, arg.Stock, arg.Description)
	var i Product
	err := row.Scan(
		&i.ProductCode,
		&i.Name,
		&i.Stock,
		&i.Description,
	)
	return i, err
}

const deleteProductByCode = `-- name: DeleteProductByCode :exec
delete from products WHERE product_code = $1
`

func (q *Queries) DeleteProductByCode(ctx context.Context, productCode int32) error {
	_, err := q.db.ExecContext(ctx, deleteProductByCode, productCode)
	return err
}

const deleteProducts = `-- name: DeleteProducts :exec
delete  from products
`

func (q *Queries) DeleteProducts(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteProducts)
	return err
}

const getProduct = `-- name: GetProduct :one
SELECT product_code, name, stock, description FROM products WHERE $1 = products.name
`

func (q *Queries) GetProduct(ctx context.Context, name string) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProduct, name)
	var i Product
	err := row.Scan(
		&i.ProductCode,
		&i.Name,
		&i.Stock,
		&i.Description,
	)
	return i, err
}

const getUserByCode = `-- name: GetUserByCode :one
SELECT product_code, name, stock, description FROM products WHERE $1 = products.product_code
`

func (q *Queries) GetUserByCode(ctx context.Context, productCode int32) (Product, error) {
	row := q.db.QueryRowContext(ctx, getUserByCode, productCode)
	var i Product
	err := row.Scan(
		&i.ProductCode,
		&i.Name,
		&i.Stock,
		&i.Description,
	)
	return i, err
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE products
SET name = $1, stock = $2,description = $3
WHERE product_code = $4
RETURNING product_code, name, stock, description
`

type UpdateProductParams struct {
	Name        string
	Stock       int32
	Description sql.NullString
	ProductCode int32
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, updateProduct,
		arg.Name,
		arg.Stock,
		arg.Description,
		arg.ProductCode,
	)
	var i Product
	err := row.Scan(
		&i.ProductCode,
		&i.Name,
		&i.Stock,
		&i.Description,
	)
	return i, err
}

const updateProductDescription = `-- name: UpdateProductDescription :one
UPDATE products SET description = $1
WHERE product_code = $2
RETURNING product_code, name, stock, description
`

type UpdateProductDescriptionParams struct {
	Description sql.NullString
	ProductCode int32
}

func (q *Queries) UpdateProductDescription(ctx context.Context, arg UpdateProductDescriptionParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, updateProductDescription, arg.Description, arg.ProductCode)
	var i Product
	err := row.Scan(
		&i.ProductCode,
		&i.Name,
		&i.Stock,
		&i.Description,
	)
	return i, err
}

const updateProductName = `-- name: UpdateProductName :one
UPDATE products SET name = $1
WHERE product_code = $2
RETURNING product_code, name, stock, description
`

type UpdateProductNameParams struct {
	Name        string
	ProductCode int32
}

func (q *Queries) UpdateProductName(ctx context.Context, arg UpdateProductNameParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, updateProductName, arg.Name, arg.ProductCode)
	var i Product
	err := row.Scan(
		&i.ProductCode,
		&i.Name,
		&i.Stock,
		&i.Description,
	)
	return i, err
}

const updateProductStock = `-- name: UpdateProductStock :one
UPDATE products SET stock = $1
WHERE product_code = $2
RETURNING product_code, name, stock, description
`

type UpdateProductStockParams struct {
	Stock       int32
	ProductCode int32
}

func (q *Queries) UpdateProductStock(ctx context.Context, arg UpdateProductStockParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, updateProductStock, arg.Stock, arg.ProductCode)
	var i Product
	err := row.Scan(
		&i.ProductCode,
		&i.Name,
		&i.Stock,
		&i.Description,
	)
	return i, err
}
