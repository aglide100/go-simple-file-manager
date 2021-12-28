package dir

type DirManger interface {
	GetDirectoryList(string) ([]string, error)
	CreateNewDirectory(string string) error
}