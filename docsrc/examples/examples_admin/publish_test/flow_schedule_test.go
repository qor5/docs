package publish_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/qor5/admin/v3/utils/testflow"
	"github.com/qor5/web/v3/multipartestutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var dateTimeLayout = "2006-01-02 15:04:05"

type FlowSchedule struct {
	*Flow

	// local vars
	ScheduledStartAt time.Time
	ScheduledEndAt   time.Time
}

func TestFlowSchedule(t *testing.T) {
	dataSeed.TruncatePut(SQLDB)

	flowSchedule(t, &FlowSchedule{
		Flow: &Flow{
			db: DB, h: PresetsBuilder,
			ID: "6_2024-05-22-v01",
		},
	})
}

func flowSchedule(t *testing.T, f *FlowSchedule) {
	flowSchedule_Step00_Event_presets_DetailingDrawer(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	// start < end < now
	f.ScheduledStartAt = time.Now().AddDate(0, 0, -3)
	f.ScheduledEndAt = time.Now().AddDate(0, 0, -2)
	flowSchedule_Step01_Event_publish_eventSchedulePublishDialog(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	flowSchedule_Step02_Event_publish_eventSchedulePublish(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	// now < start < end
	f.ScheduledStartAt = time.Now().AddDate(0, 0, 2)
	f.ScheduledEndAt = time.Now().AddDate(0, 0, 3)
	flowSchedule_Step03_Event_publish_eventSchedulePublishDialog(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	flowSchedule_Step04_Event_publish_eventSchedulePublish(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	// end < start < now
	f.ScheduledStartAt = time.Now().AddDate(0, 0, -2)
	f.ScheduledEndAt = time.Now().AddDate(0, 0, -3)
	flowSchedule_Step05_Event_publish_eventSchedulePublishDialog(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})

	// TODO: 目前是通过的，应该报错才对
	flowSchedule_Step06_Event_publish_eventSchedulePublish(t, f).Then(func(t *testing.T, w *httptest.ResponseRecorder, r *http.Request) {
		// assert.Contains(t, w.Body.String(), "xx")
	})
}

func flowSchedule_Step00_Event_presets_DetailingDrawer(t *testing.T, f *FlowSchedule) *testflow.Then {
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

func flowSchedule_Step01_Event_publish_eventSchedulePublishDialog(t *testing.T, f *FlowSchedule) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("publish_eventSchedulePublishDialog").
		// Query("id", "6_2024-05-22-v01").
		Query("id", f.ID).
		Query("overlay", "dialog").
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
	assert.Equal(t, "publish_PortalSchedulePublishDialog", resp.UpdatePortals[0].Name)
	assert.Nil(t, resp.Data)
	assert.Empty(t, resp.RunScript)

	return testflow.NewThen(t, w, r)
}

func flowSchedule_Step02_Event_publish_eventSchedulePublish(t *testing.T, f *FlowSchedule) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("publish_eventSchedulePublish").
		// Query("id", "6_2024-05-22-v01").
		Query("id", f.ID).
		Query("overlay", "dialog").
		// AddField("ScheduledEndAt", "2024-05-22 00:00").
		AddField("ScheduledEndAt", f.ScheduledEndAt.Format(dateTimeLayout)).
		// AddField("ScheduledStartAt", "2024-05-21 00:00").
		AddField("ScheduledStartAt", f.ScheduledStartAt.Format(dateTimeLayout)).
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
	assert.Equal(t, "publish_PortalSchedulePublishDialog", resp.UpdatePortals[0].Name)
	assert.Nil(t, resp.Data)
	assert.Empty(t, resp.RunScript)

	return testflow.NewThen(t, w, r)
}

func flowSchedule_Step03_Event_publish_eventSchedulePublishDialog(t *testing.T, f *FlowSchedule) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("publish_eventSchedulePublishDialog").
		// Query("id", "6_2024-05-22-v01").
		Query("id", f.ID).
		Query("overlay", "dialog").
		// AddField("ScheduledEndAt", "2024-05-22 00:00").
		AddField("ScheduledEndAt", f.ScheduledEndAt.Format(dateTimeLayout)).
		// AddField("ScheduledStartAt", "2024-05-21 00:00").
		AddField("ScheduledStartAt", f.ScheduledStartAt.Format(dateTimeLayout)).
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
	assert.Equal(t, "publish_PortalSchedulePublishDialog", resp.UpdatePortals[0].Name)
	assert.Nil(t, resp.Data)
	assert.Empty(t, resp.RunScript)

	return testflow.NewThen(t, w, r)
}

func flowSchedule_Step04_Event_publish_eventSchedulePublish(t *testing.T, f *FlowSchedule) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("publish_eventSchedulePublish").
		// Query("id", "6_2024-05-22-v01").
		Query("id", f.ID).
		Query("overlay", "dialog").
		// AddField("ScheduledEndAt", "2024-05-27 00:00").
		AddField("ScheduledEndAt", f.ScheduledEndAt.Format(dateTimeLayout)).
		// AddField("ScheduledStartAt", "2024-05-26 00:00").
		AddField("ScheduledStartAt", f.ScheduledStartAt.Format(dateTimeLayout)).
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
	assert.Equal(t, "publish_PortalSchedulePublishDialog", resp.UpdatePortals[0].Name)
	assert.Nil(t, resp.Data)
	assert.Empty(t, resp.RunScript)

	return testflow.NewThen(t, w, r)
}

func flowSchedule_Step05_Event_publish_eventSchedulePublishDialog(t *testing.T, f *FlowSchedule) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("publish_eventSchedulePublishDialog").
		// Query("id", "6_2024-05-22-v01").
		Query("id", f.ID).
		Query("overlay", "dialog").
		// AddField("ScheduledEndAt", "2024-05-27 00:00").
		AddField("ScheduledEndAt", f.ScheduledEndAt.Format(dateTimeLayout)).
		// AddField("ScheduledStartAt", "2024-05-26 00:00").
		AddField("ScheduledStartAt", f.ScheduledStartAt.Format(dateTimeLayout)).
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
	assert.Equal(t, "publish_PortalSchedulePublishDialog", resp.UpdatePortals[0].Name)
	assert.Nil(t, resp.Data)
	assert.Empty(t, resp.RunScript)

	return testflow.NewThen(t, w, r)
}

func flowSchedule_Step06_Event_publish_eventSchedulePublish(t *testing.T, f *FlowSchedule) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products-version-list-dialog").
		EventFunc("publish_eventSchedulePublish").
		// Query("id", "6_2024-05-22-v01").
		Query("id", f.ID).
		Query("overlay", "dialog").
		// AddField("ScheduledEndAt", "2024-05-21 00:00").
		AddField("ScheduledEndAt", f.ScheduledEndAt.Format(dateTimeLayout)).
		// AddField("ScheduledStartAt", "2024-05-22 00:00").
		AddField("ScheduledStartAt", f.ScheduledStartAt.Format(dateTimeLayout)).
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
	assert.Equal(t, "publish_PortalSchedulePublishDialog", resp.UpdatePortals[0].Name)
	assert.Nil(t, resp.Data)
	assert.Empty(t, resp.RunScript)

	return testflow.NewThen(t, w, r)
}
