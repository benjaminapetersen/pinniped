// This go.mod file is generated by ./hack/codegen.sh.
module go.pinniped.dev/generated/1.23/client

go 1.13

require (
	go.pinniped.dev/generated/1.23/apis v0.0.0
	k8s.io/apimachinery v0.23.7
	k8s.io/client-go v0.23.7
)

replace go.pinniped.dev/generated/1.23/apis => ../apis
