package netutil

import (
	"testing"
)

func TestInternalIP(t *testing.T) {
	t.Log(InternalIP())
}

func TestInternalIPOld(t *testing.T) {
	t.Log(InternalIPOld())
}

func TestInternalIPv4(t *testing.T) {
	t.Log(InternalIPv4())
}

func TestInternalIPv6(t *testing.T) {
	t.Log(InternalIPv6())
}
