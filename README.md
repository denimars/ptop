# ptop

`ptop` is a terminal-based process and resource monitor written in Go. It provides real-time information about CPU, memory, and running processes in a user-friendly terminal UI.

## Features
- Real-time CPU and memory usage updates
- Live process list with details
- Keyboard navigation and input handling

## Requirements
- Go 1.18 or newer
- Terminal emulator compatible with advanced keyboard interactions

## Installation
1. Clone this repository:
   ```sh
   git clone <your-repo-url>
   cd ptop
   ```
2. Download dependencies and build the application:
   ```sh
   go mod download
   go build -o ptop
   ```

## Usage
Run the application from your terminal:
```sh
./ptop
```

### Keyboard Shortcuts
- Use the keyboard to navigate and interact with the UI.
- Press `q` or `Ctrl+C` to quit the application.

## Project Structure
- `main.go`: Entry point of the application.
- `monitor/`: Contains modules for UI, process monitoring, input handling, and update loops.

## Contribution
Feel free to submit issues or pull requests to improve this project!

## License
This project is licensed under the MIT License.
