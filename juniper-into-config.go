package main

import "fmt"

func main() {
    ifaceName := "ge-0/0/0"
    vlanTag := 100
    ipv4Addr := "192.168.1.1/24"
    description := "Sample interface configuration"

    template := `
interfaces {
    {{ .InterfaceName }} {
        unit 0 {
            vlan-id {{ .VlanTag }};
            family inet {
                address {{ .IPv4Address }};
            }
        }
        description "{{ .Description }}";
    }
}`

    data := struct {
        InterfaceName string
        VlanTag       int
        IPv4Address   string
        Description   string
    }{
        InterfaceName: ifaceName,
        VlanTag:       vlanTag,
        IPv4Address:   ipv4Addr,
        Description:   description,
    }

    config := generateConfig(template, data)
    fmt.Println(config)
}

func generateConfig(template string, data interface{}) string {
    t := template.Must(template.New("").Parse(template))
    var output []byte
    t.Execute(&output, data)
    return string(output)
}
