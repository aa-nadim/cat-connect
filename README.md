# CatConnect

This project is a web application that integrates with The CatAPI to display cat images and information. The application uses `Go channels` for making API calls to The CatAPI. This allows for efficient, non-blocking operations when fetching cat data.


## Features

- Browse cat images from The Cat API
- Filter cats by breed
- View detailed information about each cat
- Add cats to favorites
- View all favorite cats
- Vote and unvote for cats
- Responsive design for various screen sizes

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go (version 1.16 or later)
- Git

## Installation

1. **Install Dependencies:**
   Ensure you have Go installed, along with the Beego CLI tool (`bee`). If you don't have `bee` installed, you can do so with:

   ```bash
   go install github.com/beego/bee/v2@latest
   ```
   For Linux:
    ```bash
    echo "export PATH=\$PATH:\$(go env GOPATH)/bin" >> ~/.bashrc
    source ~/.bashrc
    ```
    make directory
    ```bash
    mkdir -p ~/go/src/        "Windows===> mkdir C:\Users\Your_User_Name\go\src\"
    cd ~/go/src/              "Windows===> cd C:\Users\Your_User_Name\go\src\"
    ```

   Then, install project dependencies:

   ```bash
   go mod tidy
   ```

2. **Clone the Repository:**

   ```bash
   git clone https://github.com/aa-nadim/cat-connect.git
   cd cat-connect
   ```



3. **Set up your configuration:**
   Create a `conf/app.conf` file in your project root with the following content:

   ```
   appname = cat-connect
   httpport = PORT ADDRESS
   runmode = dev
   cat_api_key = YOUR_CAT_API_KEY_HERE
   StaticDir = static:static
   ```

   Replace `YOUR_CAT_API_KEY_HERE` with the API key you received from The Cat API.

4. **Generating an API Key**

    To use The Cat API, you need to generate an API key:

    1. Visit <https://thecatapi.com/>
    2. Click on the "GET YOUR API KEY" button
    3. Fill out the registration form with your email address
    4. Check your email for a message from The Cat API containing your API key
    5. Copy this API key and add it to your `conf/app.conf` file


5. **Run the Application:**

    ```bash
    go mod tidy
    bee run
    ```
6. Open your browser and navigate to `http://localhost:8080/`



## Unit Tests

```bash
go test ./... -coverprofile=coverage.out
# Display total coverage percentage
go tool cover -func=coverage.out | grep total: | awk '{print $3}'

go tool cover -html=coverage.out -o coverage.html
open coverage.html

```


