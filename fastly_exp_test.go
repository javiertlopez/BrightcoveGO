package main

import "testing"

func Test_decode(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"success",
			args{
				"NWZmY2U4MWFfY2ViMDViZWRkOWYwZTRkY2NkYTIxOWE3NjRhYmNlMjczZWViMWNmYWQ2ZTc4OTQzYmViNTE0YWMwZjBlMjFjNQ%3D%3D",
			},
			"2021-01-12 00:06:50 +0000 UTC",
			false,
		},
		{
			"error (not base64)",
			args{
				"NWZmY2U4MWFfY2ViM",
			},
			"",
			true,
		},
		{
			"error (not hex)",
			args{
				"aGVsbG9fdGhlcmU=",
			},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decode(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
