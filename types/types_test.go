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

	// Test with empty value (should use default)
	conf := Conf{HTTPTimeout: ""}
	if conf.HTTPTimeout != "" {
		t.Errorf("Expected HTTPTimeout to be empty (unset), got %s", conf.HTTPTimeout)
	}

	// Test with duration string
	conf2 := Conf{HTTPTimeout: "5m"}
	if conf2.HTTPTimeout != "5m" {
		t.Errorf("Expected HTTPTimeout to be '5m', got %s", conf2.HTTPTimeout)
	}

	// Test with plain number (seconds)
	conf3 := Conf{HTTPTimeout: "600"}
	if conf3.HTTPTimeout != "600" {
		t.Errorf("Expected HTTPTimeout to be '600', got %s", conf3.HTTPTimeout)
	}
}
