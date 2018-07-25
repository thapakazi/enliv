package library

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/parnurzeal/gorequest"
)

type Job struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	Remarks     string `json:"remarks"`
	PublisedAt  string `json:"published_at"`
}

func ListJobs() []Job {
	jobsAPI := os.Getenv("MOCK_API_URL")

	_, body, err := gorequest.New().Get(jobsAPI).End()
	if err != nil {
		fmt.Println("ERROR: fetching jobs list", err)
	}

	// unmarshall to strcut
	var jsonBlob = []byte(body)
	var jobs = []Job{}
	e := json.Unmarshal(jsonBlob, &jobs)
	if e != nil {
		fmt.Println("ERROR: on unmarshal of get jobs response json", err)
	}
	return jobs
}

func PrintJobs(jobs []Job) {

	// TODO: output to table; move it to cli
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Description", "Remarks", "PublishedAt"})
	// table.SetBorder(false)
	// table.SetCenterSeparator("")
	// table.SetColumnSeparator(",")
	// table.SetRowSeparator(" ")
	for _, v := range jobs {
		jobID := fmt.Sprintf("%v", v.Id)
		row := []string{jobID, v.Title, v.Description, v.Remarks, v.PublisedAt}
		table.Append(row)
	}
	table.Render() // Send output
}
