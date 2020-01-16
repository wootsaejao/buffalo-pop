package newapp

import (
	"os"
	"testing"

	"github.com/gobuffalo/genny/v2/gentest"
	"github.com/gobuffalo/pop/v5"
	"github.com/stretchr/testify/require"
)

func Test_New(t *testing.T) {
	r := require.New(t)

	for _, d := range pop.AvailableDialects {
		run := gentest.NewRunner()
		err := run.Chdir(os.TempDir(), func() error {
			g, err := New(&Options{
				Prefix:  "foo",
				Dialect: d,
			})
			r.NoError(err)
			run.WithGroup(g)

			r.NoError(run.Run())

			res := run.Results()
			r.Len(res.Commands, 0)
			r.Len(res.Files, 4)
			return nil
		})
		r.NoError(err)
	}
}
