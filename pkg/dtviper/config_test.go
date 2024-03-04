package dtviper

import (
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/spf13/viper"
	"os"
	"reflect"
	"testing"
)

type TestConfig struct {
	Server configStruct.Base
}

func TestConfig_UnmarshalStruct(t *testing.T) {
	type fields struct {
		Viper *viper.Viper
	}
	type args struct {
		stc any
	}
	tests := struct {
		name   string
		fields fields
		args   args
	}{
		name: "DOUTOK_TEST",
		fields: fields{
			Viper: ConfigInit("DOUTOK_TEST", "test").Viper,
		},
		args: args{
			stc: TestConfig{},
		},
	}
	t.Run(tests.name, func(t *testing.T) {
		v := &Config{
			Viper: tests.fields.Viper,
		}
		err := os.Setenv("DOUTOK_TEST_SERVER_PORT", "9090")
		if err != nil {
			t.Error("设置环境变量失败")
			return
		}
		v.UnmarshalStructTags(reflect.TypeOf(tests.args.stc), "")
		v.UnmarshalStruct(&tests.args.stc)
		config, ok := tests.args.stc.(TestConfig)
		if !ok {
			t.Error("类型转换失败")
			return
		}
		if config.Server.Name != "DouTokTest" {
			t.Errorf("Server.Name is %s, expect %s", config.Server.Name, "DouTokTest")
		}
		if config.Server.Address != "localhost" {
			t.Errorf("Server.Address is %s, expect %s", config.Server.Address, "localhost")
		}
		if config.Server.Port != 9090 {
			t.Errorf("Server.Port is %d, expect %d", config.Server.Port, 9090)
		}
	})
}
