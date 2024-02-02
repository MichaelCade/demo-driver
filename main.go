package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Example: Loading HTML templates from a 'templates' directory
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	// Define a route to display lab information
	router.GET("/labs", func(c *gin.Context) {
		// Retrieve lab information from Pulumi or a database
		// You can fetch details like lab name, status, etc.
		// Display the information in the HTML template
		c.HTML(http.StatusOK, "labs.html", gin.H{
			"title": "Choose your Demo Driver",
			"labs":  getLabInformation(),
		})
	})

	// Define a route to create a new lab
	router.POST("/createLab", func(c *gin.Context) {
		// Extract parameters from the request, e.g., lab name, cloud provider, etc.
		labName := c.PostForm("labName")
		cloudProvider := c.PostForm("cloudProvider")

		// Call the Pulumi function to create the lab
		createLab(labName, cloudProvider)

		// Redirect to the lab information page after creating the lab
		c.Redirect(http.StatusSeeOther, "/labs")
	})

	// Define a route to destroy a lab
	router.POST("/destroyLab", func(c *gin.Context) {
		// Extract parameters from the request, e.g., lab name
		labName := c.PostForm("labName")

		// Call the Pulumi function to destroy the lab
		destroyLab(labName)

		// Redirect to the lab information page after destroying the lab
		c.Redirect(http.StatusSeeOther, "/labs")
	})

	// Run the web server
	router.Run(":8080")
}

// Example function to retrieve lab information (replace with your implementation)
func getLabInformation() []string {
	// Retrieve lab information from Pulumi or a database
	// Return details like lab name, status, etc.
	return []string{"Lab 1", "Lab 2", "Lab 3"}
}

// Example function to create a lab (replace with your Pulumi implementation)
func createLab(labName string, cloudProvider string) {
	// Call Pulumi to create the lab using the provided parameters
	// Add your Pulumi code here
}

// Example function to destroy a lab (replace with your Pulumi implementation)
func destroyLab(labName string) {
	// Call Pulumi to destroy the lab using the provided lab name
	// Add your Pulumi code here
}
