-- +goose Up
-- +goose StatementBegin
CREATE SEQUENCE categories_seq
START WITH 1
INCREMENT BY 1;

CREATE TABLE categories (
    id int DEFAULT nextval('categories_seq'::regclass) NOT NULL,
    name text NOT NULL,
    updated_at timestamp with time zone DEFAULT NOW() NOT NULL,
    created_at timestamp with time zone DEFAULT NOW() NOT NULL,
    PRIMARY KEY (id)
);

CREATE SEQUENCE goods_seq
START WITH 1
INCREMENT BY 1;

CREATE TABLE goods (
    id int DEFAULT nextval('goods_seq'::regclass) NOT NULL,
    name text NOT NULL,
    removed bool DEFAULT false,
    updated_at timestamp with time zone DEFAULT NOW() NOT NULL,
    created_at timestamp with time zone DEFAULT NOW() NOT NULL,
    PRIMARY KEY (id),
);

CREATE TABLE relation {
    goods_id int,
    category_id int,
    FOREIGN KEY(goods_id) REFERENCES goods(id)
    FOREIGN KEY(category_id) REFERENCES categories(id)
}

CREATE INDEX goods_id_idx ON goods(id);
CREATE INDEX goods_name_idx ON goods(name);
CREATE INDEX categories_id_idx ON categories(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
