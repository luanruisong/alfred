package alfred

import "testing"

func TestBase(t *testing.T) {
	mgr := NewMgr()
	mgr.AddItem(NewDefItem("I am title", "I am sub title"))
	mgr.AddItem(NewDefItem("I am title2", "I am sub title2"))

	mgr.PrintJson()
}
