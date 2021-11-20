package split

import (
	"reflect"
	"strconv"
	"testing"
)
//func TestSplit(t *testing.T)  {
//	got := split("fsfexsfg","f")
//	res := []string{"s","exs","g"}
//	if !reflect.DeepEqual(got,res){
//		fmt.Println()
//		t.Error("res:",res,"\n","got:",got)
//	}
//}

func TestSplit(t *testing.T)  {
	type test struct {
		input string
		sep string
		want []string
	}

	tests := []test{
		{
			input: "adfwfgwe",
			sep: "f",
			want: []string{"ad","w","gwe"},
		},
		{
			input: "rgewfqef",
			sep: "f",
			want: []string{"rgew","qe"},
		},
	}
	for name,te := range tests{
		//fmt.Println(name)
		t.Run(strconv.Itoa(name), func(t *testing.T) {
			got := split(te.input,te.sep)
			if !reflect.DeepEqual(got,te.want){
				t.Error("index:",name,"res:",te.want,"\n","got:",got)
			}
		})

	}
}