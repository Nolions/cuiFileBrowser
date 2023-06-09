package gui

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

type KeyBinder struct {
	Key      gocui.Key
	Action   func(g *gocui.Gui, v *gocui.View) error
	ViewName string
}

// hotkey binder action
func (g *GUI) keyBinders() []KeyBinder {
	var binders []KeyBinder
	binders = append(binders, KeyBinder{
		Key:      gocui.MouseLeft,
		ViewName: DirListViewName,
		Action:   g.dirItemOnClick,
	})

	binders = append(binders, KeyBinder{
		Key:      gocui.KeyTab,
		ViewName: "",
		Action:   g.nextView,
	})

	binders = append(binders, KeyBinder{
		Key:      gocui.MouseLeft,
		ViewName: SearchBarButtonViewName,
		Action:   g.searchBtnOnClick,
	})

	return binders
}

// 關閉程式
func (g *GUI) quit(gui *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// 設置content layout內容
func (g *GUI) dirItemOnClick(gui *gocui.Gui, v *gocui.View) error {
	//取得點擊項目
	_, cy := v.Cursor()
	_, err := v.Line(cy)
	if err != nil {
		log.Printf("dirItemOnClick(), v line error:%s", err.Error())
		_ = ""
	}

	// TODO 測試用取得字串用
	//out, err := gui.View(FileListViewName)
	//if err != nil {
	//	return err
	//}
	//out.Clear()
	//_, _ = fmt.Fprintln(out, "content: "+s)

	return nil
}

func (g *GUI) searchBtnOnClick(gui *gocui.Gui, v *gocui.View) error {
	// 取得指定路徑
	sv, err := gui.View(SearchBarInputViewName)
	if err != nil {
		return err
	}

	// TODO 測試用取得字串用
	_, cy := sv.Cursor()
	s, err := sv.Line(cy)

	out, err := gui.View(FileListViewName)
	if err != nil {
		return err
	}
	out.Clear()

	_, _ = fmt.Fprintln(out, "searchBtnOnClick: "+s)

	return nil
}

func (g *GUI) nextView(gui *gocui.Gui, v *gocui.View) error {
	nextIndex := (active + 1) % len(viewArr)
	name := viewArr[nextIndex]

	if _, err := g.focus(name); err != nil {
		return err
	}

	if nextIndex == 0 || nextIndex == 3 {
		g.Gui.Cursor = true
	} else {
		g.Gui.Cursor = false
	}

	active = nextIndex
	return nil
}
