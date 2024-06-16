# Adapter to create Allure reports from Ginkgo machine-readable reports

<mark>**This version works with Ginko >=2.5.0**<mark>

The way how this library works is by taking Ginkgo own machine-readable report and creating 
a corresponding Allure report structure from that.

More details about Ginkgo reporting model
can be found in [official Ginkgo documentation](https://onsi.github.io/ginkgo/#generating-reports-programmatically)

More details about Allure reports can be found in [official Allure documentation](https://docs.qameta.io/allure-report/)

### Usage:

- Add this block in your Ginkgo suite file.
```go
var _ = ReportAfterSuite("This description does nothing", func(report Report) {
    _ = allure.FromGinkgoReport(report)
})
```
- Set `REPORTS_DIR` environment variable to the desired location of the `allure-results` folder. 
If not set, `allure-results` will be created in the same folder as your `*_suite_test.go` file. 
Thus, if you have several suites, a separate results folder will be created for each of them.
- Now you can run your Ginkgo test as usual, result files will be created automatically.
- When test run is finished you can generate Allure report with allure-cli, e.g. `allure generate -c ./allure-results`
More about this can be found in Allure docs.

### Features:  
Usage examples can be found in `pkg/examples`. 
You can run these commands from repository root
```bash
export REPORTS_DIR=$(pwd) 
ginkgo -r -p --keep-going --randomize-all --randomize-suites .
allure generate -c ./allure-results
```
Sample report will be generated in the allure-report folder.

*To be continued...*

