# hellomock

install go mock tools
```bash
$ go get github.com/golang/mock/gomock
$ go install github.com/golang/mock/mockgen
```

mockgen
```bash
$ mockgen -source=hellomock.go > mock_hellomock/mock_talker.go
```

test compare
```bash
$ go test -v . -run TestCompany_Meeting
$ go test -v . -run TestCompany_Meeting2
```
