
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE sup_src ADD COLUMN `click` INT(11) NOT NULL DEFAULT 0;
ALTER TABLE sup_dem_src ADD COLUMN `click` INT(11) NOT NULL DEFAULT 0;
ALTER TABLE demand_report ADD COLUMN `click` INT(11) NOT NULL DEFAULT 0;
ALTER TABLE exchange_report ADD COLUMN `click` INT(11) NOT NULL DEFAULT 0;
ALTER TABLE supplier_report ADD COLUMN `click` INT(11) NOT NULL DEFAULT 0;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE sup_src DROP COLUMN `click`;
ALTER TABLE sup_dem_src DROP COLUMN `click`;
ALTER TABLE demand_report DROP COLUMN `click`;
ALTER TABLE exchange_report DROP COLUMN `click`;
ALTER TABLE supplier_report DROP COLUMN `click`;

