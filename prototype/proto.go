package prototype

import "fmt"

// Inode 每个文件和文件夹都可用一个inode接口来表示。 inode接口中同样也有clone克隆功能。
type Inode interface {
	print(string)
	clone() Inode
}

// File 具体原型
type File struct {
	name string
}

// print indentation打印的时候添加缩进，结果更清晰
func (f *File) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *File) clone() Inode {
	// 注意这里使用的是*File实现了print和clone接口，所以clone返回的interface对象必须加上&
	return &File{name: f.name + "_clone"}
}

type Folder struct {
	children []Inode
	name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *Folder) clone() Inode {
	newInode := &Folder{name: f.name + "_clone"}
	var tmpChild = make([]Inode, 0)
	for _, i := range f.children {
		tmpChild = append(tmpChild, i.clone())
	}
	newInode.children = tmpChild
	return newInode
}
