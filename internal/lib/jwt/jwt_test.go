package jwt

import (
	"sso/internal/domain/models"
	"testing"
	"time"
)

func TestNewToken(t *testing.T) {
	cases := []struct {
		name     string
		user     models.User
		app      models.App
		duration time.Duration
		want     string
	}{
		{
			name: "Success",
			user: models.User{
				ID:    1,
				Email: "asd@asd.asd",
			},
			duration: time.Hour,
			app: models.App{
				ID:     2,
				Secret: "secret",
			},
			want: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfaWQiOjIsImVtYWlsIjoiYXNkQGFzZC5hc2QiLCJleHAiOjE3MzA5Njk2MDUsInVpZCI6MX0.LkZi8FYtZ75j26hlOhK3GI12hVYLmzI06M6B_lS-3Ho",
		},
		{
			name: "Invalid email",
			user: models.User{
				ID:    1,
				Email: "asd@asd.asd",
			},
			duration: time.Hour,
			app: models.App{
				ID:     2,
				Secret: "secret",
			},
			want: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfaWQiOjIsImVtYWlsIjoiYXNkQGFzZC5hc2QiLCJleHAiOjE3MzA5Njk2MDUsInVpZCI6MX0.LkZi8FYtZ75j26hlOhK3GI12hVYLmzI06M6B_lS-3Ho",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := NewToken(tc.user, tc.app, tc.duration)
			if err != nil {
				t.Errorf("NewToken() error = %v", err)
			}
			if got != tc.want {
				t.Errorf("NewToken() = %v, want %v", got, tc.want)
			}
			/* got, err := NewToken(tt.user, tt.app, tt.duration)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewToken() = %v, want %v", got, tt.want)
			} */
		})
	}
}
