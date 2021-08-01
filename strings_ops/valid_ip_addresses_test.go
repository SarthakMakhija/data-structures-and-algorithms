package stringsops_test

import (
	stringsops "github.com/SarthakMakhija/data-structures-and-algorithms/strings_ops"
	"reflect"
	"testing"
)

func TestValidIPAddressesFrom(t *testing.T) {

	ips := stringsops.ValidIPAddressesFrom("192168")
	expected := []string{"1.9.2.168", "1.9.21.68", "1.9.216.8", "1.92.1.68", "1.92.16.8", "19.2.1.68", "19.2.16.8", "19.21.6.8", "192.1.6.8"}

	if !reflect.DeepEqual(ips, expected) {
		t.Fatalf("Expected %v, received %v", expected, ips)
	}
}
