//package redis
//
//import (
//	"context"
//	"github.com/reallovelei/ggg/framework/provider/config"
//	"github.com/reallovelei/ggg/framework/provider/log"
//	tests "github.com/reallovelei/ggg/test"
//	. "github.com/smartystreets/goconvey/convey"
//	"testing"
//	"time"
//)
//
//func TestGGGService_Load(t *testing.T) {
//	container := tests.InitBaseContainer()
//	container.Bind(&config.GGGConfigProvider{})
//	container.Bind(&log.GGGLogServiceProvider{})
//
//	Convey("test get client", t, func() {
//		GGGRedis, err := NewGGGRedis(container)
//		So(err, ShouldBeNil)
//		service, ok := GGGRedis.(*GGGRedis)
//		So(ok, ShouldBeTrue)
//		client, err := service.GetClient(WithConfigPath("redis.write"))
//		So(err, ShouldBeNil)
//		So(client, ShouldNotBeNil)
//		ctx := context.Background()
//		err = client.Set(ctx, "foo", "bar", 1*time.Hour).Err()
//		So(err, ShouldBeNil)
//		val, err := client.Get(ctx, "foo").Result()
//		So(err, ShouldBeNil)
//		So(val, ShouldEqual, "bar")
//		err = client.Del(ctx, "foo").Err()
//		So(err, ShouldBeNil)
//	})
//}
