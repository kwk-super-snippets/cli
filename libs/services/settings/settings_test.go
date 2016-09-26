package settings

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/smartystreets/assertions/should"
	"testing"
	"bitbucket.com/sharingmachine/kwkcli/libs/models"
	"bitbucket.com/sharingmachine/kwkcli/libs/services/system"
	"os"
)

func Test_System(t *testing.T) {
	Convey("Settings test", t, func() {
		Convey(`Should create, update, delete and get a setting`, func() {
			sys := system.New()
			s := New(sys, "test_settings")
			key := "user"
			expected := models.User{
				Email:"richard@kwk.co",
				Token:"asdfsdfsdfuiu",
			}
			s.Upsert(key, expected)
			actual := &models.User{}
			s.Get(key, actual)
			So(*actual, should.Resemble, expected)

			expectedUpdate := models.User{
				Email:"richard@kwk.io",
				Token:"asdfsdfsdfuiu",
			}
			s.Upsert(key, expectedUpdate)
			s.Get(key, actual)
			So(*actual, should.Resemble, expectedUpdate)

			err := s.Delete(key)
			So(err, ShouldBeNil)

			actual = &models.User{}
			err = s.Get(key, actual)
			So(err.Error(), should.Equal, "Not found.")
			p, _ := sys.GetDirPath("test_settings")
			os.RemoveAll(p)
		})
	})
}


