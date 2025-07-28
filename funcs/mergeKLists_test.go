package funcs

import (
	"testing"
)

func TestMerge2List(t *testing.T) {

	n3 := &ListNode{Val: 5, Next: nil}
	n2 := &ListNode{Val: 4, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}

	n33 := &ListNode{Val: 4, Next: nil}
	n22 := &ListNode{Val: 3, Next: n33}
	n11 := &ListNode{Val: 1, Next: n22}

	nv3 := &ListNode{Val: 5, Next: nil}
	nv2 := &ListNode{Val: 4, Next: nv3}
	nv1 := &ListNode{Val: 1, Next: nv2}

	n222 := &ListNode{Val: 6, Next: nil}
	n111 := &ListNode{Val: 2, Next: n222}

	type args struct {
		arr1 *ListNode
		arr2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				arr1: n1,
				arr2: n11,
			},
			want: nil,
		},
		{
			name: "test2",
			args: args{
				arr1: nv1,
				arr2: n111,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Merge2List(tt.args.arr1, tt.args.arr2)
			if got != nil {
				t.Logf("got: %s", got.String())
			}
			// if got := Merge2List(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Merge2List() = %v, want %v", got, tt.want)
			// }
		})
	}
}
