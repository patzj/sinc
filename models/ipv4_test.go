package models

import "testing"

func TestIpv4IsBefore(t *testing.T) {
	var hostA, hostB *IPv4
	var actual, expected bool

	// A is before B
	hostA, _ = NewIPv4("172.16.0.1")
	hostB, _ = NewIPv4("172.16.0.2")

	actual = hostA.IsBefore(*hostB)
	expected = true

	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}

	// A is after B
	hostA, _ = NewIPv4("172.16.0.2")
	hostB, _ = NewIPv4("172.16.0.1")

	actual = hostA.IsBefore(*hostB)
	expected = false

	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}

	// A is equal to B
	hostA, _ = NewIPv4("172.16.0.1")
	hostB, _ = NewIPv4("172.16.0.1")

	actual = hostA.IsBefore(*hostB)
	expected = false

	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}

}

func TestIpv4IsAfter(t *testing.T) {
	var hostA, hostB *IPv4
	var actual, expected bool

	// A is before B
	hostA, _ = NewIPv4("172.16.0.1")
	hostB, _ = NewIPv4("172.16.0.2")

	actual = hostA.IsAfter(*hostB)
	expected = false

	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}

	// A is after B
	hostA, _ = NewIPv4("172.16.0.2")
	hostB, _ = NewIPv4("172.16.0.1")

	actual = hostA.IsAfter(*hostB)
	expected = true

	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}

	// A is equal to B
	hostA, _ = NewIPv4("172.16.0.1")
	hostB, _ = NewIPv4("172.16.0.1")

	actual = hostA.IsAfter(*hostB)
	expected = false

	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}

}

func TestIpv4IsEqual(t *testing.T) {
	var hostA, hostB *IPv4
	var actual, expected bool

	// A is before B
	hostA, _ = NewIPv4("172.16.0.1")
	hostB, _ = NewIPv4("172.16.0.2")

	actual = hostA.IsEqual(*hostB)
	expected = false

	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}

	// A is after B
	hostA, _ = NewIPv4("172.16.0.2")
	hostB, _ = NewIPv4("172.16.0.1")

	actual = hostA.IsEqual(*hostB)
	expected = false

	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}

	// A is equal to B
	hostA, _ = NewIPv4("172.16.0.1")
	hostB, _ = NewIPv4("172.16.0.1")

	actual = hostA.IsEqual(*hostB)
	expected = true

	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}

}
