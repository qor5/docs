package examples_admin

// @snippet_begin(LoginBasicUsage)
import (
	"net/http"
	"os"

	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	plogin "github.com/qor5/admin/v3/login"
	"github.com/qor5/admin/v3/presets"
	. "github.com/qor5/ui/v3/vuetify"
	"github.com/qor5/web/v3"
	"github.com/qor5/x/v3/login"
	. "github.com/theplant/htmlgo"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name    string
	Address string

	login.UserPass
	login.OAuthInfo
	login.SessionSecure
}

func serve() {
	DB := ExampleDB()

	pb := presets.New()
	lb := plogin.New(pb).
		DB(DB).
		UserModel(&User{}).
		Secret(os.Getenv("LOGIN_SECRET")).
		OAuthProviders(
			&login.Provider{
				Goth: google.New(os.Getenv("LOGIN_GOOGLE_KEY"), os.Getenv("LOGIN_GOOGLE_SECRET"), os.Getenv("BASE_URL")+"/auth/callback?provider=google"),
				Key:  "google",
				Text: "Google",
			},
			&login.Provider{
				Goth: github.New(os.Getenv("LOGIN_GITHUB_KEY"), os.Getenv("LOGIN_GITHUB_SECRET"), os.Getenv("BASE_URL")+"/auth/callback?provider=github"),
				Key:  "github",
				Text: "Login with Github",
			},
		)
	pb.ProfileFunc(func(ctx *web.EventContext) HTMLComponent {
		return A(Text("logout")).Href(lb.LogoutURL)
	})

	r := http.NewServeMux()
	r.Handle("/", pb)
	lb.Mount(r)

	mux := http.NewServeMux()
	mux.Handle("/", lb.Middleware()(r))
	http.ListenAndServe(":8080", nil)
}

// @snippet_end

func loginPieces() {
	// @snippet_begin(LoginEnableUserPass)
	type User struct {
		gorm.Model

		login.UserPass
	}
	// @snippet_end

	var loginBuilder *login.Builder
	var count int
	// @snippet_begin(LoginSetMaxRetryCount)
	loginBuilder.MaxRetryCount(count)
	// @snippet_end

	var enable bool
	// @snippet_begin(LoginSetTOTP)
	loginBuilder.TOTP(enable, login.TOTPConfig{
		Issuer: "Issuer",
	})
	// @snippet_end

	// @snippet_begin(LoginSetRecaptcha)
	loginBuilder.Recaptcha(enable, login.RecaptchaConfig{
		SiteKey:   "SiteKey",
		SecretKey: "SecretKey",
	})
	// @snippet_end
}

func loginPiece2() {
	// @snippet_begin(LoginEnableOAuth)
	type User struct {
		gorm.Model

		login.OAuthInfo
	}
	// @snippet_end
}

func loginPiece3() {
	// @snippet_begin(LoginEnableSessionSecure)
	type User struct {
		gorm.Model

		login.UserPass
		login.OAuthInfo
		login.SessionSecure
	}
	// @snippet_end
}

func loginPiece4() {
	var loginBuilder *login.Builder
	// @snippet_begin(LoginCustomizePage)
	loginBuilder.LoginPageFunc(func(ctx *web.EventContext) (r web.PageResponse, err error) {
		r.Body = Text("This is login page")
		return
	})
	// @snippet_end

	var mux *http.ServeMux
	var loginPage http.Handler

	// @snippet_begin(LoginCustomizePage2)
	loginBuilder.LoginPageURL("/custom-login-page")
	loginBuilder.MountAPI(mux)
	mux.Handle("/custom-login-page", loginPage)
	// @snippet_end
}

func loginPiece5() {
	// @snippet_begin(LoginOpenChangePasswordDialog)
	VBtn("Change Password").OnClick(plogin.OpenChangePasswordDialogEvent)
	// @snippet_end

	var userModelBuilder *presets.ModelBuilder
	// @snippet_begin(LoginChangePasswordInEditing)
	userModelBuilder.Editing().Field("Password").
		SetterFunc(func(obj interface{}, field *presets.FieldContext, ctx *web.EventContext) (err error) {
			u := obj.(*User)
			if v := ctx.R.FormValue(field.Name); v != "" {
				u.Password = v
				u.EncryptPassword()
			}
			return nil
		})
	// @snippet_end
}
