
# backend-microservices

This repository showcases a microservices architecture with three distinct services: Department, User, and Point. <br/>
The project emphasizes code reusability through a common container for initializing essential components like logging, database connections, and environment variables. 
- Department service exposes a REST API endpoint /api/departments/v1/departments for retrieving department data, 
- User service provides a similar endpoint /api/user/v1/users for user information.
- Point service leverages gRPC to offer user point data, which is then consumed by the User service, demonstrating inter-service communication.



This design promotes modularity and scalability across the services.
