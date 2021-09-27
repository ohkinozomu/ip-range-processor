package process

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type DatadogSyntheticsIPRange struct {
	Version    int    `json:"version"`
	Modified   string `json:"modified"`
	Synthetics struct {
		PrefixesIpv4           []string      `json:"prefixes_ipv4"`
		PrefixesIpv6           []interface{} `json:"prefixes_ipv6"`
		PrefixesIpv4ByLocation struct {
			AwsApNortheast1 []string `json:"aws:ap-northeast-1"`
			AwsApNortheast2 []string `json:"aws:ap-northeast-2"`
			AwsApSouth1     []string `json:"aws:ap-south-1"`
			AwsApSoutheast1 []string `json:"aws:ap-southeast-1"`
			AwsApSoutheast2 []string `json:"aws:ap-southeast-2"`
			AwsCaCentral1   []string `json:"aws:ca-central-1"`
			AwsEuCentral1   []string `json:"aws:eu-central-1"`
			AwsEuNorth1     []string `json:"aws:eu-north-1"`
			AwsEuWest1      []string `json:"aws:eu-west-1"`
			AwsEuWest2      []string `json:"aws:eu-west-2"`
			AwsEuWest3      []string `json:"aws:eu-west-3"`
			AwsSaEast1      []string `json:"aws:sa-east-1"`
			AwsUsEast2      []string `json:"aws:us-east-2"`
			AwsUsWest1      []string `json:"aws:us-west-1"`
			AwsUsWest2      []string `json:"aws:us-west-2"`
			AzureEastus     []string `json:"azure:eastus"`
		} `json:"prefixes_ipv4_by_location"`
		PrefixesIpv6ByLocation struct {
		} `json:"prefixes_ipv6_by_location"`
	} `json:"synthetics"`
}

func getJson() ([]byte, error) {
	// https://docs.datadoghq.com/synthetics/guide/identify_synthetics_bots/?tab=singleandmultistepapitests#ip-addresses
	url := "https://ip-ranges.datadoghq.com/synthetics.json"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	j, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func getAllIPsFromJson(j []byte) ([]string, error) {
	var dsir DatadogSyntheticsIPRange
	err := json.Unmarshal(j, &dsir)
	if err != nil {
		return nil, err
	}
	ss := [][]string{
		dsir.Synthetics.PrefixesIpv4,
		dsir.Synthetics.PrefixesIpv4ByLocation.AwsApNortheast1,
		dsir.Synthetics.PrefixesIpv4ByLocation.AwsApNortheast2,
		dsir.Synthetics.PrefixesIpv4ByLocation.AwsApSouth1,
		dsir.Synthetics.PrefixesIpv4ByLocation.AwsApSoutheast1,
		dsir.Synthetics.PrefixesIpv4ByLocation.AwsApSoutheast2,
		dsir.Synthetics.PrefixesIpv4ByLocation.AwsCaCentral1,
		dsir.Synthetics.PrefixesIpv4ByLocation.AwsEuCentral1,
		dsir.Synthetics.PrefixesIpv4ByLocation.AwsEuNorth1,
		dsir.Synthetics.PrefixesIpv4ByLocation.AwsEuWest1,
		dsir.Synthetics.PrefixesIpv4ByLocation.AwsEuWest2,
		dsir.Synthetics.PrefixesIpv4ByLocation.AwsEuWest3,
		dsir.Synthetics.PrefixesIpv4ByLocation.AwsSaEast1,
		dsir.Synthetics.PrefixesIpv4ByLocation.AwsUsEast2,
		dsir.Synthetics.PrefixesIpv4ByLocation.AwsUsWest1,
		dsir.Synthetics.PrefixesIpv4ByLocation.AwsUsWest2,
		dsir.Synthetics.PrefixesIpv4ByLocation.AzureEastus,
	}

	var IPs []string
	for _, s := range ss {
		IPs = append(IPs, s...)
	}
	return IPs, nil
}

func Process() error {
	json, err := getJson()
	if err != nil {
		return err
	}
	IPs, err := getAllIPsFromJson(json)
	if err != nil {
		return err
	}

	tmpl := `
{{define "base"}}
{{range $index, $ip := .IPs}}
ip_set_descriptors {
  type  = "IPV4"
  value = "{{$ip}}"
}
{{end}}
{{end}}
`
	params := map[string][]string{
		"IPs": IPs,
	}

	t, err := template.New("base").Parse(tmpl)
	if err != nil {
		return err
	}
	t.Execute(os.Stdout, params)
	return nil
}
