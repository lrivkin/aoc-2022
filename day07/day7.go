package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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

type Directory struct {
	name     string
	fullPath string
	parent   *Directory
	files    []*File
	size     uint64
}

var (
	fs []map[string]*Directory
)

func getPathName(parent *Directory, name string) string {
	if name == "/" {
		return name
	}
	if parent.fullPath == "/" {
		return fmt.Sprintf("/%s", name)
	}
	return fmt.Sprintf("%s/%s", parent.fullPath, name)
}

func parse(path string) {
	in, _ := os.ReadFile(path)
	inStr := string(in)
	workDir := &Directory{
		name:     "/",
		fullPath: "/",
		parent:   nil,
		files:    []*File{},
		size:     0,
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
					fullPath := getPathName(workDir, dir)
					d, ok := fs[level][fullPath]
					if !ok {
						fmt.Printf("looking for %s, dir=%s, fs[%d]=%v\n", fullPath, dir, level, fs[level])
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
					fullPath := getPathName(workDir, dir)
					d := &Directory{
						name:     dir,
						fullPath: fullPath,
						parent:   workDir,
						files:    []*File{},
					}
					fs[level][fullPath] = d
				}
				// fmt.Printf("added subdir %s to dir %s, fs=%v\n", dir, workDir.name, fs)
			}
		}
	}
}

func calculate_sizes() {
	for i := len(fs) - 1; i >= 0; i-- {
		for _, dir := range fs[i] {
			// sum up files
			fmt.Printf("dir=%s (%s):", dir.name, dir.fullPath)
			var files uint64 = 0
			for _, file := range dir.files {
				files += file.size
			}
			dir.size += files
			fmt.Printf(" files=%d", files)
			// add size to parent
			if dir.parent != nil {
				dir.parent.size += dir.size
				fmt.Printf(" parent=%s", dir.parent.name)
			}
			fmt.Printf(" size=%d\n", dir.size)
		}
		fmt.Println()
	}
}

func part1() uint64 {
	var total uint64
	total = 0
	for _, l := range fs {
		for _, d := range l {
			if d.size <= 100000 {
				total += d.size
			}
		}
	}
	fmt.Printf("Part1 = %d\n", total)
	return total
}

func part2() uint64 {
	fmt.Printf("\nPart 2\n")
	var total_space uint64 = 70000000
	homedir := fs[0]["/"]
	space_left := total_space - homedir.size
	needed := 30000000 - space_left
	fmt.Printf("total size= %d, need to delete= %d\n", homedir.size, needed)

	var space_freed uint64
	found := false
	for i := len(fs) - 1; i >= 0; i-- {
		for _, dir := range fs[i] {
			if dir.size >= needed {
				fmt.Printf("delete dir=%s size=%d\n", dir.name, dir.size)
				if !found {
					space_freed = dir.size
					found = true
				}
				if dir.size < space_freed {
					space_freed = dir.size
				}
			}
		}
	}
	fmt.Printf("Part2 = %d\n", space_freed)
	// if the home dir is giant?
	return space_freed
}

func main() {
	parse("test.txt")
	calculate_sizes()
	part1()
	part2()

	fmt.Printf("\nTrying with my real input\n")
	parse("input.txt")
	calculate_sizes()
	part1()
	part2()
}
