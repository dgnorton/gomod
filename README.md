# gomod
Tool for diffing dep Gopkg.lock and Go modules go.sum files

## Usage
```
go build ./cmd/difflocksum
./difflocksum <path/Gopkg.lock> <path/go.sum>
```

### Example output snippet
```
Dependency                                         LockVer  ModVer
cloud.google.com/go                                v0.47.0  v0.51.0
github.com/andreyvit/diff                                   v0.0.0-20170406064948-c7f18ee00883/go.mod
github.com/bmizerany/pat                                    v0.0.0-20170815010413-6226ea591a40
github.com/prometheus/client_model                          v0.0.0-20190129233127-fd36f4220a90/go.mod
go.uber.org/atomic                                 v1.3.1   v1.3.2/go.mod
gonum.org/v1/netlib                                         v0.0.0-20181029234149-ec6d1f5cefe6/go.mod
google.golang.org/appengine                        v1.2.0   v1.4.0/go.mod
github.com/golang/glog                                      v0.0.0-20160126235308-23def4e6c14b/go.mod
github.com/google/go-cmp                           v0.2.0   v0.2.0/go.mod
google.golang.org/genproto                                  v0.0.0-20190418145605-e7d98fc518a7/go.mod
```
