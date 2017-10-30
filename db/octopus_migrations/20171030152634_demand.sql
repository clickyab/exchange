
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE demands drop win_url;
ALTER TABLE demands MODIFY test_mode BOOL DEFAULT FALSE ;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE demands add win_url VARCHAR(300) NULL DEFAULT NULL;
ALTER TABLE demands MODIFY test_mode INT DEFAULT 0 ;


