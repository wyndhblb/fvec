//go:generate python3 genproto.py
//go:generate protoc --proto_path=vendor/:. --gogofaster_out=. vepr.proto
//go:generate protoc-go-inject-tag -input=./vepr.pb.go
//go:generate msgp -tests -o vepr_msgp.go --file vepr.pb.go
//go:generate msgp -tests -o tags_msgp.go --file tags.go
//go:generate msgp -tests -o name_list_msgp.go --file name.go
// XX:generate msgp -tests -o vepr_list_msgp.go --file vepr.go
// XX:generate easyjson -all vepr_list_msgp.go
//go:generate easyjson vepr.pb.go
//go:generate easyjson -all vepr.go
//go:generate gofmt -w -s .

package vepr
