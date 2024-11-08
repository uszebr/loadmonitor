# Load Monitor

- Job: entity that loads proccessor and memory
  - complexity level of proccessor load
  - memoryLoad level of memory(in bytes) taken during job execution(cleared after job is done or canceled)
- JobProducer creating jobs to no buffer channel
- WorkerPool executing Jobs with limited adjustable quantity of workers(gorutines) return jobProccessed channel
  