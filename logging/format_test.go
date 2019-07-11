package logging

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func TestLogFormatMarshalYAML(t *testing.T) {
	var l Format
	err := l.Set("json")
	require.NoError(t, err)

	// Test the non-pointed to Format, as people might embed it.
	y, err := yaml.Marshal(l)
	require.NoError(t, err)
	require.Equal(t, []byte("json\n"), y)

	// And the pointed to Format.
	y, err = yaml.Marshal(&l)
	require.NoError(t, err)
	require.Equal(t, []byte("json\n"), y)
}
