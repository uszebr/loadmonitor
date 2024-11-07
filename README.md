# Load Monitor

Emulating loading on GoLang using dynamic quantity of gorutines with dynamic amount of calculations for each with ability to change memory consumption for each gorutine. CleanUp memory after gorutine is finished. 
Load Monitor can help with calculating appropriate quantity of gorutines for different workload and system.

- Job: entity that loads proccessor and memory
  - complexity: level of proccessor load
  - memoryLoad: level of memory(in bytes) taken during job execution(cleared after job is done or canceled)
- JobProducer: creating jobs to no buffer channel
- WorkerPoo:l executing Jobs with limited adjustable quantity of workers(gorutines) return jobProccessed channel


 # Dependencies
 - Go
    - Gin
    - Templ
  -  HTMX
  -  Bootstrap
