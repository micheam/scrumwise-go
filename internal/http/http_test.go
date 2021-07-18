package http

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_banner(t *testing.T) {
	v := "vXXX.YY.Z"
	got := string(banner(v))
	want := `
██╗    ██╗██╗███████╗███████╗███╗   ███╗ █████╗ ███╗   ██╗
██║    ██║██║██╔════╝██╔════╝████╗ ████║██╔══██╗████╗  ██║
██║ █╗ ██║██║███████╗█████╗  ██╔████╔██║███████║██╔██╗ ██║
██║███╗██║██║╚════██║██╔══╝  ██║╚██╔╝██║██╔══██║██║╚██╗██║
╚███╔███╔╝██║███████║███████╗██║ ╚═╝ ██║██║  ██║██║ ╚████║
 ╚══╝╚══╝ ╚═╝╚══════╝╚══════╝╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝
                                                 vXXX.YY.Z`
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("banner mismatch (-want, +got):%s\n", diff)
	}
}
