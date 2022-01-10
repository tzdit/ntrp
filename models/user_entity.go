package models

import (
	"time"
)

//User defines class user data
type User struct {
	ID              string    `json:"id,omitempty" form:"id"`
	HospitalID      string    `json:"hospital_id,omitempty" form:"hospital_id"`
	Name            string    `json:"name,omitempty" form:"name"`
	Email           string    `json:"email,omitempty" form:"email"`
	EmailVerifiedAt time.Time `json:"email_verified_at,omitempty" form:"email_verified_at"`
	Password        string    `json:"password,omitempty" form:"password"`
	LoginAttempt    int32     `json:"login_attempt,omitempty" form:"login_attempt"`
	RememberToken   string    `json:"remember_token,omitempty" form:"remember_token"`
	Hospital        *Hospital `json:"hospital"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	DeletedAt       time.Time `json:"deleted_at,omitempty"`
}
