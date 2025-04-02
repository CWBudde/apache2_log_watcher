# apache2watcher

A simple CLI tool written in Go to monitor Apache2 access logs and send notifications (via Signal or Email) when someone visits your site. Mostly written with the help of AI for development speed reasons. So don't expect too much from it.

---

## Features

- Watch Apache access logs in real-time
- Send alerts via **Signal** or **Email**
- Filter log lines using `--grep`
- Configure everything via a simple YAML file

---

## Installation

```bash
git clone https://github.com/yourusername/apache2watcher.git
cd apache2watcher
go build -o apache2watcher
```
## Usage

```bash
./apache2watcher watch --config=config.yaml --grep="GET /"
```

## Configuration

The configuration file is in YAML format. Here's an example:

```yaml
channel: signal

# Signal settings
signal_from: "your number"
signal_to: "target number"

# Email settings (if using email channel)
smtp_server: "smtp.example.com"
smtp_port: 587
smtp_user: "you@example.com"
smtp_pass: "yourpassword"
email_from: "you@example.com"
email_to: "target@example.com"
```

## Requirements

- Apache2 with access log enabled (/var/log/apache2/access.log)
- signal-cli installed and configured (if using Signal)
- SMTP credentials (if using Email)