
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

ALTER TABLE suppliers ADD COLUMN "test_mode" INT DEFAULT 1;
ALTER TABLE demands ADD COLUMN "test_mode" INT DEFAULT 1;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE suppliers DROP COLUMN "test_mode";
ALTER TABLE demands DROP COLUMN "test_mode";

