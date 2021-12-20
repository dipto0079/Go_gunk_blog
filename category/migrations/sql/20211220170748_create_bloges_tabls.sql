-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS blogs(
    id SERIAL NOT NULL,
    cat_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    image TEXT NOT NULL,

    PRIMARY KEY(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS blogs;
-- +goose StatementEnd