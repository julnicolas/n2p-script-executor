package executor

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

// WorkQueue is the struct for the work queue
type WorkQueue struct {
	numWorkers      int
	pendingWorkChan chan string
	ResultsChan     chan ExecutionResult
	doneChan        chan interface{}
	Wg              *sync.WaitGroup
}

// NewWorkQueue returns a new instance of WorkQueue
func NewWorkQueue(capacity int, maxWorkers int, totalScripts int) *WorkQueue {
	wq := &WorkQueue{
		numWorkers:      maxWorkers,
		pendingWorkChan: make(chan string, capacity),
		ResultsChan:     make(chan ExecutionResult, capacity),
		doneChan:        make(chan interface{}, 1),
		Wg:              &sync.WaitGroup{},
	}
	return wq
}

// SubmitTask adds a new script execution task to the channel
func (w *WorkQueue) SubmitTask(scriptPath string) {
	log.Debug("Submiting script ", scriptPath, " to be executed...")
	w.Wg.Add(1)
	w.pendingWorkChan <- scriptPath
	log.Debug("Script submitted")
}

// Process starts the workers that will process jobs
func (w *WorkQueue) Process() {
	for n := 1; n <= w.numWorkers; n++ {
		log.Debug("Starting Execution Worker #", n)
		go w.execWorker(n)
	}
}

func (w *WorkQueue) execWorker(id int) {
	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case script := <-w.pendingWorkChan:
			log.Debugf("[Worker #%d] Running script %s", id, script)
			metric, err := RunScript(script, "exitcode", 15)
			if err != nil {
				log.Errorf("[Worker #%d] Encountered error executing script %s (Exit Code: %v, Error: %v)", id, script, metric, err)
			} else {
				log.Debugf("[Worker #%d] Script %s completed execution (Exit Code: %v)", id, script, metric)
			}
			w.ResultsChan <- ExecutionResult{
				ScriptPath: script,
				Metric:     metric,
				Error:      err,
			}
		case <-w.doneChan:
			log.Debugf("[Worker #%d] Shutting down", id)
			return
		case <-ticker.C:
			log.Tracef("[Worker #%d] Waiting for work", id)
		}
	}
}
