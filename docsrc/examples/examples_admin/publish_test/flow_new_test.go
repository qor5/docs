package publish_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/qor5/admin/v3/utils/testflow"
	"github.com/qor5/docs/v3/docsrc/examples/examples_admin"
	"github.com/qor5/web/v3/multipartestutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type FlowNew struct {
	*Flow

	// local vars
	Name  string
	Price int
}

func TestFlowNew(t *testing.T) {
	// empty db
	dataEmpty.TruncatePut(SQLDB)

	flowNew(t, &FlowNew{
		Flow: &Flow{
			db: DB, h: PresetsBuilder,
		},
		Name:  "FirstWithPublishProduct",
		Price: 123,
	})
}

func flowNew(t *testing.T, f *FlowNew) {
	flowNew_Step00_Event_presets_New(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
		// TODO: 需要确认没有 ControlBar 的显示
	})

	previous := time.Now()

	flowNew_Step01_Event_presets_Update(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
		// TODO: 额外校验
		// TODO: 需要确认要求执行了 reload
	})

	flowNew_Step02_Event___reload__(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
		// TODO: 其他校验
	})

	{
		var m examples_admin.WithPublishProduct
		assert.NoError(t, f.db.Where("created_at > ?", previous).Order("created_at DESC").First(&m).Error)
		assert.Equal(t, f.Name, m.Name)
		assert.Equal(t, f.Price, m.Price)
	}
}

func flowNew_Step00_Event_presets_New(t *testing.T, f *FlowNew) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products").
		EventFunc("presets_New").
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
		testflow.OpenRightDrawer("New WithPublishProduct"),
	)

	return testflow.NewThen(t, w, r)
}

func flowNew_Step01_Event_presets_Update(t *testing.T, f *FlowNew) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products").
		EventFunc("presets_Update").
		// AddField("Name", "FirstWithPublishProduct").
		AddField("Name", f.Name).
		// AddField("Price", "123").
		AddField("Price", fmt.Sprint(f.Price)).
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
	assert.Equal(t, "vars.presetsRightDrawer = false; vars.presetsMessage = { show: true, message: \"Successfully Updated\", color: \"success\"}", resp.RunScript)

	return testflow.NewThen(t, w, r)
}

func flowNew_Step02_Event___reload__(t *testing.T, f *FlowNew) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products").
		EventFunc("__reload__").
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
	assert.Empty(t, resp.RunScript)

	return testflow.NewThen(t, w, r)
}
