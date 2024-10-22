package reserve

import (
	"context"
	"reflect"
	"testTaskLamoda/internal/api/responses"
	"testTaskLamoda/internal/storage/mocks"
	"testing"
)

func TestReserve_Create(t *testing.T) {
	type args struct {
		ctx  context.Context
		skus []string
	}
	tests := []struct {
		name    string
		args    args
		want    responses.ReserveCreateResult
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "first case",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := mocks.NewStorage(t)
			r := &Reserve{
				storage: store,
			}
			store.
				On("CreateReserve", tt.args.ctx, tt.args.skus).
				Return(nil, nil)
			got, err := r.Create(tt.args.ctx, tt.args.skus)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reserve.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reserve.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
