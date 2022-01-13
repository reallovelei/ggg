package id

import (
	tests "github.com/reallovelei/ggg/test"
	"testing"

	"github.com/reallovelei/ggg/framework/contract"
	"github.com/reallovelei/ggg/framework/provider/config"
	. "github.com/smartystreets/goconvey/convey"
)

func TestConsoleLog_Normal(t *testing.T) {
	Convey("test ggg console log normal case", t, func() {
		c := tests.InitBaseContainer()
		c.Bind(&config.GGGConfigProvider{})

		err := c.Bind(&GGGIDProvider{})
		So(err, ShouldBeNil)

		idService := c.MustMake(contract.IDKey).(contract.IDService)
		xid := idService.NewID()
		So(xid, ShouldNotBeEmpty)
	})
}
