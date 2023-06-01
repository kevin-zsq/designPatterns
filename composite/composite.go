package composite

import "fmt"

type Component interface {
	search(string)
}

type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *File) getName() string {
	return f.name
}

type Folder struct {
	components []Component
	name       string
}

func (fr *Folder) search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, fr.name)
	for _, composite := range fr.components {
		composite.search(keyword)
	}
}

func (fr *Folder) add(c Component) {
	fr.components = append(fr.components, c)
}
