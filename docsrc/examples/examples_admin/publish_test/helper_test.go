package publish_test

import (
	"fmt"
	"strings"
)

func ContainsVersionBar(body string) bool {
	return strings.Contains(body, "presets_OpenListingDialog") && strings.Contains(body, "-version-list-dialog")
}

func mustSplitIDVersion(expr string) []string {
	segs := strings.Split(expr, "_")
	if len(segs) < 2 {
		panic(fmt.Errorf("invalid expr %q", expr))
	}
	return segs[0:2]
}

func mustIDVersion(expr string) (string, string) {
	segs := mustSplitIDVersion(expr)
	return segs[0], segs[1]
}
