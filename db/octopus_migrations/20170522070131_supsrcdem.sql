
-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE supsrcdem
(
    source VARCHAR(50),
    request INT,
    impression INT,
    `show` INT,
    imp_bid INT,
    show_bid INT,
    win INT,
    CONSTRAINT supsrcdem_suppliers_id_fk FOREIGN KEY (supplier) REFERENCES suppliers (id),
    CONSTRAINT supsrcdem_demands_id_fk FOREIGN KEY (demand) REFERENCES demands (id),
    CONSTRAINT supsrcdem_timetavle_id_fk FOREIGN KEY (time) REFERENCES timetavle (id)
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE supsrc CASCADE ;