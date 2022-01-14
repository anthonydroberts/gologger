package process

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/anthonydroberts/gologger/terminal"
	"github.com/mitchellh/go-ps"
)

func RunCommand(arg string, silent bool, gologgerSilent bool) (string, string, string, int, time.Time) {
	shell := getExecutingShellCmd()

	execCmd := exec.Command(shell, "-c", arg)

	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)

	if silent == false {
		execCmd.Stdout = io.MultiWriter(os.Stdout, stdout)
		execCmd.Stderr = io.MultiWriter(os.Stderr, stderr)
	} else {
		execCmd.Stdout = io.MultiWriter(stdout)
		execCmd.Stderr = io.MultiWriter(stderr)
	}
	execCmd.Stdin = os.Stdin

	startTime := time.Now()
	if !gologgerSilent {
		terminal.Msg("print", fmt.Sprintf("[%s -c %s]\n", shell, arg))
	}

	execCmd.Run()

	exit := execCmd.ProcessState.ExitCode()

	if !gologgerSilent {
		if exit == 0 {
			terminal.Msg("success", fmt.Sprintf("[%s -c %s] exited with 0", shell, arg))
		} else {
			terminal.Msg("fail", fmt.Sprintf("[%s -c %s] exited with %d", shell, arg, exit))
		}
	}

	return arg, stdout.String(), stderr.String(), exit, startTime
}

func getParentProcess(pid int) ps.Process {
	ppid := -1
	var pProcess ps.Process = nil
	procs, err := ps.Processes()
	if err != nil {
		fmt.Println("Error reading process tree!")
		panic(err)
	}

	for _, p := range procs {
		if p.Pid() == pid {
			ppid = p.PPid()
			break
		}
	}

	for _, p := range procs {
		if p.Pid() == ppid {
			pProcess = p
			break
		}
	}

	if ppid == -1 || pProcess == nil {
		log.Fatalf("Error retrieving PPID of process %s", string(rune(pid)))
	}

	return pProcess
}

func getExecutingShellCmd() string {
	// Get the ps.Process of Go's parent process, which should be the shell's process
	// Process hierarchy: Shell -> Go -> Gologger
	shellProc := getParentProcess(os.Getppid())
	shellStr := shellProc.Executable()

	return strings.TrimSuffix(shellStr, filepath.Ext(shellStr))
}
