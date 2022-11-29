package httptomd

import (
	"testing"
)

var ttValidateURL = []struct {
	name    string
	input   string
	wantErr bool
	wantURL string
}{
	{
		name:    "invalid url",
		input:   " ",
		wantErr: true,
	},
	{
		name:    "valid url without http",
		input:   "google.com",
		wantErr: false,
		wantURL: "http://google.com",
	},
	{
		name:    "valid url with http",
		input:   "http://google.com",
		wantErr: false,
		wantURL: "http://google.com",
	},
}

func TestValidateURL(t *testing.T) {
	for _, tt := range ttValidateURL {
		t.Run(tt.name, func(t *testing.T) {
			newURL, err := validateURL(tt.input)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("validateURL -> expected error %v, got %v", tt.wantErr, err)
				}
			} else if newURL != tt.wantURL {
				t.Errorf("validateURL -> expected %v, got %v", tt.wantURL, newURL)
			}
		})
	}
}
