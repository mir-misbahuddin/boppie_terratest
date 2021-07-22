package test
import (
        "testing"
        "github.com/gruntwork-io/terratest/modules/terraform"
        "github.com/stretchr/testify/assert"
)
func TestTerraformproject(t *testing.T) {
        // Run test in parallel
        t.Parallel()
        // Get the Project Id to use
        // projectId := gcp.GetGoogleProjectIDFromEnvVar(t)
        bucketName := "grunt-works02"
        // Get the name of the Cloud Router
        bucketLocation := "US"
        // Get the name of the NAT Instance
        bucketProject := "golang-misbah03"
        // Get Compute Address Name
        //computeName := "hopper-address"
        // Get the name of the network
        //networkName := "test-network"
        // Get the name of the Region
        //regionName := "us-central1"
        terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
                // The path to Terraform code
                TerraformDir: "../modules",
                // Variables to pass to our Terraform code using -var options
                Vars: map[string]interface{}{
                        "name":           bucketName,
                        "location":  bucketLocation,
                        "project":     bucketProject,
                        //"name_address": computeName,
                        //"network_name": networkName,
                        //"region":       regionName,
                 },
        })
        // At the end of the test, run `terraform destroy` to clean up any resources that were created
        defer terraform.Destroy(t, terraformOptions)
        // This will run `terraform init` and `terraform apply` and fail the test if there are any errors
        terraform.InitAndApply(t, terraformOptions)
        // Run `terraform output` to get the values of output variables for nat instance name
        cloudBucketName := terraform.Output(t, terraformOptions, "bucket_name")
        cloudBucketLocation := terraform.Output(t, terraformOptions, "bucket_location")
        cloudBucketProject := terraform.Output(t, terraformOptions, "bucket_project")
        // Verify we're getting back the outputs we expect
        assert.Equal(t, bucketName, cloudBucketName)
        assert.Equal(t, bucketLocation, cloudBucketLocation)
        assert.Equal(t, bucketProject, cloudBucketProject)
}
