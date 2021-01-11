package codec

import (
	"reflect"
	"testing"
)

func TestDecoder(t *testing.T) {
	type hello struct {
		Msg string
	}
	data, _ := Encoder(&hello{"hello"})

	type args struct {
		b []byte
		s interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Decoder Test",
			args: args{
				b: data,
				s: &hello{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Decoder(tt.args.b, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("Decoder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "Encoder Test",
			args:    args{s: "hello"},
			want:    []byte{8, 12, 0, 5, 104, 101, 108, 108, 111},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encoder(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encoder() got = %v, want %v", got, tt.want)
			}
		})
	}
}
