package publish_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/qor5/admin/v3/publish"
	"github.com/qor5/admin/v3/utils/testflow"
	"github.com/qor5/docs/v3/docsrc/examples/examples_admin"
	"github.com/qor5/web/v3/multipartestutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/theplant/gofixtures"
	"gorm.io/gorm"
)

var dataEmptyForFlowNew = gofixtures.Data(gofixtures.Sql(``, []string{"with_publish_products"}))

type FlowNew struct {
	*Flow

	Name  string
	Price int
}

func TestFlowNew(t *testing.T) {
	dataEmptyForFlowNew.TruncatePut(SQLDB)

	flowNew(t, &FlowNew{
		Flow: &Flow{
			db: DB, h: PresetsBuilder,
		},
		Name:  "TestProduct",
		Price: 234,
	})
}

func flowNew(t *testing.T, f *FlowNew) {
	previous := time.Now()

	flowNew_Step00_Event_presets_New(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		assert.False(t, containsVersionBar(w.Body.String()))
	})

	// PushStateNotNil inside ensures its interaction: reload
	flowNew_Step01_Event_presets_Update(t, f)

	var m examples_admin.WithPublishProduct
	assert.NoError(t, f.db.Where("created_at > ?", previous).Order("created_at DESC").First(&m).Error)
	assert.Equal(t, m.Version.Version, m.Version.VersionName)
	assert.Empty(t, m.Version.ParentVersion)

	// for compare
	m.Model = gorm.Model{}
	m.Version = publish.Version{}
	assert.Equal(t, examples_admin.WithPublishProduct{
		Name:     f.Name,
		Price:    f.Price,
		Status:   publish.Status{Status: publish.StatusDraft},
		Schedule: publish.Schedule{},
		Version:  publish.Version{},
	}, m)

	flowNew_Step02_Event___reload__(t, f)
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
		AddField("Name", f.Name).
		AddField("Price", fmt.Sprint(f.Price)).
		BuildEventFuncRequest()

	w := httptest.NewRecorder()
	f.h.ServeHTTP(w, r)

	var resp multipartestutils.TestEventResponse
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Empty(t, resp.PageTitle)
	assert.False(t, resp.Reload)
	assert.NotNil(t, resp.PushState)
	assert.False(t, resp.PushState.MyMergeQuery)
	assert.Empty(t, resp.PushState.MyURL)
	assert.Empty(t, resp.PushState.MyStringQuery)
	assert.Empty(t, resp.PushState.MyClearMergeQueryKeys)
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
