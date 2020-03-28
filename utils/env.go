package utils

import (
	"fmt"
	"reflect"

	"github.com/sirupsen/logrus"

	"github.com/deanishe/go-env"
)

type Environ struct {
	DBHost          string `env:"DB_HOST" def:""`
	DBUser          string `env:"DB_USER" def:"test"`
	DBPass          string `env:"DB_PASS" def:"falsepassword"`
	DBDatabase      string `env:"DB_NAME" def:""`
	RedisHost       string `env:"REDIS_HOST" def:"localhost:6379"`
	RedisPass       string `env:"REDIS_PASS" def:""`
	AppHost         string `env:"APP_HOST" def:":8000"`
	AppSecret       string `env:"APP_SECRET" def:"fake secret"`
	GinMode         string `env:"GIN_MODE" def:"debug"`
	NewRelicLicense string `env:"NEWRELIC_KEY" def:""`
	NewRelicAppName string `env:"NEWRELIC_APPNAME" def:""`
	NewRelicVerbose string `env:"NEWRELIC_VERBOSE" def:"true"`
	Pprof           string `env:"PPROF" def:"true"`
	GCPCredential   string `env:"GOOGLE_APPLICATION_CREDENTIALS" def:""`
	GCPProjectID    string `env:"GCP_PROJECT_ID" def:""`
}

func InitEnv(log *logrus.Logger) Environ {
	var environ Environ

	if err := env.Bind(&environ); err == nil {
		t := environ

		t2 := Environ{}
		s := reflect.ValueOf(&t).Elem()
		s2 := reflect.ValueOf(&t2).Elem()
		typeOfT := s.Type()
	
		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			f2 := s2.Field(i)
			f2.Set(reflect.Value(f))

			if f.Interface() == "" {
				log.Warn(fmt.Sprintf("No Environment Variables found for %s, using default value:%s", typeOfT.Field(i).Tag.Get("env"), typeOfT.Field(i).Tag.Get("def")))
				f2.SetString(typeOfT.Field(i).Tag.Get("def"))
			}

		}

		return t2
	}
	
	return environ
}