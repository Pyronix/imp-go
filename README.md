# IMP Go

![GitHub Workflow Status](https://github.com/Pyronix/imp-go/workflows/CI/badge.svg)
[![codecov](https://codecov.io/gh/Pyronix/imp-go/branch/main/graph/badge.svg)](https://codecov.io/gh/Pyronix/imp-go)

As a apart of the elective course "Modellbasierte Softwareentwicklung"   
at the Hochschule Karlsruhe – University of Applied Sciences (HKA) this group project aims at implementing a
- Typechecker
- Evaluator
- Parser

for a given simple imperative language (IMP).

Details for IMP can be found here:  
[https://sulzmann.github.io/ModelBasedSW/imp.html#(1)](https://sulzmann.github.io/ModelBasedSW/imp.html#(1))

## Testing
### Export Test Coverage as HTML

- Use go test with the `-coverprofile` flag
- Use go tool cover to generate html

Command:
```
go test -coverprofile cov.out <path> && go tool cover -html=cov.out -o cov.out.html
```
