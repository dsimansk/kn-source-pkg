module github.com/maximilien/kn-source-pkg

go 1.15

require (
	github.com/maxbrunsfeld/counterfeiter/v6 v6.3.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	gotest.tools/v3 v3.0.3
	k8s.io/api v0.19.7
	k8s.io/client-go v0.19.7
	knative.dev/client v0.20.1-0.20210219184703-18ff59604195
	knative.dev/pkg v0.0.0-20210219164703-8a9bf766d36d
)

replace github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.3
