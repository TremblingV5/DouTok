package configurator

import (
	"context"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/TremblingV5/DouTok/pkg/dtviper"
)

func TestLoadFromFiles(t *testing.T) {
	v := dtviper.ConfigInit("DOUTOK_UNIT_TEST", "test", "./")

	err := loadFromFile(context.Background(), 6, v)
	require.EqualError(t, err, ErrNotPtrOfStruct)

	type config struct {
		Name  string `configPath:"Path"`
		Embed struct {
			Age int `configPath:"Path1.Path2"`
		}
	}

	cfg := &config{}
	err = loadFromFile(context.Background(), cfg, v)
	require.NoError(t, err)
	require.Equal(t, cfg.Name, "test")
	require.Equal(t, cfg.Embed.Age, 12)
}

func TestLoadFromEnv(t *testing.T) {
	type config struct {
		Name  string `env:"Name"`
		Embed struct {
			Age int `env:"Age"`
		}
	}

	os.Setenv("Name", "name") //nolint
	os.Setenv("Age", "12")    //nolint

	cfg := &config{}
	err := loadFromEnv(context.Background(), cfg)
	require.NoError(t, err)
	require.Equal(t, cfg.Name, "name")
	require.Equal(t, cfg.Embed.Age, 12)
}

// Testing for confirm that loading configuration from files have a higher priority.
func TestLoad(t *testing.T) {
	type config struct {
		Name  string `env:"Name" configPath:"Path"`
		Embed struct {
			Age int `env:"Age" configPath:"Path1.Path2"`
		}
	}

	os.Setenv("Name", "name") //nolint

	cfg := &config{}
	err := Load(context.Background(), cfg, "DOUTOK_UNIT_TEST", "test")
	require.NoError(t, err)
	require.Equal(t, cfg.Name, "test")
	require.Equal(t, cfg.Embed.Age, 12)
}

func TestInitConfigWithNonExistFileName(t *testing.T) {
	configPath, err := getConfigFilesPath("Non-Exist")
	require.ErrorContains(t, err, constants.ErrConfigFileNotFound)
	require.Equal(t, configPath, "")
}
