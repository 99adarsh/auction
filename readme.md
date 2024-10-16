# auction
**auction** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

Any user can create a new auction for any item with minimum starting price and for specified number of block duration. Different users can put their bid which should be higher than the earlier bid or the minimum specified amount for the auction.

Once the block height reaches the auctionEndHeight, the auction is then finalized and removed from the active auction list.

## Get started

This blockchain is built using ignite CLI version v0.27.2. Run the below command to start the auction blockchain.

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

## Steps to interact with the auction module

There are two ways in which a user can interact with the auction module, one is by creating auction and the other is by participating in the auction.

### 1. Create auction

```sh
auctiond tx auction create-auction mobilephone 1000 200 --from alice
```

It will create an auction for mobilephone with initial price 1000 and the auction will get finalized after 200 blocks.

### 2. Participate in auction (Bidding)

```sh
auctiond tx auction place-bid mobilephone179 1200 --from bob
```

It will add a bid for the auction-id mobilephone179 with the amount 1200. The bid is placed from the bob account.

### Query Active Auctions
To list all the active auctions,

```sh
auctiond q auction list-active-auctions-list
```

It will provide you with the list of auction-ids with auctionEndHeight. To query the info for any particular auction, 

```sh
auctiond q auction show-auction-info mobile18
```

It will return with the auction-info, giving auction end height, currentHighestBid and the bidder etc.
```
auctionInfo:
  auctionEndHeight: "818"
  auctionId: mobile18
  currentHighestBid: "1200"
  currentHighestBidder: cosmos17hz34cs4cx2tqvd3stvchgtkvvjrjpcusrxftt
  itemName: mobile
  startingPrice: "1000"
```

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/auction@latest! | sudo bash
```
`username/auction` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
