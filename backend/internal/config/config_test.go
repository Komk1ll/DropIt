package config

import "testing"

func TestLoadUsesDefaultPort(t *testing.T) {
	t.Setenv("PORT", "")

	cfg := Load()

	if cfg.Port != defaultPort {
		t.Fatalf("expected default port %q, got %q", defaultPort, cfg.Port)
	}
	if cfg.Addr() != ":"+defaultPort {
		t.Fatalf("expected default addr %q, got %q", ":"+defaultPort, cfg.Addr())
	}
}

func TestLoadUsesEnvPort(t *testing.T) {
	t.Setenv("PORT", "9090")

	cfg := Load()

	if cfg.Port != "9090" {
		t.Fatalf("expected port %q, got %q", "9090", cfg.Port)
	}
	if cfg.Addr() != ":9090" {
		t.Fatalf("expected addr %q, got %q", ":9090", cfg.Addr())
	}
}