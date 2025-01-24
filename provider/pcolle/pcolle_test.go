package pcolle

import (
	"testing"

	"github.com/HDuoDuo/metatube/provider/internal/testkit"
)

func TestPcolle_GetMovieInfoByID(t *testing.T) {
	testkit.Test(t, New, []string{
		"156785614478ab480db",
	})
}
