// Copyright 2019-2022 Graham Clark. All rights reserved.  Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// A gowid test app which exercises the list, edit and framed widgets.
package main

import (
	"fmt"

	"github.com/blacknon/gowid"
	"github.com/blacknon/gowid/examples"
	"github.com/blacknon/gowid/widgets/edit"
	"github.com/blacknon/gowid/widgets/framed"
	"github.com/blacknon/gowid/widgets/list"
	log "github.com/sirupsen/logrus"
)

//======================================================================

func main() {

	f := examples.RedirectLogger("widgets7.log")
	defer f.Close()

	palette := gowid.Palette{}

	edits := make([]gowid.IWidget, 0)

	for i := 0; i < 40; i++ {
		edits = append(edits, edit.New(edit.Options{Caption: fmt.Sprintf("Cap%d:", i+1), Text: "abcde1111111222222222222222223333333333444444444"}))
	}

	walker := list.NewSimpleListWalker(edits)
	lb := list.New(walker)
	fr := framed.New(lb)

	app, err := gowid.NewApp(gowid.AppArgs{
		View:    fr,
		Palette: &palette,
		Log:     log.StandardLogger(),
	})
	examples.ExitOnErr(err)

	app.SimpleMainLoop()
}

//======================================================================
// Local Variables:
// mode: Go
// fill-column: 110
// End:
