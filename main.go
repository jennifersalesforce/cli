// +build !test

package main

import (
	"os"
	"runtime"

	"github.com/pkg/profile"
)

var herokuProfile = getProfile()

func getProfile() interface {
	Stop()
} {
	runtime.SetCPUProfileRate(10000000)
	return profile.Start(profile.TraceProfile, profile.ProfilePath("."), profile.NoShutdownHook)
	//return profile.Start(profile.CPUProfile, profile.ProfilePath("."), profile.NoShutdownHook)
}

func main() {
	defer handlePanic()
	runtime.GOMAXPROCS(1) // more procs causes runtime: failed to create new OS thread on Ubuntu

	// handle sigint
	handleSignal(os.Interrupt, func() {
		if !swallowSigint {
			ShowCursor()
			herokuProfile.Stop()
			os.Exit(1)
		}
	})

	Start(os.Args...)
	herokuProfile.Stop()
	Exit(0)
}
