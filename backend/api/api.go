package api

import (
	"encoding/json"
	"go-vue-docker/config"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func ServerGo(cfg config.Config) {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173, http://localhost:8083", // Lista de orígenes permitidos
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",               // Métodos permitidos
	}))

	app.Get("/api/posts", func(c *fiber.Ctx) error {
		posts, err := fetchPosts()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Unable to fetch posts",
			})
		}
		return c.JSON(posts)
	})

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": cfg.SSH_PORT})
	})

	app.Get("/api/data", func(c *fiber.Ctx) error {
		// Connect to the remote server via SSH and retrieve data
		data, err := fetchDataViaSSH(cfg)
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

func fetchDataViaSSH(cfg config.Config) (*Data, error) {
	// Replace these with your server's details
	host := cfg.SSH_HOST
	port := cfg.SSH_PORT
	username := cfg.SSH_USERNAME
	password := cfg.SSH_PASS

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

func fetchPosts() ([]Post, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var posts []Post
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		return nil, err
	}
	return posts, nil
}
