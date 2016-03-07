// Package nessusExporter is used to export results from Nessus when they are
// ready.
package nessusExporter

import (
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

// ExportResultPipeline is used to check the SQLite database for running scans,
// and begin querying Nessus to find out if a scan is ready to be downloaded to
// the local machine for processing.
func (e *Exporter) ExportResultPipeline() error {
	rows, err := e.sqliteDB.Query("SELECT * FROM active_scans ORDER BY request_id DESC;")
	if err != nil {
		return err
	}

	var wg *sync.WaitGroup

	for rows.Next() {
		wg.Add(1)
		var launchedScanRow launchedScanDBRow
		rows.Scan(&launchedScanRow.requestID, &launchedScanRow.method, &launchedScanRow.scanUUID, &launchedScanRow.scanID, &launchedScanRow.scanStartTime)
		go func(launchedScanRow launchedScanDBRow, e *Exporter) {
			exportedFileResponse, err := e.apiClient.ExportScan(e.httpClient, launchedScanRow.scanID, `{"format":"csv"}`)
			if err != nil {
				return
			}

			readyToExport := false
			for readyToExport {
				status, err := e.apiClient.ScanExportStatus(e.httpClient, launchedScanRow.scanID, exportedFileResponse.File)
				if err != nil {
					return
				}

				if status.Status == "ready" {
					readyToExport = true
					continue
				}

				time.Sleep(1000)
			}

			scanResults, err := e.apiClient.DownloadScan(e.httpClient, launchedScanRow.scanID, exportedFileResponse.File)
			if err != nil {
				return
			}
			filename := fmt.Sprintf("Scanner_%s-RequestID_%d-Method_%s-ScanId_%d-Time_%s.csv", getLocalIPAddress(), launchedScanRow.requestID, launchedScanRow.method, launchedScanRow.scanID, launchedScanRow.scanStartTime)
			filepath := fmt.Sprintf("%s/%s", e.fileLocations.resultsDirectory, filename)
			err = ioutil.WriteFile(filepath, []byte(scanResults), 0644)
			if err != nil {
				return
			}
		}(launchedScanRow, e)
	}

	wg.Wait()

	return nil
}