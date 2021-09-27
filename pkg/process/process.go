package process

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

func Process() error {
	// https://docs.datadoghq.com/synthetics/guide/identify_synthetics_bots/?tab=singleandmultistepapitests#ip-addresses
	url := "https://ip-ranges.datadoghq.com/synthetics.json"

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(byteArray))
	return nil
}
