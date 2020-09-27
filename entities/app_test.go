package entities

import (
	"fmt"
	"testing"
)

func TestAppAS_Fill(t *testing.T) {
	app := AppAS{TrackId: 1261944766}
	err := app.Fill()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("应用信息：%+v\n", app)
}
