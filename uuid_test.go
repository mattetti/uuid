package uuid

import (
	"regexp"
	"testing"
)

func TestUUID(t *testing.T) {
	uuid := GenUUID()
	if len(uuid) != 43 {
		t.Errorf("generated uuid expected of lenght %d but instead was %d", 43, len(uuid))
	}
	t.Logf("uuid[%s]\n", uuid)
}

func TestUuidFormat(t *testing.T) {
	var format = regexp.MustCompile(`[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`)
	uuid := GenUUID()
	if !format.MatchString(uuid) {
		t.Errorf("%s doesn't match %v", uuid, format)
	}
}

func BenchmarkUUID(b *testing.B) {
	m := make(map[string]int, 1000)
	for i := 0; i < b.N; i++ {
		uuid := GenUUID()
		b.StopTimer()
		c := m[uuid]
		if c > 0 {
			b.Fatalf("duplicate uuid[%s] count %d", uuid, c)
		}
		m[uuid] = c + 1
		b.StartTimer()
	}
}
