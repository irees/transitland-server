package workers

import (
	"log"
	"os"
	"testing"

	"github.com/interline-io/transitland-server/internal/testutil"
)

func TestMain(m *testing.M) {
	if a, ok := testutil.CheckTestDB(); !ok {
		log.Print(a)
		return
	}
	os.Exit(m.Run())
}
