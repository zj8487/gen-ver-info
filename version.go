package main

import "time"

const (
	//Version indicate version from version branch"
	Version              = "1e34bcc"
	//VersionControlSystem indicate version mananger vender"
	VersionControlSystem = "git"
)
var (
	//VersionTime indicate commit time in version branch"
	VersionTime      = time.Date(2016, 2, 16, 22, 40, 0, 0, time.UTC)
	//VersionBuildTime  indicate build time"
	VersionBuildTime = time.Date(2016, 3, 28, 3, 25, 48, 0, time.UTC)
)
