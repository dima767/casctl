// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// uptimeCmd represents the uptime command
var uptimeCmd = &cobra.Command{
	Use:   "uptime [URL of CAS server (with trailing slash)]",
	Short: "Display an uptime of a running CAS server.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			erAndExit("[ssosessioins] needs a URL of CAS server")
		}
		casServerBaseUrl = args[0]
		displayCasServerUptime()
	},
}

func init() {
	RootCmd.AddCommand(uptimeCmd)
}

type AvailabilityReport struct {
	UpTime int
}

func (s *casReportingService) getCasServerAvailability() (*AvailabilityReport, *http.Response, error) {
	availability := &AvailabilityReport{}
	resp, err := s.sling.New().Path("status/stats/getAvailability").ReceiveSuccess(availability)
	return availability, resp, err
}

func displayCasServerUptime() {
	casReportingService := newCasReportingService(nil)
	availability, resp, _ := casReportingService.getCasServerAvailability()
	checkResponseAndExitIfNecessary(resp)

	uptimeDuration, _ := time.ParseDuration(strconv.Itoa(availability.UpTime) + "s")
	uptimeString := fmt.Sprintf("CAS server %s uptime: %v", casServerBaseUrl, uptimeDuration)
	fmt.Println()
	greenPrintln(uptimeString)
}
