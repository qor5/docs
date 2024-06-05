package examples_admin

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/qor5/admin/v3/presets"
	"github.com/qor5/admin/v3/presets/gorm2op"
	"github.com/qor5/web/v3/multipartestutils"
	"github.com/theplant/gofixtures"
	"github.com/theplant/testingutils"
)

var l10nData = gofixtures.Data(gofixtures.Sql(`
INSERT INTO public.l10n_models (id, created_at, updated_at, deleted_at, title, locale_code) VALUES (1, 
'2024-06-04 23:27:40.442281 +00:00', '2024-06-04 23:27:40.442281 +00:00', null, 'My model title', 'International');


`, []string{"l10n_models"}))

func TestLocalization(t *testing.T) {
	pb := presets.New().DataOperator(gorm2op.DataOperator(TestDB))
	LocalizationExample(pb, TestDB)

	cases := []multipartestutils.TestCase{
		{
			Name:  "Index Page",
			Debug: true,
			ReqFunc: func() *http.Request {
				l10nData.TruncatePut(SqlDB)
				return httptest.NewRequest("GET", "/l10n-models", nil)
			},
			ExpectPageBodyContainsInOrder: []string{"My model title", "International"},
		},
		{
			Name:  "Localize dialog",
			Debug: true,
			ReqFunc: func() *http.Request {
				l10nData.TruncatePut(SqlDB)
				req := multipartestutils.NewMultipartBuilder().
					PageURL("/l10n-models?__execute_event__=l10n_LocalizeEvent&id=1_International").
					BuildEventFuncRequest()
				return req
			},
			ExpectPortalUpdate0ContainsInOrder: []string{"China", "Japan"},
		},
		{
			Name:  "Localize to China and Japan",
			Debug: true,
			ReqFunc: func() *http.Request {
				l10nData.TruncatePut(SqlDB)
				req := multipartestutils.NewMultipartBuilder().
					PageURL("/l10n-models?__execute_event__=l10n_DoLocalizeEvent&id=1_International&localize_from=International").
					AddField("localize_to", "China").
					AddField("localize_to", "Japan").
					BuildEventFuncRequest()
				return req
			},
			EventResponseMatch: func(t *testing.T, er *multipartestutils.TestEventResponse) {
				var localeCodes []string
				TestDB.Raw("SELECT locale_code FROM l10n_models ORDER BY locale_code").Scan(&localeCodes)
				if diff := testingutils.PrettyJsonDiff(
					[]string{"China", "International", "Japan"},
					localeCodes); diff != "" {
					t.Error(diff)
				}
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			multipartestutils.RunCase(t, c, pb)
		})
	}
}
