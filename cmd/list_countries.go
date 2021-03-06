package cmd

import (
	"context"
	"github.com/ProspectOne/perfops-cli/cmd/internal"
	"github.com/ProspectOne/perfops-cli/perfops"
	"net/http"
)

func runCountriesCmd(c *perfops.Client) error {
	var res *[]perfops.Country

	ctx := context.Background()
	u := c.BasePath + "/analytics/dns/countries"

	f := internal.NewFormatter(debug)
	f.StartSpinner()

	req, _ := http.NewRequest("GET", u, nil)
	req = req.WithContext(ctx)

	err := c.DoRequest(req, &res);
	f.StopSpinner()

	if err != nil {
		return err
	}

	internal.PrintOutputJSON(res)

	return nil
}
