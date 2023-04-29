package postgres

import (
	"bootcamp/article/models"
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
)

func Test_authorRepo_GetMinPublisher(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	tests := []struct {
		name     string
		fields   fields
		wantResp []models.MostPublisher
		wantErr  bool
	}{
		// TODO: Add test cases.
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := authorRepo{
				db: tt.fields.db,
			}
			gotResp, err := r.GetMinPublisher()
			if (err != nil) != tt.wantErr {
				t.Errorf("authorRepo.GetMinPublisher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("authorRepo.GetMinPublisher() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
