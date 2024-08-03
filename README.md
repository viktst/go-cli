## Go-cli

A terminal-based app retrieving interesting facts about numbers.

### Features

- Retrieve intriguing facts about a specified number by choose the category of fact (math, trivia, date, year) or opt for a random number.

### Usage
1. Clone the repo
   ```bash
   git clone https://github.com/viktst/go-cli
   
2. Install dependencies

    ```bash
    go mod tidy

3. Execute the `root.go` file to access the help menu, also u can use `fetch --help` to view all available commands and their descriptions

### Flags
- `--number`, -n int: The number to fetch a fact about. If this is provided, `--random` cannot be used. (ex. `-n 1010`)

- `--type`, `-t` string: The type of fact to fetch. Valid values are math, trivia, date, year. Default is math. (ex. `-t trivia`)

- `--random`, `-r` bool: Fetch a random fact. If this is used, `--number` should not be provided.

### Error handling
- Conflicting Flags: If both `--random` and `--number` are provided, an error will be displayed.

- Missing Flags: If neither `--random` nor `--number` is provided, an error will be displayed.

- Invalid Fact Type: If an invalid fact type is specified, an error will be displayed.
