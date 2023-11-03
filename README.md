# Task Manager using Go and htmx

![Screenshot](https://github.com/DevitoDbug/taskMaster/blob/master/screenshots/screenshot1.png?raw=true)

This is a simple Task Manager web application built with Go and htmx, demonstrating the principle of minimal JavaScript in frontend development. It allows you to create, mark tasks as either pending or completed, and delete them.

## Prerequisites

- Docker: Ensure you have Docker installed to easily run this application in a container.

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/DevitoDbug/your-task-manager.git
   ```

2. Build and run the Docker container:
```bash
    cd your-task-manager
    docker build -t task-manager .
    docker run -p 8000:8000 task-manager:latest
```
You can now access the Task Manager in your web browser by opening http://localhost:8000.

## Usage
Create a new task by typing it in the input field and clicking "Add Task."
Mark tasks as "Completed" or "Pending" by clicking the appropriate buttons.
Delete tasks by clicking the "Delete" button.

## Contributing
If you would like to contribute to this project, feel free to open issues or pull requests on the GitHub repository.

## License
This project is licensed under the MIT License.

