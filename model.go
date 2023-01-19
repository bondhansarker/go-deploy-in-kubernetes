package main

import "time"

type User struct {
	ID                  uint      `json:"id" gorm:"index"`
	Email               string    `json:"email" gorm:"unique"`
	Password            *string   `json:"password,omitempty"`
	Name                string    `json:"name"`
	UserName            string    `json:"user_name" gorm:"index"`
	Phone               *string   `json:"phone"`
	Website             *string   `json:"website"`
	Bio                 *string   `json:"bio"`
	Gender              *string   `json:"gender"`
	ProfilePic          *string   `json:"profile_pic"`
	ProfilePicExtension *string   `json:"profile_pic_extension"`
	Verified            *bool     `json:"verified" gorm:"type:tinyint(1);default:0;not null;"`
	IsAdmin             *bool     `json:"is_admin" gorm:"type:tinyint(1);default:0;not null"`
	IsActive            *bool     `json:"is_active" gorm:"type:tinyint(1);default:1;not null"`
	IsDeleted           *bool     `json:"is_deleted" gorm:"type:tinyint(1);default:0;not null"`
	IsBlocked           *bool     `json:"is_blocked" gorm:"type:tinyint(1);default:0;not null"`
	DownloadCount       *int64    `json:"download_count" gorm:"type:bigint;default:0;not null"`
	UploadCount         *int64    `json:"upload_count" gorm:"type:bigint;default:0;not null"`
	LoginProvider       string    `json:"login_provider"`
	AppleIdentityToken  string    `json:"apple_identity_token"`
	CreatedAt           time.Time `json:"created_at" gorm:"index"`
	UpdatedAt           time.Time `json:"updated_at"`
}
