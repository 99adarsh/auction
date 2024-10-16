package keeper_test

import (
	"testing"
	"time"

	"auction/x/auction/keeper"

	simapp "auction/app"
	"auction/x/auction/types"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

const (
	balAlice = 500000000000
	balBob   = 600000000000
	balCarol = 700000000000
)

type KeeperTestSuite struct {
	suite.Suite

	legacyAmino *codec.LegacyAmino
	ctx         sdk.Context
	app         *simapp.App
	msgServer   types.MsgServer
	queryClient types.QueryClient
}

var auctionModuleAddress string

func (suite *KeeperTestSuite) SetupTest() {
	// app := simapp.InitSideTestApp(initChain)
	app := simapp.Setup(suite.T())
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Time: time.Now().UTC()})

	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
	app.BankKeeper.SetParams(ctx, banktypes.DefaultParams())
	stakingParams := stakingtypes.DefaultParams()
	stakingParams.MinCommissionRate = sdk.OneDec()
	app.StakingKeeper.SetParams(ctx, stakingtypes.DefaultParams())

	auctionModuleAddress = app.AccountKeeper.GetModuleAddress(types.ModuleName).String()

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.AuctionKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	suite.legacyAmino = app.LegacyAmino()
	suite.ctx = ctx
	suite.app = app
	suite.queryClient = queryClient
	suite.msgServer = keeper.NewMsgServerImpl(app.AuctionKeeper)

	// Set Coins
	suite.setupSuiteWithBalances()
}

func TestKeeperSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func makeBalance(address string, balance int64, denom string) banktypes.Balance {
	return banktypes.Balance{
		Address: address,
		Coins: sdk.Coins{
			sdk.Coin{
				Denom:  denom,
				Amount: sdk.NewInt(balance),
			},
		},
	}
}

func getBankGenesis() *banktypes.GenesisState {
	coins := []banktypes.Balance{
		makeBalance(types.Alice, balAlice, simapp.DefaultBondDenom),
		makeBalance(types.Alice, balAlice, simapp.ETH),
		makeBalance(types.Alice, balAlice, simapp.WETH),
	}

	params := banktypes.DefaultParams()
	params.DefaultSendEnabled = true
	state := banktypes.NewGenesisState(
		params,
		coins,
		addAll(coins),
		[]banktypes.Metadata{}, []banktypes.SendEnabled{
			{Denom: simapp.DefaultBondDenom, Enabled: true},
			{Denom: simapp.WETH, Enabled: true},
			{Denom: simapp.ETH, Enabled: true},
		})

	return state
}

func addAll(balances []banktypes.Balance) sdk.Coins {
	total := sdk.NewCoins()
	for _, balance := range balances {
		total = total.Add(balance.Coins...)
	}
	return total
}

func (suite *KeeperTestSuite) setupSuiteWithBalances() {
	// suite.app.StakingKeeper.InitGenesis(suite.ctx, getStakeGenesis())
	suite.app.BankKeeper.InitGenesis(suite.ctx, getBankGenesis())
}
