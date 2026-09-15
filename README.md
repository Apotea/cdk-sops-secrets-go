# cdk-sops-secrets-go

Go bindings for [dbsystel/cdk-sops-secrets](https://github.com/dbsystel/cdk-sops-secrets),
generated with jsii.

```bash
go get github.com/Apotea/cdk-sops-secrets-go/cdksopssecrets/v2
```

```go
cdksopssecrets.NewSopsSecret(stack, jsii.String("Secret"), &cdksopssecrets.SopsSecretProps{
    SopsFilePath: jsii.String("secret.sops.yaml"),
})
```

The construct API matches upstream; their
[README](https://github.com/dbsystel/cdk-sops-secrets#readme) is the reference.

## Versions

Tags mirror the upstream release they were generated from: `cdksopssecrets/v2.8.5` carries
`cdk-sops-secrets@2.8.5`, including that release's Lambda binary.

`.github/workflows/release.yml` generates a version: it packs the published npm package, adds a
jsii Go target, runs `jsii-pacmak`, synthesizes a stack through the result, then tags the
generated sources. Bindings are generated against the lowest `aws-cdk-lib` the upstream release
supports, so the module's floor stays low — Go resolves to the maximum across a build.

Upstream tracks Go support in
[dbsystel/cdk-sops-secrets#1434](https://github.com/dbsystel/cdk-sops-secrets/issues/1434). This
repo is archived once they publish Go bindings themselves.

## License

Apache-2.0, as upstream. Generated sources carry upstream's `LICENSE`.
