//go:generate ../fvecgen/fvecgen --classname=GenTT --pkgname=main --format=col1=s,col2=mssi,col3=c,col4=f,col5=si,col6=msss,col7=li,col8=b,col9=ss,col10=msd,col11=c,col12=sid > tester.go

package main

import (
	"fmt"
	"strings"
)

func main(){

	cls := NewGenTT()
	fmt.Println(strings.Join(cls.CassandraCreateStatement("mykey", "mtable"), "\n"))
	fmt.Println(strings.Join(cls.CassandraSelectQueries("mtable"), "\n"))
}