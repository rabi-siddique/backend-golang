## Go vs Express Comparison

| Feature         | Go HTTP Server                        | Express Server (Node.js)                |
| --------------- | ------------------------------------- | --------------------------------------- |
| **Setup**       | Minimal, built-in `net/http` package  | Requires `express` package installation |
| **Port**        | Can specify, e.g., `:8080`            | Can specify, e.g., `3000`               |
| **Routing**     | Uses `http.HandleFunc()`              | Uses `app.get()`, `app.post()`, etc.    |
| **Dependency**  | None, part of Go standard library     | Needs `express` package                 |
| **Performance** | Often faster for simple HTTP handling | Slightly slower but highly flexible     |

## JSON Response Handling: Go vs Express

| Feature             | Go JSON Response Example                 | Express JSON Response Example                    |
| ------------------- | ---------------------------------------- | ------------------------------------------------ |
| **Library/Package** | `encoding/json`                          | `express`                                        |
| **Response Method** | `json.NewEncoder(w).Encode(response)`    | `res.json(data)`                                 |
| **Struct Fields**   | Must be exported (capitalized)           | Automatically uses object properties             |
| **Error Handling**  | Manual error checking when encoding JSON | Automatic error handling and response management |
