package row_struct

import (
	"database/sql"
	"testing"
)

func TestFieldScan(t *testing.T) {
	type args struct {
		rows  *sql.Rows
		field interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "row_struct",
			args: args{
				rows:  nil,
				field: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Scan(tt.args.rows, tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("FieldScan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
