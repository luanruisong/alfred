package alfred

import "github.com/luanruisong/alfred/model"

func NewMgr() *model.Mgr {
	return &model.Mgr{
		Items: make([]*model.Item, 0),
	}
}

func NewIcon(path string) *model.Icon {
	return NewIconByType("", path)
}

func NewIconByType(iconType, path string) *model.Icon {
	return &model.Icon{
		Type: iconType,
		Path: path,
	}
}

func NewMod() model.Mod {
	return NewModByMap(make(map[string]string))
}

func NewModByMap(m map[string]string) model.Mod {
	return model.Mod(m)
}

func NewItem() *model.Item {
	return &model.Item{}
}

func NewDefItem(title, subTitle string) *model.Item {
	return &model.Item{
		Title:    title,
		Subtitle: subTitle,
	}
}
