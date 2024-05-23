package publish_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/qor5/web/v3/multipartestutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestFlowDuplicate(t *testing.T) {
	dataSeed.TruncatePut(SQLDB)

	flowDuplicate(t, PresetsBuilder, DB)
}

func flowDuplicate(t *testing.T, h http.Handler, db *gorm.DB) {
	flowDuplicate_Step00_Event_presets_DetailingDrawer(t, h).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	flowDuplicate_Step01_Event_publish_EventDuplicateVersion(t, h).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	flowDuplicate_Step02_Event___reload__(t, h).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})
}

func flowDuplicate_Step00_Event_presets_DetailingDrawer(t *testing.T, h http.Handler) *multipartestutils.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products").
		EventFunc("presets_DetailingDrawer").
		Query("id", "6_2024-05-22-v01").
		BuildEventFuncRequest()

	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

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

	multipartestutils.OpenRightDrawer("WithPublishProduct 6_2024-05-22-v01")

	return multipartestutils.NewThen(t, w, r)
}

func flowDuplicate_Step01_Event_publish_EventDuplicateVersion(t *testing.T, h http.Handler) *multipartestutils.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products").
		EventFunc("publish_EventDuplicateVersion").
		Query("id", "6_2024-05-22-v01").
		BuildEventFuncRequest()

	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

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

	return multipartestutils.NewThen(t, w, r)
}

func flowDuplicate_Step02_Event___reload__(t *testing.T, h http.Handler) *multipartestutils.Then {
	// TODO: should handle now func or calc now new version name
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products/6_2024-05-23-v01").
		EventFunc("__reload__").
		BuildEventFuncRequest()

	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

	var resp multipartestutils.TestEventResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, "WithPublishProduct 6_2024-05-23-v01 - Admin", resp.PageTitle)
	assert.True(t, resp.Reload)
	assert.Nil(t, resp.PushState)
	assert.Empty(t, resp.RedirectURL)
	assert.Empty(t, resp.ReloadPortals)
	assert.Empty(t, resp.UpdatePortals)
	assert.Nil(t, resp.Data)
	assert.Empty(t, resp.RunScript)

	return multipartestutils.NewThen(t, w, r)
}
