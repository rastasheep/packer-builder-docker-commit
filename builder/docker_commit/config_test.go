package docker_commit

import (
	"testing"
)

func testConfig() map[string]interface{} {
	return map[string]interface{}{
		"export_path": "foo",
		"commit_tag":  "foo/bar:baz",
		"image":       "bar",
	}
}

func testConfigStruct(t *testing.T) *Config {
	c, warns, errs := NewConfig(testConfig())
	if len(warns) > 0 {
		t.Fatalf("bad: %#v", len(warns))
	}
	if errs != nil {
		t.Fatalf("bad: %#v", errs)
	}

	return c
}

func testConfigErr(t *testing.T, warns []string, err error) {
	if len(warns) > 0 {
		t.Fatalf("bad: %#v", warns)
	}
	if err == nil {
		t.Fatal("should error")
	}
}

func testConfigOk(t *testing.T, warns []string, err error) {
	if len(warns) > 0 {
		t.Fatalf("bad: %#v", warns)
	}
	if err != nil {
		t.Fatalf("bad: %s", err)
	}
}

func TestConfigPrepare_exportPath(t *testing.T) {
	raw := testConfig()

	// No export path
	delete(raw, "export_path")
	c, warns, errs := NewConfig(raw)
	testConfigOk(t, warns, errs)
	if c.Export {
		t.Fatal("should not export")
	}

	// Good export path
	raw["export_path"] = "good"
	c, warns, errs = NewConfig(raw)
	testConfigOk(t, warns, errs)
	if !c.Export {
		t.Fatal("should export")
	}
}

func TestConfigPrepare_image(t *testing.T) {
	raw := testConfig()

	// No image
	delete(raw, "image")
	_, warns, errs := NewConfig(raw)
	testConfigErr(t, warns, errs)

	// Good image
	raw["image"] = "path"
	_, warns, errs = NewConfig(raw)
	testConfigOk(t, warns, errs)
}

func TestConfigPrepare_commit(t *testing.T) {
	raw := testConfig()

	// No tag
	delete(raw, "commit_tag")
	_, warns, errs := NewConfig(raw)
	testConfigErr(t, warns, errs)

	// Good tag
	raw["commit_tag"] = "rastasheep/ubuntu:12.04"
	_, warns, errs = NewConfig(raw)
	testConfigOk(t, warns, errs)
}

func TestConfigPrepare_pull(t *testing.T) {
	raw := testConfig()

	// No pull set
	delete(raw, "pull")
	c, warns, errs := NewConfig(raw)
	testConfigOk(t, warns, errs)
	if !c.Pull {
		t.Fatal("should pull by default")
	}

	// Pull set
	raw["pull"] = false
	c, warns, errs = NewConfig(raw)
	testConfigOk(t, warns, errs)
	if c.Pull {
		t.Fatal("should not pull")
	}
}
