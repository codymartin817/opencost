package config

import "testing"

// Repro for https://github.com/opencost/opencost/issues/4029: UnmarshalJSON must not
// stack-overflow via infinite self-recursion when using the confUnmarshaller trick.
func TestConfigurationsUnmarshalJSON_NoRecursion(t *testing.T) {
	data := []byte(`{"aws":{"athena":[{"account":"123456789012","authorizer":{"authorizer":{"authorizerType":"AWSServiceAccount"},"authorizerType":"AWSAssumeRole","roleARN":"arn:aws:iam::123456789012:role/opencost-cur-reader"},"bucket":"s3://example-athena-results-123456789012","database":"athenacurcfn_example","region":"us-east-1","table":"exampletable","workgroup":"primary"}]}}`)

	c := &Configurations{}
	if err := c.UnmarshalJSON(data); err != nil {
		t.Fatalf("UnmarshalJSON returned error: %v", err)
	}
}

func TestConfigurationsUnmarshalJSON_NewFormat(t *testing.T) {
	data := []byte(`{}`)

	c := &Configurations{}
	if err := c.UnmarshalJSON(data); err != nil {
		t.Fatalf("UnmarshalJSON returned error: %v", err)
	}
}
