package publish_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/qor5/admin/v3/utils/testflow"
	"github.com/qor5/web/v3/multipartestutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type FlowDuplicate struct {
	*Flow

	// local vars
	DuplicateID string
}

func TestFlowDuplicate(t *testing.T) {
	dataSeed.TruncatePut(SQLDB)

	flowDuplicate(t, &FlowDuplicate{
		Flow: &Flow{
			db: DB, h: PresetsBuilder,
			ID: "6_2024-05-22-v01",
		},
	})
}

func flowDuplicate(t *testing.T, f *FlowDuplicate) {
	flowDuplicate_Step00_Event_presets_DetailingDrawer(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	flowDuplicate_Step01_Event_publish_EventDuplicateVersion(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
		var err error
		f.DuplicateID, err = getNextVersion(f.ID)
		assert.NoError(t, err)
	})

	flowDuplicate_Step02_Event___reload__(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})
}

func flowDuplicate_Step00_Event_presets_DetailingDrawer(t *testing.T, f *FlowDuplicate) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products").
		EventFunc("presets_DetailingDrawer").
		// Query("id", "6_2024-05-22-v01").
		Query("id", f.ID).
		BuildEventFuncRequest()

	w := httptest.NewRecorder()
	f.h.ServeHTTP(w, r)

	var resp multipartestutils.TestEventResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Empty(t, resp.PageTitle)
	assert.False(t, resp.Reload)
	assert.Nil(t, resp.PushState)
	assert.Empty(t, resp.RedirectURL)
	assert.Empty(t, resp.ReloadPortals)
	assert.Len(t, resp.UpdatePortals, 1)
	assert.Equal(t, "presets_RightDrawerPortalName", resp.UpdatePortals[0].Name)
	assert.Nil(t, resp.Data)
	assert.Equal(t, "setTimeout(function(){ vars.presetsRightDrawer = true }, 100)", resp.RunScript)

	testflow.Validate(t, w, r,
		// multipartestutils.OpenRightDrawer("WithPublishProduct 6_2024-05-22-v01")
		testflow.OpenRightDrawer("WithPublishProduct "+f.ID),
	)

	return testflow.NewThen(t, w, r)
}

func flowDuplicate_Step01_Event_publish_EventDuplicateVersion(t *testing.T, f *FlowDuplicate) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products").
		EventFunc("publish_EventDuplicateVersion").
		// Query("id", "6_2024-05-22-v01").
		Query("id", f.ID).
		BuildEventFuncRequest()

	w := httptest.NewRecorder()
	f.h.ServeHTTP(w, r)

	var resp multipartestutils.TestEventResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Empty(t, resp.PageTitle)
	assert.False(t, resp.Reload)
	assert.NotNil(t, resp.PushState)
	assert.Empty(t, resp.RedirectURL)
	assert.Empty(t, resp.ReloadPortals)
	assert.Empty(t, resp.UpdatePortals)
	assert.Nil(t, resp.Data)
	assert.Equal(t, "vars.presetsMessage = { show: true, message: \"Successfully Created\", color: \"success\"}", resp.RunScript)

	return testflow.NewThen(t, w, r)
}

func flowDuplicate_Step02_Event___reload__(t *testing.T, f *FlowDuplicate) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		// PageURL("/samples/publish/with-publish-products/6_2024-05-23-v01").
		PageURL("/samples/publish/with-publish-products/" + f.DuplicateID).
		EventFunc("__reload__").
		BuildEventFuncRequest()

	w := httptest.NewRecorder()
	f.h.ServeHTTP(w, r)

	var resp multipartestutils.TestEventResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	// assert.Equal(t, "WithPublishProduct 6_2024-05-23-v01 - Admin", resp.PageTitle)
	assert.Equal(t, "WithPublishProduct "+f.DuplicateID+" - Admin", resp.PageTitle)
	assert.True(t, resp.Reload)
	assert.Nil(t, resp.PushState)
	assert.Empty(t, resp.RedirectURL)
	assert.Empty(t, resp.ReloadPortals)
	assert.Empty(t, resp.UpdatePortals)
	assert.Nil(t, resp.Data)
	assert.Empty(t, resp.RunScript)

	return testflow.NewThen(t, w, r)
}

// TODO: maybe version name be renamed
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
