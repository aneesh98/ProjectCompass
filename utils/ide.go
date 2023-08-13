package utils

type IDE struct {
	IdeName        string `json:"ideName"`
	ExecutablePath string `json:"executablePath"`
}

func (i *IDE) SetIdeName(name string) {
	i.IdeName = name
}

func (i *IDE) SetExecutablePath(path string) {
	i.ExecutablePath = ProcessPath(path)
}
