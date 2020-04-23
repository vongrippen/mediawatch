package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// MediaFile is used by pop to map your media_files database table to your go code.
type MediaFile struct {
	ID          uuid.UUID     `json:"id" db:"id"`
	CreatedAt   time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" db:"updated_at"`
	Size        nulls.Int64   `json:"size" db:"size" csv:"size"`
	Pathname    string        `json:"pathname" db:"pathname" csv:"pathname"`
	Filename    string        `json:"filename" db:"filename" csv:"filename"`
	ContentType nulls.String  `json:"content_type" db:"content_type" csv:"content_type"`
	VideoCodec  nulls.String  `json:"video_codec" db:"video_codec" csv:"video_codec"`
	AudioCodec  nulls.String  `json:"audio_codec" db:"audio_codec" csv:"audio_codec"`
	Duration    nulls.Float64 `json:"duration_seconds" db:"duration_seconds" csv:"duration_seconds"`
	Format      nulls.String  `json:"container_format" db:"container_format" csv:"container_format"`
	Width       nulls.Int     `json:"width" db:"width" csv:"width"`
	Height      nulls.Int     `json:"height" db:"height" csv:"height"`
	Show        nulls.String  `json:"show" db:"show" csv:"show"`
	Season      nulls.String  `json:"season" db:"season" csv:"season"`
}

// String is not required by pop and may be deleted
func (m MediaFile) String() string {
	jm, _ := json.Marshal(m)
	return string(jm)
}

// MediaFiles is not required by pop and may be deleted
type MediaFiles []MediaFile

// String is not required by pop and may be deleted
func (m MediaFiles) String() string {
	jm, _ := json.Marshal(m)
	return string(jm)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (m *MediaFile) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (m *MediaFile) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (m *MediaFile) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
