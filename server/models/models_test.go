package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetStatusMap(t *testing.T) {
	gotMap := GetStatusMap()

	var fail bool
	var notFoundStatuses []Status
	for _, status := range StatusTypes {
		if s, ok := gotMap[status.Name]; !ok || s != status.ID {
			notFoundStatuses = append(notFoundStatuses, status)
			fail = true
		}
	}

	assert.False(t, fail, "status(es) not found in status map: %v", notFoundStatuses)
}
