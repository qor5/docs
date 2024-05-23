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

func TestFlowPublishAndUnpublish(t *testing.T) {
	dataSeed.TruncatePut(SQLDB)

	flowPublishAndUnpublish(t, PresetsBuilder, DB)
}

func flowPublishAndUnpublish(t *testing.T, h http.Handler, db *gorm.DB) {
	flowPublishAndUnpublish_Step00_Event_presets_DetailingDrawer(t, h).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	flowPublishAndUnpublish_Step01_Event_publish_EventPublish(t, h).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	flowPublishAndUnpublish_Step02_Event_presets_DetailingDrawer(t, h).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	flowPublishAndUnpublish_Step03_Event_publish_EventUnpublish(t, h).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})
}

func flowPublishAndUnpublish_Step00_Event_presets_DetailingDrawer(t *testing.T, h http.Handler) *multipartestutils.Then {
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
	assert.Equal(t, resp.UpdatePortals[0].Name, "presets_RightDrawerPortalName")
	assert.Nil(t, resp.Data)
	assert.Equal(t, resp.RunScript, "setTimeout(function(){ vars.presetsRightDrawer = true }, 100)")

	multipartestutils.OpenRightDrawer("WithPublishProduct 6_2024-05-22-v01")

	return multipartestutils.NewThen(t, w, r)
}

func flowPublishAndUnpublish_Step01_Event_publish_EventPublish(t *testing.T, h http.Handler) *multipartestutils.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products").
		EventFunc("publish_EventPublish").
		Query("id", "6_2024-05-22-v01").
		BuildEventFuncRequest()

	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

	var resp multipartestutils.TestEventResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, resp.PageTitle, "Listing WithPublishProducts - Admin")
	assert.True(t, resp.Reload)
	assert.Nil(t, resp.PushState)
	assert.Empty(t, resp.RedirectURL)
	assert.Empty(t, resp.ReloadPortals)
	assert.Empty(t, resp.UpdatePortals)
	assert.Nil(t, resp.Data)
	assert.Equal(t, resp.RunScript, "vars.presetsMessage = { show: true, message: \"success\", color: \"success\"}")

	return multipartestutils.NewThen(t, w, r)
}

func flowPublishAndUnpublish_Step02_Event_presets_DetailingDrawer(t *testing.T, h http.Handler) *multipartestutils.Then {
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
	assert.Equal(t, resp.UpdatePortals[0].Name, "presets_RightDrawerPortalName")
	assert.Nil(t, resp.Data)
	assert.Equal(t, resp.RunScript, "setTimeout(function(){ vars.presetsRightDrawer = true }, 100)")

	multipartestutils.OpenRightDrawer("WithPublishProduct 6_2024-05-22-v01")

	return multipartestutils.NewThen(t, w, r)
}

func flowPublishAndUnpublish_Step03_Event_publish_EventUnpublish(t *testing.T, h http.Handler) *multipartestutils.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products").
		EventFunc("publish_EventUnpublish").
		Query("id", "6_2024-05-22-v01").
		BuildEventFuncRequest()

	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

	var resp multipartestutils.TestEventResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, resp.PageTitle, "Listing WithPublishProducts - Admin")
	assert.True(t, resp.Reload)
	assert.Nil(t, resp.PushState)
	assert.Empty(t, resp.RedirectURL)
	assert.Empty(t, resp.ReloadPortals)
	assert.Empty(t, resp.UpdatePortals)
	assert.Nil(t, resp.Data)
	assert.Equal(t, resp.RunScript, "vars.presetsMessage = { show: true, message: \"success\", color: \"success\"}")

	return multipartestutils.NewThen(t, w, r)
}
