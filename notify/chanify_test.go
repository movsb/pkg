package notify

import "testing"

func TestChanify(t *testing.T) {
	c := NewOfficialChanify(``)
	c.Send(`title`, `content`, true)
}
