
# File Explorer Application

Welcome to the File Explorer Application! This project is a web-based tool designed to navigate and manage files and directories on your system. With a user-friendly interface and efficient backend functionality, you can easily list, search, and explore your files.

## Features

- **Directory Navigation**: Navigate through directories and view subdirectories.
- **File Searching**: Search for files by name and get complete paths of all matching files.
- **Dynamic Rendering**: Use HTMX for smooth page transitions without full page reloads.
- **Responsive Design**: The application is built with Tailwind CSS, ensuring a responsive layout across devices.

## Tech Stack

- **Backend**: Go (Golang)
- **Frontend**: HTML, CSS (Tailwind CSS), HTMX
- **File System Management**: Custom `fakeStorage` implementation for simulating file system operations.

## Getting Started

### Prerequisites

- Go 1.16 or later
- A web browser

### Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd file-explorer
   ```

2. Install dependencies (if any):
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
    make build
   ```

4. Run the application:
   ```bash
    make watch
   ```


5. Open your browser and navigate to `http://localhost:8080` to access the application.

## Usage

### Listing Files and Directories

- You can view the entire directory structure starting from the root. Simply click on the directory names to navigate into them.

### Searching for Files

- Use the search form to input the name of the file you want to find. The application will display all matching files along with their full paths.

## Contributing

Contributions are welcome! If you have suggestions for improvements or new features, please open an issue or submit a pull request.

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/YourFeature`).
3. Make your changes and commit them (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Open a pull request.

## Contact

For any questions or inquiries, feel free to reach out to me on [LinkedIn](your-linkedin-url) or via email at [your-email@example.com].