
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

ALTER TABLE suppliers ADD COLUMN click_mode VARCHAR(10) NOT NULL DEFAULT 'none';

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE suppliers DROP COLUMN click_mode;
