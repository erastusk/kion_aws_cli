# Kion AWS CLI Tool

A command-line interface tool for managing AWS credentials through the Kion cloud access management platform. This tool simplifies the process of obtaining temporary AWS credentials by providing an interactive interface to select and configure AWS roles managed by Kion.

## Overview

The Kion AWS CLI tool integrates with Kion (formerly cloudtamer.io) to:
- Fetch available AWS roles from your Kion account
- Present an interactive menu for role selection
- Retrieve temporary AWS credentials for the selected role
- Configure your local AWS credentials file automatically

This eliminates the need to manually copy and paste credentials from the Kion web interface.

## Features

- **Interactive Role Selection**: Browse and select from all available AWS roles using a user-friendly menu
- **Automatic Credential Management**: Fetches temporary credentials and configures your AWS CLI automatically
- **Cross-Platform Support**: Works on Windows, macOS, and Linux
- **Secure**: Uses temporary credentials that expire, improving security over long-term access keys
- **Kion Integration**: Seamless integration with Kion's cloud access management

## Prerequisites

- Go 1.21.4 or higher (for building from source)
- A Kion account with access to AWS roles
- Kion API token with appropriate permissions
- AWS CLI installed and configured (for using the credentials)

## Installation

### Building from Source

1. Clone the repository:
```bash
git clone https://github.com/erastusk/kion_aws_cli.git
cd kion_aws_cli
```

2. Install dependencies:
```bash
go mod download
```

3. Build the binary:
```bash
./build.sh
```

The compiled binary will be available in the `bin/` directory as `br` (or `brcp` depending on configuration).

### Alternative: Direct Go Build

```bash
go build -o bin/br main.go
```

## Configuration

### Setting Up Your Kion API Token

Before using the tool, you must configure your Kion API token:

1. Log in to your Kion portal
2. Navigate to your user settings and generate an API token
3. Open the file `internal/cli.go`
4. Replace `<YOUR API TOKEN>` with your actual token:

```go
const token = "Bearer your_actual_token_here"
```

**Security Note**: In a production environment, you should store the token in an environment variable or secure configuration file rather than hardcoding it.

### Recommended: Environment Variable Configuration

For better security, modify the code to read the token from an environment variable:

```bash
export KION_API_TOKEN="your_token_here"
```

Then update `internal/cli.go` to read from `os.Getenv("KION_API_TOKEN")`.

## Usage

### Basic Usage

Run the login command to select an AWS role and configure credentials:

```bash
./bin/br aws login
```

This will:
1. Fetch all available AWS roles from Kion
2. Display an interactive menu
3. Wait for you to select a role
4. Retrieve temporary credentials
5. Display the credentials (currently prints to console)

### Example Output

```
? Select AWS Role:
  ▸ 123456789012 AdminRole Production-Account
    234567890123 ReadOnlyRole Development-Account
    345678901234 PowerUserRole Staging-Account

aws_access_key_id = ASIAXXXXXXXXXXX
aws_secret_access_key = xxxxxxxxxxxxxxxxxxxxxxxxxxxx
aws_session_token = IQoJb3JpZ2luX2VjE...
```

## Project Structure

```
kion_aws_cli/
├── main.go                 # Application entry point
├── go.mod                  # Go module dependencies
├── build.sh                # Build script
├── README.md               # This file
├── LICENSE                 # License information
├── Notes                   # Development notes
├── bin/                    # Compiled binaries
│   └── br                  # Main executable
├── cmd/                    # Command definitions
│   ├── root.go            # Root command setup
│   └── aws/
│       ├── aws.go         # AWS command group
│       └── login.go       # Login command implementation
└── internal/              # Internal packages
    ├── awsconfig.go       # AWS credentials file management
    ├── cli.go             # Kion API client functions
    └── models.go          # Data models for API responses
```

## Commands

### Root Command
```bash
brcp
```
Displays help information about the tool.

### AWS Command Group
```bash
brcp aws
```
Displays help for AWS-related commands.

### Login Command
```bash
brcp aws login
```
Interactive login to select and configure AWS credentials.

## Development

### Dependencies

This project uses the following main dependencies:
- **Cobra**: CLI framework for building command-line applications
- **Survey**: Interactive prompt library for terminal user interfaces
- **Go Standard Library**: For HTTP client, JSON parsing, file I/O, etc.

### Code Organization

- **cmd/**: Contains Cobra command definitions and CLI logic
- **internal/**: Contains internal packages not meant to be imported by other projects
  - API client functions
  - Data models
  - Configuration management

### Adding New Commands

1. Create a new command file in `cmd/` or `cmd/aws/`
2. Define a new `cobra.Command` instance
3. Register it with the appropriate parent command in the `init()` function

### Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Known Issues and TODOs

### Current Limitations

1. **Credentials Not Written to File**: The `UpdateFile` function in `internal/awsconfig.go` currently only prints credentials to stdout. It needs to be implemented to actually write to the AWS credentials file.

2. **Hardcoded API Token**: The Kion API token is hardcoded in the source. This should be moved to environment variables or a secure configuration file.

3. **Limited Error Handling**: Error handling could be improved throughout the codebase.

4. **No Credential Expiration Tracking**: The tool doesn't track when temporary credentials expire.

5. **Module Name Mismatch**: The import paths use `brcp` but the go.mod file specifies `github.com/erastusk/kion_aws_cli`.

### Planned Improvements

- [ ] Implement actual file writing in `UpdateFile()` function
- [ ] Add environment variable support for API token
- [ ] Add non-interactive mode for automation
- [ ] Implement credential expiration tracking and automatic renewal
- [ ] Add support for multiple profiles
- [ ] Add unit tests
- [ ] Improve error handling and validation
- [ ] Add configuration file support (e.g., YAML or TOML)
- [ ] Add credential caching
- [ ] Support for custom Kion API endpoints

## API Endpoints

The tool interacts with the following Kion API endpoints:

- `GET /api/v3/me/cloud-access-role` - Retrieve available cloud access roles
- `GET /api/v3/account/{id}` - Get account details and labels
- `POST /api/v3/temporary-credentials` - Request temporary AWS credentials

## Security Considerations

- **Temporary Credentials**: This tool uses temporary AWS credentials which automatically expire, providing better security than long-term access keys.
- **Token Storage**: Store your Kion API token securely. Consider using environment variables or a secrets manager.
- **Credentials File**: The AWS credentials file (`~/.aws/credentials`) contains sensitive data. Ensure proper file permissions (typically `600` on Unix systems).
- **HTTPS**: All API communication with Kion uses HTTPS for encryption in transit.

## Troubleshooting

### "Could not determine os"
The tool detected an unsupported operating system. Supported platforms: Windows, Linux, macOS.

### "No such file or directory" for credentials file
The AWS credentials file doesn't exist. Create the directory and file:
```bash
mkdir -p ~/.aws
touch ~/.aws/credentials
```

### API Authentication Errors
- Verify your Kion API token is valid and properly configured
- Check that your token has the necessary permissions in Kion
- Ensure the token is prefixed with "Bearer " in the code

### Connection Errors
- Verify you can reach the Kion API endpoint from your network
- Check for firewall or proxy settings that might block the connection

## License

See the [LICENSE](LICENSE) file for details.

## Support

For issues, questions, or contributions, please open an issue on the GitHub repository.

## Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) CLI framework
- Interactive prompts powered by [Survey](https://github.com/AlecAivazis/survey)
- Integrates with [Kion](https://kion.io/) cloud access management platform

---

**Note**: This tool is designed for Broadridge Cloud Platform (BCP) integration with Kion. Adjust the Kion API endpoint in `internal/cli.go` if you're using a different Kion instance.
