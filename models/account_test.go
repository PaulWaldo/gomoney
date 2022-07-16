package models

import (
	"reflect"
	"testing"
)

// func TestCreateAccountType(t *testing.T) {
// 	expectedName := "abc123"
// 	c := NewAccount(expectedName, Checking)
// 	if c.Name != expectedName {
// 		t.Errorf("Expecting name to be %s, got %s", expectedName, c.Name)
// 	}
// }

func TestAccountTypeFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    AccountType
		wantErr bool
	}{
		{
			name:    "checking",
			args:    args{"checking"},
			want:    Checking,
			wantErr: false,
		},
		{
			name:    "savings",
			args:    args{"savings"},
			want:    Savings,
			wantErr: false,
		},
		{
			name:    "creditCard",
			args:    args{"creditCard"},
			want:    CreditCard,
			wantErr: false,
		},
		{
			name:    "unknown account type",
			args:    args{"badvalue"},
			want:    Unknown,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AccountTypeFromString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountTypeFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountTypeFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestAccountType_String(t *testing.T) {
// 	type fields struct {
// 		slug string
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   string
// 	}{
// 		{
// 			name:   "checking slug",
// 			fields: fields{AccountTypeChecking},
// 			want:   "checking",
// 		},
// 		{
// 			name:   "creditCard slug",
// 			fields: fields{AccountTypeCreditCard},
// 			want:   "creditCard",
// 		},
// 		{
// 			name:   "checking slug",
// 			fields: fields{AccountTypeChecking},
// 			want:   "checking",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := AccountType{
// 				slug: tt.fields.slug,
// 			}
// 			if got := r.String(); got != tt.want {
// 				t.Errorf("AccountType.String() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
