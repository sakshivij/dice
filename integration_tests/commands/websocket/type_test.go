package websocket

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestType(t *testing.T) {
	exec := NewWebsocketCommandExecutor()

	testCases := []TestCase{
		{
			name:     "TYPE with invalid number of arguments",
			commands: []string{"TYPE"},
			expected: []interface{}{"ERR wrong number of arguments for 'type' command"},
		},
		{
			name:     "TYPE for non-existent key",
			commands: []string{"TYPE k1"},
			expected: []interface{}{"none"},
		},
		{
			name:     "TYPE for key with String value",
			commands: []string{"SET k1 v1", "TYPE k1"},
			expected: []interface{}{"OK", "string"},
		},
		{
			name:     "TYPE for key with List value",
			commands: []string{"LPUSH k1 v1", "TYPE k1"},
			expected: []interface{}{float64(1), "list"},
		},
		{
			name:     "TYPE for key with Set value",
			commands: []string{"SADD k1 v1", "TYPE k1"},
			expected: []interface{}{float64(1), "set"},
		},
		{
			name:     "TYPE for key with Hash value",
			commands: []string{"HSET k1 field1 v1", "TYPE k1"},
			expected: []interface{}{float64(1), "hash"},
		},
		{
			name:     "TYPE for key with value created from SETBIT command",
			commands: []string{"SETBIT k1 1 1", "TYPE k1"},
			expected: []interface{}{float64(0), "string"},
		},
		// TODO: uncomment when bitop is added
		// {
		// 	name:     "TYPE for key with value created from SETOP command",
		// 	commands: []string{"SET key1 foobar", "SET key2 abcdef", "TYPE dest"},
		// 	expected: []interface{}{"OK", "OK", "string"},
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			conn := exec.ConnectToServer()
			exec.FireCommandAndReadResponse(conn, "FLUSHDB")
			for i, cmd := range tc.commands {
				result, err := exec.FireCommandAndReadResponse(conn, cmd)
				assert.Nil(t, err)
				assert.Equal(t, tc.expected[i], result)
			}
		})
	}
}
