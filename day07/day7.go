package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	Rex        = regexp.MustCompile(`(\$ cd \S+\n)+\$ ls\n[^$]+`)
)

type File struct {
	name string
	size uint64
}

func (f *File) String() string {
	return fmt.Sprintf("%s: %d", f.name, f.size)
}

func (f *File) getSize() uint64 {

	return f.size
}

type Directory struct {
	name     string
	parent   *Directory
	files    []*File
	children []*Directory
	size     uint64
	done     bool
}

func (d *Directory) PrintAll(level int) {
	indent := strings.Repeat("  ", level)
	fmt.Printf("%s%s\n", indent, d.name)
	indent += "  "
	fmt.Printf("%s%v\n", indent, d.files)
	for _, subDir := range d.children {
		subDir.PrintAll(level + 1)
	}

}

func parse(path string) map[string]*Directory {
	in, _ := os.ReadFile(path)
	inStr := string(in)
	workDir := &Directory{
		name:     "/",
		parent:   nil,
		files:    []*File{},
		children: []*Directory{},
		size:     0,
	}
	fs := map[string]*Directory{"/": workDir}
	for _, block := range Rex.FindAllString(inStr, -1) {
		for _, line := range strings.Split(block, "\n") {
			if matches := CdRegexp.FindStringSubmatch(line); matches != nil {
				dir := matches[1]
				if dir == ".." {
					workDir = workDir.parent
				} else {
					workDir = fs[dir]
				}
			} else if matches := FileRegexp.FindStringSubmatch(line); matches != nil {
				sizeStr := matches[1]
				size, _ := strconv.ParseUint(sizeStr, 10, 0)
				name := matches[2]
				f := &File{
					name: name,
					size: size,
				}
				workDir.files = append(workDir.files, f)
				workDir.size += size
			} else if matches := DirRegexp.FindStringSubmatch(line); matches != nil {
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
			}
		}
		if len(workDir.children) == 0 {
			workDir.done = true
			if workDir.parent != nil {
				workDir.parent.size += workDir.size
			}
		}
	}
	return fs
}
func main() {
	test := parse("test.txt")
	test["/"].PrintAll(0)
	for _, dir := range test {
		fmt.Printf("%s size= %d\n", dir.name, dir.size)
	}
	// test["/"].GetSize()
	// fmt.Printf("\nMy Input\n")
	// myInput := parse("input.txt")
	// // myInput["/"].GetSize()
	// myInput["/"].PrintAll(0)
}
