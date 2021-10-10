package array

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestArray_Insert(t *testing.T) {
	type fields struct {
		data []int
	}
	type args struct {
		index  int
		values []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test",
			fields: fields{data: []int{1, 2, 3}},
			args: args{
				index:  1,
				values: []int{4, 5, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array{
				data: tt.fields.data,
			}
			a.Insert(tt.args.index, tt.args.values...)
			assert.Equal(t, a.Size(), 6)
		})
	}
}

func TestArray_Delete(t *testing.T) {
	type fields struct {
		data []int
	}
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "test",
			fields:  fields{data: []int{0,1,2,3,4,5,6,7,8,9,0}},
			args:    args{
				start: 4,
				end:   9,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array{
				data: tt.fields.data,
			}
			if err := a.Delete(tt.args.start, tt.args.end); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func BenchmarkArray_Delete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := Array{data: []int{1,2,3,4,5,6,7,8,9,0}}
		_ = a.Delete(2,7)
	}
}

func TestArray_Remove(t *testing.T) {
	type fields struct {
		data []int
	}
	type args struct {
		item int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "test", fields: fields{ data: []int{1,2,3,9,9,9,4,9,5,6,7,8,9}},
			args: args{ item: 9} },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Array{
				data: tt.fields.data,
			}
			a.Remove(tt.args.item)
		})
	}
}