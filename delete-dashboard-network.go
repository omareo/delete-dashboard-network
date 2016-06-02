package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Network struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organizationid"`
	Type           string `json:"type"`
	Name           string `json:"name"`
	TimeZone       string `json:"timeZone"`
	Tags           string `json:"tags"`
}

func stripchars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}

func main() {

	DASHBOARD_API_SHARD_ID := os.Getenv("DASHBOARD_API_SHARD_ID")
	DASHBOARD_API_ORG_ID := os.Getenv("DASHBOARD_API_ORG_ID")
	DASHBOARD_API_KEY := os.Getenv("DASHBOARD_API_KEY")

	var err error

	// Fire off http request to get list of Networks and return results to console

	getNetUrl := "https://" + DASHBOARD_API_SHARD_ID + ".meraki.com/api/v0/organizations/" + DASHBOARD_API_ORG_ID + "/networks"
	getNetworks, err := http.NewRequest("GET", getNetUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{}
	getNetworks.Header.Add("X-Cisco-Meraki-API-Key", DASHBOARD_API_KEY)
	requestResult, err := client.Do(getNetworks)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(requestResult.Body)
	requestResult.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var networks []Network
	json.Unmarshal(body, &networks)

	for n := range networks {
		fmt.Printf(networks[n].Name)
		fmt.Println()
	}

	// get user input, select network to drop
	fmt.Println("Type in the name you would like to remove")

	r := bufio.NewReader(os.Stdin)
	const endOfString = '\n'

	userInputNetworkName, err := r.ReadString(endOfString)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fixedUserInputNetworkName := stripchars(userInputNetworkName, "\n")

	for n := range networks {
		//confirmation of user selection
		if fixedUserInputNetworkName == networks[n].Name {
			fixedUserInputNetworkName := stripchars(userInputNetworkName, "\n")
			fmt.Println("the following network will be removed:")
			fmt.Printf(networks[n].Name)
			fmt.Println()
			for n := range networks {

				if fixedUserInputNetworkName == networks[n].Name {
					fmt.Println("Please confirm with Y")

					b := bufio.NewReader(os.Stdin)
					const endOfString = '\n'

					userInputChoice, err := b.ReadString(endOfString)
					fixeduserInputChoice := stripchars(userInputChoice, "\n")

					if err != nil {
						log.Fatal(err)
						os.Exit(1)
					}

					if fixeduserInputChoice == "Y" {
						//build url
						deletNetUrl := "https://" + DASHBOARD_API_SHARD_ID + ".meraki.com/api/v0/networks/" + networks[n].ID

						// Fire off http DELETE
						deleteNetwork, err := http.NewRequest("DELETE", deletNetUrl, nil)
						if err != nil {
							log.Fatal(err)
						}

						client := http.Client{}
						deleteNetwork.Header.Add("X-Cisco-Meraki-API-Key", DASHBOARD_API_KEY)
						requestResult, err := client.Do(deleteNetwork)
						if err != nil {
							log.Fatal(err)
						}
						requestResult.Body.Close()
						fmt.Println("Network has been deleted.")

					}
				}
			}
		}
	}
}
