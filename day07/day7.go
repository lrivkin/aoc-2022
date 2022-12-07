package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/lrivkin/aoc-2022/utils"
)

func part1() int {
	return 0
}
func part2() int {
	return 0
}

var (
	CdRegexp   = regexp.MustCompile(`^\$ cd (\S+)$`)
	LS         = "$ ls"
	DirRegexp  = regexp.MustCompile(`^dir (\S+)$`)
	FileRegexp = regexp.MustCompile(`^(\d+) (\S+)$`)
)

type File struct {
	name string
	size uint64
}

func (f *File) String() string {
	return fmt.Sprintf("%s: size=%d", f.name, f.size)
}

func (f *File) getSize() uint64 {
	return f.size
}

type Directory struct {
	name     string
	parent   *Directory
	files    []*File
	children []*Directory
	size     *uint64
}

func (d *Directory) String() string {
	children := []string{}
	for _, subDir := range d.children {
		children = append(children, subDir.name)
	}
	return fmt.Sprintf("%s:\n\tFiles= %v\n\tSubdirectories= %v", d.name, d.files, children)
}

func (d *Directory) PrintAll(level int) {
	indent := strings.Repeat("\t", level)
	fmt.Printf("%s%s\n", indent, d.name)
	indent += "\t"
	fmt.Printf("%sFiles= %v\n", indent, d.files)
	for _, subDir := range d.children {
		subDir.PrintAll(level + 1)
	}

}

func (d *Directory) GetSize() *uint64 {
	if d.size != nil {
		// fmt.Printf("Dir %s size= %d", d.name, *d.size)
		return d.size
	}
	dirSize := uint64(0)
	for _, f := range d.files {
		dirSize += f.getSize()
	}
	for _, subDir := range d.children {
		dirSize += *subDir.GetSize()
	}
	d.size = &dirSize
	fmt.Printf("Dir %s size= %d\n", d.name, *d.size)

	return &dirSize
}

func parseFS(path string) map[string]*Directory {
	lines, _ := utils.ReadLines(path)

	fs := map[string]*Directory{}
	workDir := &Directory{
		name:     "/",
		parent:   nil,
		files:    []*File{},
		children: []*Directory{},
	}
	fs["/"] = workDir

	for _, l := range lines[1:] {
		// fmt.Println(l)
		if l == LS {
			// get the next lines
			// fmt.Printf("ls\n")
		} else if matches := CdRegexp.FindStringSubmatch(l); matches != nil {
			// get the directory
			dir := matches[1]
			if dir == ".." {
				workDir = workDir.parent
				continue
			}
			d, ok := fs[dir]
			if !ok {
				// create this directory for the first time
				d = &Directory{
					name:     dir,
					parent:   workDir,
					files:    []*File{},
					children: []*Directory{},
				}
				fs[dir] = d
			}
			// move into that folder
			workDir = d
		} else if matches := DirRegexp.FindStringSubmatch(l); matches != nil {
			dir := matches[1]
			d, ok := fs[dir]
			if !ok {
				d = &Directory{
					name:     dir,
					parent:   workDir,
					files:    []*File{},
					children: []*Directory{},
				}
				fs[dir] = d
			}
			// add this directory to child set
			workDir.children = append(workDir.children, d)
			// fmt.Printf("(from ls) dir = %s\n", dir)
		} else if matches := FileRegexp.FindStringSubmatch(l); matches != nil {
			sizeStr := matches[1]
			size, _ := strconv.ParseUint(sizeStr, 10, 0)
			name := matches[2]
			f := &File{
				name: name,
				size: size,
			}
			workDir.files = append(workDir.files, f)
		}
	}
	return fs
}
func main() {
	test := parseFS("test.txt")
	test["/"].PrintAll(0)
	test["/"].GetSize()

	fmt.Printf("\nMy Input\n")
	myInput := parseFS("input.txt")
	myInput["/"].GetSize()
	// myInput["/"].PrintAll(0)
}
