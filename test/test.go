// package main
// import (
//     "encoding/json"
//     "fmt"
// )

// type Person struct {
//     Name string
//     Age  int
//     City string
// }

// func main() {
//     p := Person{Name: "Prabhat", Age: 22, City: "Delhi"}

//     // Marshal: Go struct → JSON
//     var jsonData, err = json.Marshal(p)
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }

//     fmt.Println(string(jsonData))



	
// }

// func unmarshell(){
	
	
// }

package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

type User struct {
    Login        string `json:"login"`
    Name         string `json:"name"`
    PublicRepos  int    `json:"public_repos"`
    Followers    int    `json:"followers"`
    Following    int    `json:"following"`
    AvatarURL    string `json:"avatar_url"`
}

func main() {
    url := "https://api.github.com/users/Prabhat0571"

    // 1️⃣ Make GET request
    res, err := http.Get(url)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer res.Body.Close()

    // 2️⃣ Read response body
    body, err := io.ReadAll(res.Body)
    if err != nil {
        fmt.Println("Error reading body:", err)
        return
    }

    // 3️⃣ Create variable and Unmarshal
    var user User
    json.Unmarshal(body, &user)

    // 4️⃣ Print the result
    fmt.Printf("Login: %s\nName: %s\nRepos: %d\nFollowers: %d\nFollowing: %d\n",
        user.Login, user.Name, user.PublicRepos, user.Followers, user.Following)
}
