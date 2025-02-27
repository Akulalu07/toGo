# toGo - CLI for Task Management

![Gopher](gopher.png)

toGo is a simple and convenient CLI tool for managing notes and tasks. With it, you can save and delete notes and tasks, making it an ideal assistant for organizing your time.

## Features
- Add, delete, and view tasks
- Simple command-line interface

## Prerequisites
Make sure you have the following installed:
- `git`
- `bash`

## Installation

To install toGo, simply download the repository and run the `build.sh` script:

```bash
git clone https://github.com/Akulalu07/toGo.git
cd toGo
./build.sh
```

## Usage

Here are some examples of commands you can use with toGo:

- **Add a task:** 
  ```bash
  toGo add "Buy milk"
  ```
- **Delete a task:** 
  ```bash
  toGo del 1
  ```
- **View all tasks:** 
  ```bash
  toGo list
  ```

## Error Handling
If you encounter issues with the database, you can try running:
```bash
sudo toGo [command]
```
(Note: This command is intended to display the logo.)

## Contributing
If you'd like to contribute to toGo, please fork the repository and submit a pull request.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

Let me know if you need any further assistance!

```
Feel free to adjust the suggestions based on your specific needs and the functionality of your tool!
```