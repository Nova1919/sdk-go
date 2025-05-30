package env

import "testing"

func TestGetLogEnv(t *testing.T) {
	logEnv := GetLogEnv()
	t.Logf("%+v", logEnv)
}

func TestGetActorEnv(t *testing.T) {
	actorEnv := GetActorEnv()
	t.Logf("%+v", actorEnv)
}
