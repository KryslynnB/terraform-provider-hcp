package clients

import (
	"context"

	"github.com/hashicorp/cloud-sdk-go/clients/cloud-consul-service/preview/2020-08-26/client/consul_service"
	consulmodels "github.com/hashicorp/cloud-sdk-go/clients/cloud-consul-service/preview/2020-08-26/models"
	sharedmodels "github.com/hashicorp/cloud-sdk-go/clients/cloud-shared/v1/models"
)

// CreateSnapshot  will make a call to the Consul service to initiate the create Consul
// snapshot workflow.
func CreateSnapshot(ctx context.Context, client *Client, res *sharedmodels.HashicorpCloudLocationLink,
	snapshotName string) (*consulmodels.HashicorpCloudConsul20200826CreateSnapshotResponse, error) {

	p := consul_service.NewCreateSnapshotParams()
	p.Context = ctx
	p.ResourceLocationOrganizationID = res.Location.OrganizationID
	p.ResourceLocationProjectID = res.Location.ProjectID
	p.Body = &consulmodels.HashicorpCloudConsul20200826CreateSnapshotRequest{
		Name:     snapshotName,
		Resource: res,
	}

	resp, err := client.Consul.CreateSnapshot(p, nil)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

// GetSnapshotByID gets a Consul snapshot by its ID
func GetSnapshotByID(ctx context.Context, client *Client, loc *sharedmodels.HashicorpCloudLocationLocation,
	snapshotID string) (*consulmodels.HashicorpCloudConsul20200826GetSnapshotResponse, error) {

	p := consul_service.NewGetSnapshotParams()
	p.Context = ctx
	p.LocationOrganizationID = loc.OrganizationID
	p.LocationProjectID = loc.ProjectID
	p.SnapshotID = snapshotID

	resp, err := client.Consul.GetSnapshot(p, nil)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

// DeleteSnapshotByID deletes a Consul snapshot by its ID
func DeleteSnapshotByID(ctx context.Context, client *Client, loc *sharedmodels.HashicorpCloudLocationLocation,
	snapshotID string) (*consulmodels.HashicorpCloudConsul20200826DeleteSnapshotResponse, error) {

	p := consul_service.NewDeleteSnapshotParams()
	p.Context = ctx
	p.LocationOrganizationID = loc.OrganizationID
	p.LocationProjectID = loc.ProjectID
	p.SnapshotID = snapshotID

	resp, err := client.Consul.DeleteSnapshot(p, nil)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

// RenameSnapshotByID renames a Consul snapshot by its ID
func RenameSnapshotByID(ctx context.Context, client *Client, loc *sharedmodels.HashicorpCloudLocationLocation,
	snapshotID string, snapshotName string) (*consulmodels.HashicorpCloudConsul20200826UpdateSnapshotResponse, error) {

	p := consul_service.NewUpdateSnapshotParams()
	p.Context = ctx
	p.SnapshotLocationOrganizationID = loc.OrganizationID
	p.SnapshotLocationProjectID = loc.ProjectID
	p.SnapshotID = snapshotID

	p.Body = &consulmodels.HashicorpCloudConsul20200826Snapshot{
		Name: snapshotName,
	}

	resp, err := client.Consul.UpdateSnapshot(p, nil)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
