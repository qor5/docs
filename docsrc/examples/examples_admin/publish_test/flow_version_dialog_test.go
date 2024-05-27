package publish_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/qor5/admin/v3/utils/testflow"
	"github.com/qor5/docs/v3/docsrc/examples/examples_admin"
	"github.com/qor5/web/v3/multipartestutils"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/theplant/gofixtures"
)

var dataSeedForVersionDialog = gofixtures.Data(gofixtures.Sql(`
INSERT INTO "public"."with_publish_products" ("id", "created_at", "updated_at", "deleted_at", "name", "price", "status", "online_url", "scheduled_start_at", "scheduled_end_at", "actual_start_at", "actual_end_at", "version", "version_name", "parent_version") VALUES ('1', '2024-05-26 13:12:06.408234+00', '2024-05-26 13:12:06.408234+00', NULL, 'FirstProduct', '123', 'draft', '', NULL, NULL, NULL, NULL, '2024-05-26-v01', '2024-05-26-v01', ''),
('1', '2024-05-26 13:13:09.768116+00', '2024-05-26 13:13:09.764082+00', NULL, 'FirstProduct', '123', 'draft', '', NULL, NULL, NULL, NULL, '2024-05-26-v02', '2024-05-26-v02', '2024-05-26-v01'),
('1', '2024-05-26 13:13:11.858454+00', '2024-05-26 13:13:11.855648+00', NULL, 'FirstProduct', '123', 'draft', '', NULL, NULL, NULL, NULL, '2024-05-26-v03', '2024-05-26-v03', '2024-05-26-v02'),
('1', '2024-05-26 13:13:14.463547+00', '2024-05-26 13:14:47.64948+00', NULL, 'FirstProduct', '123', 'draft', '', NULL, NULL, NULL, NULL, '2024-05-26-v04', '2024-05-26-x04', '2024-05-26-v03'),
('1', '2024-05-26 13:13:16.56434+00', '2024-05-26 13:14:39.705527+00', NULL, 'FirstProduct', '123', 'draft', '', NULL, NULL, NULL, NULL, '2024-05-26-v05', '2024-05-26-x05', '2024-05-26-v04'),
('1', '2024-05-26 13:13:18.256404+00', '2024-05-26 13:14:43.729016+00', NULL, 'FirstProduct', '123', 'draft', '', NULL, NULL, NULL, NULL, '2024-05-26-v06', '2024-05-26-x06', '2024-05-26-v05');
`, []string{"with_publish_products"}))

type FlowVersionDialog struct {
	*Flow
}

func TestFlowVersionDialog(t *testing.T) {
	dataSeedForVersionDialog.TruncatePut(SQLDB)
	flowVersionDialog(t, &FlowVersionDialog{
		Flow: &Flow{
			db: DB, h: PresetsBuilder,
		},
	})
}

func flowVersionDialog(t *testing.T, f *FlowVersionDialog) {
	displayID := "1_2024-05-26-v06"

	models := []*examples_admin.WithPublishProduct{}
	id, _ := mustIDVersion(displayID)
	require.NoError(t, f.db.Where("id = ?", id).Order("version DESC").Find(&models).Error)
	assert.Len(t, models, 6)

	flowVersionDialog_Step00_Event_presets_DetailingDrawer(t, f).ThenValidate(
		// ensure current display id would be set
		testflow.ContainsInOrderAtUpdatePortal(0, "<v-chip", fmt.Sprintf(`vars.publish_VarCurrentDisplayID = %q`, displayID), "</v-chip>"),
	)

	selectID := displayID
	dislayModels := models
	ensureListDisplay := func() testflow.ValidatorFunc {
		return testflow.Combine(
			// ensure list head display
			testflow.ContainsInOrderAtUpdatePortal(0,
				// 确认 tab 显示
				"<v-tabs",
				"active_filter_tab", "all", "f_all", "f_select_id", selectID, "All Versions",
				"active_filter_tab", "online_versions", "f_online_versions", "f_select_id", selectID, "Online Versions",
				"active_filter_tab", "named_versions", "f_named_versions", "f_select_id", selectID, "Named Versions",
				"</v-tabs>",
				// 确认列名显示
				"<tr>", "<th>Version</th>", "<th>State</th>", "<th>Start at</th>", "<th>End at</th>", "<th>Unread Notes</th>", "<th>Option</th>", "</tr>",
			// TODO: 需要确认 vars.publish_VarCurrentDisplayID 是否匹配
			),
			// ensure list content display
			testflow.WrapEvent(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request, e multipartestutils.TestEventResponse) {
				subs := regexp.MustCompile(`<tr[\s\S]+?<td>[\s\S]+?<v-radio :model-value='([^']+)'\s*:true-value='([^']+)'[\s\S]+?</v-radio>\s*([^<]+)?\s*</div>[\s\S]+?</tr>`).FindAllStringSubmatch(e.UpdatePortals[0].Body, -1)
				assert.Len(t, subs, len(dislayModels))
				for i, sub := range subs {
					// 确认只有被选择项目被打标，确认显示的是版本名称，而非原始版本号
					modelValue, _ := strconv.Unquote(sub[1])
					trueValue, _ := strconv.Unquote(sub[2])
					assert.Equal(t, dislayModels[i].PrimarySlug(), modelValue)
					assert.Equal(t, selectID, trueValue)
					assert.Equal(t, dislayModels[i].Version.VersionName, sub[3])
				}
			}),
		)
	}

	// 打开版本列表
	flowVersionDialog_Step01_Event_presets_OpenListingDialog(t, f).ThenValidate(ensureListDisplay())

	// 切换选择
	selectID = "1_2024-05-26-v05"
	flowVersionDialog_Step02_Event_presets_UpdateListingDialog(t, f).ThenValidate(ensureListDisplay())

	// 切换 tab 至 named_version
	namedModels := lo.Filter(models, func(item *examples_admin.WithPublishProduct, index int) bool {
		return item.Version.VersionName != item.Version.Version
	})
	dislayModels = namedModels
	flowVersionDialog_Step03_Event_presets_UpdateListingDialog(t, f).ThenValidate(ensureListDisplay())

	// 切换选择
	selectID = "1_2024-05-26-v04"
	flowVersionDialog_Step04_Event_presets_UpdateListingDialog(t, f).ThenValidate(ensureListDisplay())

	// 关键词 A
	dislayModels = lo.Filter(namedModels, func(item *examples_admin.WithPublishProduct, index int) bool {
		return strings.Contains(item.Version.VersionName, "2025") // TODO: 需要优化
	})
	flowVersionDialog_Step05_Event_presets_UpdateListingDialog(t, f).ThenValidate(ensureListDisplay())

	// 关键词 B
	dislayModels = lo.Filter(namedModels, func(item *examples_admin.WithPublishProduct, index int) bool {
		return strings.Contains(item.Version.VersionName, "2024") // TODO: 需要优化
	})
	flowVersionDialog_Step06_Event_presets_UpdateListingDialog(t, f).ThenValidate(ensureListDisplay())

	// 选中当前显示
	selectID = displayID
	flowVersionDialog_Step07_Event_presets_UpdateListingDialog(t, f).ThenValidate(ensureListDisplay())

	// 确认选择，即点击 Save
	flowVersionDialog_Step08_Event_publish_eventSelectVersion(t, f)

	// 再次打开版本列表
	dislayModels = models
	flowVersionDialog_Step09_Event_presets_OpenListingDialog(t, f).ThenValidate(ensureListDisplay())

	// 选择非当前显示
	selectID = "1_2024-05-26-v05"
	flowVersionDialog_Step10_Event_presets_UpdateListingDialog(t, f).ThenValidate(ensureListDisplay())

	// 确认选择
	flowVersionDialog_Step11_Event_publish_eventSelectVersion(t, f)

	// 被要求打开新选择的 Drawer
	displayID = selectID
	flowVersionDialog_Step12_Event_presets_DetailingDrawer(t, f)
	// TODO: 这点需要 check
}

func flowVersionDialog_Step00_Event_presets_DetailingDrawer(t *testing.T, f *FlowVersionDialog) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products").
		EventFunc("presets_DetailingDrawer").
		Query("id", "1_2024-05-26-v06").
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
		testflow.OpenRightDrawer("WithPublishProduct 1_2024-05-26-v06"),
	)

	return testflow.NewThen(t, w, r)
}

func flowVersionDialog_Step01_Event_presets_OpenListingDialog(t *testing.T, f *FlowVersionDialog) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("presets_OpenListingDialog").
		Query("select_id", "1_2024-05-26-v06").
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

func flowVersionDialog_Step02_Event_presets_UpdateListingDialog(t *testing.T, f *FlowVersionDialog) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("presets_UpdateListingDialog").
		Query("select_id", "1_2024-05-26-v05").
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

func flowVersionDialog_Step03_Event_presets_UpdateListingDialog(t *testing.T, f *FlowVersionDialog) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("presets_UpdateListingDialog").
		Query("active_filter_tab", "named_versions").
		Query("f_named_versions", "1").
		Query("f_select_id", "1_2024-05-26-v05").
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

func flowVersionDialog_Step04_Event_presets_UpdateListingDialog(t *testing.T, f *FlowVersionDialog) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("presets_UpdateListingDialog").
		Query("active_filter_tab", "named_versions").
		Query("f_named_versions", "1").
		Query("f_select_id", "1_2024-05-26-v05").
		Query("select_id", "1_2024-05-26-v04").
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

func flowVersionDialog_Step05_Event_presets_UpdateListingDialog(t *testing.T, f *FlowVersionDialog) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("presets_UpdateListingDialog").
		Query("active_filter_tab", "named_versions").
		Query("f_named_versions", "1").
		Query("f_select_id", "1_2024-05-26-v05").
		Query("keyword", "2025").
		Query("select_id", "1_2024-05-26-v04").
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

func flowVersionDialog_Step06_Event_presets_UpdateListingDialog(t *testing.T, f *FlowVersionDialog) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("presets_UpdateListingDialog").
		Query("active_filter_tab", "named_versions").
		Query("f_named_versions", "1").
		Query("f_select_id", "1_2024-05-26-v05").
		Query("keyword", "2024").
		Query("select_id", "1_2024-05-26-v04").
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

func flowVersionDialog_Step07_Event_presets_UpdateListingDialog(t *testing.T, f *FlowVersionDialog) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("presets_UpdateListingDialog").
		Query("active_filter_tab", "named_versions").
		Query("f_named_versions", "1").
		Query("f_select_id", "1_2024-05-26-v05").
		Query("keyword", "2024").
		Query("select_id", "1_2024-05-26-v06").
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

func flowVersionDialog_Step08_Event_publish_eventSelectVersion(t *testing.T, f *FlowVersionDialog) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products").
		EventFunc("publish_eventSelectVersion").
		Query("select_id", "1_2024-05-26-v06").
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
	assert.Equal(t, "vars.presetsListingDialog = false; if (!!vars.publish_VarCurrentDisplayID && vars.publish_VarCurrentDisplayID != \"1_2024-05-26-v06\") { vars.presetsRightDrawer = false;plaid().vars(vars).locals(locals).form(form).eventFunc(\"presets_DetailingDrawer\").query(\"id\", \"1_2024-05-26-v06\").go() }", resp.RunScript)

	return testflow.NewThen(t, w, r)
}

func flowVersionDialog_Step09_Event_presets_OpenListingDialog(t *testing.T, f *FlowVersionDialog) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("presets_OpenListingDialog").
		Query("select_id", "1_2024-05-26-v06").
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

func flowVersionDialog_Step10_Event_presets_UpdateListingDialog(t *testing.T, f *FlowVersionDialog) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("presets_UpdateListingDialog").
		Query("select_id", "1_2024-05-26-v05").
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

func flowVersionDialog_Step11_Event_publish_eventSelectVersion(t *testing.T, f *FlowVersionDialog) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products").
		EventFunc("publish_eventSelectVersion").
		Query("select_id", "1_2024-05-26-v05").
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
	assert.Equal(t, "vars.presetsListingDialog = false; if (!!vars.publish_VarCurrentDisplayID && vars.publish_VarCurrentDisplayID != \"1_2024-05-26-v05\") { vars.presetsRightDrawer = false;plaid().vars(vars).locals(locals).form(form).eventFunc(\"presets_DetailingDrawer\").query(\"id\", \"1_2024-05-26-v05\").go() }", resp.RunScript)

	return testflow.NewThen(t, w, r)
}

func flowVersionDialog_Step12_Event_presets_DetailingDrawer(t *testing.T, f *FlowVersionDialog) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products").
		EventFunc("presets_DetailingDrawer").
		Query("id", "1_2024-05-26-v05").
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
		testflow.OpenRightDrawer("WithPublishProduct 1_2024-05-26-v05"),
	)

	return testflow.NewThen(t, w, r)
}
