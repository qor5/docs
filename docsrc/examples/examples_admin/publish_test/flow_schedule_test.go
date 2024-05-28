package publish_test

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/qor5/admin/v3/utils/testflow"
	"github.com/qor5/docs/v3/docsrc/examples/examples_admin"
	"github.com/qor5/web/v3/multipartestutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/theplant/gofixtures"
)

var dataSeedForFlowSchedule = gofixtures.Data(gofixtures.Sql(`
INSERT INTO "public"."with_publish_products" ("id", "created_at", "updated_at", "deleted_at", "name", "price", "status", "online_url", "scheduled_start_at", "scheduled_end_at", "actual_start_at", "actual_end_at", "version", "version_name", "parent_version") VALUES ('6', '2024-05-28 06:42:41.620394+00', '2024-05-28 06:42:41.620394+00', NULL, 'FirstProduct', '456', 'draft', '', NULL, NULL, NULL, NULL, '2024-05-28-v01', '2024-05-28-v01', '');
`, []string{"with_publish_products"}))

type FlowSchedule struct {
	*Flow

	ID               string
	ScheduledStartAt time.Time
	ScheduledEndAt   time.Time
}

func TestFlowSchedule(t *testing.T) {
	// TODO: Missing some expect definitions, need to optimize as needed, current logic definitely needs optimization, not yet confirmed
	testCases := []struct {
		desc    string
		startAt time.Time
		endAt   time.Time
	}{
		{
			desc:    "empty",
			startAt: time.Time{},
			endAt:   time.Time{},
		},
		{
			desc:    "start < end < now",
			startAt: time.Now().AddDate(0, 0, -3),
			endAt:   time.Now().AddDate(0, 0, -2),
		},
		{
			desc:    "now < start < end",
			startAt: time.Now().AddDate(0, 0, 2),
			endAt:   time.Now().AddDate(0, 0, 3),
		},
		{
			desc:    "start < now < end",
			startAt: time.Now().AddDate(0, 0, -1),
			endAt:   time.Now().AddDate(0, 0, -1),
		},
		{
			desc:    "end < start < now",
			startAt: time.Now().AddDate(0, 0, -2),
			endAt:   time.Now().AddDate(0, 0, -3),
		},
		{
			desc:    "now < start < no end",
			startAt: time.Now().AddDate(0, 0, 2),
			endAt:   time.Time{},
		},
		{
			desc:    "now < no start < end",
			startAt: time.Time{},
			endAt:   time.Now().AddDate(0, 0, 2),
		},
	}
	for _, c := range testCases {
		t.Run(c.desc, func(t *testing.T) {
			dataSeedForFlowSchedule.TruncatePut(SQLDB)
			flowSchedule(t, &FlowSchedule{
				Flow:             &Flow{db: DB, h: PresetsBuilder},
				ID:               "6_2024-05-28-v01",
				ScheduledStartAt: c.startAt,
				ScheduledEndAt:   c.endAt,
			})
		})
	}
}

func scheduledTimeFormat(t *time.Time) string {
	if t == nil || t.IsZero() {
		return ""
	}
	return t.Format("2006-01-02 15:04") // minute
}

func flowSchedule(t *testing.T, f *FlowSchedule) {
	id, ver := MustIDVersion(f.ID)
	db := f.db.Where("id = ? AND version = ?", id, ver)

	var prev examples_admin.WithPublishProduct
	require.NoError(t, db.First(&prev).Error)

	prevScheduledStartAt := scheduledTimeFormat(prev.ScheduledStartAt)
	prevScheduledEndAt := scheduledTimeFormat(prev.ScheduledEndAt)

	// Ensure the schedule button is displayed; if not, it implies that the state is incorrect
	flowSchedule_Step00_Event_presets_DetailingDrawer(t, f).ThenValidate(testflow.ContainsInOrderAtUpdatePortal(0, "publish_eventSchedulePublishDialog"))

	flowSchedule_Step01_Event_publish_eventSchedulePublishDialog(t, f).ThenValidate(
		testflow.ContainsInOrderAtUpdatePortal(0,
			fmt.Sprintf(`[form, {"ScheduledStartAt":%q}]`, prevScheduledStartAt), // Ensure the original planned time is correct
			fmt.Sprintf(`[form, {"ScheduledEndAt":%q}]`, prevScheduledEndAt),
			fmt.Sprintf(`.query("id", %q)`, f.ID), // Ensure the correct project is being operated on
		),
	)
	flowSchedule_Step02_Event_publish_eventSchedulePublish(t, f)
	{
		var m examples_admin.WithPublishProduct
		require.NoError(t, db.First(&m).Error)
		assert.Equal(t, scheduledTimeFormat(&f.ScheduledStartAt), scheduledTimeFormat(m.ScheduledStartAt))
		assert.Equal(t, scheduledTimeFormat(&f.ScheduledEndAt), scheduledTimeFormat(m.ScheduledEndAt))
		// TODO: What about the impact of resources not in draft state?
		// TODO: Ensure no other data has been modified
		prev = m
	}
}

func flowSchedule_Step00_Event_presets_DetailingDrawer(t *testing.T, f *FlowSchedule) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products").
		EventFunc("presets_DetailingDrawer").
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
		testflow.OpenRightDrawer("WithPublishProduct "+f.ID),
	)

	return testflow.NewThen(t, w, r)
}

func flowSchedule_Step01_Event_publish_eventSchedulePublishDialog(t *testing.T, f *FlowSchedule) *testflow.Then {
	r := multipartestutils.NewMultipartBuilder().
		PageURL("/samples/publish/with-publish-products").
		EventFunc("publish_eventSchedulePublishDialog").
		Query("id", f.ID).
		Query("overlay", "dialog").
		AddField("ScheduledStartAt", scheduledTimeFormat(&f.ScheduledStartAt)).
		AddField("ScheduledEndAt", scheduledTimeFormat(&f.ScheduledEndAt)).
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
		PageURL("/samples/publish/with-publish-products").
		EventFunc("publish_eventSchedulePublish").
		Query("id", f.ID).
		Query("overlay", "dialog").
		AddField("ScheduledStartAt", scheduledTimeFormat(&f.ScheduledStartAt)).
		AddField("ScheduledEndAt", scheduledTimeFormat(&f.ScheduledEndAt)).
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
