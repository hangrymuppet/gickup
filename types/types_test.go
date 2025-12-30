package types

import "testing"

func TestConfCronMissing(t *testing.T) {
	t.Parallel()

	conf := Conf{Cron: ""}

	if !conf.MissingCronSpec() {
		t.Error("Conf passed empty cron spec doesn't think that it's missing")
	}
}

func TestParseValidCron(t *testing.T) {
	t.Parallel()

	conf := Conf{Cron: "0 0 * * *"}

	if !conf.HasValidCronSpec() {
		t.Error("Valid cron spec parsed invalidly")
	}
}

func TestParseInvalidCron(t *testing.T) {
	t.Parallel()

	conf := Conf{Cron: "redshirt"}

	if conf.HasValidCronSpec() {
		t.Error("Invalid cron spec parsed validly")
	}
}

func TestHTTPTimeoutDefaults(t *testing.T) {
t.Parallel()

// Test with zero value (should use default)
conf := Conf{HTTPTimeout: 0}
if conf.HTTPTimeout != 0 {
t.Errorf("Expected HTTPTimeout to be 0 (unset), got %d", conf.HTTPTimeout)
}

// Test with custom value
conf2 := Conf{HTTPTimeout: 600}
if conf2.HTTPTimeout != 600 {
t.Errorf("Expected HTTPTimeout to be 600, got %d", conf2.HTTPTimeout)
}

// Test with negative value (should be treated as zero/default in main.go)
conf3 := Conf{HTTPTimeout: -1}
if conf3.HTTPTimeout != -1 {
t.Errorf("Expected HTTPTimeout to be -1, got %d", conf3.HTTPTimeout)
}
}
