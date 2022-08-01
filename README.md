# HW 2 Worker Pool implementation

```text
DONE! Initial implementation, time elapsed: 112.24 seconds for 100 / 1 
DONE! With pool, single worker, time elapsed: 101.22 seconds for 100 / 1 
DONE! With pool, couple workers, time elapsed: 50.72 seconds for 100 / 2 
DONE! With pool, 10% workers, time elapsed: 10.27 seconds for 100 / 10 
DONE! With pool, 20% workers, time elapsed: 5.23 seconds for 100 / 20 
DONE! With pool, half queue workers, time elapsed: 2.30 seconds for 100 / 50 
DONE! With pool, full queue workers, time elapsed: 1.37 seconds for 100 / 100 
DONE! With pool, exceed queue workers, time elapsed: 1.40 seconds for 100 / 1000 
```
|                     | ETA     |                        |
|---------------------|---------|------------------------|
| not optimized       | 112.24s | ********************** |
| single worker       | 101.22s | ********************   |
| 2 workers           | 50.720s | **********             |
| 10% workers, 10     | 10.27s  | **                     |
| 20% workers, 20     | 5.23s   | *                      |
| 50% workers, 50     | 2.30s   | .                      |
| 100% workers, 100   | 1.37s   | .                      |
| 1000% workers, 1000 | 1.40s   | .                      |
