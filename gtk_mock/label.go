package gtk_mock

type MockLabel struct {
	*MockWidget
}

func (*MockLabel) SetLabel(v1 string) {
}

func (*MockLabel) SetText(v1 string) {
}

func (*MockLabel) SetSelectable(v1 bool) {
}
