
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE fbuser (
    id serial NOT NULL,
    name varchar(25),
    interest varchar(25),
    likes varchar(25),
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE fbuser;
-- SQL section 'Down' is executed when this migration is rolled back

