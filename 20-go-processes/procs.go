package main

import "os"
import "fmt"
import "io"
import osexec "os/exec"

func main() {
	// FindProcess()
	// LaunchProcess()
	ExecProcess()
}

func ExecProcess() {
	procName := "pwsh"
	args := []string{"-Command", "$FirstName = Read-Host 'Please enter your name'; Write-Host -Object \"Your first name was entered as $FirstName\""}
	proc := osexec.Command(procName, args...)

	stdout, outerr := proc.StdoutPipe()
	stdin, _ := proc.StdinPipe()
	fmt.Println(outerr)
	// err := proc.Run()
	err := proc.Start()

	fmt.Println(stdout)

	var input string
	fmt.Scanln(&input)
	input = input + "\n"
	stdinbytes, stdinerr := stdin.Write([]byte(input))
	fmt.Printf("We wrote %v bytes into the external process\n", stdinbytes)
	fmt.Printf("Was there a stdin error? %v\n", stdinerr)

	dataoutput, readerr := io.ReadAll(stdout)
	fmt.Println(string(dataoutput), readerr)

	proc.Wait()
	fmt.Println(err)
	fmt.Println("Process ran")
}

func LaunchProcess() {
	procName := "/usr/bin/ls"
	argList := []string{"-lGa"}

	procName = "/snap/powershell/234/opt/powershell/pwsh"
	argList = []string{"-Command", "Start-Sleep -Seconds 1"}

	attr := os.ProcAttr{Dir: ""}
	proc, err := os.StartProcess(procName, argList, &attr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(proc)
	procState, err := proc.Wait()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Has process exited? %v\n", procState.Exited())
	fmt.Printf("How much user CPU time was used? %v\n", procState.UserTime())
}

func FindProcess() {
	proc, err := os.FindProcess(181798)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", proc)
	proc.Kill()
	fmt.Println("Process was killed (hopefully)")
}
