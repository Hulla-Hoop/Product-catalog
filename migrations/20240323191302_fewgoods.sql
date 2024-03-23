-- +goose Up
-- +goose StatementBegin
INSERT INTO goods(name) VALUES ('Spagetti');
INSERT INTO goods(name) VALUES ('Apple');
INSERT INTO goods(name) VALUES ('Apple Pie');

INSERT INTO categories(name) VALUES ('Flour products');
INSERT INTO categories(name) VALUES ('Bakery products');
INSERT INTO categories(name) VALUES ('Fruits');
INSERT INTO categories(name) VALUES ('Fresh');
INSERT INTO categories(name) VALUES ('For boil');


INSERT INTO relation(goods_id,category_id) VALUES (1,5);
INSERT INTO relation(goods_id,category_id) VALUES (1,1);
INSERT INTO relation(goods_id,category_id) VALUES (2,3);
INSERT INTO relation(goods_id,category_id) VALUES (2,4);
INSERT INTO relation(goods_id,category_id) VALUES (3,2);
INSERT INTO relation(goods_id,category_id) VALUES (3,3);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM relation WHERE goods_id=1;
DELETE FROM relation WHERE goods_id=2;
DELETE FROM relation WHERE goods_id=3;

DELETE FROM categories WHERE id=1;
DELETE FROM categories WHERE id=2;
DELETE FROM categories WHERE id=3;
DELETE FROM categories WHERE id=4;
DELETE FROM categories WHERE id=5;

DELETE FROM goods WHERE id=1;
DELETE FROM goods WHERE id=2;
DELETE FROM goods WHERE id=3;


-- +goose StatementEnd
