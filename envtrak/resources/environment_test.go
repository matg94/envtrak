package resources

import "testing"
import "github.com/stretchr/testify/assert"

func TestEnvironmentGetYaml(t *testing.T) {
    env := &Environment {
        Dtap: "dtap",
        BusinessId: "test",
        Alias: "use1",
        Region: "us-east-1",
    }

    expectedYaml := "dtap: dtap\nbusiness-id: test\nalias: use1\nregion: us-east-1\n"

    assert.Equal(t, expectedYaml, env.GetYaml(), "Expected yaml to match")
}

func TestEnvironmentGetJson(t *testing.T) {

    env := &Environment {
        Dtap: "dtap",
        BusinessId: "test",
        Alias: "use1",
        Region: "us-east-1",
    }
    expectedJson := "{\"dtap\":\"dtap\",\"businessId\":\"test\",\"alias\":\"use1\",\"region\":\"us-east-1\"}"
    assert.Equal(t, expectedJson, env.GetJson(), "Expected json to match")
} 

func TestEnvironmentGetRow(t *testing.T) {

    env := &Environment {
        Dtap: "dtap",
        BusinessId: "test",
        Alias: "use1",
        Region: "us-east-1",
    }
    
    expectedRow := []string{"test", "use1", "dtap", "us-east-1"}
    assert.Equal(t, expectedRow, env.GetRow(), "Expected row to match")
}
