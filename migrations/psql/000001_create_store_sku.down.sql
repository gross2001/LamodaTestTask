BEGIN;
DROP SEQUENCE public.stores_id_seq;

DROP TABLE store;
DROP TABLE sku;
DROP TABLE sku_store;

COMMIT;