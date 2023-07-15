module github.com/yyt030/myfaker

go 1.20

require (
	github.com/go-ini/ini v1.67.0
	github.com/go-sql-driver/mysql v1.7.1
	github.com/gosuri/uiprogress v0.0.1
	github.com/hashicorp/go-version v1.2.0
	github.com/icrowley/fake v0.0.0-20221112152111-d7b7e2276db2
	github.com/kr/pretty v0.3.1
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.9.3
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)

replace github.com/icrowley/fake v0.0.0-20221112152111-d7b7e2276db2 => github.com/yyt030/fake v0.0.0-20230715164623-8e2e405ec08c

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137 // indirect
	github.com/corpix/uarand v0.2.0 // indirect
	github.com/gosuri/uilive v0.0.4 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
)
