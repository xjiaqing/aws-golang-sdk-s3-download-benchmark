### Run
```
go run main.go --cpc <concurrency per vCPU> --key <download file key> --bucket <download file bucket> --region <s3 bucket region code>
    --cpc default 1
    --region default us-west-2
```


### Environment

- Region: us-west-2

### Instance Type
| Instance Type | vCPUs | Memory | Bandwidth |
|:-------------:|:-----:|:------:|:---------:|
| [c6in.8xlarge](https://aws.amazon.com/cn/ec2/instance-types/) | 32 | 64G |  50Gbps |

### Test Result

#### Partition Size: 64M
| concurrency	| concurrency per vCPU |	cost time (seconds) |	bandwidth (Gbps) |
|:----:|:------:|:---------:|:-----:|
| 64   |    2	| 122.07736	| 33.55 |
| 96   |	3	| 93.13486	| 43.98 |
| 128  |	4	| 92.02816	| 44.51 |
| 160  |	5	| 93.26285	| 43.92 |
| 192  |    6	| 90.65462	| 45.18 |
| 224  |    7	| 92.71951	| 44.18 |
| 256  |	8	| 90.95211	| 45.03 |
| 288  |	9	| 90.54326	| 45.24 | 
| 320  |	10	| 91.31812	| 44.85 |
| 640  |	20	| 92.706	| 44.18 |
| 960  |	30	| 92.8664	| 44.11 | 
| 1280 |	40	| 90.23964	| 45.39 |
| 1600 |	50	| 90.75476	| 45.13 |
| 1920 |	60	| 93.27354	| 43.91 |
| 2240 |	70	| 91.02245	| 45    |
| 2560 |	80	| 93.48433	| 43.81 |
| 2880 |	90	| 90.63142	| 45.19 |
| 3200 |	100	| 92.24451	| 44.4  |
| 6400 |	200	| 104.04163	| 39.37 |
| 8000 |	250	| 121.26304	| 33.78 |
| 9600 |    300	| 109.26268	| 37.49 |
| 12800 |	400	| 113.98811	| 35.93 |
