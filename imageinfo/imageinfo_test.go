package imageinfo

import "testing"

func TestFromArgs(t *testing.T) {
	args := []string{"", "", "", "", "100", "200", "./file"}
	info, err := FromArgs(args)

	if err != nil {
		t.Error(err)
	}
	if info.Width != 100 {
		t.Errorf("got %d; want 100", info.Width)
	}
	if info.Height != 200 {
		t.Errorf("got %d; want 200", info.Height)
	}
	if info.TargetFile != "./file" {
		t.Errorf("go %s; want ./file", info.TargetFile)
	}
}
