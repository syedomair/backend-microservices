# Repository Overview

This repository showcases a robust microservices architecture comprising three distinct services: **Department**, **User**, and **Point**. The project emphasizes code reusability through a common container that initializes essential components such as logging, database connections, and environment variables.

## Services
- Department Service:
  Exposes a REST API endpoint: `/api/departments/v1/departments` for retrieving department data.
- User Service:
  Provides a REST API endpoint: `/api/users/v1/users` for accessing user information.
- Point Service:
  Utilizes gRPC to deliver user point data, which is consumed by the User service, effectively demonstrating inter-service communication.

This design promotes modularity and scalability across the services.

## Key Features

### CI/CD Integration:
The repository includes CI/CD workflows located in `.github/workflows`, which automate the deployment process to AWS Elastic Container Registry (ECR) and Elastic Container Service (ECS) servers. This ensures seamless updates and efficient management of service deployments.

### Performance Monitoring
- **Prometheus Metrics**: Integrated Prometheus metrics allow users to monitor the performance of each service in real-time. This feature provides insights into system health and resource utilization.
  
- **Memory Profiling with pprof**: 
  The project includes pprof for memory monitoring, enabling developers to analyze memory usage and optimize performance effectively.

### Testing Framework
- **Integration Testing**: 
  The system performs integration testing using a mock database running in a test Docker container. This setup ensures that all services interact correctly and maintain data integrity during operations.

- **Unit Testing**: 
  Comprehensive unit tests cover all code components, ensuring high code quality and reliability. Each service is rigorously tested to validate functionality and catch potential issues early in the development cycle.

### Design Patterns and Architectural Choices
- **Concurrency Pattern**: Utilized in `service/user_service/user/user_service` to execute multiple database queries and gRPC calls concurrently using Go's errgroup, enhancing the performance of the GetAllUserStatistics method by leveraging parallel processing.

- **Dependency Injection Pattern**: Utilized in `lib/container/container.go` to manage logging, database connections, and environment variables by injecting these dependencies into a central container, promoting modularity and flexibility in the application architecture.

- **Singleton Pattern**: Implemented in `lib/container/db.go` through synchronized lazy initialization (sync.Mutex + instance check) in PostgresAdapter and MySQLAdapter, ensuring only one database connection instance is created per adapter while maintaining thread safety.

- **Adapter Pattern**: Used in `lib/container/db.go` to create a unified database interface (Db) with concrete implementations (PostgresAdapter and MySQLAdapter), enabling seamless switching between database providers without modifying client code.

- **Factory Pattern**: Utilized in `lib/container/db.go` through the NewDBConnectionAdapter function, which acts as a factory method to create instances of different database adapters (PostgresAdapter, MySQLAdapter) based on the specified database type, encapsulating object creation logic and promoting flexibility in database provider selection.

- **External Configuration Pattern**: Utilized in `lib/container/container.go` to manage and validate essential configuration through environment variables, ensuring centralized and type-safe access to settings while promoting flexibility and ease of deployment across different environments.

- **Decorator Pattern**: Utilized in `lib/response/response.go` to dynamically add behaviors to response handlers, such as setting headers or handling different response types, without altering the underlying handler implementation.

- **Middleware Pattern**: Utilized in `lib/router/router.go` to chain multiple handlers that add functionalities like logging, request ID management, and Prometheus metrics collection, enhancing the HTTP request processing pipeline with modular and reusable components. 

- **Object Pool Pattern**: Implemented in `lib/container/connection.go` to manage a pool of reusable gRPC client connections, optimizing resource usage and improving performance by reducing the overhead of repeatedly creating and destroying connections.


## Conclusion
This microservices architecture not only demonstrates best practices in software design but also incorporates essential features for modern application development, such as CI/CD, performance monitoring, and robust testing frameworks. By leveraging these technologies, developers can build scalable, maintainable, and high-performing applications.
