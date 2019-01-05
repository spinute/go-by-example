# プログラムを実行すると5つの仕事が色々なワーカーに実行される様子を確認できる。
# 仕事は合わせて5秒かかるはずだが、プログラムは2秒ほどしか使わない。
# これは、平行に処理を行うワーカーが3ついるためである。
$ time go run worker-pools.go 
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5

real	0m2.358s
