
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE suppliers drop COLUMN click_mode;
ALTER TABLE suppliers MODIFY test_mode BOOL NOT NULL DEFAULT FALSE ;


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE suppliers ADD COLUMN click_mode VARCHAR(10) NOT NULL DEFAULT 'none';
ALTER TABLE suppliers MODIFY test_mode INT(11) NOT NULL DEFAULT 0 ;

