package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

var current string
var curev termbox.Event

func mouse_button_num(k termbox.Key) int {
	switch k {
	case termbox.MouseLeft:
		return 0
	case termbox.MouseMiddle:
		return 1
	case termbox.MouseRight:
		return 2
	}
	return 0
}

func redraw_all() {
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)

	tbprint(0, 0, termbox.ColorMagenta, coldef, "Press 'q' to quit")
	tbprint(0, 1, coldef, coldef, current)
	tbprint(0, 3, coldef, coldef, fmt.Sprintf("%d", curev.N))
        
        tbprint(0,4, termbox.ColorRed,    coldef,   "Meow");
        tbprint(0,5, termbox.ColorGreen,  coldef,   "Meow");
        tbprint(0,6, termbox.ColorBlue,   coldef,   "Meow");
        tbprint(0,7, termbox.ColorYellow, coldef,   "Meow");

        for i := 0; i < 10; i++ {
            line, isPrefix, err := hexfile.ReadLine();
            tbprint(0,8 + i, termbox.ColorBlue, coldef,   line);
        }


	termbox.Flush()
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputAlt | termbox.InputMouse)
	redraw_all()

	data := make([]byte, 0, 64)
mainloop:
	for {
		if cap(data)-len(data) < 32 {
			newdata := make([]byte, len(data), len(data)+32)
			copy(newdata, data)
			data = newdata
		}
		beg := len(data)
		d := data[beg : beg+32]
		switch ev := termbox.PollRawEvent(d); ev.Type {
		case termbox.EventRaw:
			data = data[:beg+ev.N]
			current = fmt.Sprintf("%q", data)
			if current == `"q"` {
				break mainloop
			}

			curev = termbox.ParseEvent(data)
			if curev.N > 0 {
				copy(data, data[curev.N:])
				data = data[:len(data)-curev.N]
			}
		case termbox.EventError:
			panic(ev.Err)
		}
		redraw_all()
	}
}
