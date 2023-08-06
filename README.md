# Ustimator

![Ustimator](./assets/2587060.jpeg)

Ustimator is a CLI utility for work with uniswap V2 pools. You can quicly get info about pool reserves and estimate output of a swap

## Build CLI
To build project run `make build`

### Pair info
To get uniswap V2 pair info provide address of a pool contract

For example `0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852` is the address of ETH/USDT pair

```bash
 ustimator pair info 0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852
 ```
### Estimate swap output
To estimate swap outupt provide address of a pool contract, addresses of input and output tokens and input amount

```bash
ustimator pair estimate 0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852 0xdAC17F958D2ee523a2206206994597C13D831ec7 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2 6e8
 ```

### Ethereum RPC endpoint
In may need to specify ethereum RPC url to connect to the network

For this purpose use flag `--eth-rpc`

```bash
ustimator pair info 0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852 --eth-rpc https://mainnet.infura.io/v3/YOUR_API_KEY
 ```