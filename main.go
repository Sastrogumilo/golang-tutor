package main

//import routes
import (
	"belajar_golang/routes"
	"log"
	"os"
	"sync"

	_ "net/http/pprof"

	"github.com/joho/godotenv"
)

var once sync.Once

func main() {

	/**
	 * Load .env file
	 */

	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}

	})

	// go func() {
	// 	http.ListenAndServe("localhost:6060", nil)
	// }()

	port := os.Getenv("PORT")

	/**
	 * Just Pick one of these, just curious, for now I will use gin
	 */

	/**
	 * Gin
	 */
	// r := routes.Routes()
	// r.Run(":" + port)

	/**
	C:\Users\Belldandy>autocannon  http://127.0.0.1:6969
	Running 10s test @ http://127.0.0.1:6969
	10 connections

	┌─────────┬──────┬──────┬───────┬──────┬─────────┬─────────┬───────┐
	│ Stat    │ 2.5% │ 50%  │ 97.5% │ 99%  │ Avg     │ Stdev   │ Max   │
	├─────────┼──────┼──────┼───────┼──────┼─────────┼─────────┼───────┤
	│ Latency │ 0 ms │ 0 ms │ 3 ms  │ 5 ms │ 0.33 ms │ 0.97 ms │ 20 ms │
	└─────────┴──────┴──────┴───────┴──────┴─────────┴─────────┴───────┘
	┌───────────┬────────┬────────┬─────────┬─────────┬─────────┬─────────┬────────┐
	│ Stat      │ 1%     │ 2.5%   │ 50%     │ 97.5%   │ Avg     │ Stdev   │ Min    │
	├───────────┼────────┼────────┼─────────┼─────────┼─────────┼─────────┼────────┤
	│ Req/Sec   │ 2993   │ 2993   │ 16095   │ 20047   │ 13470.2 │ 7249.79 │ 2993   │
	├───────────┼────────┼────────┼─────────┼─────────┼─────────┼─────────┼────────┤
	│ Bytes/Sec │ 443 kB │ 443 kB │ 2.38 MB │ 2.97 MB │ 1.99 MB │ 1.07 MB │ 443 kB │
	└───────────┴────────┴────────┴─────────┴─────────┴─────────┴─────────┴────────┘

	Req/Bytes counts sampled once per second.
	# of samples: 10

	135k requests in 10.02s, 19.9 MB read
	*/
	// ====================================================================================================
	/**
	 * Fiber
	 */

	r := routes.RoutesV2Fiber()
	r.Listen(":" + port)

	/**
	C:\Users\Belldandy>autocannon  http://127.0.0.1:6969
	Running 10s test @ http://127.0.0.1:6969
	10 connections

	┌─────────┬──────┬──────┬───────┬──────┬─────────┬─────────┬───────┐
	│ Stat    │ 2.5% │ 50%  │ 97.5% │ 99%  │ Avg     │ Stdev   │ Max   │
	├─────────┼──────┼──────┼───────┼──────┼─────────┼─────────┼───────┤
	│ Latency │ 0 ms │ 0 ms │ 0 ms  │ 0 ms │ 0.01 ms │ 0.08 ms │ 13 ms │
	└─────────┴──────┴──────┴───────┴──────┴─────────┴─────────┴───────┘
	┌───────────┬─────────┬─────────┬─────────┬─────────┬─────────┬─────────┬─────────┐
	│ Stat      │ 1%      │ 2.5%    │ 50%     │ 97.5%   │ Avg     │ Stdev   │ Min     │
	├───────────┼─────────┼─────────┼─────────┼─────────┼─────────┼─────────┼─────────┤
	│ Req/Sec   │ 31887   │ 31887   │ 36127   │ 36767   │ 35730.4 │ 1340.15 │ 31873   │
	├───────────┼─────────┼─────────┼─────────┼─────────┼─────────┼─────────┼─────────┤
	│ Bytes/Sec │ 4.24 MB │ 4.24 MB │ 4.81 MB │ 4.89 MB │ 4.75 MB │ 179 kB  │ 4.24 MB │
	└───────────┴─────────┴─────────┴─────────┴─────────┴─────────┴─────────┴─────────┘

	Req/Bytes counts sampled once per second.
	# of samples: 10

	357k requests in 10.02s, 47.5 MB read

	*/
	//====================================================================================================
}
