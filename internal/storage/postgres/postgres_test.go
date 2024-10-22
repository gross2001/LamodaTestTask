package postgres

import (
	"reflect"
	"testTaskLamoda/internal/storage/models"
	"testing"
)

func Test_chooseStore(t *testing.T) {
	sku1_Main_Out := models.SkuStore{Id: 1, Sku: "test sku1", Store_id: 1, Total_quantity: 0, Reserved: 0}
	sku1_Main_Free := models.SkuStore{Id: 1, Sku: "test sku1", Store_id: 1, Total_quantity: 10, Reserved: 9}
	sku1_Sec_Free := models.SkuStore{Id: 1, Sku: "test sku1", Store_id: 2, Total_quantity: 10, Reserved: 9}
	sku1_Main_Free_Status := models.SkuStoreStatus{SkuStoreInfo: sku1_Main_Free, Status: "Ok"}
	sku1_Sec_Free_Status := models.SkuStoreStatus{SkuStoreInfo: sku1_Sec_Free, Status: "Ok"}
	sku2_SecStore_Out := models.SkuStore{Id: 2, Sku: "test sku2", Store_id: 2, Total_quantity: 10, Reserved: 10}

	type args struct {
		skusInfo []models.SkuStore
	}
	tests := []struct {
		name string
		args args
		want map[string]models.SkuStoreStatus
	}{
		{
			name: "All SKUs are not available to reserve",
			args: args{[]models.SkuStore{sku1_Main_Out, sku2_SecStore_Out}},
			want: map[string]models.SkuStoreStatus{},
		},
		{
			name: "SKU are from two store",
			args: args{[]models.SkuStore{sku1_Sec_Free, sku1_Main_Free, sku1_Sec_Free}},
			want: map[string]models.SkuStoreStatus{"test sku1": sku1_Main_Free_Status},
		},
		{
			name: "Sku from main store (out) and second store (free)",
			args: args{[]models.SkuStore{sku1_Main_Out, sku1_Sec_Free}},
			want: map[string]models.SkuStoreStatus{"test sku1": sku1_Sec_Free_Status},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chooseStoreToReserve(tt.args.skusInfo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chooseStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_chooseStoreToDeleteReserve(t *testing.T) {
	sku1_Main_Out := models.SkuStore{Id: 1, Sku: "test sku1", Store_id: 1, Total_quantity: 0, Reserved: 0}
	sku1_Main_Free := models.SkuStore{Id: 1, Sku: "test sku1", Store_id: 1, Total_quantity: 10, Reserved: 1}
	sku1_Sec_Free := models.SkuStore{Id: 1, Sku: "test sku1", Store_id: 2, Total_quantity: 10, Reserved: 9}
	sku1_Main_Free_Status := models.SkuStoreStatus{SkuStoreInfo: sku1_Main_Free, Status: "Ok"}
	sku1_Sec_Free_Status := models.SkuStoreStatus{SkuStoreInfo: sku1_Sec_Free, Status: "Ok"}
	sku2_SecStore_Out := models.SkuStore{Id: 2, Sku: "test sku2", Store_id: 2, Total_quantity: 10, Reserved: 0}

	type args struct {
		skusInfo []models.SkuStore
	}
	tests := []struct {
		name string
		args args
		want map[string]models.SkuStoreStatus
	}{
		{
			name: "All SKUs are not available to delete",
			args: args{[]models.SkuStore{sku1_Main_Out, sku2_SecStore_Out}},
			want: map[string]models.SkuStoreStatus{},
		},
		{
			name: "SKU are from two store",
			args: args{[]models.SkuStore{sku1_Main_Free, sku1_Sec_Free, sku1_Main_Free}},
			want: map[string]models.SkuStoreStatus{"test sku1": sku1_Sec_Free_Status},
		},
		{
			name: "Sku from sec store (out) and main store (free)",
			args: args{[]models.SkuStore{sku1_Main_Free, sku2_SecStore_Out}},
			want: map[string]models.SkuStoreStatus{"test sku1": sku1_Main_Free_Status},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chooseStoreToDeleteReserve(tt.args.skusInfo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chooseStoreToDeleteReserve() = %v, want %v", got, tt.want)
			}
		})
	}
}
