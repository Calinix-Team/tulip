package internal

import (
	"fmt"
	"os"
	"os/exec"
	"errors"
	"path/filepath"
	"github.com/otiai10/copy"
	"log"
	"io/ioutil"
	"os/user"
	// "github.com/fatih/color"
	// "github.com/Jguer/aur"
)

func ShIns(pkg []string) error {
	// args := strings.Join(pkg, " ")
	arg := []string{"pacman", "-S", "--noconfirm", "--needed"}
	args := append(arg, pkg...)
    cmd := exec.Command("pkexec", args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr;
    err := cmd.Run()
	if err!=nil {
		return err
	}else{
		return nil
	}
}

func ShRem(pkg []string) error{
	// args := strings.Join(pkg, " ")
	arg := []string{"pacman", "-R", "--noconfirm"}
	args := append(arg, pkg...)
    cmd := exec.Command("pkexec", args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
    err := cmd.Run()
	if err!=nil {
		return err
	}else{
		return nil
	}
}

func PacmanInstall(args []string) {
	// pkg := args[0]
	archs:= []string{"Yes! Do it!", "No! Abort!"}
	arch := Select("Are you sure you want to install the package/packages?", archs)
	if arch=="Yes! Do it!" {
	err := ShIns(args)
	if err!=nil{
		fmt.Fprintln(os.Stderr)
	}
}else if arch=="No! Abort!" {
	fmt.Println("Aborting...")
	panic("Aborting...")
}
}

func PacmanRemove(args []string) {
	// pkg := args[0]
	archs:= []string{"Yes! Do it!", "No! Abort!"}
	arch := Select("Are you sure you want to remove the package/packages?", archs)
	if arch=="Yes! Do it!" {
	err := ShRem(args)
	if err!=nil{
		fmt.Fprintln(os.Stderr)
	}
}else if arch=="No! Abort!" {
	fmt.Println("Aborting...")
	panic("Aborting...")
}
}

func BuildFromAUR(pkg string , builddir string, keep bool) {
	usr, _ := user.Current()
	dir := usr.HomeDir
	if builddir=="" {
		dirpath := filepath.Join(dir, ".tulip", "tmp")
		pkgpath := filepath.Join(dir, ".tulip", "tmp", pkg)
		os.MkdirAll(dirpath, os.ModePerm)
		copy.Copy(pkg, pkgpath)
		cm := exec.Command("git", "clone", fmt.Sprintf("https://aur.archlinux.org/%v.git", pkg))
		cm.Stdout = os.Stdout
		cm.Dir = dirpath
		cm.Run()
		cm = exec.Command("makepkg", "-sfi")
		cm.Stdout = os.Stdout
		cm.Stderr = os.Stderr
		cm.Stdin = os.Stdin
		cm.Dir = pkgpath
		cm.Run()
		if !keep {
			os.RemoveAll(pkgpath)
		}
	} else {
		dirpath := builddir
		pkgpath := filepath.Join(builddir, pkg)
		os.MkdirAll(dirpath, os.ModePerm)
		copy.Copy(pkg, pkgpath)
		cm := exec.Command("git", "clone", fmt.Sprintf("https://aur.archlinux.org/%v.git", pkg))
		cm.Stdout = os.Stdout
		cm.Stderr = os.Stderr
		cm.Dir = dirpath
		cm.Run()
		cm = exec.Command("makepkg", "-si")
		cm.Stdout = os.Stdout
		cm.Stderr = os.Stderr
		cm.Stdin = os.Stdin
		cm.Dir = pkgpath
		cm.Run()
		if !keep {
			os.RemoveAll(pkgpath)
		}
	}
}

func CopyFile(src, dst string) (error) {
	bytesRead, err := ioutil.ReadFile(src)

    if err != nil {
        log.Fatal(err)
    }

    err = ioutil.WriteFile(dst, bytesRead, 0644)

    if err != nil {
        return errors.New("copy failed")
    }
    return nil
}

func WalkMatch(root, pattern string) ([]string, error) {
    var matches []string
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() {
            return nil
        }
        if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
            return err
        } else if matched {
            matches = append(matches, path)
        }
        return nil
    })
    if err != nil {
        return nil, err
    }
    return matches, nil
}