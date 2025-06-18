#!/bin/bash

# Exit on error
set -e

echo "[*] Beginning setup..."

# Install Go if not already installed
if ! command -v go &> /dev/null; then
    echo "[*] Installing Go..."
    sudo apt update && sudo apt install golang-go -y
else
    echo "[*] Go is already installed"
fi

# Add Go Binary path to ~/.bashrc if not present
if ! grep -q 'export PATH="$HOME/go/bin:$PATH"' ~/.bashrc; then
  echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.bashrc
  echo "[+] Added \$HOME/go/bin to PATH in ~/.bashrc"
else
  echo "[*] PATH already updated in ~/.bashrc"
fi

# Update ~/.bashrc to apply changes in current session
export PATH="$HOME/go/bin:$PATH"

# Install the Go program
echo "[*] Installing gocowsay..."
go install github.com/nehjoshi/gocowsay/cmd/gocowsay@latest

# Check if fortune is installed
if ! command -v fortune &> /dev/null; then
    echo "[*] Installing 'fortune' package..."
    sudo apt update && sudo apt install fortune -y
else
    echo "[*] 'fortune' is already installed"
fi

echo "[*] Running: fortune | gocowsay"
fortune | gocowsay
