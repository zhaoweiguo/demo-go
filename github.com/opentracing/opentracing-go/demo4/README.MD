

# formatter
$ go run ./lesson04/exercise/formatter/formatter.go
2017/09/25 17:43:43 Initializing logging reporter
2017/09/25 17:44:02 Reporting span 719c2d0b77869cb3:79e7ce843e340d93:536316b3383042c8:1

# publisher
$ go run ./lesson04/exercise/publisher/publisher.go
2017/09/25 17:43:46 Initializing logging reporter
Bonjour, Bryan!
2017/09/25 17:44:02 Reporting span 719c2d0b77869cb3:64846658fbbf5e3e:4970dc35776b6b8a:1

# client
$ go run ./lesson04/exercise/client/hello.go Bryan Bonjour
2017/09/25 17:44:02 Initializing logging reporter
2017/09/25 17:44:02 Reporting span 719c2d0b77869cb3:536316b3383042c8:719c2d0b77869cb3:1
2017/09/25 17:44:02 Reporting span 719c2d0b77869cb3:4970dc35776b6b8a:719c2d0b77869cb3:1
2017/09/25 17:44:02 Reporting span 719c2d0b77869cb3:719c2d0b77869cb3:0:1



