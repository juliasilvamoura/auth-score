# API Auth-Score

Authentication and authorization API with dynamic permission management and intelligent token invalidation.

The API provides a robust access control system where JWT tokens are automatically invalidated when user permissions change, ensuring enhanced security and granular control over authorizations.

## Architectural and Technology Decisions

- **Go (Golang)**: Chosen for its excellent performance in REST APIs, strong typing, and concurrency capabilities. Go also offers a robust standard library and a moderate learning curve.

- **PostgreSQL**: Relational database chosen for its reliability, ACID support, and excellent integration with Go through GORM. Ideal for storing structured data such as users and permissions.

- **Redis**: Used as a distributed store for token blacklisting, offering:
  - High performance for invalidated token queries
  - Automatic token expiration through TTL
  - Low latency for real-time validations

- **JWT (JSON Web Tokens)**: Implemented for stateless authentication, allowing:
  - Horizontal API scalability
  - Reduced database queries
  - Secure storage of user claims

- **Docker**: Containerization to ensure:
  - Consistent development environment
  - Easy deployment and scalability
  - Service isolation (API, PostgreSQL, Redis)

# Requirements:

<ul>
  <li>Use Go </li>
  <li>Use GitHub for code storage</li>
  <li>Use PostgreSQL</li>
  <li>Use Docker</li>
</ul>

# Environment Prerequisites:

Install Go 1.23.2
<a href="https://go.dev/doc/install">Go</a>

# Setup Instructions:
<ul>
<li>Run command: <i>go mod tidy</i> to install dependencies</li>
<li>Run command: <i>docker-compose up -d</i> to start containers</li>
<li>Run command: <i>go run src/cmd/seed/main.go</i> to seed the database</li>
<li>Run command: <i>go run .</i> to start the application</li>
<li>Run command: <i>docker exec -it auth-score-redis redis-cli</i> to access Redis CLI</l>
<li>Run command: <i>KEYS blacklist:*</i> to view blacklisted tokens</l>
<li>Run command: <i>TTL blacklist:<token></i> to check token expiration time</l>
</ul>

The server will be available at `http://localhost:8080`. 

# Technologies:
<p align="center">
<img width="65px" height="65px" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/goland/goland-original.svg" />
<img width="65px" height="65px" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/github/github-original-wordmark.svg" />
<img width="65px" height="65px" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/postgresql/postgresql-original-wordmark.svg" />
<img width="65px" height="65px" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/redis/redis-original.svg" />
</p>

# Dev:

I'm Julia, a Computer Engineer graduated from IFSP. Currently, I work as a Full-Stack Developer with a strong interest in Project Management and the Financial Market.

# Social Media:

<ul>
<li><a href="https://www.linkedin.com/in/julia-m-9abba9110/" target="_blank"><i>Linkedin</i></a></li>
</ul>