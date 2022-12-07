package main

import (
	"fmt"
	"regexp"
	"strconv"

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
	size int
}

func (f *File) String() string {
	return fmt.Sprintf("%s: size=%d", f.name, f.size)
}

func (f *File) getSize() int {
	return f.size
}

type Directory struct {
	name     string
	parent   *Directory
	files    []*File
	children []*Directory
}

func (d *Directory) String() string {
	return fmt.Sprintf("%s: %v", d.name, d.files)
}

func (d *Directory) getSize() int {
	size := 0
	for _, f := range d.files {
		size += f.getSize()
	}
	for _, subDir := range d.children {
		size += subDir.getSize()
	}
	return size
}

func main() {
	lines, _ := utils.ReadLines("test.txt")
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
			var d *Directory
			if _, ok := fs[dir]; !ok {
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
			size, _ := strconv.Atoi(sizeStr)
			name := matches[2]
			f := &File{
				name: name,
				size: size,
			}
			workDir.files = append(workDir.files, f)
		}
	}
	for _, dir := range fs {
		fmt.Printf("%v\n", dir)
	}

}
