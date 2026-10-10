package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const baseURL = "https://jsonplaceholder.typicode.com"

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Website  string `json:"website"`
}

var client = &http.Client{Timeout: 12 * time.Second}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("======================================")
	fmt.Println("   GO API CLIENT — WEBSITE DATA DEMO")
	fmt.Println("======================================")
	fmt.Println("This program retrieves live JSON responses from a public REST API.")
	fmt.Println("Source: https://jsonplaceholder.typicode.com")
	fmt.Println()

	for {
		fmt.Println("Choose an action:")
		fmt.Println("1. Show the 10 latest posts")
		fmt.Println("2. Find a post by ID")
		fmt.Println("3. Show posts by user ID")
		fmt.Println("4. Show user information by ID")
		fmt.Println("0. Exit")
		fmt.Print("> ")

		choice := readLine(reader)
		fmt.Println()

		switch choice {
		case "1":
			showLatestPosts()
		case "2":
			id := askPositiveInt(reader, "Enter post ID (1-100): ")
			if id > 0 {
				showPost(id)
			}
		case "3":
			id := askPositiveInt(reader, "Enter user ID (1-10): ")
			if id > 0 {
				showUserPosts(id)
			}
		case "4":
			id := askPositiveInt(reader, "Enter user ID (1-10): ")
			if id > 0 {
				showUser(id)
			}
		case "0":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Unknown option. Enter 1, 2, 3, 4 or 0.")
		}
		fmt.Println()
	}
}

func showLatestPosts() {
	var posts []Post
	if err := getJSON(baseURL+"/posts", &posts); err != nil {
		printAPIError(err)
		return
	}
	if len(posts) > 10 {
		posts = posts[:10]
	}
	fmt.Printf("Retrieved %d posts:\n\n", len(posts))
	for _, post := range posts {
		printPost(post)
	}
}

func showPost(id int) {
	var post Post
	url := fmt.Sprintf("%s/posts/%d", baseURL, id)
	if err := getJSON(url, &post); err != nil {
		printAPIError(err)
		return
	}
	if post.ID == 0 {
		fmt.Println("No post found with that ID.")
		return
	}
	printPost(post)
}

func showUserPosts(userID int) {
	var posts []Post
	url := fmt.Sprintf("%s/posts?userId=%d", baseURL, userID)
	if err := getJSON(url, &posts); err != nil {
		printAPIError(err)
		return
	}
	if len(posts) == 0 {
		fmt.Println("No posts found for this user.")
		return
	}
	fmt.Printf("Posts for user %d (%d found):\n\n", userID, len(posts))
	for _, post := range posts {
		printPost(post)
	}
}

func showUser(id int) {
	var user User
	url := fmt.Sprintf("%s/users/%d", baseURL, id)
	if err := getJSON(url, &user); err != nil {
		printAPIError(err)
		return
	}
	if user.ID == 0 {
		fmt.Println("No user found with that ID.")
		return
	}

	fmt.Println("User information")
	fmt.Println("----------------")
	fmt.Printf("ID:       %d\n", user.ID)
	fmt.Printf("Name:     %s\n", user.Name)
	fmt.Printf("Username: %s\n", user.Username)
	fmt.Printf("Email:    %s\n", user.Email)
	fmt.Printf("Phone:    %s\n", user.Phone)
	fmt.Printf("Website:  %s\n", user.Website)
}

func getJSON(url string, target any) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("could not create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Go-API-Client/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
		return fmt.Errorf("API returned HTTP %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	decoder := json.NewDecoder(io.LimitReader(resp.Body, 2*1024*1024))
	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("could not decode JSON response: %w", err)
	}
	return nil
}

func printPost(post Post) {
	fmt.Printf("[%d] %s\n", post.ID, post.Title)
	fmt.Printf("    User ID: %d\n", post.UserID)
	fmt.Printf("    %s\n\n", post.Body)
}

func askPositiveInt(reader *bufio.Reader, prompt string) int {
	fmt.Print(prompt)
	value := readLine(reader)
	n, err := strconv.Atoi(value)
	if err != nil || n <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return 0
	}
	return n
}

func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func printAPIError(err error) {
	fmt.Println("Could not retrieve data from the API.")
	fmt.Println("Details:", err)
	fmt.Println("Check your internet connection and try again.")
}
