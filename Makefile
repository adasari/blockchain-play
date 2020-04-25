sync:
	bazel run //:gazelle
	go mod tidy                         
	bazel run //:gazelle -- update-repos -from_file=go.mod
