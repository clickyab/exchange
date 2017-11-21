
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE suppliers ADD currency VARCHAR(5) DEFAULT 'IRR' NOT NULL;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE suppliers DROP currency;

