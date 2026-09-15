package main

import (
	"github.com/Apotea/cdk-sops-secrets-go/cdksopssecrets/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()
	app := awscdk.NewApp(&awscdk.AppProps{Outdir: jsii.String("cdk.out")})
	stack := awscdk.NewStack(app, jsii.String("SmokeStack"), nil)
	cdksopssecrets.NewSopsSecret(stack, jsii.String("Secret"), &cdksopssecrets.SopsSecretProps{
		SopsFilePath: jsii.String("secret.sops.yaml"),
	})
	app.Synth(nil)
}
