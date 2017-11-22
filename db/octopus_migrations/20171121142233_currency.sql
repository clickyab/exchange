
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE suppliers MODIFY COLUMN floor_cpm FLOAT NOT NULL ;
ALTER TABLE suppliers MODIFY COLUMN soft_floor_cpm FLOAT NOT NULL ;
ALTER TABLE supplier_report MODIFY COLUMN earn FLOAT NOT NULL ;
ALTER TABLE sup_src MODIFY COLUMN profit FLOAT NOT NULL ;
ALTER TABLE sup_src MODIFY COLUMN deliver_bid FLOAT NOT NULL ;
ALTER TABLE sup_dem_src MODIFY COLUMN profit FLOAT NOT NULL ;
ALTER TABLE sup_dem_src MODIFY COLUMN deliver_bid FLOAT NOT NULL ;
ALTER TABLE exchange_report MODIFY COLUMN earn FLOAT NOT NULL ;
ALTER TABLE exchange_report MODIFY COLUMN spent FLOAT NOT NULL ;
ALTER TABLE exchange_report MODIFY COLUMN income FLOAT NOT NULL ;
ALTER TABLE demand_report MODIFY COLUMN profit FLOAT NOT NULL ;
ALTER TABLE demand_report MODIFY COLUMN deliver_bid FLOAT NOT NULL ;


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE suppliers MODIFY COLUMN floor_cpm INT NOT NULL ;
ALTER TABLE suppliers MODIFY COLUMN soft_floor_cpm INT NOT NULL ;
ALTER TABLE supplier_report MODIFY COLUMN earn INT NOT NULL ;
ALTER TABLE sup_src MODIFY COLUMN profit INT NOT NULL ;
ALTER TABLE sup_src MODIFY COLUMN deliver_bid INT NOT NULL ;
ALTER TABLE sup_dem_src MODIFY COLUMN profit INT NOT NULL ;
ALTER TABLE sup_dem_src MODIFY COLUMN deliver_bid INT NOT NULL ;
ALTER TABLE exchange_report MODIFY COLUMN earn INT NOT NULL ;
ALTER TABLE exchange_report MODIFY COLUMN spent INT NOT NULL ;
ALTER TABLE exchange_report MODIFY COLUMN income INT NOT NULL ;
ALTER TABLE demand_report MODIFY COLUMN profit INT NOT NULL ;
ALTER TABLE demand_report MODIFY COLUMN deliver_bid INT NOT NULL ;

