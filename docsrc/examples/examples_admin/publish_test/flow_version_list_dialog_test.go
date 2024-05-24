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

type FlowVersionListDialog struct {
	*FlowDuplicate
}

func TestFlowVersionListDialog(t *testing.T) {
	dataSeed.TruncatePut(SQLDB)

	flowVersionListDialog(t, &FlowVersionListDialog{
		FlowDuplicate: &FlowDuplicate{
			Flow: &Flow{
				db: DB, h: PresetsBuilder,
				ID: "6_2024-05-22-v01",
			},
		},
	})
}

func flowVersionListDialog(t *testing.T, f *FlowVersionListDialog) {
	// duplicate one
	flowDuplicate(t, f.FlowDuplicate)

	// open duplicate drawer
	flowVersionListDialog_Step00_Event_presets_DetailingDrawer(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	// open list dialog
	flowVersionListDialog_Step01_Event_presets_OpenListingDialog(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	// online_version tab
	flowVersionListDialog_Step02_Event_presets_UpdateListingDialog(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	// named_version tab
	flowVersionListDialog_Step03_Event_presets_UpdateListingDialog(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	// all tab
	flowVersionListDialog_Step04_Event_presets_UpdateListingDialog(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	// keyword 2025
	flowVersionListDialog_Step05_Event_presets_UpdateListingDialog(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	// keyword 2024
	flowVersionListDialog_Step06_Event_presets_UpdateListingDialog(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	// select
	flowVersionListDialog_Step07_Event_publish_eventSelectVersion(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	// click save
	flowVersionListDialog_Step08_Event_presets_DetailingDrawer(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})
}

func flowVersionListDialog_Step00_Event_presets_DetailingDrawer(t *testing.T, f *FlowVersionListDialog) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products").
		EventFunc("presets_DetailingDrawer").
		// Query("id", "6_2024-05-22-v02").
		Query("id", f.DuplicateID).
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
		// testflow.OpenRightDrawer("WithPublishProduct 6_2024-05-22-v02"),
		testflow.OpenRightDrawer("WithPublishProduct "+f.DuplicateID),
	)
	return testflow.NewThen(t, w, r)
}

func flowVersionListDialog_Step01_Event_presets_OpenListingDialog(t *testing.T, f *FlowVersionListDialog) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("presets_OpenListingDialog").
		// Query("select_id", "6_2024-05-22-v02").
		Query("select_id", f.DuplicateID).
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
	assert.Equal(t, "presets_listingDialogPortalName", resp.UpdatePortals[0].Name)
	assert.Nil(t, resp.Data)
	assert.Equal(t, "setTimeout(function(){ vars.presetsListingDialog = true }, 100)", resp.RunScript)

	return testflow.NewThen(t, w, r)
}

func flowVersionListDialog_Step02_Event_presets_UpdateListingDialog(t *testing.T, f *FlowVersionListDialog) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("presets_UpdateListingDialog").
		Query("active_filter_tab", "online_version").
		Query("f_online_version", "1").
		// Query("f_select_id", "6_2024-05-22-v01").
		Query("f_select_id", f.ID).
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
	assert.Equal(t, "listingDialogContentPortal", resp.UpdatePortals[0].Name)
	assert.Nil(t, resp.Data)
	assert.Equal(t, "\nvar listingDialogElem = document.getElementById('listingDialog'); \nif (listingDialogElem.offsetHeight > parseInt(listingDialogElem.style.minHeight || '0', 10)) {\n    listingDialogElem.style.minHeight = listingDialogElem.offsetHeight+'px';\n};", resp.RunScript)

	return testflow.NewThen(t, w, r)
}

func flowVersionListDialog_Step03_Event_presets_UpdateListingDialog(t *testing.T, f *FlowVersionListDialog) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("presets_UpdateListingDialog").
		Query("active_filter_tab", "named_versions").
		Query("f_named_versions", "1").
		// Query("f_select_id", "6_2024-05-22-v01").
		Query("f_select_id", f.ID).
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
	assert.Equal(t, "listingDialogContentPortal", resp.UpdatePortals[0].Name)
	assert.Nil(t, resp.Data)
	assert.Equal(t, "\nvar listingDialogElem = document.getElementById('listingDialog'); \nif (listingDialogElem.offsetHeight > parseInt(listingDialogElem.style.minHeight || '0', 10)) {\n    listingDialogElem.style.minHeight = listingDialogElem.offsetHeight+'px';\n};", resp.RunScript)

	return testflow.NewThen(t, w, r)
}

func flowVersionListDialog_Step04_Event_presets_UpdateListingDialog(t *testing.T, f *FlowVersionListDialog) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("presets_UpdateListingDialog").
		Query("active_filter_tab", "all").
		Query("f_all", "1").
		// Query("f_select_id", "6_2024-05-22-v01").
		Query("f_select_id", f.ID).
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
	assert.Equal(t, "listingDialogContentPortal", resp.UpdatePortals[0].Name)
	assert.Nil(t, resp.Data)
	assert.Equal(t, "\nvar listingDialogElem = document.getElementById('listingDialog'); \nif (listingDialogElem.offsetHeight > parseInt(listingDialogElem.style.minHeight || '0', 10)) {\n    listingDialogElem.style.minHeight = listingDialogElem.offsetHeight+'px';\n};", resp.RunScript)

	return testflow.NewThen(t, w, r)
}

func flowVersionListDialog_Step05_Event_presets_UpdateListingDialog(t *testing.T, f *FlowVersionListDialog) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("presets_UpdateListingDialog").
		Query("active_filter_tab", "all").
		Query("f_all", "1").
		// Query("f_select_id", "6_2024-05-22-v01").
		Query("f_select_id", f.ID).
		Query("keyword", "2025").
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
	assert.Equal(t, "listingDialogContentPortal", resp.UpdatePortals[0].Name)
	assert.Nil(t, resp.Data)
	assert.Equal(t, "\nvar listingDialogElem = document.getElementById('listingDialog'); \nif (listingDialogElem.offsetHeight > parseInt(listingDialogElem.style.minHeight || '0', 10)) {\n    listingDialogElem.style.minHeight = listingDialogElem.offsetHeight+'px';\n};", resp.RunScript)

	return testflow.NewThen(t, w, r)
}

func flowVersionListDialog_Step06_Event_presets_UpdateListingDialog(t *testing.T, f *FlowVersionListDialog) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("presets_UpdateListingDialog").
		Query("active_filter_tab", "all").
		Query("f_all", "1").
		// Query("f_select_id", "6_2024-05-22-v01").
		Query("f_select_id", f.ID).
		Query("keyword", "2024").
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
	assert.Equal(t, "listingDialogContentPortal", resp.UpdatePortals[0].Name)
	assert.Nil(t, resp.Data)
	assert.Equal(t, "\nvar listingDialogElem = document.getElementById('listingDialog'); \nif (listingDialogElem.offsetHeight > parseInt(listingDialogElem.style.minHeight || '0', 10)) {\n    listingDialogElem.style.minHeight = listingDialogElem.offsetHeight+'px';\n};", resp.RunScript)

	return testflow.NewThen(t, w, r)
}

func flowVersionListDialog_Step07_Event_publish_eventSelectVersion(t *testing.T, f *FlowVersionListDialog) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("publish_eventSelectVersion").
		// Query("select_id", "6_2024-05-22-v01").
		Query("select_id", f.ID).
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
	assert.Empty(t, resp.UpdatePortals)
	assert.Nil(t, resp.Data)
	assert.Equal(t, "vars.presetsListingDialog = false;plaid().vars(vars).locals(locals).form(form).eventFunc(\"presets_DetailingDrawer\").queries({\"id\":[\"6_2024-05-22-v01\"]}).go()", resp.RunScript)

	return testflow.NewThen(t, w, r)
}

func flowVersionListDialog_Step08_Event_presets_DetailingDrawer(t *testing.T, f *FlowVersionListDialog) *testflow.Then {
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
		// testflow.OpenRightDrawer("WithPublishProduct 6_2024-05-22-v01"),
		testflow.OpenRightDrawer("WithPublishProduct "+f.ID),
	)
	return testflow.NewThen(t, w, r)
}
