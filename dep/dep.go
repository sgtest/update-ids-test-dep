package dep

import (
	"github.com/sgtest/update-ids-test-base/base"
	"github.com/sgtest/update-ids-test-otherbase/otherbase"
)

func Bar() {
	base.Foo()
	otherbase.Qux()
}
