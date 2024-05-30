# go-fusion-pbx
This is a RPA (Robotic Process Automation) project that uses the FusionPBX application to automate some tasks.

This is a starting point for the project, and it will be updated as the project progresses.

## Requirements

- FusionPBX instaled and configured. I'm currently using the version 5.1.0
- configuration file as follows:
```json
{
    "Username": "admin",
    "Password": "SuperSecretPassword",
    "BaseLink": "https://voip01.example.com",
}
```

## Usage

This project is used as a library, so you can import it and use it in your own project.

```go
package main

import (
    "github.com/pericles-luz/go-fusion-pbx"
)
```

## Functionalities

- [x] Add an agent to a queue: `Callcenter.AddAgent`
- [x] Remove an agent from a queue: `Callcenter.RemoveAgent`
- [ ] Create a new agent: `Callcenter.CreateAgent`
- [ ] Delete an agent: `Callcenter.DeleteAgent`
- [ ] Create a new queue: `Callcenter.CreateQueue`
- [ ] Delete a queue: `Callcenter.DeleteQueue`

The unchecked functionalities are under development.

To use the functionalities, follow the example given in the test file `gear_test.go`.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.