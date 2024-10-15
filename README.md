# Server Stats Service

## Overview
Server Stats Service is a lightweight web application that provides real-time server statistics such as CPU usage, memory usage, and system information. It is built using Go with the `gopsutil` package for system information collection and serves the stats via a web interface accessible at `http://localhost:8080/stats`.

## Features
- **CPU Usage**: Monitor the current CPU load percentage with a visual progress bar.
- **Memory Usage**: Track used and available memory, both in bytes and as a percentage, along with a progress bar.
- **System Uptime**: Display the uptime of the server in hours.
- **Platform Info**: Show system platform and version.

## Tech Stack
- **Backend**: [Go](https://golang.org)
- **System Info Library**: [gopsutil](https://github.com/shirou/gopsutil)
- **Web Server**: Go’s built-in `net/http` package

## Getting Started

### Prerequisites
To run this service, you'll need:
- **Go**: Version 1.16 or newer. [Install Go](https://golang.org/doc/install)
- **Ubuntu or Windows** server with network access.

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/server-stats-service.git
   cd server-stats-service
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the project:
   ```bash
   go build
   ```

4. Run the application:
   ```bash
   ./server-stats-service
   ```

   The service will be available at `http://localhost:8080/stats`.

### Systemd Setup (Linux)

To keep the service running after a system reboot, follow these steps:

1. Create a service file:
   ```bash
   sudo nano /etc/systemd/system/server-stats.service
   ```

2. Add the following configuration:
   ```ini
   [Unit]
   Description=Server Stats Service
   After=network.target

   [Service]
   ExecStart=/path/to/your/server-stats-service
   Restart=always

   [Install]
   WantedBy=multi-user.target
   ```

3. Enable and start the service:
   ```bash
   sudo systemctl enable server-stats
   sudo systemctl start server-stats
   ```

### OpenVPN/Localhost Access
To make the service accessible via `http://localhost:8080/stats` for users connected through a VPN, ensure that the correct network settings and routing are configured. You may also need to adjust firewall rules to allow access.

### Example
![screenshot](screenshot.png)

## Usage
Access the server statistics on your web browser:
```bash
http://your-server-ip:8080/stats
```

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing
Contributions are welcome! Please create a new issue or submit a pull request.

## Acknowledgments
- [gopsutil](https://github.com/shirou/gopsutil) for system info collection.
```

Feel free to update it with your specific details, such as repository links and paths.
