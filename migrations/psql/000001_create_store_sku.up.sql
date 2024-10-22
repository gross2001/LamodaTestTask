--  пока не понимаю зачем используется SEQUENCE, увидел в статье Ламоды, может пойму потом
-- нужна ли здесь транзакция

BEGIN;
CREATE SEQUENCE stores_id_seq INCREMENT BY 1 MINVALUE 1 START 1;

CREATE TABLE store (
  id INT NOT NULL,
  name VARCHAR(180) NOT NULL,
  is_available BOOLEAN NOT NULL DEFAULT NULL,
  PRIMARY KEY(id)
);

CREATE TABLE sku (
  sku VARCHAR(180) NOT NULL, --UNIQUE
  name VARCHAR(180) NOT NULL,
  size VARCHAR(180) NOT NULL,
  PRIMARY KEY(sku)
);

CREATE TABLE sku_store (
  id INT NOT NULL,
  sku VARCHAR(180) NOT NULL,
  store_id INT NOT NULL,
  total_quantity INT NOT NULL CHECK (quantity >= 0),
  reserved INT NOT NULL CHECK (reserved >= 0),
  PRIMARY KEY(id),
  UNIQUE (sku, store_id)
);

COMMIT;