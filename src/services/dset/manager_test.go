package dset

import (
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestManager(t *testing.T) {
	convey.Convey("test set", t, func() {
		d := GetDistributedSet("test_key")
		convey.So(d.Key(), convey.ShouldEqual, "test_key")
		d.Add("1", "2")
		convey.So(len(d.Members()), convey.ShouldEqual, 0)
		d.Save(2 * time.Second)
		convey.So(d.Members(), convey.ShouldResemble, []string{"1", "2"})
		e := GetDistributedSet("test_key")
		convey.So(d.Members(), convey.ShouldResemble, e.Members())
		time.Sleep(3 * time.Second)
		f := GetDistributedSet("test_key")
		convey.So(len(f.Members()), convey.ShouldEqual, 0)
	})
}
