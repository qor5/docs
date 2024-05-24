package publish_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/qor5/admin/v3/utils/testflow"
	"github.com/qor5/web/v3/multipartestutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type FlowPublishAndUnpublish struct {
	*Flow
}

func TestFlowPublishAndUnpublish(t *testing.T) {
	dataSeed.TruncatePut(SQLDB)

	flowPublishAndUnpublish(t, &FlowPublishAndUnpublish{
		Flow: &Flow{
			db: DB, h: PresetsBuilder,
			ID: "6_2024-05-22-v01",
		},
	})
}

func flowPublishAndUnpublish(t *testing.T, f *FlowPublishAndUnpublish) {
	flowPublishAndUnpublish_Step00_Event_presets_DetailingDrawer(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	flowPublishAndUnpublish_Step01_Event_publish_EventPublish(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	flowPublishAndUnpublish_Step02_Event_presets_DetailingDrawer(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	flowPublishAndUnpublish_Step03_Event_publish_EventUnpublish(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})
}

func flowPublishAndUnpublish_Step00_Event_presets_DetailingDrawer(t *testing.T, f *FlowPublishAndUnpublish) *testflow.Then {
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
		// testflow.OpenRightDrawer("WithPublishProduct 6_2024-05-22-v01")
		testflow.OpenRightDrawer("WithPublishProduct "+f.ID),
	)

	return testflow.NewThen(t, w, r)
}

func flowPublishAndUnpublish_Step01_Event_publish_EventPublish(t *testing.T, f *FlowPublishAndUnpublish) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products").
		EventFunc("publish_EventPublish").
		// Query("id", "6_2024-05-22-v01").
		Query("id", f.ID).
		BuildEventFuncRequest()

	w := httptest.NewRecorder()
	f.h.ServeHTTP(w, r)

	var resp multipartestutils.TestEventResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, "Listing WithPublishProducts - Admin", resp.PageTitle)
	assert.True(t, resp.Reload)
	assert.Nil(t, resp.PushState)
	assert.Empty(t, resp.RedirectURL)
	assert.Empty(t, resp.ReloadPortals)
	assert.Empty(t, resp.UpdatePortals)
	assert.Nil(t, resp.Data)
	assert.Equal(t, "vars.presetsMessage = { show: true, message: \"success\", color: \"success\"}", resp.RunScript)

	return testflow.NewThen(t, w, r)
}

func flowPublishAndUnpublish_Step02_Event_presets_DetailingDrawer(t *testing.T, f *FlowPublishAndUnpublish) *testflow.Then {
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
		// testflow.OpenRightDrawer("WithPublishProduct 6_2024-05-22-v01")
		testflow.OpenRightDrawer("WithPublishProduct "+f.ID),
	)

	return testflow.NewThen(t, w, r)
}

func flowPublishAndUnpublish_Step03_Event_publish_EventUnpublish(t *testing.T, f *FlowPublishAndUnpublish) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products").
		EventFunc("publish_EventUnpublish").
		// Query("id", "6_2024-05-22-v01").
		Query("id", f.ID).
		BuildEventFuncRequest()

	w := httptest.NewRecorder()
	f.h.ServeHTTP(w, r)

	var resp multipartestutils.TestEventResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, "Listing WithPublishProducts - Admin", resp.PageTitle)
	assert.True(t, resp.Reload)
	assert.Nil(t, resp.PushState)
	assert.Empty(t, resp.RedirectURL)
	assert.Empty(t, resp.ReloadPortals)
	assert.Empty(t, resp.UpdatePortals)
	assert.Nil(t, resp.Data)
	assert.Equal(t, "vars.presetsMessage = { show: true, message: \"success\", color: \"success\"}", resp.RunScript)

	return testflow.NewThen(t, w, r)
}
