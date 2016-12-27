package rpc

import (
	"bitbucket.com/sharingmachine/kwkcli/libs/models"
	"bitbucket.com/sharingmachine/kwkcli/libs/services/settings"
	"github.com/smartystreets/assertions/should"
	. "github.com/smartystreets/goconvey/convey"
	"google.golang.org/grpc/metadata"
	"testing"
)

func Test_RPC(t *testing.T) {
	Convey("RPC", t, func() {

		Convey(`Upgrade`, func() {
			Convey(`Given the current settings has a profile (signed in user) should add token to context`, func() {
				t := &settings.SettingsMock{}
				token := "sometoken234234234"
				t.GetHydrateWith = &models.User{Token: token}
				h := NewHeaders(t)
				md, _ := metadata.FromContext(h.GetContext())
				So(md[models.TokenHeaderName][0], should.Equal, token)
			})

			Convey(`Given the current settings does not have a profile (signed in user) should not add token to context`, func() {
				t := &settings.SettingsMock{}
				h := NewHeaders(t)
				md, _ := metadata.FromContext(h.GetContext())
				So(md[models.TokenHeaderName][0], should.BeEmpty)
			})
		})

	})
}
