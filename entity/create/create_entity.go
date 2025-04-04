package create

import (
	"fmt"
	"math/big"
	"os"
	"os/signal"

	"github.com/Golem-Base/golembase-demo/account/pkg/useraccount"
	"github.com/Golem-Base/golembase-demo/pkg/address"
	"github.com/Golem-Base/golembase-demo/pkg/defaults"
	"github.com/Golem-Base/golembase-demo/pkg/storagetx"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/urfave/cli/v2"
)

func Create() *cli.Command {

	cfg := struct {
		nodeURL string
		data    string
		ttl     uint64
	}{}
	return &cli.Command{
		Name:  "create",
		Usage: "Create a new entity",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "node-url",
				Usage:       "The URL of the node to connect to",
				Value:       defaults.NodeURL,
				EnvVars:     []string{"NODE_URL"},
				Destination: &cfg.nodeURL,
			},
			&cli.StringFlag{
				Name:        "data",
				Usage:       "data for the create operation",
				Value:       "this is a test",
				EnvVars:     []string{"ENTITY_DATA"},
				Destination: &cfg.data,
			},
			&cli.Uint64Flag{
				Name:        "ttl",
				Usage:       "ttl for the create operation",
				Value:       100,
				EnvVars:     []string{"ENTITY_TTL"},
				Destination: &cfg.ttl,
			},
		},
		Action: func(c *cli.Context) error {

			ctx, cancel := signal.NotifyContext(c.Context, os.Interrupt)
			defer cancel()

			userAccount, err := useraccount.Load()
			if err != nil {
				return fmt.Errorf("failed to load user account: %w", err)
			}

			// Connect to the geth node
			client, err := ethclient.DialContext(ctx, cfg.nodeURL)
			if err != nil {
				return fmt.Errorf("failed to connect to node: %w", err)
			}
			defer client.Close()

			// Get the chain ID
			chainID, err := client.ChainID(ctx)
			if err != nil {
				return fmt.Errorf("failed to get chain ID: %w", err)
			}

			// Get the nonce for the sender account
			nonce, err := client.PendingNonceAt(ctx, userAccount.Address)
			if err != nil {
				return fmt.Errorf("failed to get nonce: %w", err)
			}

			// Create the storage transaction
			storageTx := &storagetx.StorageTransaction{
				Create: []storagetx.Create{
					{
						TTL:     c.Uint64("ttl"),
						Payload: []byte(c.String("data")),
						StringAnnotations: []storagetx.StringAnnotation{
							{
								Key:   "foo",
								Value: "bar",
							},
						},
					},
				},
			}

			// Encode the storage transaction
			txData, err := rlp.EncodeToBytes(storageTx)
			if err != nil {
				return fmt.Errorf("failed to encode storage tx: %w", err)
			}

			// Create the GolemBaseUpdateStorageTx
			tx := &types.DynamicFeeTx{
				ChainID:   chainID,
				Nonce:     nonce,
				Gas:       1_000_000,
				Data:      txData,
				To:        &address.GolemBaseStorageProcessorAddress,
				GasTipCap: big.NewInt(1e9), // 1 Gwei
				GasFeeCap: big.NewInt(5e9), // 5 Gwei
			}

			// Use the London signer since we're using a dynamic fee transaction
			signer := types.LatestSignerForChainID(chainID)

			// return nil, fmt.Errorf("signer: %#v", signer)

			// Create and sign the transaction
			signedTx, err := types.SignNewTx(userAccount.PrivateKey, signer, tx)
			if err != nil {
				return fmt.Errorf("failed to sign transaction: %w", err)
			}

			txHash := signedTx.Hash()

			err = client.SendTransaction(ctx, signedTx)
			if err != nil {
				return fmt.Errorf("failed to send tx: %w", err)
			}

			receipt, err := bind.WaitMinedHash(ctx, client, txHash)
			if err != nil {
				return fmt.Errorf("failed to wait for tx: %w", err)
			}

			if receipt.Status != types.ReceiptStatusSuccessful {
				return fmt.Errorf("tx failed")
			}

			for _, log := range receipt.Logs {
				if log.Topics[0] == storagetx.GolemBaseStorageEntityCreated {
					fmt.Println("Entity created", "key", log.Topics[1])
				}
			}

			return nil
		},
	}
}

func pointerOf[T any](v T) *T {
	return &v
}
