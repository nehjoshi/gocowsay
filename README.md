# 🐮 gocowsay

A Go-powered clone of `cowsay` that prints your messages in a speech balloon followed by an ASCII cow. Easily pipe in text from tools like [`fortune`](https://man7.org/linux/man-pages/man6/fortune.6.html), or supply your own message!

```
$ fortune | gocowsay
 __________________________
/ I feel like I'm diagonally \
\ parked in a parallel universe. /
 --------------------------
         \  ^__^
          \ (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

---

## 🚀 Installation

### ✅ Option 1: `go install` (Quickest)

```bash
go install github.com/nehjoshi/gocowsay/cmd/gocowsay@latest
```

Make sure your `$GOPATH/bin` or `$HOME/go/bin` is in your `PATH`. If not, add it to `~/.bashrc` by adding the following line to the end of the file:

```bash
export PATH="$HOME/go/bin:$PATH"
```

Then run:

```bash
fortune | gocowsay
```

If the installation fails, you may not have Go or Fortune installed on your system. If so, go to Option #4 to resolve this.

---

### 🛠️ Option 2: Clone & Run Locally (Development)

```bash
git clone https://github.com/nehjoshi/gocowsay.git
cd gocowsay
echo "hello world" | go run ./cmd/gocowsay
```

To build a local binary:

```bash
go build -o gocowsay ./cmd/gocowsay
./gocowsay < message.txt
```
---

### 🐳 Option 3: Run with Docker

If you don't want to install Go, you can run `gocowsay` using Docker. Use the following series of commands:

```bash
git clone https://github.com/nehjoshi/gocowsay.git
cd gocowsay
docker build -t gocowsay .
```

---

### Option 4: Use ```setup.sh``` (Recommended for Ubuntu)
1. Download or copy the ```setup.sh``` file to your machine.
2. Make the script executable (only needed once):
```bash
chmod u+x setup.sh
```
Run the script to automatically install all required dependencies and set up the project:
```bash
./setup.sh
```
This script is designed for Ubuntu-based distributions and will install Go (if missing), update your ```PATH```, install fortune, and build the project.

---

## 🐧 Usage Examples

```bash
echo "hello world" | gocowsay
fortune | gocowsay
```

If you're on Ubuntu or Debian:

```bash
sudo apt install fortune-mod
```

---

## 🧩 Project Structure

```
gocowsay/
├── go.mod
├── cowsay/
│   ├── balloon.go          # reusable logic
│   └── balloon_test.go     # unit tests
└── cmd/
    └── gocowsay/
        └── main.go         # CLI entry point
```

---

## ✅ TODO (Contributions Welcome!)

- [ ] Add animal variants (penguin, moose, etc.)
- [ ] Add command-line flags (`--thought`, `--animal=moose`)
- [ ] Ship as a binary release for easy download
- [ ] Unicode support improvements

---

## 📜 License

MIT — feel free to fork, improve, and share! 🚀

---

## 👤 Author

Made with ❤️ by [@nehjoshi](https://github.com/nehjoshi)
