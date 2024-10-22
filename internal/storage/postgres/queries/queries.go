package queries

const (
	GetSkuStoreQuery = `
	SELECT 
		ss.id, ss.sku, ss.store_id, ss.quantity, ss.reserved
	FROM sku_store ss
	INNER JOIN store s on ss.store_id = s.id 
	WHERE
		ss.sku = any($1) AND
		s.is_available = true
	FOR update;
	`

	CreateReserveSkuStoreQuery = `
	UPDATE
		sku_store
	SET reserved=reserved+1
	WHERE
		sku=$1 AND store_id = $2
	`

	DeleteReserveSkuStoreQuery = `
	UPDATE sku_store
	SET reserved=reserved-1
	WHERE
		sku=$1 AND store_id = $2
	`

	GetSkusInStore = `
	SELECT
		ss.id, ss.sku, ss.store_id, ss.quantity, ss.reserved
	FROM sku_store ss 
	WHERE
		ss.store_id = $1;
	`
)
