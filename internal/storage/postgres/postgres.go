package postgres

import (
	"context"
	"fmt"
	"testTaskLamoda/internal/consts"
	"testTaskLamoda/internal/storage"
	"testTaskLamoda/internal/storage/models"
	"testTaskLamoda/internal/storage/postgres/queries"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	db *pgx.Conn
}

func New(dbURL string) (*Storage, error) {
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	err = conn.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("unable to ping to database: %v", err)
	}

	return &Storage{db: conn}, nil
}

func (s *Storage) CreateReserve(ctx context.Context, skus []string) (map[string]models.SkuStoreStatus, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	skusInfo, err := getSkuStoreBySku(ctx, tx, skus)
	if err != nil {
		return nil, err
	}
	if len(skusInfo) == 0 {
		return nil, storage.ErrSkuNotFound
	}

	skuToSkuInfo := chooseStoreToReserve(skusInfo)

	for _, skuInfo := range skuToSkuInfo {
		_, err := tx.Exec(ctx, queries.CreateReserveSkuStoreQuery, skuInfo.SkuStoreInfo.Sku, skuInfo.SkuStoreInfo.Store_id)
		if err != nil {
			return nil, fmt.Errorf("unable to exec transaction to update: %w", err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to commit transaction: %w", err)
	}

	skuToSkuInfo = setStatusToNotOkSku(skus, skusInfo, skuToSkuInfo)
	return skuToSkuInfo, nil
}

func (s *Storage) StoreBalance(ctx context.Context, storeId uint) ([]models.SkuStore, error) {
	rows, err := s.db.Query(ctx, queries.GetSkusInStore, storeId)
	if err != nil {
		return nil, fmt.Errorf("unable to create query: %w", err)
	}

	skusInfo := make([]models.SkuStore, 0)
	for rows.Next() {
		var skuStore models.SkuStore
		err := rows.Scan(&skuStore.Id, &skuStore.Sku, &skuStore.Store_id, &skuStore.Total_quantity, &skuStore.Reserved)
		if err != nil {
			return nil, fmt.Errorf("unable to scan from skuStore: %w", err)
		}
		skusInfo = append(skusInfo, skuStore)
	}
	return skusInfo, nil
}

func (s *Storage) DeleteReserve(ctx context.Context, skus []string) (map[string]models.SkuStoreStatus, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	skusInfo, err := getSkuStoreBySku(ctx, tx, skus)
	if err != nil {
		return nil, err
	}

	skuToSkuInfo := chooseStoreToDeleteReserve(skusInfo)
	fmt.Println(skuToSkuInfo)

	for _, skuInfo := range skuToSkuInfo {
		_, err := tx.Exec(ctx, queries.DeleteReserveSkuStoreQuery, skuInfo.SkuStoreInfo.Sku, skuInfo.SkuStoreInfo.Store_id)
		if err != nil {
			return nil, fmt.Errorf("unable to exec transaction: %w", err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to commit transaction: %w", err)
	}
	skuToSkuInfo = setStatusToNotOkSku(skus, skusInfo, skuToSkuInfo)
	return skuToSkuInfo, nil
}

func getSkuStoreBySku(ctx context.Context, tx pgx.Tx, skus []string) ([]models.SkuStore, error) {
	rows, err := tx.Query(ctx, queries.GetSkuStoreQuery, skus)
	if err != nil {
		return nil, fmt.Errorf("unable to create query: %w", err)
	}
	skusInfo := make([]models.SkuStore, 0, len(skus))
	for rows.Next() {
		var skuStore models.SkuStore
		err := rows.Scan(&skuStore.Id, &skuStore.Sku, &skuStore.Store_id, &skuStore.Total_quantity, &skuStore.Reserved)
		if err != nil {
			return []models.SkuStore{}, fmt.Errorf("unable to scan from skuStore: %w", err)
		}
		skusInfo = append(skusInfo, skuStore)
	}
	return skusInfo, nil
}

func chooseStoreToReserve(skusInfo []models.SkuStore) map[string]models.SkuStoreStatus {
	// По одному sku может вернуться несколько записей - по количеству складов, в которых sku имеется
	// Реализуем бизнес-логику по выбору склада. Предположим, что главный склад приоритетен
	skuToSkuInfo := make(map[string]models.SkuStoreStatus, len(skusInfo))
	for _, skuInfo := range skusInfo {
		_, ok := skuToSkuInfo[skuInfo.Sku]
		if skuInfo.Total_quantity <= skuInfo.Reserved {
			continue
		}
		if !ok {
			skuToSkuInfo[skuInfo.Sku] = models.SkuStoreStatus{
				Status:       consts.StatusOk,
				SkuStoreInfo: skuInfo,
			}
			continue
		}
		if ok && (skuInfo.Store_id == consts.Main_store_ID) {
			skuToSkuInfo[skuInfo.Sku] = models.SkuStoreStatus{
				Status:       consts.StatusOk,
				SkuStoreInfo: skuInfo,
			}
		}
	}
	return skuToSkuInfo
}

func chooseStoreToDeleteReserve(skusInfo []models.SkuStore) map[string]models.SkuStoreStatus {
	skuToSkuInfo := make(map[string]models.SkuStoreStatus, len(skusInfo))
	for _, skuInfo := range skusInfo {
		_, ok := skuToSkuInfo[skuInfo.Sku]
		if skuInfo.Reserved == 0 {
			continue
		}
		if !ok {
			skuToSkuInfo[skuInfo.Sku] = models.SkuStoreStatus{
				SkuStoreInfo: skuInfo,
				Status:       consts.StatusOk,
			}
			continue
		}
		if ok && (skuInfo.Store_id != consts.Main_store_ID) {
			skuToSkuInfo[skuInfo.Sku] = models.SkuStoreStatus{
				SkuStoreInfo: skuInfo,
				Status:       consts.StatusOk,
			}
		}
	}
	return skuToSkuInfo
}

func setStatusToNotOkSku(skus []string, skusInfo []models.SkuStore, skuToSkuInfo map[string]models.SkuStoreStatus) map[string]models.SkuStoreStatus {
	// Устанавливаем статус для недоступных к заказу sku
	for _, sku := range skusInfo {
		_, ok := skuToSkuInfo[sku.Sku]
		if !ok {
			skuToSkuInfo[sku.Sku] = models.SkuStoreStatus{
				Status:       consts.StatusNotAvailable,
				SkuStoreInfo: models.SkuStore{},
			}
		}
	}
	// Устанавливаем статус для некорректных sku
	for _, sku := range skus {
		_, ok := skuToSkuInfo[sku]
		if !ok {
			skuToSkuInfo[sku] = models.SkuStoreStatus{
				Status:       consts.StatusSKUError,
				SkuStoreInfo: models.SkuStore{},
			}
		}
	}
	return skuToSkuInfo
}
