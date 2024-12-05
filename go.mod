module github.com/zhaoweiguo/demo-go

go 1.12

require (
	docker.io/go-docker v1.0.0
	git.apache.org/thrift.git v0.12.0
	github.com/BurntSushi/toml v0.3.1
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/Shopify/sarama v1.22.1
	github.com/aliyun/aliyun-log-go-sdk v0.1.12
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/codegangsta/cli v1.20.0
	github.com/coreos/bbolt v1.3.3 // indirect
	github.com/coreos/etcd v3.3.18+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dailymotion/allure-go v0.5.5
	github.com/davecgh/go-spew v1.1.1
	github.com/docker/docker v1.13.1
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/drone/drone-runtime v0.0.0-20190210191445-ad403a0ca24e
	github.com/drone/drone-yaml v1.0.6
	github.com/drone/go-scm v1.6.1-0.20190930172114-d8fff5ce7761
	github.com/drone/signal v1.0.0
	github.com/fatih/structs v1.1.0 // indirect
	github.com/gavv/httpexpect/v2 v2.1.0
	github.com/gin-contrib/sse v0.0.0-20190301062529-5545eab6dad3 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-ini/ini v1.39.0
	github.com/go-kit/kit v0.9.0
	github.com/go-redis/redis/v7 v7.0.0-beta.4
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.5.5
	github.com/google/wire v0.3.0
	github.com/googleapis/gnostic v0.2.0 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20200217142428-fce0ec30dd00 // indirect
	github.com/gorilla/mux v1.7.2
	github.com/h2non/gock v1.0.15 // indirect
	github.com/hashicorp/go-retryablehttp v0.6.2
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/influxdata/influxdb v1.7.6
	github.com/influxdata/influxdb1-client v0.0.0-20190809212627-fc22c7df067e
	github.com/jinzhu/configor v1.1.1
	github.com/jmoiron/sqlx v1.2.0
	github.com/joho/godotenv v1.3.0
	github.com/jroimartin/gocui v0.4.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/lextoumbourou/goodhosts v2.1.0+incompatible
	github.com/lib/pq v1.3.0
	github.com/lucas-clemente/quic-go v0.25.0
	github.com/mailgun/holster v3.0.0+incompatible // indirect
	github.com/mailgun/iptools v0.0.0-20170310010557-ba8d5743f678 // indirect
	github.com/mailgun/log v0.0.0-20150926000944-2f35a4607f1a // indirect
	github.com/mailgun/metrics v0.0.0-20170714162148-fd99b46995bd // indirect
	github.com/mailgun/scroll v1.2.2
	github.com/natessilva/dag v0.0.0-20180124060714-7194b8dcc5c4 // indirect
	github.com/nsf/termbox-go v0.0.0-20201124104050-ed494de23a00 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/opentracing/opentracing-go v1.1.0
	github.com/rs/zerolog v1.18.0
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.4.2
	github.com/smartystreets/goconvey v0.0.0-20190731233626-505e41936337
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.3.2
	github.com/stretchr/testify v1.5.1
	github.com/tal-tech/go-zero v1.1.7
	github.com/tealeg/xlsx v1.0.3
	github.com/tidwall/pretty v1.0.0 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20190109142713-0ad062ec5ee5 // indirect
	github.com/uber/jaeger-client-go v2.22.1+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	github.com/ugorji/go/codec v0.0.0-20181209151446-772ced7fd4c2 // indirect
	github.com/urfave/cli v1.22.5
	github.com/vardius/gollback v1.1.0
	github.com/vinzenz/yaml v0.0.0-20170920082545-91409cdd725d
	go.mongodb.org/mongo-driver v1.0.3
	golang.org/x/net v0.7.0
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4
	golang.org/x/text v0.7.0
	google.golang.org/appengine v1.6.1 // indirect
	gopkg.in/ahmetb/go-linq.v3 v3.0.0 // indirect
	gopkg.in/airbrake/gobrake.v2 v2.0.9 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15
	gopkg.in/gemnasium/logrus-airbrake-hook.v2 v2.1.2
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/ini.v1 v1.48.0 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.18.5
	k8s.io/apimachinery v0.18.5
	k8s.io/client-go v0.18.5

)

replace github.com/h2non/gock => gopkg.in/h2non/gock.v1 v1.0.14
