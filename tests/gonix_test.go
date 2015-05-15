package gonix

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

type testCase struct {
	in, want string
}

func assert(t *testing.T, name string, cmd *exec.Cmd, c testCase) {
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	got := out.String()
	if got != c.want {
		t.Errorf("%s (%q) == %q, want %q", name, c.in, got, c.want)
	}
}

func TestBase64Decode(t *testing.T) {
	name := "base64 decode"
	file, err := ioutil.TempFile(os.TempDir(), "testing_base64_decode")

	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("aG9sYSBjaGFvCg==")

	cases := []testCase{
		{"aG9sYSBjaGFvCg==", "hola chao\n"},
	}

	cmd := new(exec.Cmd)
	for _, c := range cases {
		cmd = exec.Command("../build/base64", "-d", file.Name())
		assert(t, name, cmd, c)
		cmd = exec.Command("../build/base64", "-d")
		stdin, _ := cmd.StdinPipe()
		stdin.Write([]byte(c.in))
		stdin.Close()
		assert(t, name, cmd, c)
	}
}

func TestBase64Encode(t *testing.T) {
	name := "base64 encode"
	file, err := ioutil.TempFile(os.TempDir(), "testing_base64_encode")

	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("hola")

	cases := []testCase{
		{"hola", "aG9sYQ==\n"},
	}

	cmd := new(exec.Cmd)
	for _, c := range cases {
		cmd = exec.Command("../build/base64", file.Name())
		assert(t, name, cmd, c)
		cmd = exec.Command("../build/base64")
		stdin, _ := cmd.StdinPipe()
		stdin.Write([]byte(c.in))
		stdin.Close()
		assert(t, name, cmd, c)
	}

	wrapCases := []testCase{
		{"hola", "aG9\nsYQ\n==\n"},
	}

	for _, c := range wrapCases {
		cmd = exec.Command("../build/base64", "-w", "3", file.Name())
		assert(t, name, cmd, c)
		cmd = exec.Command("../build/base64", "-w", "3")
		stdin, _ := cmd.StdinPipe()
		stdin.Write([]byte(c.in))
		stdin.Close()
		assert(t, name, cmd, c)
	}
}

func TestBasename(t *testing.T) {
	name := "basename"
	cases := []testCase{
		{"hola.txt", "hola.txt\n"},
		{"tmp/chao.txt", "chao.txt\n"},
		{"", ".\n"},
	}

	cmd := new(exec.Cmd)
	for _, c := range cases {
		cmd = exec.Command("../build/basename", c.in)
		assert(t, name, cmd, c)
	}
}

func TestDirname(t *testing.T) {
	name := "dirname"
	cases := []testCase{
		{"/tmp/hola.txt", "/tmp\n"},
		{"/tmp/", "/\n"},
		{"/tmp", "/\n"},
		{"", ".\n"},
	}

	cmd := new(exec.Cmd)
	for _, c := range cases {
		cmd = exec.Command("../build/dirname", c.in)
		assert(t, name, cmd, c)
	}
}

func TestEcho(t *testing.T) {
	name := "echo"
	cases := []testCase{
		{"hola", "hola\n"},
		{"chao", "chao\n"},
		{"hola chao", "hola chao\n"},
		{"hola chao si", "hola chao si\n"},
		{"-n hola chao si", "hola chao si"},
		{"-n -n hola chao si", "hola chao si"},
		{"-n hola -n chao si", "hola -n chao si"},
		{"", "\n"},
	}

	cmd := new(exec.Cmd)
	for _, c := range cases {
		cmd = exec.Command("../build/echo", strings.Split(c.in, " ")...)
		assert(t, name, cmd, c)
	}
}

func TestFalse(t *testing.T) {
	cmd := exec.Command("../build/false")
	output, err := cmd.CombinedOutput()
	if err != nil {
		// log.Fatal(err)
	} else {
		t.Errorf("false () == %q, want 1 ", string(output))
	}
}

func TestTrue(t *testing.T) {
	cmd := exec.Command("../build/true")
	output, err := cmd.CombinedOutput()
	if err == nil {
		// log.Fatal(err)
	} else {
		t.Errorf("true () == %q, want nil ", string(output))
	}
}

func TestMd5SumString(t *testing.T) {

	//TODO: pass input by stdin
	cases := []testCase{
		{"hola", "686f6c61d41d8cd98f00b204e9800998ecf8427e\t\"hola\"\n"},
		{"", ""},
	}

	for _, c := range cases {
		var out bytes.Buffer
		cmd := exec.Command("../build/md5sum", "-s", c.in)
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		got := out.String()
		if got != c.want {
			t.Errorf("md5sum (%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
