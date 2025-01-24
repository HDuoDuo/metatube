package provider

import (
	"net/http"

	"github.com/HDuoDuo/metatube/errors"
)

var (
	ErrInvalidID          = errors.New(http.StatusBadRequest, "invalid id")
	ErrInvalidURL         = errors.New(http.StatusBadRequest, "invalid url")
	ErrInvalidKeyword     = errors.New(http.StatusBadRequest, "invalid keyword")
	ErrInfoNotFound       = errors.New(http.StatusNotFound, "info not found")
	ErrImageNotFound      = errors.New(http.StatusNotFound, "image not found")
	ErrProviderNotFound   = errors.New(http.StatusNotFound, "provider not found")
	ErrIncompleteMetadata = errors.New(http.StatusInternalServerError, "incomplete metadata")
)
