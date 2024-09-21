// Copyright 2019-2022 Graham Clark. All rights reserved.  Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// The fourth example from the gowid tutorial.
package main

import (
	"fmt"

	"github.com/blacknon/gowid"
	"github.com/blacknon/gowid/examples"
	"github.com/blacknon/gowid/widgets/edit"
	"github.com/blacknon/gowid/widgets/text"
	tcell "github.com/gdamore/tcell/v2"
)

//======================================================================

type QuestionBox struct {
	gowid.IWidget
}

func (w *QuestionBox) UserInput(ev interface{}, size gowid.IRenderSize, focus gowid.Selector, app gowid.IApp) bool {
	res := true
	if evk, ok := ev.(*tcell.EventKey); ok {
		switch evk.Key() {
		case tcell.KeyEnter:
			w.IWidget = text.New(fmt.Sprintf("Nice to meet you, %s.\n\nPress Q to exit.", w.IWidget.(*edit.Widget).Text()))
		default:
			res = w.IWidget.UserInput(ev, size, focus, app)
		}
	}
	return res
}

func main() {
	edit := edit.New(edit.Options{Caption: "What is your name?\n"})
	qb := &QuestionBox{edit}
	app, err := gowid.NewApp(gowid.AppArgs{View: qb})
	examples.ExitOnErr(err)
	app.MainLoop(gowid.UnhandledInputFunc(gowid.HandleQuitKeys))
}

//======================================================================
// Local Variables:
// mode: Go
// fill-column: 110
// End:
