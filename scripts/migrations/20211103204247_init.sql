-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user"
(
    id         bigserial primary key,
    username   text not null unique,
    first_name text not null,
    last_name  text not null,
    email      text not null unique,
    phone      text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "user"
-- +goose StatementEnd
