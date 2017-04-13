package update

import (
	gu "github.com/inconshreveable/go-update"
	"bitbucket.com/sharingmachine/kwkcli/models"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"strings"
	"testing"
	"errors"
	"io"
	"bitbucket.com/sharingmachine/kwkcli/persist"
)

func Test_Update(t *testing.T) {
	Convey("Runner test", t, func() {
		am := &ApplierMock{}
		rm := &RemoterMock{}
		pm := &persist.PersisterMock{}
		models.SetPrefs(models.DefaultPrefs())
		models.Client.Version = "v0.0.1"

		r := Runner{
			Applier:    am.Apply,
			Rollbacker: am.RollbackError,
			Remoter:    rm,
			Persister: pm,
		}

		Convey("Given there is an update record should NOT run update.", func() {
			rm.BI = ReleaseInfo{Version: "v0.0.2"}
			pm.GetHydrates = &Record{}
			err := r.Run()
			So(err, ShouldBeNil)
			So(rm.LatestCalled, ShouldBeFalse)
			So(am.ApplyCalledWith, ShouldBeNil)
		})

		Convey(`Given the current version is equal to the latest version should NOT run update.`, func() {
			rm.BI = ReleaseInfo{Version: "v0.0.1"}
			pm.GetHydrates = &Record{}
			err := r.Run()
			So(err, ShouldBeNil)
			So(rm.LatestCalled, ShouldBeFalse)
			So(am.ApplyCalledWith, ShouldBeNil)
		})

		Convey("Given there is NOT an update record and the remote version is newer should update.", func() {
			rm.BI = ReleaseInfo{Version: "v0.0.2"}
			pm.GetReturns = persist.ErrFileNotFound
			err := r.Run()
			So(err, ShouldBeNil)
			So(rm.LatestCalled, ShouldBeTrue)
			So(am.ApplyCalledWith, ShouldNotBeNil)
		})

		Convey(`When updating Given the applier returns an error should rollback.`, func() {
			rm.BI = ReleaseInfo{Version: "v0.0.2"}
			pm.GetReturns = persist.ErrFileNotFound
			m := "Couldn't apply."
			am.ApplyErr = errors.New(m)
			err := r.Run()
			So(err.Error(), ShouldEqual, m)
			So(am.RollbackCalledWith.Error(), ShouldEqual, m)
		})

		// //TODO: Run on ad-hoc basis
		 Convey(`Test remoter info and bin downloader`, func() {
			//r := S3Remoter{}
			//ri, err := r.LatestInfo()
			//So(err, ShouldBeNil)
			//So(ri.Version, ShouldEqual, "1.2.3")
			//So(ri.Build, ShouldEqual, "12")
			//So(ri.Time, ShouldEqual, 233423423)
			//So(ri.Notes, ShouldResemble, "Feature A\nFeature B\n")
			//rdr, err := r.LatestBinary()
			//So(err, ShouldBeNil)
			//out, err := os.Create("kwk")
			//So(err, ShouldBeNil)
			//io.Copy(out, rdr)
		})

	})
}

type ApplierMock struct {
	ApplyCalledWith []interface{}
	RollbackCalledWith error
	ApplyErr        error
}

func (am *ApplierMock) Apply(update io.Reader, opts gu.Options) error {
	am.ApplyCalledWith = []interface{}{update, opts}
	if am.ApplyErr != nil {
		return am.ApplyErr
	}
	return nil
}

func (am *ApplierMock) RollbackError(err error) error {
	am.RollbackCalledWith = err
	return err
}

type RemoterMock struct {
	BI                ReleaseInfo
	LatestCalled      bool
	ReleaseInfoCalled bool
}

func (rm *RemoterMock) LatestBinary() (io.ReadCloser, error) {
	rm.LatestCalled = true
	r := strings.NewReader("This is the binary")
	return ioutil.NopCloser(r), nil
}

func (rm *RemoterMock) LatestInfo() (*ReleaseInfo, error) {
	rm.ReleaseInfoCalled = true
	return &rm.BI, nil
}

func (rm *RemoterMock) CleanUp() {

}
