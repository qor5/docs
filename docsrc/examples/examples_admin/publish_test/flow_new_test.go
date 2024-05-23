package publish_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/qor5/docs/v3/docsrc/examples/examples_admin"
	"github.com/qor5/web/v3/multipartestutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestFlowNew(t *testing.T) {
	// empty db
	dataEmpty.TruncatePut(SQLDB)

	flowNew(t, PresetsBuilder, DB)
}

func flowNew(t *testing.T, h http.Handler, db *gorm.DB) {
	flowNew_Step00_Event_presets_New(t, h).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
		// TODO: 需要确认没有 ControlBar 的显示
	})

	flowNew_Step01_Event_presets_Update(t, h).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
		// TODO: 额外校验
		// TODO: 需要确认要求执行了 reload
	})

	flowNew_Step02_Event___reload__(t, h).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
		// TODO: 其他校验
	})

	{
		var m examples_admin.WithPublishProduct
		assert.NoError(t, db.First(&m).Error)
	}
}

func flowNew_Step00_Event_presets_New(t *testing.T, h http.Handler) *multipartestutils.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products").
		EventFunc("presets_New").
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

	multipartestutils.OpenRightDrawer("New WithPublishProduct")

	return multipartestutils.NewThen(t, w, r)
}

func flowNew_Step01_Event_presets_Update(t *testing.T, h http.Handler) *multipartestutils.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products").
		EventFunc("presets_Update").
		AddField("Name", "FirstWithPublishProduct").
		AddField("Price", "123").
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
	assert.Equal(t, resp.RunScript, "vars.presetsRightDrawer = false; vars.presetsMessage = { show: true, message: \"Successfully Updated\", color: \"success\"}")

	return multipartestutils.NewThen(t, w, r)
}

func flowNew_Step02_Event___reload__(t *testing.T, h http.Handler) *multipartestutils.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products").
		EventFunc("__reload__").
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
	assert.Empty(t, resp.RunScript)

	return multipartestutils.NewThen(t, w, r)
}
