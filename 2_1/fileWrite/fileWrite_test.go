package fileWrite

import "testing"

func TestFileWrite(t *testing.T) {
    result := Hello("e205727")
    want := "Hi, e205727. Welcome!"
    if result != want {
        t.Errorf("fileWrite.Hello() = %q want %q", result, want)
    }
}