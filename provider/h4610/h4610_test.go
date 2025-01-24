package h4610

import (
	"testing"

	"github.com/HDuoDuo/metatube/provider/internal/testkit"
)

func TestH4610_GetMovieInfoByID(t *testing.T) {
	testkit.Test(t, New, []string{
		"tk0047",
		"pla0051",
		"tk0062",
		"tk0050",
	})
}
