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
- **Singleton Pattern**: Utilized in `lib/container/container.go` to ensure that only one instance of the container is created, managing logging, database connections, and environment variables efficiently.

- ** **
- ** **
- ** **


## Conclusion
This microservices architecture not only demonstrates best practices in software design but also incorporates essential features for modern application development, such as CI/CD, performance monitoring, and robust testing frameworks. By leveraging these technologies, developers can build scalable, maintainable, and high-performing applications.
