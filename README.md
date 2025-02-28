# This assignment rejected by its reviewers

They didn't share feedback why, but you can check `code_challenge.pdf` file to see the requirements and code.

I found a new mathematical expression to solve issue in the easiest way. You can investigate it on below.

# TLDR
Run the application go to http://localhost:8081/swagger/index.html and click `Try it out` button to test the api.

# API SETUP INSTRUCTIONS

1. Clone the repository
2. Run `docker compose up` to start the application
3. Hit http://localhost:8081 to access the application
4. Hit http://localhost:8080/swagger/index.html to access Swagger API DOC


# HOW TO RUN TESTS
```go
go test ./...
```

# Swagger & Testing UI

Run the application then hit http://localhost:8081/swagger/index.html to access Swagger API DOC

You can simply test api responses using swagger documentations. Click on endpoint and hit `Try it out` button to test the api.

# Weight Calculation
For order size: `501`

There are three options for packaging:
1. 1x500 + 1x250
2. 1x1000 (more items than necessary)
3. 3x250 (more packs than necessary)

For each case here is the math formula: `ΔQ = orderSize - ∑Quantity`, `ΔQ * PackSize`

Results:
1. `(ΔQ = 249) * 2 (package size) = 498`
2. `(ΔQ = 499) * 1 = 499`
3. `(ΔQ = 249) * 3 = 747`

From the above results, the first option is the best choice.