package hammy

import "testing"

func TestSetConfigDefaults(t *testing.T) {
	var cfg Config
	err := SetConfigDefaults(&cfg)
	if err == nil {
		t.Errorf("Error should not be nil")
	}
	//Mandatory fields:
	cfg.Workers.CmdLine = "/bin/ls"
	cfg.CouchbaseTriggers.ConnectTo = "http://localhost:8091/"
	cfg.CouchbaseTriggers.Bucket = "default"
	cfg.CouchbaseStates.ConnectTo = "http://localhost:8091/"
	cfg.CouchbaseStates.Bucket = "default"


	//Retry...
	err = SetConfigDefaults(&cfg)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	if cfg.Http.Addr != ":4000" {
		t.Errorf("cfg.Http.Addr = %#v, expected %#v", cfg.Http.Addr, ":4000")
	}
	if cfg.CouchbaseStates.Ttl != 86400 {
		t.Errorf("cfg.CouchbaseStates.Ttl = %#v, expected 86400", cfg.CouchbaseStates.Ttl)
	}
}
