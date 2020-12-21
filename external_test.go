package render_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/chyroc/render"
)

func Test_External(t *testing.T) {
	fmt.Println(render.Render(time.Now()))

	not := time.Now()
	render.RegisterFormat(not, func(v interface{}) string {
		return v.(time.Time).Format(time.RFC3339)
	})
	render.RegisterFormat(&not, func(v interface{}) string {
		return v.(*time.Time).Format(time.RFC3339)
	})

	fmt.Println(render.Render(time.Now()))
}
