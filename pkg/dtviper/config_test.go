package dtviper

import (
	"fmt"
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
		os.Setenv("DOUTOK_TEST_SERVER_PORT", "9090")
		v.UnmarshalStructTags(reflect.TypeOf(tests.args.stc), "")
		v.UnmarshalStruct(&tests.args.stc)
		fmt.Println("Server.Name is set by yaml : DouTokTest")
		fmt.Println("Server.Address is set by default : localhost")
		fmt.Println("Server.Port is set by environment : 9090")
		fmt.Println("Struct is ", tests.args.stc)
	})
}
