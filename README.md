<div style="text-align: center;">
  <img src="symbol.png" alt="AutoISO Logo" width="128" height="128">

  <h1>AutoISO - Ubuntu Autoinstall ISO Generator</h1>

  <a href="https://golang.org/">
    <img src="https://img.shields.io/badge/Go-1.24.5+-00ADD8?style=for-the-badge&logo=go" alt="Go Version">
  </a>
  <a href="https://gin-gonic.com/">
    <img src="https://img.shields.io/badge/Gin-Web%20Framework-00ADD8?style=for-the-badge" alt="Gin Framework">
  </a>
  <a href="LICENSE">
    <img src="https://img.shields.io/badge/License-MIT-blue?style=for-the-badge" alt="License">
  </a>
  <a href="http://localhost:8080/swagger/index.html">
    <img src="https://img.shields.io/badge/API-Swagger-85EA2D?style=for-the-badge&logo=swagger" alt="API Docs">
  </a>
</div>


AutoISO is a powerful web-based tool built with Go and Gin framework for automatically generating customized Ubuntu Server ISO images with cloud-init autoinstall configurations. It streamlines the process of creating unattended Ubuntu installations with custom configurations.

## Features

### Configuration Management
- **Interactive Configuration Builder** - Web-based interface for creating autoinstall configurations
- **YAML Configuration Support** - Import/export configurations in YAML format
- **Template System** - Pre-built templates for common installation scenarios
- **Configuration Validation** - Real-time validation of cloud-init configurations
- **Preview Mode** - Preview generated user-data before ISO creation

### ISO Generation & Processing
- **Dual Source Support** - Use local ISO files or download from Ubuntu mirrors
- **Automated Downloads** - Fetch latest Ubuntu Server images (Focal 20.04, Jammy 22.04, Noble 24.04)
- **GPG Verification** - Optional cryptographic verification of downloaded ISOs
- **Custom Package Integration** - Include additional packages in the installation
- **HWE Kernel Support** - Hardware Enablement stack for newer hardware
- **MD5 Checksum Updates** - Maintain integrity validation for modified ISOs

### Web Interface & API
- **RESTful API** - Complete HTTP API for all operations
- **Real-time Progress Tracking** - Monitor ISO generation progress with detailed logs
- **File Upload Support** - Upload custom ISO files through web interface
- **Swagger Documentation** - Interactive API documentation
- **Build Status Management** - Track multiple concurrent ISO generation tasks
- **Download Management** - Direct download of generated ISO files

### Cloud-Init Integration
- **Full cloud-init Support** - Generate complete user-data configurations
- **Network Configuration** - Advanced networking setup including static IPs, bridges, bonds
- **User Management** - Configure users, SSH keys, and authentication
- **Package Management** - Specify packages to install during deployment
- **Storage Configuration** - Custom disk partitioning and filesystem setup
- **Post-install Scripts** - Execute custom commands after installation



## Quick Start

### Prerequisites
- **Go 1.24.5+** - [Download Go](https://golang.org/dl/)
- **Linux/macOS/Windows** - Cross-platform support
- **7zip/p7zip** - For ISO extraction and creation
- **xorriso** - For ISO remastering (Linux/macOS)
- **Internet connection** - For downloading Ubuntu ISOs and packages

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd autouiso
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Build the application**
   ```bash
   go build -o autouiso
   ```

4. **Run the server**
   ```bash
   ./autouiso
   ```

5. **Access the application**
   - Web Interface: http://localhost:8080
   - API Documentation: http://localhost:8080/swagger/index.html

### Docker Deployment

```bash
# Build Docker image
docker build -t autouiso .

# Run with Docker Compose
docker-compose up -d
```

## 📚 Usage Examples

### 1. Generate User-Data Configuration

```bash
curl -X POST http://localhost:8080/userdata/generate \
  -H "Content-Type: application/json" \
  -d '{
    "config": {
      "hostname": "ubuntu-server",
      "username": "admin",
      "password": "$6$rounds=4096$saltsalt$hash",
      "timezone": "UTC",
      "packages": ["curl", "wget", "git"]
    }
  }'
```

### 2. Generate ISO from Downloaded Ubuntu Image

```bash
curl -X POST http://localhost:8080/iso/generate \
  -H "Content-Type: application/json" \
  -d '{
    "sourceType": "download",
    "codeName": "jammy",
    "destinationISO": "ubuntu-22.04-autoinstall.iso",
    "userData": "#cloud-config\n...",
    "packageList": ["htop", "vim", "git"],
    "useHWEKernel": true,
    "md5Checksum": true,
    "gpgVerify": true
  }'
```

### 3. Generate ISO from Local File

```bash
curl -X POST http://localhost:8080/iso/generate \
  -H "Content-Type: application/json" \
  -d '{
    "sourceType": "local",
    "sourceISO": "/path/to/ubuntu-22.04-server.iso",
    "destinationISO": "custom-autoinstall.iso",
    "userData": "#cloud-config\n...",
    "useHWEKernel": false,
    "md5Checksum": true
  }'
```

### 4. Monitor Build Progress

```bash
# Get build status
curl http://localhost:8080/iso/build/{buildId}/status

# Get build logs
curl http://localhost:8080/iso/build/{buildId}/logs

# Download completed ISO
curl -O http://localhost:8080/download/{buildId}
```

## 🔧 Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8080` |
| `LOG_LEVEL` | Logging level (debug, info, warn, error) | `info` |
| `TEMP_DIR` | Temporary files directory | `/tmp/iso` |
| `MAX_UPLOAD_SIZE` | Maximum ISO upload size | `10GB` |
| `DOWNLOAD_TIMEOUT` | ISO download timeout | `30m` |

### Sample Configuration File

```yaml
# examples/config.yaml
hostname: ubuntu-server
timezone: America/New_York
keyboard: us
language: en_US.UTF-8

users:
  - name: admin
    password: $6$rounds=4096$salt$hash
    groups: [sudo]
    shell: /bin/bash
    ssh_authorized_keys:
      - ssh-rsa AAAAB3... user@host

network:
  version: 2
  ethernets:
    eth0:
      dhcp4: true

packages:
  - curl
  - wget
  - git
  - htop
  - vim

storage:
  layout:
    name: direct
  config:
    - type: disk
      id: main_disk
      ptable: gpt
```

## 📡 API Endpoints

### Configuration Management
- `GET /config/default` - Get default configuration template
- `POST /config/load` - Load configuration from YAML
- `POST /config/validate` - Validate configuration

### User-Data Generation
- `POST /userdata/generate` - Generate user-data from configuration
- `POST /userdata/preview` - Preview user-data output

### ISO Operations
- `POST /iso/upload` - Upload ISO file
- `POST /iso/generate` - Generate customized ISO
- `GET /iso/build/{id}/status` - Get build status
- `GET /iso/build/{id}/logs` - Get build logs
- `GET /download/{id}` - Download generated ISO

## 🛠️ Development

### Project Structure

```
├── main.go                 # Application entry point
├── api/
│   └── handler.go         # HTTP request handlers
├── generator/
│   ├── generator.go       # Core ISO generation logic
│   ├── userdata.go       # Cloud-init user-data generation
│   └── template.go       # Configuration templates
├── config/
│   └── config.go         # Configuration structures and validation
├── utils/
│   ├── utils.go          # General utilities
│   └── passwd.go         # Password hashing utilities
└── static/               # Web UI components
    ├── index.html        # Main web interface
    ├── css/              # Stylesheets
    └── js/               # JavaScript modules
```

### Building from Source

```bash
# Install development dependencies
go mod tidy

# Run tests
go test ./...

# Build for different platforms
GOOS=linux GOARCH=amd64 go build -o autouiso-linux
GOOS=windows GOARCH=amd64 go build -o autouiso.exe
GOOS=darwin GOARCH=amd64 go build -o autouiso-macos
```

### Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📖 Documentation

- [User-Data Configuration Guide](doc/user-data.md)
- [Network Configuration](doc/network.md)
- [Storage Configuration](doc/storage.md)
- [Package Management](doc/apt.md)
- [Troubleshooting Guide](doc/problem.md)

## 🔍 Supported Ubuntu Versions

- **Ubuntu 24.04 LTS (Noble Numbat)** - Latest LTS release
- **Ubuntu 22.04 LTS (Jammy Jellyfish)** - Current LTS
- **Ubuntu 20.04 LTS (Focal Fossa)** - Previous LTS

## ⚠️ System Requirements

### Server Requirements
- **RAM**: Minimum 2GB, Recommended 4GB+
- **Storage**: 10GB+ free space for temporary files
- **CPU**: Any modern x64 processor

### Supported Platforms
- **Linux** - Primary development and testing platform
- **macOS** - Full functionality with Homebrew dependencies
- **Windows** - Basic functionality (some features may require WSL)

## 🤝 Support & Community

- **Issues**: [GitHub Issues](https://github.com/your-repo/autouiso/issues)
- **Discussions**: [GitHub Discussions](https://github.com/your-repo/autouiso/discussions)
- **Documentation**: [Wiki](https://github.com/your-repo/autouiso/wiki)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Ubuntu Team for the excellent server distribution
- Canonical for cloud-init technology
- Gin framework for the robust web framework
- The Go community for outstanding tooling and libraries

---

<div align="center">
  Made with ❤️ by the AutoISO Team
</div>
