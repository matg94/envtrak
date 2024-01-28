package resources

import "testing"
import "github.com/stretchr/testify/assert"

func TestServiceToYaml(t *testing.T) {
	service := &Service{
		AppId: "test-app",
		Version: &VVersion{
			Value: "v11",
		},
	}

	expectedYaml := "app-id: test-app\nversion:\n    value: v11\n"
	assert.Equal(t, expectedYaml, ToYaml(service), "Expected yaml to match")
}

func TestServiceToJson(t *testing.T) {

	service := &Service{
		AppId: "test-app",
		Version: &VVersion{
			Value: "v11",
		},
	}

	expectedJson := `{"appId":"test-app","version":{"value":"v11"}}`
	assert.Equal(t, expectedJson, ToJson(service), "Expected json to match")
}

func TestServiceGetRow(t *testing.T) {
	service := &Service{
		AppId: "test-app",
		Version: &VVersion{
			Value: "v11",
		},
	}

	expectedRow := []string{"test-app", "v11"}
	assert.Equal(t, expectedRow, service.GetRow(), "Expected row to match")
}
