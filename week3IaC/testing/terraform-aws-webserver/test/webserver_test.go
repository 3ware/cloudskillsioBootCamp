package test


// Import the required libraries
import (
	"testing"
	"fmt"
	"time"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"	
)

// function to define path and variables
func TestTerraformWebserverExample(t *testing.T) {

	// The values to pass to the Terraform CLI. t for test
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		// The path to the example terraform code
		TerraformDir: "../examples/webserver",

		// Variables to pass to the terraform code using -var options
		Vars: map[string]interface{}{
			"region": "eu-west-2",
			"servername": "testwebserver",
		},
	})
	// Call terraform library from gruntwork import. 
	// Run init and apply function from library.
	// Use the path and variables defined in the function above
	terraform.InitAndApply(t, terraformOptions)

	// Run a destroy at the end of the test.
	// defer postpones the execution of a function until the surrounding function returns, either normally or through a panic.
	defer terraform.Destroy(t, terraformOptions)

	// Define public IP variable by reading the outputs.tf file
	publicIP := terraform.Output(t, terraformOptions, "public_ip")

	// Define rule URL variable. Formart as a string, replace %s with publicIP 
	url := fmt.Sprintf("http://%s:8080", publicIP)

	// test, url variable, nil = no tls, 200 = http response code, try 30 times, every 5 seconds
	http_helper.HttpGetWithRetry(t, url, nil, 200, "I made a Terraform Module!", 30, 5*time.Second)
}
