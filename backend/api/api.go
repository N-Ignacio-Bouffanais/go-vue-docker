package api

import (
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/ssh"
)

type Client struct {
	baseURL string
}

type ApiResponse struct {
	Message string `json:"message"`
}
type Data struct {
	// Define your data structure here
	Items []string `json:"items"`
}

func NewClient(baseURL string) *Client {
	return &Client{baseURL: baseURL}
}

func ServerGo() {
	app := fiber.New()

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hola mundo"})
	})

	app.Get("/api/data", func(c *fiber.Ctx) error {
		// Connect to the remote server via SSH and retrieve data
		data, err := fetchDataViaSSH()
		if err != nil {
			log.Println("Error fetching data:", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch data",
			})
		}

		// Return the data as JSON
		return c.JSON(data)
	})

	log.Fatal(app.Listen(":8083"))

}

func fetchDataViaSSH() (*Data, error) {
	// Replace these with your server's details
	host := "your-ssh-host"
	port := "22"
	username := "your-ssh-username"
	password := "your-ssh-password"

	// Connect to the SSH server
	client, err := ssh.Dial("tcp", host+":"+port, &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	})
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// Example command to execute on the remote server
	command := "your-command-to-fetch-data"

	// Run the command on the remote server
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	output, err := session.CombinedOutput(command)
	if err != nil {
		return nil, err
	}

	// Parse the output into your desired data structure
	var data Data
	data.Items = parseOutput(string(output))

	return &data, nil
}

func parseOutput(output string) []string {
	lines := strings.Split(output, "\n")

	// Return the parsed data
	return lines
}
