package main

import (
	"encoding/json"
	"github.com/mitchellh/packer/packer"
	"github.com/mitchellh/packer/packer/plugin"
	"io/ioutil"
	"strings"
)

type PostProcessor struct {
	AMI string `json:"ami"`
}

func (p *PostProcessor) Configure(raws ...interface{}) error {
	return nil
}

func (p *PostProcessor) PostProcess(ui packer.Ui, artifact packer.Artifact) (packer.Artifact, bool, error) {
	amiIDAndRegion := strings.Split(artifact.Id(), ":")
	amiID := amiIDAndRegion[1]
	b, err := json.Marshal(PostProcessor{amiID})
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("packer_ami.json", b, 0644)
	if err != nil {
		panic(err)
	}

	return artifact, true, nil
}

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}

	server.RegisterPostProcessor(new(PostProcessor))
	server.Serve()
}
