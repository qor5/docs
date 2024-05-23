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

func TestFlowVersionListDialog(t *testing.T) {
	dataSeed.TruncatePut(SQLDB)

	flowVersionListDialog(t, PresetsBuilder, DB)
}

func flowVersionListDialog(t *testing.T, h http.Handler, db *gorm.DB) {
	flowVersionListDialog_Step00_Event_presets_DetailingDrawer(t, h).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	flowVersionListDialog_Step01_Event_presets_OpenListingDialog(t, h).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	// online_version tab
	flowVersionListDialog_Step02_Event_presets_UpdateListingDialog(t, h).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	// named_version tab
	flowVersionListDialog_Step03_Event_presets_UpdateListingDialog(t, h).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	// all tab
	flowVersionListDialog_Step04_Event_presets_UpdateListingDialog(t, h).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	// keyword 2025
	flowVersionListDialog_Step05_Event_presets_UpdateListingDialog(t, h).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	// keyword 2024
	flowVersionListDialog_Step06_Event_presets_UpdateListingDialog(t, h).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	// select
	flowVersionListDialog_Step07_Event_publish_eventSelectVersion(t, h).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	// click save
	flowVersionListDialog_Step08_Event_presets_DetailingDrawer(t, h).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})
}

func flowVersionListDialog_Step00_Event_presets_DetailingDrawer(t *testing.T, h http.Handler) *multipartestutils.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products").
		EventFunc("presets_DetailingDrawer").
		Query("id", "6_2024-05-22-v02").
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

	multipartestutils.OpenRightDrawer("WithPublishProduct 6_2024-05-22-v02")

	return multipartestutils.NewThen(t, w, r)
}

func flowVersionListDialog_Step01_Event_presets_OpenListingDialog(t *testing.T, h http.Handler) *multipartestutils.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("presets_OpenListingDialog").
		Query("select_id", "6_2024-05-22-v02").
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
	assert.Equal(t, resp.UpdatePortals[0].Name, "presets_listingDialogPortalName")
	assert.Nil(t, resp.Data)
	assert.Equal(t, resp.RunScript, "setTimeout(function(){ vars.presetsListingDialog = true }, 100)")

	return multipartestutils.NewThen(t, w, r)
}

func flowVersionListDialog_Step02_Event_presets_UpdateListingDialog(t *testing.T, h http.Handler) *multipartestutils.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("presets_UpdateListingDialog").
		Query("active_filter_tab", "online_version").
		Query("f_online_version", "1").
		Query("f_select_id", "6_2024-05-22-v01").
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
	assert.Equal(t, resp.UpdatePortals[0].Name, "listingDialogContentPortal")
	assert.Nil(t, resp.Data)
	assert.Equal(t, resp.RunScript, "\nvar listingDialogElem = document.getElementById('listingDialog'); \nif (listingDialogElem.offsetHeight > parseInt(listingDialogElem.style.minHeight || '0', 10)) {\n    listingDialogElem.style.minHeight = listingDialogElem.offsetHeight+'px';\n};")

	return multipartestutils.NewThen(t, w, r)
}

func flowVersionListDialog_Step03_Event_presets_UpdateListingDialog(t *testing.T, h http.Handler) *multipartestutils.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("presets_UpdateListingDialog").
		Query("active_filter_tab", "named_versions").
		Query("f_named_versions", "1").
		Query("f_select_id", "6_2024-05-22-v01").
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
	assert.Equal(t, resp.UpdatePortals[0].Name, "listingDialogContentPortal")
	assert.Nil(t, resp.Data)
	assert.Equal(t, resp.RunScript, "\nvar listingDialogElem = document.getElementById('listingDialog'); \nif (listingDialogElem.offsetHeight > parseInt(listingDialogElem.style.minHeight || '0', 10)) {\n    listingDialogElem.style.minHeight = listingDialogElem.offsetHeight+'px';\n};")

	return multipartestutils.NewThen(t, w, r)
}

func flowVersionListDialog_Step04_Event_presets_UpdateListingDialog(t *testing.T, h http.Handler) *multipartestutils.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("presets_UpdateListingDialog").
		Query("active_filter_tab", "all").
		Query("f_all", "1").
		Query("f_select_id", "6_2024-05-22-v01").
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
	assert.Equal(t, resp.UpdatePortals[0].Name, "listingDialogContentPortal")
	assert.Nil(t, resp.Data)
	assert.Equal(t, resp.RunScript, "\nvar listingDialogElem = document.getElementById('listingDialog'); \nif (listingDialogElem.offsetHeight > parseInt(listingDialogElem.style.minHeight || '0', 10)) {\n    listingDialogElem.style.minHeight = listingDialogElem.offsetHeight+'px';\n};")

	return multipartestutils.NewThen(t, w, r)
}

func flowVersionListDialog_Step05_Event_presets_UpdateListingDialog(t *testing.T, h http.Handler) *multipartestutils.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("presets_UpdateListingDialog").
		Query("active_filter_tab", "all").
		Query("f_all", "1").
		Query("f_select_id", "6_2024-05-22-v01").
		Query("keyword", "2025").
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
	assert.Equal(t, resp.UpdatePortals[0].Name, "listingDialogContentPortal")
	assert.Nil(t, resp.Data)
	assert.Equal(t, resp.RunScript, "\nvar listingDialogElem = document.getElementById('listingDialog'); \nif (listingDialogElem.offsetHeight > parseInt(listingDialogElem.style.minHeight || '0', 10)) {\n    listingDialogElem.style.minHeight = listingDialogElem.offsetHeight+'px';\n};")

	return multipartestutils.NewThen(t, w, r)
}

func flowVersionListDialog_Step06_Event_presets_UpdateListingDialog(t *testing.T, h http.Handler) *multipartestutils.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("presets_UpdateListingDialog").
		Query("active_filter_tab", "all").
		Query("f_all", "1").
		Query("f_select_id", "6_2024-05-22-v01").
		Query("keyword", "2024").
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
	assert.Equal(t, resp.UpdatePortals[0].Name, "listingDialogContentPortal")
	assert.Nil(t, resp.Data)
	assert.Equal(t, resp.RunScript, "\nvar listingDialogElem = document.getElementById('listingDialog'); \nif (listingDialogElem.offsetHeight > parseInt(listingDialogElem.style.minHeight || '0', 10)) {\n    listingDialogElem.style.minHeight = listingDialogElem.offsetHeight+'px';\n};")

	return multipartestutils.NewThen(t, w, r)
}

func flowVersionListDialog_Step07_Event_publish_eventSelectVersion(t *testing.T, h http.Handler) *multipartestutils.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("publish_eventSelectVersion").
		Query("select_id", "6_2024-05-22-v01").
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
	assert.Empty(t, resp.UpdatePortals)
	assert.Nil(t, resp.Data)
	assert.Equal(t, resp.RunScript, "vars.presetsListingDialog = false;plaid().vars(vars).locals(locals).form(form).eventFunc(\"presets_DetailingDrawer\").queries({\"id\":[\"6_2024-05-22-v01\"]}).go()")

	return multipartestutils.NewThen(t, w, r)
}

func flowVersionListDialog_Step08_Event_presets_DetailingDrawer(t *testing.T, h http.Handler) *multipartestutils.Then {
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
