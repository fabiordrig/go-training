---
marp: true
theme: default
class:
  - lead
  - invert
---

# **Technology Choices: Rust and Hexagonal Architecture**

---

## Why Rust?

- Performance comparable to C/C++
- Memory safety without garbage collection
- Rich type system and ownership model
- Concurrency support

---

## Why Hexagonal Architecture?

- Promotes a clean separation of concerns
- Facilitates testability and maintainability
- Enables flexibility in external service integration
- Supports evolutionary design

---

# **CI/CD Pipeline**

---

## Continuous Integration (CI)

- Unit and integration testing
- Code quality checks
- Merge request validation

---

## Continuous Delivery (CD)

- Automated deployment to staging
- Health checks and rollback mechanisms
- Final review before production deployment

---

# **Testing Practices**

---

## Integration Testing with External APIs

- **Edge Case Handling**: For each new edge case not covered in documentation, we created specific tests to ensure correct parsing and functionality.
- **Real-World Scenarios**: Tests simulated real-world interactions with external APIs to validate our system's robustness.

---

## Database Integration Testing

- **In-Memory Database Mocking**: To optimize test speed and reduce costs, we mocked the database in memory.
- **Efficiency in CI**: This approach accelerated CI processes on GitHub Actions, making it more cost-effective and faster.

---

# **Learnings and Reflections**

---

## Key Learnings

- **Importance of Raw Data**: Keeping raw data is crucial when documentation is inaccurate.
- **Partial Data Display**: Showing even partial data can accelerate analysis.
- **Necessity of Retry Mechanisms**: Retries minimize manual intervention when data sources are offline.
- **Understanding Source Update Frequency**: Knowing update frequency of each data source optimizes resources.

---

## Reflections

- **Closer Collaboration with Internal Team**: Would engage more with the internal team early on, leveraging their knowledge of data source update frequencies.
- **Investing in Flexible Parsing**: Spending more time developing a parser that can handle various data types, e.g., treating strings where numbers are expected.
- **Avoid Rigid Enums and Constants**: Given the lack of standardization in data, a more flexible approach to parsing rules would be beneficial, avoiding strict enums and constants in code.
