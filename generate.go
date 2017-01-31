//go:generate python3 genproto.py
//go:generate protoc --proto_path=vendor/:. --gogofaster_out=. fvec.proto
//go:generate protoc-go-inject-tag -input=./fvec.pb.go
//go:generate msgp -tests -o fvec_msgp.go --file fvec.pb.go
//go:generate msgp -tests -o tags_msgp.go --file tags.go
//go:generate msgp -tests -o name_list_msgp.go --file name.go
// XX:generate msgp -tests -o fvec_list_msgp.go --file fvec.go
// XX:generate easyjson -all fvec_list_msgp.go
//go:generate easyjson fvec.pb.go
//go:generate easyjson fvec.go
//go:generate gofmt -w -s .

package fvec
