package uname

import (
	"io/ioutil"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func out(name string, arg string) (string, error) {
	cmd := exec.Command(name, arg)

	// pipe that will be connected to the command's standard output
	stdout, _ := cmd.StdoutPipe()

	// start command
	err := cmd.Start()
	if err != nil {
		return "", err
	}

	// read standard output
	buf, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}

	// print standard output
	return strings.TrimRight(string(buf), "\n"), err
}

func TestUname(t *testing.T) {
	uname := "uname"

	u, err := New()
	assert.NoErrorf(t, err, "uname syscall")
	machine := u.Machine()
	sysname := u.Sysname()
	nodename := u.Nodename()
	kernelVersion := u.KernelVersion()
	kernelRelease := u.KernelRelease()

	tests := []struct {
		name     string
		result   string
		unameArg string
	}{
		{
			"Machine",
			machine,
			"-m",
		},
		{
			"Sysname",
			sysname,
			"-s",
		},
		{
			"Nodename",
			nodename,
			"-n",
		},
		{
			"KernelVersion",
			kernelVersion,
			"-v",
		},
		{
			"KernelRelease",
			kernelRelease,
			"-r",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want, err := out(uname, tt.unameArg)
			assert.NoErrorf(t, err, "uname execute")
			assert.Equal(t, want, tt.result)
		})
	}
}
