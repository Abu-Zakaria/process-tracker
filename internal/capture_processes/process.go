package capture_processes

import (
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/mitchellh/go-ps"
)

type Process struct {
	Pid        int    `json:"pid"`
	Executable string `json:"executable"`
}

func GetProcessMemory(process Process) MemoryStatus {
	memory := GetMemory(process)

	return MemoryStatus{
		process,
		memory,
	}
}

func GetAllProcesses() []Process {
	processes, err := ps.Processes()
	if err != nil {
		log.Fatal(err)
	}

	my_processes := []Process{}

	for _, process := range processes {
		my_processes = append(my_processes, Process{
			Pid:        process.Pid(),
			Executable: GetExecutable(process.Pid()),
		})
	}

	return my_processes
}

func GetExecutable(pid int) string {
	switch OS {
	case "darwin":
		return getExecutableDarwin(pid)
	case "linux":
		return getExecutableLinux(pid)
	case "windows":
		return getExecutableWindows(pid)
	default:
		return ""
	}
}

func getExecutableDarwin(pid int) string {
	cmd := exec.Command("ps", "-p", strconv.Itoa(pid), "-o", "command=")
	var output strings.Builder
	cmd.Stdout = &output

	cmd.Run()

	return strings.TrimRight(strings.TrimLeft(output.String(), " "), "\n")
}

func getExecutableLinux(pid int) string {
	// TODO
	return ""
}

func getExecutableWindows(pid int) string {
	// TODO
	return ""
}
