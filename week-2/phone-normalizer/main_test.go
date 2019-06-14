package main

import (
	"reflect"
	"testing"
)

func Test_phoneNormalizer(t *testing.T) {
	type args struct {
		phones []string
	}
	tests := []struct {
		name  string
		args  args
		wantM map[string]int
	}{
		{
			name: "test normalize phones",
			args: args{
				phones: []string{
					"1234567890",
					"123 456 7891",
					"(123) 456 7892",
					"(123) 456-7893",
					"123-456-7894",
					"123-456-7890",
					"1234567892",
					"(123)456-7892",
				},
			},
			wantM: map[string]int{
				"1234567890": 2,
				"1234567891": 1,
				"1234567892": 3,
				"1234567893": 1,
				"1234567894": 1,
			},
		},
		{
			name: "test normalize phones",
			args: args{
				phones: []string{
					"083123456",
					"084 1234509",
					"(084) 1234509",
					"(084)-123-4509",
				},
			},
			wantM: map[string]int{
				"083123456":  1,
				"0841234509": 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotM := phoneNormalizer(tt.args.phones); !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("phoneNormalizer() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}
