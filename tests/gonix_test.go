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

func TestBase64Decode(t *testing.T) {
	file, err := ioutil.TempFile(os.TempDir(), "testing_base64_decode")

	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("aG9sYSBjaGFvCg==")

	//TODO: test passing input by Stdin
	cases := []struct {
		in, want string
	}{
		{"aG9sYSBjaGFvCg==", "hola chao\n"},
	}

	for _, c := range cases {
		var out bytes.Buffer
		cmd := exec.Command("../build/base64", "-d", file.Name())
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		got := out.String()
		if got != c.want {
			t.Errorf("base64 (%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestBase64Encode(t *testing.T) {
	file, err := ioutil.TempFile(os.TempDir(), "testing_base64_encode")

	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("hola")

	//TODO: test passing input by stdin
	cases := []struct {
		in, want string
	}{
		{"hola", "aG9sYQ==\n"},
	}

	for _, c := range cases {
		var out bytes.Buffer
		cmd := exec.Command("../build/base64", file.Name())
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		got := out.String()
		if got != c.want {
			t.Errorf("base64 (%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestBasename(t *testing.T) {

	cases := []struct {
		in, want string
	}{
		{"hola.txt", "hola.txt\n"},
		{"tmp/chao.txt", "chao.txt\n"},
		{"", ".\n"},
	}
	for _, c := range cases {
		var out bytes.Buffer
		cmd := exec.Command("../build/basename", c.in)
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		got := out.String()
		if got != c.want {
			t.Errorf("basename (%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestDirname(t *testing.T) {

	cases := []struct {
		in, want string
	}{
		{"/tmp/hola.txt", "/tmp\n"},
		{"/tmp/", "/\n"},
		{"/tmp", "/\n"},
		{"", ".\n"},
	}
	for _, c := range cases {
		var out bytes.Buffer
		cmd := exec.Command("../build/dirname", c.in)
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		got := out.String()
		if got != c.want {
			t.Errorf("dirname (%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestEcho(t *testing.T) {

	cases := []struct {
		in, want string
	}{
		{"hola", "hola\n"},
		{"chao", "chao\n"},
		{"hola chao", "hola chao\n"},
		{"hola chao si", "hola chao si\n"},
		{"-n hola chao si", "hola chao si"},
		{"-n -n hola chao si", "hola chao si"},
		{"-n hola -n chao si", "hola -n chao si"},
		{"", "\n"},
	}
	for _, c := range cases {
		var out bytes.Buffer
		cmd := exec.Command("../build/echo", strings.Split(c.in, " ")...)
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		got := out.String()
		if got != c.want {
			t.Errorf("echo (%q) == %q, want %q", c.in, got, c.want)
		}
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
	cases := []struct {
		in, want string
	}{
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
