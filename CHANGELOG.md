
## 0.4.0
* Ensured that `TYPE` and `HELP` metadata lines are only added once for every metric name as being defined multiple times was causing issues for the node_exporter textfile collector module.
* Renamed the `n2p_script_exec_lastrun` metric to `n2p_script_exec_last_execution` as that metrics should only be used with an associated script label.  The version without the script label is used to indicate the last time the application was executed.
* Ensured that if Prometheus encoded series are provided (via `raw_series` option), then an empty label value will cause that label itself to be droped from the resulting final metrics.  
* Updated `script.RunScript` function so that a result of no series is encountered with the `raw_series` option, an error will be returned to consider the execution of that script failed.


## 0.3.0
* Added missing re-initialization of metric labels at the end of each script execution as the absence of this was causing the labels to be added to each next script execution result.  This caused series to contain labels that did not belong to them.
* Specified `ScriptPath` to the `ExecutonResult` struct even when an error is encountered. 
* Added additional test scripts
* Added new `script_loaded` metric (which translates to `n2p_script_exec_script_loaded`) with a value of `1` for each script which is configured to be executed. 
* Added new `script_last_run_success` metrics (`n2p_script_exec_script_last_run_success`) metric to indicate if the latest execution of a configured script was successful

## 0.2.1
* Added new `raw_series` output type which directly accepts a text representation of prometheus series
* Added new `type` and `help` settings in configuration to give the ability to specify the corresponding `# TYPE` and `# HELP` metric comments

## 0.2.0
* Scripts are now defined via a configuration which also allows you to indicate how to parse a scripts execution result. 
* Added individual checkpoint metric per script so that users have the ability to tell if a given script execution result is stale.
* Renamed output series prefix

## 0.1.0
* Initial Release