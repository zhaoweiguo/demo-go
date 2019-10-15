package main

import (
	"fmt"
	"github.com/drone/drone-yaml/yaml"
	"github.com/google/go-cmp/cmp"
)

func main() {
	tests := getData()
	for i, test := range tests {
		fmt.Println("====", i)
		got, err := yaml.ParseRawString(test.data)
		if err != nil {
			fmt.Println(err)
		}
		if diff := cmp.Diff(got, test.want); diff != "" {
			fmt.Println(diff)
		}
	}

}

func getData() []struct {
	data string
	want []*yaml.RawResource
} {

	tests := []struct {
		data string
		want []*yaml.RawResource
	}{
		//
		// empty document returns nil resources.
		//
		{
			data: "",
			want: nil,
		},
		//
		// single document files.
		//
		{
			data: "kind: pipeline\nfoo: bar",
			want: []*yaml.RawResource{
				{Kind: "pipeline", Data: []byte("kind: pipeline\nfoo: bar\n")},
			},
		},
		//
		// multi-document files.
		//
		{
			data: "kind: a\nb: c\n---\nkind: d\ne: f",
			want: []*yaml.RawResource{
				{Kind: "a", Data: []byte("kind: a\nb: c\n")},
				{Kind: "d", Data: []byte("kind: d\ne: f\n")},
			},
		},
	}
	return tests
}
