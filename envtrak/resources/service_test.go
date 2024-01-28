package resources

import "testing"
import "github.com/stretchr/testify/assert"

func TestServiceGetYaml(t *testing.T) {
    service := &Service{
        AppId: "test-app",
        Version: &VVersion{
            Value: "v11",
        },
    }

    expectedYaml := "app-id: test-app\nversion:\n    value: v11\n"
    assert.Equal(t, expectedYaml, service.GetYaml(), "Expected yaml to match")
}

func TestServiceGetJson(t *testing.T) {

    service := &Service{
        AppId: "test-app",
        Version: &VVersion{
            Value: "v11",
        },
    }

    expectedJson := `{"appId":"test-app","version":{"value":"v11"}}`
    assert.Equal(t, expectedJson, service.GetJson(), "Expected json to match")
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
