e2e-local:
	timestamp=$(date +'%Y%m%d%H%M')
    report_filename="${timestamp}_linodego_test_report.xml"
	go test -v 2>&1 ./... | go-junit-report -set-exit-code > "$report_filename"
