package eav_test

import (
	"testing"

	. "clickyab.com/exchange/services/eav"
	"clickyab.com/exchange/services/eav/mock"

	"time"

	"reflect"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	Register(mock.NewMockStore)
	Convey("Test keyvalue store", t, func() {
		store := NewEavStore("test_key")
		So(store.Key(), ShouldEqual, "test_key")
		Convey("check set and get", func() {
			store.SetSubKey("test", "test_val")
			So(store.SubKey("test"), ShouldEqual, "test_val")
			store.SetSubKey("another", "2")
			So(store.SubKey("another"), ShouldEqual, "2")
			So(store.Save(time.Hour), ShouldBeNil)
			So(reflect.DeepEqual(store.AllKeys(), map[string]string{"test": "test_val", "another": "2"}), ShouldBeTrue)
		})
	})
}