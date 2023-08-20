package hello

import "testing"

func TestSayHello(t *testing.T) {
	subsets := []struct {
		items  []string
		result string
	}{
		{
			result: "hello, world!",
		},
		{
			items:  []string{"Mukesh"},
			result: "hello, Mukesh",
		},
		{
			items:  []string{"Mukesh", "Mahato"},
			result: "hello, Mukesh, Mahato!",
		},
	}

	for _, st := range subsets {
		if s := Say(st.items); s != st.result {
			t.Errorf("wanted %s {%v}, got %s", st.result, st.items, s)
		}
	}
}
