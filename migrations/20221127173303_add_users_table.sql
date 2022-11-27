-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id int primary key,
    username varchar(40),
    password varchar(50)
);
INSERT INTO users (id, username, password) VALUES (1, 'admin', 'admindoesnotlikegoodpasswords');
INSERT INTO users (id, username, password) VALUES (2, 'alice', 'alicepasswordisstrong');
INSERT INTO users (id, username, password) VALUES (3, 'bob', 'bobpasswordlovesroses');
INSERT INTO users (id, username, password) VALUES (4, 'john', 'johnpasswordrunsdistance');
INSERT INTO users (id, username, password) VALUES (5, 'eliza', 'elizapasswordlikesreading');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
