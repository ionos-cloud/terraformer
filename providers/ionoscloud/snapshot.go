package ionoscloud

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/providers/ionoscloud/helpers"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"log"
)

type SnapshotGenerator struct {
	Service
}

func (g *SnapshotGenerator) InitResources() error {
	client := g.generateClient()
	cloudAPIClient := client.CloudAPIClient
	resourceType := "ionoscloud_snapshot"

	snapshots, _, err := cloudAPIClient.SnapshotsApi.SnapshotsGet(context.TODO()).Depth(1).Execute()
	if err != nil {
		return err
	}
	if snapshots.Items == nil {
		log.Printf("[WARNING] expected a response containing snapshots but received 'nil' instead.")
		return nil
	}
	for _, snapshot := range *snapshots.Items {
		if snapshot.Properties == nil || snapshot.Properties.Name == nil {
			log.Printf("[WARNING] 'nil' values in the response for the snapshot with ID: %s, skipping this resource.", *snapshot.Id)
			continue
		}
		g.Resources = append(g.Resources, terraformutils.NewResource(
			*snapshot.Id,
			*snapshot.Properties.Name+"-"+*snapshot.Id,
			resourceType,
			helpers.Ionos,
			map[string]string{},
			[]string{},
			map[string]interface{}{}))
	}
	return nil
}
