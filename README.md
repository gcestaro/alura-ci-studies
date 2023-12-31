# alura-ci-studies

- See: 
  - ```.github/workflows/go.yml```
  - ```docker-compose.yml```

- [Self hosted runners](https://docs.github.com/pt/actions/hosting-your-own-runners/managing-self-hosted-runners/about-self-hosted-runners)

- Difference between CMD and ENTRYPOINT in Dockerfile:
  - [StackOverFlow](https://stackoverflow.com/questions/21553353/what-is-the-difference-between-cmd-and-entrypoint-in-a-dockerfile)
  - [Educative IO](https://www.educative.io/answers/how-are-cmd-and-entrypoint-different-in-a-dockerfile).
  - The main difference between CMD and ENTRYPOINT in Docker is that execution-wise, the default arguments provided by CMD can get overridden, whereas those provided by ENTRYPOINT cannot.

## Running locally

- Prepare env: ```docker-compose up -d```
- Run app: ```go run main.go```
- Run tests: ```go test main_test.go```
- Build: ```go build main.go```
- Docker build: ```docker build .```
