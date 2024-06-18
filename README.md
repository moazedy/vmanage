# Vehicle Data Management System

## Overview
This project serves as a robust example of a system designed to handle CRUD operations for vehicle data. It is developed using GoLang and follows best practices in software architecture and design.

## Stack
•  [**Language**](https://www.bing.com/search?form=SKPBOT&q=Language): GoLang

•  [**Database**](https://www.bing.com/search?form=SKPBOT&q=Database): PostgreSQL


## Configuration
Ensure that the PostgreSQL database is set up with the required username and password. The configuration details must be specified in the following file:
•  `configs/vmanage.yaml`


## Project Structure
•  [**Interfaces**](https://www.bing.com/search?form=SKPBOT&q=Interfaces): All interfaces are located in the `/pkg/module/vmanage/` directory.

•  [**Implementations**](https://www.bing.com/search?form=SKPBOT&q=Implementations): The implementations of the defined interfaces can be found in the `/internal/module/` directory.

•  [**Utilities**](https://www.bing.com/search?form=SKPBOT&q=Utilities): Common utilities used across the project are located in `/pkg/infra/`.


## Design Patterns and Principles
•  The project utilizes a [**generic repository factory pattern**](https://www.bing.com/search?form=SKPBOT&q=generic%20repository%20factory%20pattern) to automate the creation of the repository layer.

•  Adheres to the [**GoLang standard layout**](https://www.bing.com/search?form=SKPBOT&q=GoLang%20standard%20layout) for directory arrangement.

•  Implements [**SOLID principles**](https://www.bing.com/search?form=SKPBOT&q=SOLID%20principles) to ensure a clean and maintainable codebase.


## API and Authentication
•  Communication with the client is established through a [**REST API**](https://www.bing.com/search?form=SKPBOT&q=REST%20API).

•  Incorporates an [**OAuth2 mock implementation**](https://www.bing.com/search?form=SKPBOT&q=OAuth2%20mock%20implementation) for authentication and authorization.


## Running the Project
To run the project, execute the following command:

make run


## Data Transfer Objects (DTOs)
The request and response structures are defined and can be found in the following directory:
•  `/pkg/module/vmanage/application/dto/`


## Contributing
Contributions to the project are welcome. Please ensure to maintain the coding standards and follow the project's design patterns.

---

For more information or if you encounter any issues, please contact with moazedy@gmail.com.
