package repository

import (
	"testing"
)

func TestInMemoryRepository_Store(t *testing.T) {
	type fields struct {
		data map[string]string
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &InMemoryRepository{
				data: tt.fields.data,
			}
			if err := r.Store(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInMemoryRepository_Get(t *testing.T) {
	type fields struct {
		data map[string]string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				data: map[string]string{
					"key": "10",
				},
			},
			args: args{
				key: "key",
			},
			want:    "111",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &InMemoryRepository{
				data: tt.fields.data,
			}
			got, err := r.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("InMemoryRepository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InMemoryRepository.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
