
# Blocify Blockchain Explorer

Blocify is a simple blockchain explorer that allows users to explore and visualize blockchain data, including blocks, transactions, and addresses.

## Features

- View the latest blocks in the blockchain.
- Explore individual block details, including hash, timestamp, and data.
- Responsive web interface for easy access on desktop and mobile devices.

## File Structure

```
blocify/
├── README.md              # Project documentation
├── blocify.go             # Entry point of the Golang application
├── blocify.py             # Entry point of the Python application
├── blockchain/            # Blockchain data structure and functions
│   ├── blockchain.go      # Golang implementation
│   └── blockchain.py      # Python implementation
├── explorer/              # Explorer logic and web server
│   ├── explorer.go        # Golang implementation
│   └── explorer.py        # Python implementation
├── templates/             # HTML templates for web interface
│   ├── index.html         # Main page template
│   └── block.html         # Block details page template
└── static/                 # Static assets (CSS, JS, images)
    ├── style.css          # CSS styles
    └── script.js          # JavaScript logic
```

## Getting Started

To run the Blocify blockchain explorer project:

1. Install dependencies:
   - For the Golang version, ensure you have `gorilla/mux` installed (`go get github.com/gorilla/mux`).
   - For the Python version, ensure you have Flask installed (`pip install flask`).

2. Compile/run the Golang version:
   ```
   go build blocify.go
   ./blocify
   ```

3. Run the Python version:
   ```
   python blocify.py
   ```

4. Access the explorer:
   - Golang version: http://localhost:8080
   - Python version: http://localhost:5000

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
