package internal

import (
	"fmt"
	"os"
	"os/exec"
	// "github.com/Jguer/aur"
)

func Shellout(pkg []string) error {
	// args := strings.Join(pkg, " ")
	arg := []string{"pacman", "-S", "--noconfirm"}
	args := append(arg, pkg...)
	fmt.Println("pkexec", args)
    cmd := exec.Command("pkexec", args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
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
	err := Shellout(args)
	if err!=nil{
		fmt.Fprintln(os.Stderr, err)
	}
}else if arch=="No! Abort!" {
	fmt.Println("Aborting...")
	panic("Aborting...")
}
}