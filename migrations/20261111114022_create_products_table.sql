-- +goose Up
CREATE TABLE products (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    category_id UUID REFERENCES categories(id),
    name        VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    image       TEXT NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    row_column  SMALLINT NOT NULL DEFAULT 0
);

-- +goose Down
DROP TABLE products;
