
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

ALTER TABLE `demands` CHANGE COLUMN white_countrie white_countries TEXT;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE `demands` CHANGE COLUMN white_countries white_countrie TEXT;

