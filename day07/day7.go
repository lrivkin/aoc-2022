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
	name   string
	parent *Directory
	files  []*File
	size   uint64
	done   bool
}

var (
	fs []map[string]*Directory
)

func parse(path string) {
	in, _ := os.ReadFile(path)
	inStr := string(in)
	workDir := &Directory{
		name:   "/",
		parent: nil,
		files:  []*File{},
		size:   0,
	}

	fs = []map[string]*Directory{{"/": workDir}}
	level := 0
	for _, block := range Rex.FindAllString(inStr, -1) {
		for _, line := range strings.Split(block, "\n") {
			if matches := CdRegexp.FindStringSubmatch(line); matches != nil {
				dir := matches[1]
				if dir == ".." {
					level--
					workDir = workDir.parent
					// fmt.Printf("moved to %s\n", workDir.name)
				} else {
					d, ok := fs[level][dir]
					if !ok {
						fmt.Printf("fs[%d]=%v\n", level, fs[level])
						panic("this dir should exist but doesnt")
					}
					workDir = d
					// fmt.Printf("moved to %s\n", workDir.name)
					level++
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
			} else if matches := DirRegexp.FindStringSubmatch(line); matches != nil {
				dir := matches[1]
				// fmt.Printf("adding subdir %s to dir %s, fs=%v\n", dir, workDir.name, fs)
				// create new child dir
				if level >= len(fs) {
					newMap := map[string]*Directory{}
					fs = append(fs, newMap)
				}
				l := fs[level]
				_, ok := l[dir]
				if !ok {
					d := &Directory{
						name:   dir,
						parent: workDir,
						files:  []*File{},
					}
					fs[level][dir] = d
				}
				// fmt.Printf("added subdir %s to dir %s, fs=%v\n", dir, workDir.name, fs)
			}
		}
	}
	fmt.Printf("fs=%v\n", fs)
}

func calculate_sizes() {
	for i := len(fs) - 1; i >= 0; i-- {
		for _, dir := range fs[i] {
			// sum up files
			for _, file := range dir.files {
				dir.size += file.size
			}
			// add size to parent
			if i > 0 {
				dir.parent.size += dir.size
			}
			fmt.Printf("dir= %s size= %d\n", dir.name, dir.size)
		}
	}
}

func main() {
	parse("test.txt")
	calculate_sizes()
	fmt.Printf("\nTrying with my real input\n")
	parse("input.txt")
	calculate_sizes()
	// for _, dir := range test {
	// 	fmt.Printf("%s size= %d\n", dir.name, dir.size)
	// }
	// test["/"].GetSize()
	// fmt.Printf("\nMy Input\n")
	// myInput := parse("input.txt")
	// // myInput["/"].GetSize()
	// myInput["/"].PrintAll(0)
}
