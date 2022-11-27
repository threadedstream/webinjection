-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
    id int primary key,
    name text,
    description text
);
INSERT INTO products (id, name, description) VALUES (1, 'milk', 'milky way');
INSERT INTO products (id, name, description) VALUES (2, 'bread', 'bready way');
INSERT INTO products (id, name, description) VALUES (3, 'butter', 'buttery way');
INSERT INTO products (id, name, description) VALUES (4, 'cheese', 'cheesy way');
INSERT INTO products (id, name, description) VALUES (5, 'peanut', 'peanuty way');
INSERT INTO products (id, name, description) VALUES (6, 'nuts', 'nutty way');
INSERT INTO products (id, name, description) VALUES (7, 'pie', 'piey way');
INSERT INTO products (id, name, description) VALUES (8, 'cake', 'caky way');
INSERT INTO products (id, name, description) VALUES (9, 'coffee', 'coffee way');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
