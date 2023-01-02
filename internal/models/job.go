package models

import (
	"encoding/json"
	"job-board-api/pkg/utils"
)

type JobParams struct {
	JobID       string `json:"job_id"`
	Description string `json:"description"`
	Location    string `json:"location"`
	FullTime    string `json:"full_time"`
}

type Job struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	URL         string `json:"url"`
	CreatedAt   string `json:"created_at"`
	Company     string `json:"company"`
	CompanyURL  string `json:"company_url"`
	Location    string `json:"location"`
	Title       string `json:"title"`
	Description string `json:"description"`
	HowToApply  string `json:"how_to_apply"`
	CompanyLogo string `json:"company_logo"`
}

func GetJobs(params JobParams) ([]Job, error) {
	paramsMap := map[string]string{
		"job_id":      "",
		"description": "",
		"location":    "",
		"full_time":   "",
	}
	var jobs []Job
	if params.JobID != "" {
		paramsMap["job_id"] = params.JobID
	}
	if params.Description != "" {
		paramsMap["description"] = params.Description
	}
	if params.Location != "" {
		paramsMap["location"] = params.Location
	}
	resp, err := utils.GetRequestJob("http://dev3.dansmultipro.co.id/api/recruitment/positions.json", "http://dev3.dansmultipro.co.id/api/recruitment/positions", paramsMap)
	if err != nil {
		return nil, err
	}
	if params.JobID != "" {
		var job Job
		err = json.Unmarshal([]byte(resp), &job)
		if err != nil {
			return nil, err
		}
		return []Job{job}, nil
	} else {
		err = json.Unmarshal([]byte(resp), &jobs)
		if err != nil {
			return nil, err
		}
		if params.FullTime != "" && len(jobs) > 0 && params.JobID == "" {
			var fullTimeJobs = []Job{}
			for _, job := range jobs {
				if params.FullTime == "true" {
					if job.Type == "Full Time" {
						fullTimeJobs = append(fullTimeJobs, job)
					}
				} else {
					if job.Type != "Full Time" {
						fullTimeJobs = append(fullTimeJobs, job)
					}
				}
			}
			return fullTimeJobs, nil
		}
	}
	return jobs, nil
}
