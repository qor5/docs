package publish_test

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func containsVersionBar(body string) bool {
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

func getNextVersion(currentVersion string) (string, error) {
	parts := strings.Split(currentVersion, "_")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid version format")
	}

	id := parts[0]
	dateVersionPart := parts[1]
	dateVersion := strings.Split(dateVersionPart, "-")
	if len(dateVersion) != 4 {
		return "", fmt.Errorf("invalid date-version part format")
	}

	dateStr, versionStr := strings.Join(dateVersion[0:3], "-"), dateVersion[3]
	versionNumberStr := strings.TrimPrefix(versionStr, "v")
	versionNumber, err := strconv.Atoi(versionNumberStr)
	if err != nil {
		return "", fmt.Errorf("invalid version number")
	}

	currentDate := time.Now().UTC().Format("2006-01-02")

	var nextVersion string
	if dateStr == currentDate {
		nextVersion = fmt.Sprintf("%s_%s-v%02d", id, currentDate, versionNumber+1)
	} else {
		nextVersion = fmt.Sprintf("%s_%s-v01", id, currentDate)
	}

	return nextVersion, nil
}
