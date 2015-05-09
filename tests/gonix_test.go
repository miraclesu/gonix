package gonix

import (
  "bytes"
  "log"
  "os/exec"
  "testing"
)

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
    {"hola", "hola \n"},
    {"chao", "chao \n"},
    {"hola chao", "hola chao \n"},
    {"hola chao si", "hola chao si \n"},
    {"", " \n"},
  }
  for _, c := range cases {
    var out bytes.Buffer
    cmd := exec.Command("../build/echo", c.in)
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
