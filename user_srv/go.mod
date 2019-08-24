module oauth_srv

go 1.12

require (
	github.com/SermoDigital/jose v0.9.2-0.20161205224733-f6df55f235c2 // indirect
	github.com/asaskevich/govalidator v0.0.0-20180720115003-f9ffefc3facf // indirect
	github.com/elazarl/go-bindata-assetfs v1.0.0 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/micro/go-micro v1.9.1
	github.com/micro/go-plugins v1.2.0
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	go.uber.org/zap v1.10.0
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce // indirect
	utils v0.0.0
)

replace (
	github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
	github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.0-20190108154635-47c0da630f72
	github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
	utils => ../utils
)
