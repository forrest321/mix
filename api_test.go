// Package atomix is a sequence-based Go-native audio mixer
package atomix

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/outrightmental/go-atomix/bind"
)

func TestAPI_Debug(t *testing.T) {
	// TODO: Test API Debug
}

func TestAPI_Configure(t *testing.T) {
	// TODO: Test API Configure
}

func TestAPI_Teardown(t *testing.T) {
	// TODO: Test API Teardown
}

func TestAPI_Spec(t *testing.T) {
	// TODO: Test API Spec
}

func TestAPI_SetFire(t *testing.T) {
	testAPISetup()
	fire := SetFire("lib/S16.wav", time.Duration(0), 0, 1.0, 0)
	assert.NotNil(t, fire)
}

func TestAPI_SetSoundsPath(t *testing.T) {
	// TODO: Test API SetSoundsPath
}

func TestAPI_Start(t *testing.T) {
	Start()
}

func TestAPI_StartAt(t *testing.T) {
	StartAt(time.Now().Add(1 * time.Second))
}

func TestAPI_GetStartTime(t *testing.T) {
	startExpect := time.Now().Add(1 * time.Second)
	StartAt(startExpect)
	startActual := GetStartTime()
	assert.Equal(t, startExpect, startActual)
}

func TestAPI_AudioCallback(t *testing.T) {
	// TODO: Test API AudioCallback
}

//
// Test Components
//

func testAPISetup() {
	Configure(bind.AudioSpec{
		Freq:     44100,
		Format:   bind.AudioF32,
		Channels: 1,
	})
}
