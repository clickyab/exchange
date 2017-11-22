
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE demands MODIFY COLUMN minute_limit FLOAT NOT NULL ;
ALTER TABLE demands MODIFY COLUMN hour_limit FLOAT NOT NULL ;
ALTER TABLE demands MODIFY COLUMN day_limit FLOAT NOT NULL ;
ALTER TABLE demands MODIFY COLUMN week_limit FLOAT NOT NULL ;
ALTER TABLE demands add COLUMN currencies VARCHAR(255) DEFAULT '[]'  NOT NULL;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE demands MODIFY COLUMN minute_limit INT NOT NULL ;
ALTER TABLE demands MODIFY COLUMN hour_limit INT NOT NULL ;
ALTER TABLE demands MODIFY COLUMN day_limit INT NOT NULL ;
ALTER TABLE demands MODIFY COLUMN week_limit INT NOT NULL ;
ALTER TABLE demands DROP COLUMN currencies ;



