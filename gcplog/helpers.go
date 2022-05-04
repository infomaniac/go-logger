package gcplog

import (
	"context"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

// CurrentProject returns the current GCP project ID where the application is running.
func CurrentProject() (string, error) {
	ctx := context.Background()
	credentials, err := google.FindDefaultCredentials(ctx, compute.ComputeScope)
	if err != nil {
		return "", err
	}
	return credentials.ProjectID, nil
}
